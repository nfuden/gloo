package internal

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/solo-io/gloo/test/gomega/assertions"
	testmatchers "github.com/solo-io/gloo/test/gomega/matchers"
	"github.com/solo-io/gloo/test/gomega/transforms"
	"github.com/solo-io/go-utils/stats"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/gloosnapshot"
	defaults2 "github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	glooKube2e "github.com/solo-io/gloo/test/kube2e"

	"github.com/solo-io/solo-projects/test/kube2e"
	"github.com/solo-io/solo-projects/test/services"

	"github.com/solo-io/gloo/test/helpers"

	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/retries"

	"github.com/solo-io/go-utils/testutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"

	gatewayv1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/core/matchers"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

	kubeerrors "k8s.io/apimachinery/pkg/api/errors"

	errors "github.com/rotisserie/eris"
	extauthapi "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"

	"github.com/solo-io/k8s-utils/testutils/helper"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	gatewayv2 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gateway/pkg/defaults"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/k8s-utils/kubeutils"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"k8s.io/client-go/rest"

	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	. "github.com/solo-io/gloo/test/kube2e"
	"github.com/solo-io/solo-projects/projects/gloo/pkg/plugins/extauth"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

type cleanupFunc func()

type ExtAuthTestInputs struct {
	TestHelper *helper.SoloTestHelper

	// ShouldTestLDAP should be set to true if you're
	// testing in a cluster with an LDAP server available
	// and want to run LDAP tests.
	ShouldTestLDAP bool
}

func RunExtAuthTests(inputs *ExtAuthTestInputs) {
	var testHelper *helper.SoloTestHelper

	var _ = Describe("External auth", func() {

		const (
			gatewayPort = int(80)
			response401 = "HTTP/1.1 401 Unauthorized"
			response403 = "HTTP/1.1 403 Forbidden"
			response200 = "HTTP/1.1 200 OK"
		)

		var (
			ctx        context.Context
			cancel     context.CancelFunc
			cfg        *rest.Config
			kubeClient kubernetes.Interface

			gatewayClient        gatewayv2.GatewayClient
			proxyClient          gloov1.ProxyClient
			virtualServiceClient gatewayv1.VirtualServiceClient
			authConfigClient     extauthapi.AuthConfigClient
			settingsClient       gloov1.SettingsClient
			origSettings         *gloov1.Settings // used to capture & restore initial Settings so each test can modify them

			resourceClientset *glooKube2e.KubeResourceClientSet
			snapshotWriter    helpers.SnapshotWriter
			glooResources     *gloosnapshot.ApiSnapshot

			err error
		)

		// Credentials must be in the <username>:<password> format
		buildAuthHeader := func(credentials string) map[string]string {
			encodedCredentials := base64.StdEncoding.EncodeToString([]byte(credentials))
			return map[string]string{
				"Authorization": fmt.Sprintf("Basic %s", encodedCredentials),
			}
		}
		curlAndAssertResponse := func(path string, headers map[string]string, expectedResponseSubstring string) {
			testHelper.CurlEventuallyShouldRespond(helper.CurlOpts{
				Protocol:          "http",
				Path:              path,
				Method:            "GET",
				Headers:           headers,
				Host:              defaults.GatewayProxyName,
				Service:           defaults.GatewayProxyName,
				Port:              gatewayPort,
				ConnectionTimeout: 3,    // this is important, as the first curl call sometimes hangs indefinitely
				Verbose:           true, // this is important, as curl will only output status codes with verbose output
			}, expectedResponseSubstring, 1, 2*time.Minute)
		}

		// This just registers the clients that we will need during the tests
		BeforeEach(func() {
			testHelper = inputs.TestHelper
			ctx, cancel = context.WithCancel(context.Background())

			cfg, err = kubeutils.GetConfig("", "")
			Expect(err).NotTo(HaveOccurred())

			resourceClientset, err = glooKube2e.NewKubeResourceClientSet(ctx, cfg)
			Expect(err).NotTo(HaveOccurred())

			kubeClient = resourceClientset.KubeClients()
			authConfigClient = resourceClientset.AuthConfigClient()
			gatewayClient = resourceClientset.GatewayClient()
			virtualServiceClient = resourceClientset.VirtualServiceClient()
			proxyClient = resourceClientset.ProxyClient()
			settingsClient = resourceClientset.SettingsClient()

			// Not all tests in this file use the SnapshotWriter or glooResources snapshot
			// The idea is to progressively move towards this model, because it's used in oss
			// and is easier to set up/teardown resources consistently
			snapshotWriter = helpers.NewSnapshotWriter(resourceClientset)
			glooResources = &gloosnapshot.ApiSnapshot{}
		})

		BeforeEach(func() {
			origSettings, err = settingsClient.Read(testHelper.InstallNamespace, "default", clients.ReadOpts{Ctx: ctx})
			Expect(err).NotTo(HaveOccurred(), "Should be able to read initial settings")
		})

		AfterEach(func() {
			currentSettings, err := settingsClient.Read(testHelper.InstallNamespace, "default", clients.ReadOpts{Ctx: ctx})
			Expect(err).NotTo(HaveOccurred(), "Should be able to read current settings")

			if origSettings.Metadata.ResourceVersion != currentSettings.Metadata.ResourceVersion {
				origSettings.Metadata.ResourceVersion = currentSettings.Metadata.ResourceVersion // so we can overwrite settings
				_, err = settingsClient.Write(origSettings, clients.WriteOpts{Ctx: ctx, OverwriteExisting: true})
				Expect(err).ToNot(HaveOccurred())
			}
			cancel()
		})

		// Only run LDAP tests if explicitly set to true. Otherwise cluster may not have an LDAP server running.
		if inputs.ShouldTestLDAP {
			Describe("authenticate requests via LDAP", func() {

				var (
					extAuthConfigProto   *extauthapi.ExtAuthExtension
					disableGroupChecking bool
					ldapConfig           = func(namespace string, disableGroupChecking bool) *extauthapi.Ldap {
						return &extauthapi.Ldap{
							Address:        fmt.Sprintf("ldap.%s.svc.cluster.local:389", namespace),
							UserDnTemplate: "uid=%s,ou=people,dc=solo,dc=io",
							AllowedGroups: []string{
								"cn=managers,ou=groups,dc=solo,dc=io",
							},
							DisableGroupChecking: disableGroupChecking,
						}
					}
				)

				JustBeforeEach(func() {

					By("make sure we can still reach the LDAP server", func() {
						// Make sure we can query the LDAP server
						testHelper.CurlEventuallyShouldRespond(helper.CurlOpts{
							Protocol:          "ldap",
							Path:              "/",
							Method:            "GET",
							Service:           fmt.Sprintf("ldap.%s.svc.cluster.local", testHelper.InstallNamespace),
							Port:              389,
							ConnectionTimeout: 3,
							Verbose:           true,
						}, "OpenLDAProotDSE", 1, time.Minute)
					})

					By("create an LDAP-secured route to the test upstream", func() {

						virtualHostPlugins := &gloov1.VirtualHostOptions{
							Extauth: extAuthConfigProto,
						}

						kube2e.WriteVirtualService(ctx, testHelper, virtualServiceClient, virtualHostPlugins, nil, nil)

						defaultGateway := defaults.DefaultGateway(testHelper.InstallNamespace)
						Eventually(func() (*gatewayv2.Gateway, error) {
							return gatewayClient.Read(testHelper.InstallNamespace, defaultGateway.Metadata.Name, clients.ReadOpts{})
						}, "15s", "0.5s").Should(Not(BeNil()))
					})
				})

				AfterEach(func() {
					kube2e.DeleteVirtualService(virtualServiceClient, testHelper.InstallNamespace, "vs", clients.DeleteOpts{Ctx: ctx, IgnoreNotExist: true})
				})

				allTests := func(groupCheckingDisabled bool) {
					It("works as expected ", func() {
						By("returns 401 if no authentication header is provided", func() {
							curlAndAssertResponse(kube2e.TestMatcherPrefix, nil, response401)
						})

						By("returns 401 if the user is unknown", func() {
							curlAndAssertResponse(kube2e.TestMatcherPrefix, buildAuthHeader("john:doe"), response401)
						})

						By("returns 200 if the user belongs to one of the allowed groups", func() {
							curlAndAssertResponse(kube2e.TestMatcherPrefix, buildAuthHeader("rick:rickpwd"), response200)
						})

						if groupCheckingDisabled {
							By("returns 200 if the user does not belong to the allowed groups but group checking is disabled", func() {
								curlAndAssertResponse(kube2e.TestMatcherPrefix, buildAuthHeader("marco:marcopwd"), response200)
							})
						} else {
							By("returns 403 if the user does not belong to the allowed groups", func() {
								curlAndAssertResponse(kube2e.TestMatcherPrefix, buildAuthHeader("marco:marcopwd"), response403)
							})
						}
					})
				}

				BeforeEach(func() {
					authConfig, err := authConfigClient.Write(&extauthapi.AuthConfig{
						Metadata: &core.Metadata{
							Name:      "ldap",
							Namespace: testHelper.InstallNamespace,
						},
						Configs: []*extauthapi.AuthConfig_Config{{
							AuthConfig: &extauthapi.AuthConfig_Config_Ldap{
								Ldap: ldapConfig(testHelper.InstallNamespace, disableGroupChecking),
							},
						}},
					}, clients.WriteOpts{Ctx: ctx})
					Expect(err).NotTo(HaveOccurred())

					authConfigRef := authConfig.Metadata.Ref()
					extAuthConfigProto = &extauthapi.ExtAuthExtension{
						Spec: &extauthapi.ExtAuthExtension_ConfigRef{
							ConfigRef: authConfigRef,
						},
					}
				})

				AfterEach(func() {
					err := authConfigClient.Delete(testHelper.InstallNamespace, "ldap", clients.DeleteOpts{Ctx: ctx})
					Expect(err).NotTo(HaveOccurred())
				})

				Context("as a standalone deployment", func() {
					// no extra setup to do, just run the tests
					allTests(false)
				})

				Context("as a sidecar", func() {
					JustBeforeEach(func() {
						settings, err := settingsClient.Read(testHelper.InstallNamespace, "default", clients.ReadOpts{Ctx: ctx})
						Expect(err).NotTo(HaveOccurred(), "Should be able to read settings to switch ext auth mode")

						newRef := core.ResourceRef{
							Name:      extauth.SidecarUpstreamName,
							Namespace: testHelper.InstallNamespace,
						}
						extauthSettings := &extauthapi.Settings{
							ExtauthzServerRef: &newRef,
						}
						settings.Extauth = extauthSettings

						_, err = settingsClient.Write(settings, clients.WriteOpts{Ctx: ctx, OverwriteExisting: true})
						Expect(err).NotTo(HaveOccurred(), "Should be able to write new ext auth settings")
					})

					allTests(false)
				})

				Context("With DisableGroupChecking", func() {
					BeforeEach(func() {
						disableGroupChecking = true
					})

					It("works as expected ", func() {
						allTests(true)
					})
				})

			})
		}

		// These tests create a virtual host with two routes to two simple http-echo services. Each spec then proceeds to
		// define different permutations of extauth configs on the virtual host and on both routes and tests that requests
		// are allowed/denied correctly.
		// NOTE: we use BasicAuth configs both because of its simplicity and because in case a request is denied it returns
		// a 401 instead of the standard 403, allowing us to distinguish between legitimate denials and fallback behavior
		// if some error in the auth flow occurred.
		Describe("multiple authentication configs defined at different levels", func() {

			const (
				appName1    = "test-app-1"
				appName2    = "test-app-2"
				echoAppPort = int32(5678)
			)

			var (
				cleanUpFuncs []cleanupFunc
				allowUser,
				allowAdmin,
				allowUserAndAdmin *extauthapi.AuthConfig
			)

			writeAuthConfig := func(authConfig *extauthapi.AuthConfig) (*core.ResourceRef, cleanupFunc) {
				ac, err := authConfigClient.Write(authConfig, clients.WriteOpts{Ctx: ctx})
				ExpectWithOffset(1, err).NotTo(HaveOccurred())

				// Wait for auth config to be created
				helpers.EventuallyResourceStatusMatchesState(1, func() (resources.InputResource, error) {
					return authConfigClient.Read(testHelper.InstallNamespace, ac.Metadata.Name, clients.ReadOpts{Ctx: ctx})
				}, core.Status_Accepted)

				authConfigRef := ac.Metadata.Ref()
				return authConfigRef, func() {
					err := authConfigClient.Delete(ac.Metadata.Namespace, ac.Metadata.Name, clients.DeleteOpts{Ctx: ctx, IgnoreNotExist: true})
					Expect(err).NotTo(HaveOccurred())
				}
			}

			buildExtAuthExtension := func(authConfigRef *core.ResourceRef) *extauthapi.ExtAuthExtension {
				return &extauthapi.ExtAuthExtension{
					Spec: &extauthapi.ExtAuthExtension_ConfigRef{
						ConfigRef: authConfigRef,
					},
				}
			}

			getDisableAuthExtension := func() *extauthapi.ExtAuthExtension {
				return &extauthapi.ExtAuthExtension{
					Spec: &extauthapi.ExtAuthExtension_Disable{
						Disable: true,
					},
				}
			}

			writeVs := func(vs *gatewayv1.VirtualService) cleanupFunc {

				// We wrap this in a eventually because the validating webhook may reject the virtual service if one of the
				// resources the VS depends on is not yet available.
				EventuallyWithOffset(1, func() error {
					_, err := virtualServiceClient.Write(vs, clients.WriteOpts{Ctx: ctx, OverwriteExisting: true})
					return err
				}, "30s", "1s").Should(BeNil())

				// Wait for proxy to be accepted
				helpers.EventuallyResourceStatusMatchesState(1, func() (resources.InputResource, error) {
					return proxyClient.Read(testHelper.InstallNamespace, defaults.GatewayProxyName, clients.ReadOpts{Ctx: ctx})
				}, core.Status_Accepted)

				return func() {
					kube2e.DeleteVirtualService(virtualServiceClient, vs.Metadata.Namespace, vs.Metadata.Name, clients.DeleteOpts{Ctx: ctx, IgnoreNotExist: true})
				}
			}

			isRolloutComplete := func(name, namespace string) (bool, error) {
				result, err := testutils.KubectlOut("rollout", "status", fmt.Sprintf("deployment/%s", name), "-n", namespace)
				return strings.Contains(result, "successfully rolled out"), err
			}

			BeforeEach(func() {
				// Make sure to reset this to avoid errors during cleanup
				cleanUpFuncs = nil

				// Create two target http-echo deployments/services
				cleanUp1, err := createHttpEchoDeploymentAndService(ctx, kubeClient, testHelper.InstallNamespace, appName1, echoAppPort)
				Expect(err).NotTo(HaveOccurred())
				cleanUpFuncs = append(cleanUpFuncs, cleanUp1)

				cleanUp2, err := createHttpEchoDeploymentAndService(ctx, kubeClient, testHelper.InstallNamespace, appName2, echoAppPort)
				Expect(err).NotTo(HaveOccurred())
				cleanUpFuncs = append(cleanUpFuncs, cleanUp2)

				// Wait for both target http-echo deployments to be ready
				Eventually(func() (bool, error) {
					return isRolloutComplete(appName1, testHelper.InstallNamespace)
				}, "30s", "1s").Should(BeTrue())
				Eventually(func() (bool, error) {
					return isRolloutComplete(appName2, testHelper.InstallNamespace)
				}, "30s", "1s").Should(BeTrue())

				// Define the three types of auth configuration we will use
				allowUser = buildBasicAuthConfig("basic-auth-user", testHelper.InstallNamespace,
					map[string]*extauthapi.BasicAuth_Apr_SaltedHashedPassword{
						"user": {
							// generated with: `htpasswd -nbm user password`
							// `user:$apr1$0adzfifo$14o4fMw/Pm2L34SvyyA2r.`
							Salt:           "0adzfifo",
							HashedPassword: "14o4fMw/Pm2L34SvyyA2r.",
						},
					},
				)
				allowAdmin = buildBasicAuthConfig("basic-auth-admin", testHelper.InstallNamespace,
					map[string]*extauthapi.BasicAuth_Apr_SaltedHashedPassword{
						"admin": {
							// generated with: `htpasswd -nbm admin password`
							// `admin:$apr1$o6PF7xkS$O8kUlD9Xbyj6l1WPcgWxM1`
							Salt:           "o6PF7xkS",
							HashedPassword: "O8kUlD9Xbyj6l1WPcgWxM1",
						},
					},
				)
				allowUserAndAdmin = buildBasicAuthConfig("basic-auth-user-and-admin", testHelper.InstallNamespace,
					map[string]*extauthapi.BasicAuth_Apr_SaltedHashedPassword{
						"user": {
							// generated with: `htpasswd -nbm user password`
							// `user:$apr1$0adzfifo$14o4fMw/Pm2L34SvyyA2r.`
							Salt:           "0adzfifo",
							HashedPassword: "14o4fMw/Pm2L34SvyyA2r.",
						},
						"admin": {
							// generated with: `htpasswd -nbm admin password`
							// `admin:$apr1$o6PF7xkS$O8kUlD9Xbyj6l1WPcgWxM1`
							Salt:           "o6PF7xkS",
							HashedPassword: "O8kUlD9Xbyj6l1WPcgWxM1",
						},
					},
				)

				// bounce envoy, get a clean state (draining listener can break this test). see https://github.com/solo-io/solo-projects/issues/2921 for more.
				out, err := services.KubectlOut(strings.Split("rollout restart -n "+testHelper.InstallNamespace+" deploy/gateway-proxy", " ")...)
				fmt.Println(out)
				Expect(err).ToNot(HaveOccurred())
				out, err = services.KubectlOut(strings.Split("rollout status -n "+testHelper.InstallNamespace+" deploy/gateway-proxy", " ")...)
				fmt.Println(out)
				Expect(err).ToNot(HaveOccurred())
			})

			AfterEach(func() {
				// delete the resources in reverse order because we need to delete the VS that references the auth config
				for i := len(cleanUpFuncs) - 1; i >= 0; i-- {
					cleanup := cleanUpFuncs[i]
					if cleanup != nil {
						cleanup()
					}
				}
			})

			Context("single destinations", func() {

				var (
					vhPlugins                    *gloov1.VirtualHostOptions
					route1Plugins, route2Plugins *gloov1.RouteOptions
				)

				getVirtualService := func(vhPlugins *gloov1.VirtualHostOptions, pluginsForRoute1, pluginsForRoute2 *gloov1.RouteOptions) *gatewayv1.VirtualService {
					return &gatewayv1.VirtualService{
						Metadata: &core.Metadata{
							Name:      "echo-vs",
							Namespace: testHelper.InstallNamespace,
						},
						VirtualHost: &gatewayv1.VirtualHost{
							Options: vhPlugins,
							Domains: []string{"*"},
							Routes: []*gatewayv1.Route{
								{
									Options: pluginsForRoute1,
									Matchers: []*matchers.Matcher{{
										PathSpecifier: &matchers.Matcher_Prefix{
											Prefix: kube2e.TestMatcherPrefix + "/1",
										},
									}},
									Action: &gatewayv1.Route_RouteAction{
										RouteAction: &gloov1.RouteAction{
											Destination: &gloov1.RouteAction_Single{
												Single: &gloov1.Destination{
													DestinationType: &gloov1.Destination_Kube{
														Kube: &gloov1.KubernetesServiceDestination{
															Ref: &core.ResourceRef{
																Namespace: testHelper.InstallNamespace,
																Name:      appName1,
															},
															Port: uint32(echoAppPort),
														},
													},
												},
											},
										},
									},
								},
								{
									Options: pluginsForRoute2,
									Matchers: []*matchers.Matcher{{
										PathSpecifier: &matchers.Matcher_Prefix{
											Prefix: kube2e.TestMatcherPrefix + "/2",
										},
									}},
									Action: &gatewayv1.Route_RouteAction{
										RouteAction: &gloov1.RouteAction{
											Destination: &gloov1.RouteAction_Single{
												Single: &gloov1.Destination{
													DestinationType: &gloov1.Destination_Kube{
														Kube: &gloov1.KubernetesServiceDestination{
															Ref: &core.ResourceRef{
																Namespace: testHelper.InstallNamespace,
																Name:      appName2,
															},
															Port: uint32(echoAppPort),
														},
													},
												},
											},
										},
									},
								},
							},
						},
					}
				}

				BeforeEach(func() {
					// Make sure to re-initialize these shared variables before each test
					vhPlugins, route1Plugins, route2Plugins = nil, nil, nil
				})

				JustBeforeEach(func() {
					By("write virtual service and wait for it to be accepted", func() {

						// The arguments for this function will be set by each test as needed
						virtualService := getVirtualService(vhPlugins, route1Plugins, route2Plugins)

						cleanup := writeVs(virtualService)

						// Add a func to delete the VS to the AfterEach cleanups
						cleanUpFuncs = append(cleanUpFuncs, cleanup)
					})
				})

				When("no auth is configured", func() {

					It("can reach both services without providing credentials", func() {
						curlAndAssertResponse(kube2e.TestMatcherPrefix+"/1", nil, appName1)
						curlAndAssertResponse(kube2e.TestMatcherPrefix+"/2", nil, appName2)
					})
				})

				When("auth is configured only on the virtual host", func() {

					BeforeEach(func() {
						authConfigRef, cleanUpAuthConfig := writeAuthConfig(allowUser)
						cleanUpFuncs = append(cleanUpFuncs, cleanUpAuthConfig)

						extension := buildExtAuthExtension(authConfigRef)
						vhPlugins = &gloov1.VirtualHostOptions{Extauth: extension}
					})

					It("behaves as expected", func() {

						By("denying unauthenticated requests on both routes", func() {
							curlAndAssertResponse(kube2e.TestMatcherPrefix+"/1", nil, response401)
							curlAndAssertResponse(kube2e.TestMatcherPrefix+"/2", nil, response401)
						})

						By("allowing authenticated requests on both routes", func() {
							curlAndAssertResponse(kube2e.TestMatcherPrefix+"/1", buildAuthHeader("user:password"), appName1)
							curlAndAssertResponse(kube2e.TestMatcherPrefix+"/2", buildAuthHeader("user:password"), appName2)
						})
					})
				})

				When("auth is configured on the virtual host and disabled on one child route", func() {

					BeforeEach(func() {
						authConfigRef, cleanUpAuthConfig := writeAuthConfig(allowUser)
						cleanUpFuncs = append(cleanUpFuncs, cleanUpAuthConfig)

						vhPlugins = &gloov1.VirtualHostOptions{Extauth: buildExtAuthExtension(authConfigRef)}

						//  Disable auth on route 2
						route2Plugins = &gloov1.RouteOptions{Extauth: getDisableAuthExtension()}
					})

					It("behaves as expected", func() {

						By("denying unauthenticated requests on the first route", func() {
							curlAndAssertResponse(kube2e.TestMatcherPrefix+"/1", nil, response401)
						})

						By("allowing unauthenticated requests on the second route", func() {
							curlAndAssertResponse(kube2e.TestMatcherPrefix+"/2", nil, appName2)
						})

						By("allowing authenticated requests on the first route", func() {
							curlAndAssertResponse(kube2e.TestMatcherPrefix+"/1", buildAuthHeader("user:password"), appName1)
						})
					})
				})

				When("auth is configured on both the virtual host and on one child route", func() {

					BeforeEach(func() {

						// Virtual host allows both user:password and admin:password
						vhAuthConfigRef, cleanUpVhAuthConfig := writeAuthConfig(allowUserAndAdmin)

						// Route allows admin:password
						routeAuthConfigRef, cleanUpRouteAuthConfig := writeAuthConfig(allowAdmin)

						cleanUpFuncs = append(cleanUpFuncs, cleanUpVhAuthConfig, cleanUpRouteAuthConfig)

						vhPlugins = &gloov1.VirtualHostOptions{Extauth: buildExtAuthExtension(vhAuthConfigRef)}
						route2Plugins = &gloov1.RouteOptions{Extauth: buildExtAuthExtension(routeAuthConfigRef)}
					})

					It("behaves as expected", func() {

						By("denying unauthenticated requests on both routes", func() {
							curlAndAssertResponse(kube2e.TestMatcherPrefix+"/1", nil, response401)
							curlAndAssertResponse(kube2e.TestMatcherPrefix+"/2", nil, response401)
						})

						By("allowing user:password on the first route, but denying it on the second", func() {
							curlAndAssertResponse(kube2e.TestMatcherPrefix+"/1", buildAuthHeader("user:password"), appName1)
							curlAndAssertResponse(kube2e.TestMatcherPrefix+"/2", buildAuthHeader("user:password"), response401)
						})

						By("allowing admin:password on both routes", func() {
							curlAndAssertResponse(kube2e.TestMatcherPrefix+"/1", buildAuthHeader("admin:password"), appName1)
							curlAndAssertResponse(kube2e.TestMatcherPrefix+"/2", buildAuthHeader("admin:password"), appName2)
						})
					})
				})
			})

			Context("multi destination", func() {

				var (
					routePlugins               *gloov1.RouteOptions
					dest1Plugins, dest2Plugins *gloov1.WeightedDestinationOptions
				)

				getMultiDestinationVirtualService := func(
					routePlugins *gloov1.RouteOptions,
					pluginsForDest1, pluginsForDest2 *gloov1.WeightedDestinationOptions,
				) *gatewayv1.VirtualService {
					return &gatewayv1.VirtualService{
						Metadata: &core.Metadata{
							Name:      "echo-vs",
							Namespace: testHelper.InstallNamespace,
						},
						VirtualHost: &gatewayv1.VirtualHost{
							Domains: []string{"*"},
							Routes: []*gatewayv1.Route{
								{
									Matchers: []*matchers.Matcher{{
										PathSpecifier: &matchers.Matcher_Prefix{
											Prefix: kube2e.TestMatcherPrefix,
										},
									}},
									Options: routePlugins,
									Action: &gatewayv1.Route_RouteAction{
										RouteAction: &gloov1.RouteAction{
											Destination: &gloov1.RouteAction_Multi{
												Multi: &gloov1.MultiDestination{
													Destinations: []*gloov1.WeightedDestination{
														{
															Weight: &wrappers.UInt32Value{Value: 50},
															Destination: &gloov1.Destination{
																DestinationType: &gloov1.Destination_Kube{
																	Kube: &gloov1.KubernetesServiceDestination{
																		Ref: &core.ResourceRef{
																			Namespace: testHelper.InstallNamespace,
																			Name:      appName1,
																		},
																		Port: uint32(echoAppPort),
																	},
																},
															},
															Options: pluginsForDest1,
														},
														{
															Weight: &wrappers.UInt32Value{Value: 50},
															Destination: &gloov1.Destination{
																DestinationType: &gloov1.Destination_Kube{
																	Kube: &gloov1.KubernetesServiceDestination{
																		Ref: &core.ResourceRef{
																			Namespace: testHelper.InstallNamespace,
																			Name:      appName2,
																		},
																		Port: uint32(echoAppPort),
																	},
																},
															},
															Options: pluginsForDest2,
														},
													},
												},
											},
										},
									},
								},
							},
						},
					}
				}

				BeforeEach(func() {
					// Make sure to re-initialize these shared variables before each test
					routePlugins, dest1Plugins, dest2Plugins = nil, nil, nil
				})

				JustBeforeEach(func() {
					By("write virtual service and wait for it to be accepted", func() {

						// The arguments for this function will be set by each test as needed
						virtualService := getMultiDestinationVirtualService(routePlugins, dest1Plugins, dest2Plugins)

						cleanup := writeVs(virtualService)

						// Add a func to delete the VS to the AfterEach cleanups
						cleanUpFuncs = append(cleanUpFuncs, cleanup)
					})
				})

				When("auth is configured on both the route and on one weighted destination", func() {

					BeforeEach(func() {

						// Route allows both user:password and admin:password
						routeAuthConfigRef, cleanUpRouteAuthConfig := writeAuthConfig(allowUserAndAdmin)

						// Weighted destination allows admin:password
						weightedDestAuthConfigRef, cleanUpWeightedDestAuthConfig := writeAuthConfig(allowAdmin)

						cleanUpFuncs = append(cleanUpFuncs, cleanUpRouteAuthConfig, cleanUpWeightedDestAuthConfig)

						routePlugins = &gloov1.RouteOptions{Extauth: buildExtAuthExtension(routeAuthConfigRef)}
						dest2Plugins = &gloov1.WeightedDestinationOptions{Extauth: buildExtAuthExtension(weightedDestAuthConfigRef)}
					})

					It("behaves as expected", func() {

						By("consistently denying unauthenticated requests on both routes", func() {
							for i := 0; i < 5; i++ {
								curlAndAssertResponse(kube2e.TestMatcherPrefix, nil, response401)
							}
						})

						By("consistently allowing admin:password on both destinations", func() {
							for i := 0; i < 5; i++ {
								// Just look for the substring that is common to the responses from both services
								curlAndAssertResponse(kube2e.TestMatcherPrefix, buildAuthHeader("admin:password"), "test-app-")
							}
						})

						By("allowing user:password on one route, but not on the other", func() {
							// Eventually we should get both a response from service 1 and a 401
							curlAndAssertResponse(kube2e.TestMatcherPrefix, buildAuthHeader("user:password"), appName1)
							curlAndAssertResponse(kube2e.TestMatcherPrefix, buildAuthHeader("user:password"), response401)
						})
					})
				})

				When("different auth is configured on both weighted destinations", func() {

					BeforeEach(func() {

						// Weighted destination allows user:password
						weightedDest1AuthConfigRef, cleanUpWeightedDest1AuthConfig := writeAuthConfig(allowUser)

						// Weighted destination allows admin:password
						weightedDest2AuthConfigRef, cleanUpWeightedDest2AuthConfig := writeAuthConfig(allowAdmin)

						cleanUpFuncs = append(cleanUpFuncs, cleanUpWeightedDest1AuthConfig, cleanUpWeightedDest2AuthConfig)

						dest1Plugins = &gloov1.WeightedDestinationOptions{Extauth: buildExtAuthExtension(weightedDest1AuthConfigRef)}
						dest2Plugins = &gloov1.WeightedDestinationOptions{Extauth: buildExtAuthExtension(weightedDest2AuthConfigRef)}
					})

					It("behaves as expected", func() {

						By("consistently denying unauthenticated requests on both routes", func() {
							for i := 0; i < 5; i++ {
								curlAndAssertResponse(kube2e.TestMatcherPrefix, nil, response401)
							}
						})

						By("allowing user:password on one route, but not on the other", func() {
							// Eventually we should get both a response from service 1 and a 401
							curlAndAssertResponse(kube2e.TestMatcherPrefix, buildAuthHeader("user:password"), appName1)
							curlAndAssertResponse(kube2e.TestMatcherPrefix, buildAuthHeader("user:password"), response401)
						})

						By("allowing admin:password on one route, but not on the other", func() {
							// Eventually we should get both a response from service 2 and a 401
							curlAndAssertResponse(kube2e.TestMatcherPrefix, buildAuthHeader("admin:password"), appName2)
							curlAndAssertResponse(kube2e.TestMatcherPrefix, buildAuthHeader("admin:password"), response401)
						})
					})
				})
			})

			Context("health checker", func() {

				var (
					vhPlugins *gloov1.VirtualHostOptions

					pollingRunner            *pollingRunner
					pollingResponseMutex     sync.RWMutex
					pollingResponseFrequency map[string]int
				)

				getVirtualService := func(vhPlugins *gloov1.VirtualHostOptions) *gatewayv1.VirtualService {
					return &gatewayv1.VirtualService{
						Metadata: &core.Metadata{
							Name:      "echo-vs",
							Namespace: testHelper.InstallNamespace,
						},
						VirtualHost: &gatewayv1.VirtualHost{
							Options: vhPlugins,
							Domains: []string{"*"},
							Routes: []*gatewayv1.Route{
								{
									Matchers: []*matchers.Matcher{{
										PathSpecifier: &matchers.Matcher_Prefix{
											Prefix: kube2e.TestMatcherPrefix + "/1",
										},
									}},
									Options: &gloov1.RouteOptions{
										Retries: &retries.RetryPolicy{
											RetryOn:    "5xx",
											NumRetries: 2,
										},
									},
									Action: &gatewayv1.Route_RouteAction{
										RouteAction: &gloov1.RouteAction{
											Destination: &gloov1.RouteAction_Single{
												Single: &gloov1.Destination{
													DestinationType: &gloov1.Destination_Kube{
														Kube: &gloov1.KubernetesServiceDestination{
															Ref: &core.ResourceRef{
																Namespace: testHelper.InstallNamespace,
																Name:      appName1,
															},
															Port: uint32(echoAppPort),
														},
													},
												},
											},
										},
									},
								},
							},
						},
					}
				}

				updateSettingsWithFailureModeAllow := func() {
					// These tests poll the endpoint frequently and expect all 200 responses
					// A small network drop could result in a error, which by default is a 403
					// Ideally, we would retry on this, or set a StatusOnError and retry on that. However,
					// our retry policy does not yet support RetriableStatusCodes.
					// Therefore, for now, we use FailureModeAllow, which indicates that if we fail to connect
					// to the extauth server, allow the request.

					settings, err := settingsClient.Read(testHelper.InstallNamespace, "default", clients.ReadOpts{Ctx: ctx})
					Expect(err).NotTo(HaveOccurred(), "Should be able to read settings to switch ext auth StatusOnError")

					settings.Extauth.FailureModeAllow = true

					_, err = settingsClient.Write(settings, clients.WriteOpts{Ctx: ctx, OverwriteExisting: true})
					Expect(err).NotTo(HaveOccurred(), "Should be able to write new ext auth settings")
				}

				endpointPollingWorker := func() {
					response, err := testHelper.Curl(helper.CurlOpts{
						Protocol:          "http",
						Path:              kube2e.TestMatcherPrefix + "/1",
						Method:            "GET",
						Headers:           buildAuthHeader("user:password"),
						Host:              defaults.GatewayProxyName,
						Service:           defaults.GatewayProxyName,
						Port:              gatewayPort,
						ConnectionTimeout: 5,    // this is important, as the first curl call sometimes hangs indefinitely
						Verbose:           true, // this is important, as curl will only output status codes with verbose output
					})

					// Modify the response for expected results
					if err != nil {
						response = err.Error()
					} else if strings.Contains(response, response200) {
						response = response200
					}

					// Store the response in a map
					pollingResponseMutex.Lock()
					defer pollingResponseMutex.Unlock()
					_, ok := pollingResponseFrequency[response]
					if ok {
						pollingResponseFrequency[response] += 1
					} else {
						pollingResponseFrequency[response] = 1
					}
				}

				JustBeforeEach(func() {
					// Set FailureModeAllow to true, so that network errors do not cause flakes
					updateSettingsWithFailureModeAllow()

					// Persist the auth config
					authConfigRef, cleanUpAuthConfig := writeAuthConfig(allowUser)

					// Persist virtual service
					vhPlugins = &gloov1.VirtualHostOptions{
						Extauth: buildExtAuthExtension(authConfigRef),
					}
					virtualService := getVirtualService(vhPlugins)
					cleanUpVirtualService := writeVs(virtualService)

					// This polls the endpoint at an interval and stores the responses
					pollingRunner = newPollingRunner(endpointPollingWorker, time.Millisecond*10, 5)
					pollingResponseFrequency = make(map[string]int)

					// Ensure all created resources are cleaned between tests
					cleanUpFuncs = append(cleanUpFuncs, cleanUpAuthConfig, cleanUpVirtualService)
				})

				modifyDeploymentReplicas := func(replicas int32) error {
					deploymentClient := kubeClient.AppsV1().Deployments(testHelper.InstallNamespace)

					d, err := deploymentClient.Get(ctx, "extauth", metav1.GetOptions{})
					if err != nil {
						return err
					}

					d.Spec.Replicas = &replicas
					_, err = deploymentClient.Update(ctx, d, metav1.UpdateOptions{})
					return err
				}

				getDeploymentReplicas := func() (int, error) {
					deploymentClient := kubeClient.AppsV1().Deployments(testHelper.InstallNamespace)

					d, err := deploymentClient.Get(ctx, "extauth", metav1.GetOptions{})
					if err != nil {
						return 0, err
					}

					if d.Status.UpdatedReplicas != d.Status.ReadyReplicas {
						return 0, errors.New("Updated and Ready replicas do not match")
					}

					return int(d.Status.ReadyReplicas), nil
				}

				isExtAuthRolloutComplete := func() (bool, error) {
					return isRolloutComplete("extauth", "gloo-system")
				}

				modifyDeploymentEnv := func(envVar corev1.EnvVar) error {
					deploymentClient := kubeClient.AppsV1().Deployments(testHelper.InstallNamespace)

					d, err := deploymentClient.Get(ctx, "extauth", metav1.GetOptions{})
					if err != nil {
						return err
					}

					d.Spec.Template.Spec.Containers[0].Env = append(d.Spec.Template.Spec.Containers[0].Env, envVar)
					_, err = deploymentClient.Update(ctx, d, metav1.UpdateOptions{})
					return err
				}

				allPollingResponsesAre200 := func() (bool, error) {
					keys := make([]string, 0)
					for k := range pollingResponseFrequency {
						keys = append(keys, k)
					}

					if len(keys) != 1 || keys[0] != response200 {
						return false, errors.New(fmt.Sprintf("received non-200 response: %v", strings.Join(keys, ",")))
					}
					return true, nil
				}

				// This test is added as a safeguard to ensure that we're not accepting all requests blindly
				It("denies unauthenticated requests", func() {
					curlAndAssertResponse(kube2e.TestMatcherPrefix+"/1", nil, response401)
				})

				It("allows authenticated requests on route when no cluster events happen", func() {
					// Scale the extauth deployment to 1 pod and wait for it to be ready
					err = modifyDeploymentReplicas(1)
					Expect(err).NotTo(HaveOccurred())
					Eventually(isExtAuthRolloutComplete, "60s", "1s").Should(BeTrue())

					// Ensure that the upstream is reachable
					curlAndAssertResponse(kube2e.TestMatcherPrefix+"/1", buildAuthHeader("user:password"), response200)

					pollingRunner.StartPolling(ctx)

					// Do nothing for 1 second to allow time for successful polling requests
					time.Sleep(time.Second * 1)

					pollingRunner.StopPolling()

					// Expect all responses to be 200s
					Expect(allPollingResponsesAre200()).Should(BeTrue())
				})

				It("allows authenticated requests on route when extauth deployment is modified", func() {
					// There should only be 1 pod to start
					Eventually(getDeploymentReplicas, "60s", "1s").Should(Equal(1))

					// Ensure that the upstream is reachable
					curlAndAssertResponse(kube2e.TestMatcherPrefix+"/1", buildAuthHeader("user:password"), response200)

					pollingRunner.StartPolling(ctx)

					// Modify the deployment, causing the pods to be brought up again
					err := modifyDeploymentEnv(corev1.EnvVar{
						Name:  "HEALTH_CHECKER_ENV_VAR",
						Value: "VALUE",
					})
					Expect(err).NotTo(HaveOccurred())
					Eventually(isExtAuthRolloutComplete, "60s", "1s").Should(BeTrue())

					pollingRunner.StopPolling()

					// Expect all responses to be 200s
					Expect(allPollingResponsesAre200()).Should(BeTrue())
				})

				It("allows authenticated requests on route when extauth deployment is scaled up", func() {
					// There should only be 1 pod to start
					Eventually(getDeploymentReplicas, "60s", "1s").Should(Equal(1))

					// Ensure that the upstream is reachable
					curlAndAssertResponse(kube2e.TestMatcherPrefix+"/1", buildAuthHeader("user:password"), response200)

					pollingRunner.StartPolling(ctx)

					// Scale up the extauth deployment to 4 pods and wait for them all to be ready
					err := modifyDeploymentReplicas(4)
					Expect(err).NotTo(HaveOccurred())
					Eventually(isExtAuthRolloutComplete, "60s", "1s").Should(BeTrue())

					pollingRunner.StopPolling()

					// Expect all responses to be 200s
					Expect(allPollingResponsesAre200()).Should(BeTrue())
				})

				It("allows authenticated requests on route when extauth deployment is scaled down", func() {
					// There should be 4 pods (from the previous test) to start
					Eventually(getDeploymentReplicas, "60s", "1s").Should(Equal(4))

					// Ensure that the upstream is reachable
					curlAndAssertResponse(kube2e.TestMatcherPrefix+"/1", buildAuthHeader("user:password"), response200)

					pollingRunner.StartPolling(ctx)

					// Do nothing for 1 second to allow time for successful polling requests
					time.Sleep(time.Second * 1)

					// Scale down the extauth deployment to 1 pod and wait for it to be ready
					err = modifyDeploymentReplicas(1)
					Expect(err).NotTo(HaveOccurred())
					Eventually(isExtAuthRolloutComplete, "60s", "1s").Should(BeTrue())

					pollingRunner.StopPolling()

					// Expect all responses to be 200s
					Expect(allPollingResponsesAre200()).Should(BeTrue())
				})

			})

		})

		Context("xDS Updates", func() {
			// These tests validate that we are pushing xDS updates at the proper cadence
			// Previous bugs caused xDS updates to be more frequent than necessary
			// These frequent updates don't change the behavior necessarily, but they imply that
			// we are performing potentially expensive translation cycles that are unnecessary

			const (
				eastResource    = "east"
				westResource    = "west"
				xdsResponseStat = "xds_responses{type=\"type.googleapis.com/enterprise.gloo.solo.io.ExtAuthConfig\"}"
			)

			var (
				metricsRequest        *http.Request
				xdsResponseCountRegex *regexp.Regexp
				xdsResponseCount      = 0
			)

			BeforeEach(func() {
				xdsResponseCountRegex, err = regexp.Compile(fmt.Sprintf("%s ([\\d]+)", xdsResponseStat))
				Expect(err).NotTo(HaveOccurred())

				metricsRequest, err = http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:%d/metrics", stats.DefaultPort), nil)
				Expect(err).NotTo(HaveOccurred())

				acEast := getComplexAuthConfig(eastResource, defaults2.GlooSystem)
				vsEast := helpers.NewVirtualServiceBuilder().
					WithName(eastResource).
					WithNamespace(defaults2.GlooSystem).
					WithDomain(eastResource).
					WithVirtualHostOptions(&gloov1.VirtualHostOptions{
						Extauth: &extauthapi.ExtAuthExtension{
							Spec: &extauthapi.ExtAuthExtension_ConfigRef{
								ConfigRef: acEast.GetMetadata().Ref(),
							},
						},
					}).
					WithRoutePrefixMatcher("test", "/").
					WithRouteDirectResponseAction("test", &gloov1.DirectResponseAction{
						Status: http.StatusOK,
					}).
					Build()

				acWest := getComplexAuthConfig(westResource, defaults2.GlooSystem)
				vsWest := helpers.NewVirtualServiceBuilder().
					WithName(westResource).
					WithNamespace(defaults2.GlooSystem).
					WithDomain(westResource).
					WithVirtualHostOptions(&gloov1.VirtualHostOptions{
						Extauth: &extauthapi.ExtAuthExtension{
							Spec: &extauthapi.ExtAuthExtension_ConfigRef{
								ConfigRef: acWest.GetMetadata().Ref(),
							},
						},
					}).
					WithRoutePrefixMatcher("test", "/").
					WithRouteDirectResponseAction("test", &gloov1.DirectResponseAction{
						Status: http.StatusOK,
					}).
					Build()

				glooResources.AuthConfigs = extauthapi.AuthConfigList{acEast, acWest}
				glooResources.VirtualServices = gatewayv2.VirtualServiceList{vsEast, vsWest}

				err = snapshotWriter.WriteSnapshot(glooResources, clients.WriteOpts{
					Ctx:               ctx,
					OverwriteExisting: false,
				})
				Expect(err).NotTo(HaveOccurred())
			})

			AfterEach(func() {
				err = snapshotWriter.DeleteSnapshot(glooResources, clients.DeleteOpts{
					Ctx:            ctx,
					IgnoreNotExist: true,
				})
				Expect(err).NotTo(HaveOccurred())
			})

			JustBeforeEach(func() {
				// Wait for xDS responses to stabilize
				xdsMetricAssertion, statValue := assertions.IntStatisticReachesConsistentValueAssertion(xdsResponseStat, 2)
				assertions.EventuallyStatisticsMatchAssertions(assertions.DefaultStatsPortFwd, xdsMetricAssertion)

				// Mark the stable value of that metric. Tests will compare the new count to this value
				xdsResponseCount = *statValue
			})

			It("updates to non-authconfig resources do not change xDS request count", func() {
				vsRefToUpdate := glooResources.VirtualServices[0].GetMetadata().Ref()

				By("Patch the resource to trigger a snapshot sync in the Gloo controller")
				patchErr := helpers.PatchResource(
					ctx,
					vsRefToUpdate,
					func(resource resources.Resource) resources.Resource {
						resource.GetMetadata().Labels = map[string]string{
							"unique-label": "modified",
						}
						return resource
					},
					resourceClientset.VirtualServiceClient().BaseClient(),
				)
				Expect(patchErr).NotTo(HaveOccurred())

				By("Validate that the xDS response count has not changed")
				assertions.EventuallyStatisticsMatchAssertions(assertions.DefaultStatsPortFwd,
					Consistently(func(g Gomega) {
						currentXdsResponseCount := 0
						g.Expect(http.DefaultClient.Do(metricsRequest)).To(testmatchers.HaveHttpResponse(&testmatchers.HttpResponse{
							StatusCode: http.StatusOK,
							Body: WithTransform(func(body []byte) error {
								count, transformErr := transforms.IntRegexTransform(xdsResponseCountRegex)(body)
								currentXdsResponseCount = count
								return transformErr
							}, Succeed()),
						}))
						g.Expect(currentXdsResponseCount).To(Equal(xdsResponseCount), "xds response count should not change")
					}, assertions.SafeTimeToSyncStats, "1s"),
				)
			})

			// https://github.com/solo-io/solo-projects/issues/4461
			// This test exposed a potential bug in the statistics we emit from our NackDetector
			// The test is valuable so I decided to keep it to demonstrate how we should be verifying this statistic
			// For now however, it doesn't perform the assertion that the metric increases
			It("updates to authconfig resources change xDS request count", func() {
				authRefToUpdate := glooResources.AuthConfigs[0].GetMetadata().Ref()

				By("Patch the resource to trigger a snapshot sync in the Gloo controller")
				patchErr := helpers.PatchResource(
					ctx,
					authRefToUpdate,
					func(resource resources.Resource) resources.Resource {
						authConfig := resource.(*extauthapi.AuthConfig)
						// Perform a meaningful change to the AuthConfig to ensure that it will result in a different resource hash
						authConfig.Configs[1].GetOauth2().GetAccessTokenValidation().GetIntrospection().IntrospectionUrl = "introspection-url-modified"
						return authConfig
					},
					resourceClientset.AuthConfigClient().BaseClient(),
				)
				Expect(patchErr).NotTo(HaveOccurred())

				By("Validate that the xDS response count has changed")
				statStabilizesAssertion, _ := assertions.IntStatisticReachesConsistentValueAssertion(xdsResponseStat, 2)
				assertions.EventuallyStatisticsMatchAssertions(assertions.DefaultStatsPortFwd,
					// Assert that the statistic has increased
					Eventually(func(g Gomega) {
						currentXdsResponseCount := 0
						g.Expect(http.DefaultClient.Do(metricsRequest)).To(testmatchers.HaveHttpResponse(&testmatchers.HttpResponse{
							StatusCode: http.StatusOK,
							Body: WithTransform(func(body []byte) error {
								count, transformErr := transforms.IntRegexTransform(xdsResponseCountRegex)(body)
								currentXdsResponseCount = count
								return transformErr
							}, Succeed()),
						}))

						// When https://github.com/solo-io/solo-projects/issues/4461 is fixed, this should be converted
						// to xdsResponseCount + 1
						metricAtLeast := xdsResponseCount

						g.Expect(currentXdsResponseCount).To(BeNumerically(">=", metricAtLeast), "xds response count should increase for resource patch")
					}, "1m", "1s"),
					// Assert that after increasing, its stabilizes
					statStabilizesAssertion,
				)

			})

		})
		Context("applying resources", func() {
			Context("Encryption Key secret (kube secret converter)", func() {
				const (
					oauthName              = "client-secret"
					encryptionKeyName      = "encryption-key"
					correctEncryptionKey   = "an example of an encryption key1"
					incorrectEncryptionKey = "an example of an encryption"
					authConfigName         = "encryption-key-auth-config"
				)
				var (
					encryptionKeyUsed = ""
				)
				JustBeforeEach(func() {
					namespace := testHelper.InstallNamespace
					authConfig := extauthapi.AuthConfig{
						Metadata: &core.Metadata{
							Name:      authConfigName,
							Namespace: namespace,
						},
						Configs: []*extauthapi.AuthConfig_Config{
							{
								AuthConfig: &extauthapi.AuthConfig_Config_Oauth2{
									Oauth2: &extauthapi.OAuth2{
										OauthType: &extauthapi.OAuth2_OidcAuthorizationCode{
											OidcAuthorizationCode: &extauthapi.OidcAuthorizationCode{
												AppUrl:       "http://localhost",
												CallbackPath: "/callback",
												ClientSecretRef: &core.ResourceRef{
													Name:      oauthName,
													Namespace: namespace,
												},
												IssuerUrl: "https://localhost.com/oauth2/dummy",
												ClientId:  "someclientId",
												Session: &extauthapi.UserSession{
													CipherConfig: &extauthapi.UserSession_CipherConfig{
														Key: &extauthapi.UserSession_CipherConfig_KeyRef{
															KeyRef: &core.ResourceRef{
																Name:      encryptionKeyName,
																Namespace: namespace,
															},
														},
													},
													Session: &extauthapi.UserSession_Cookie{
														Cookie: &extauthapi.UserSession_InternalSession{
															KeyPrefix: "my-prefix",
														},
													},
												},
											},
										},
									},
								},
							},
						},
					}
					encryptionKey := gloov1.Secret{
						Metadata: &core.Metadata{
							Name:      encryptionKeyName,
							Namespace: namespace,
						},
						Kind: &gloov1.Secret_Encryption{
							Encryption: &gloov1.EncryptionKeySecret{
								Key: encryptionKeyUsed,
							},
						},
					}
					oauthSecret := gloov1.Secret{
						Metadata: &core.Metadata{
							Name:      oauthName,
							Namespace: namespace,
						},
						Kind: &gloov1.Secret_Oauth{
							Oauth: &extauthapi.OauthSecret{
								ClientSecret: "this is the client secret",
							},
						},
					}
					glooResources.AuthConfigs = extauthapi.AuthConfigList{&authConfig}
					glooResources.Secrets = append(glooResources.Secrets, &encryptionKey)
					glooResources.Secrets = append(glooResources.Secrets, &oauthSecret)
					err = snapshotWriter.WriteSnapshot(glooResources, clients.WriteOpts{
						Ctx:               ctx,
						OverwriteExisting: false,
					})
					Expect(err).NotTo(HaveOccurred())
				})
				AfterEach(func() {
					err = snapshotWriter.DeleteSnapshot(glooResources, clients.DeleteOpts{
						Ctx:            ctx,
						IgnoreNotExist: true,
					})
					Expect(err).NotTo(HaveOccurred())
				})
				Context("correct encryption key", func() {
					BeforeEach(func() {
						encryptionKeyUsed = correctEncryptionKey
					})
					It("it will apply a Encryption Key secret and will not Error in Gloo or Ext-Auth", func() {
						helpers.EventuallyResourceAccepted(func() (resources.InputResource, error) {
							return authConfigClient.Read(testHelper.InstallNamespace, authConfigName, clients.ReadOpts{Ctx: ctx})
						})
					})
				})
				Context("incorrect encryption key", func() {
					BeforeEach(func() {
						encryptionKeyUsed = incorrectEncryptionKey
					})
					It("it will apply a Encryption Key incorrect secret and be rejected", func() {
						helpers.EventuallyResourceRejected(func() (resources.InputResource, error) {
							return authConfigClient.Read(testHelper.InstallNamespace, authConfigName, clients.ReadOpts{Ctx: ctx})
						})
					})
				})
			})
		})
	})
}

func createHttpEchoDeploymentAndService(ctx context.Context, kubeClient kubernetes.Interface, namespace, appName string, port int32) (cleanupFunc, error) {
	_, err := kubeClient.AppsV1().Deployments(namespace).Create(ctx, &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:   appName,
			Labels: map[string]string{"app": appName},
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": appName},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"app": appName},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Name:  "http-echo",
						Image: GetHttpEchoImage(),
						Args:  []string{fmt.Sprintf("-text=%s", expectedResponse(appName))},
						Ports: []corev1.ContainerPort{{
							Name:          "http",
							ContainerPort: port,
						}},
					}},
					// important, otherwise termination lasts 30 seconds!
					TerminationGracePeriodSeconds: pointerToInt64(0),
				},
			},
		},
	}, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	_, err = kubeClient.CoreV1().Services(namespace).Create(ctx, &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:   appName,
			Labels: map[string]string{"app": appName},
		},
		Spec: corev1.ServiceSpec{
			Type:     corev1.ServiceTypeClusterIP,
			Selector: map[string]string{"app": appName},
			Ports: []corev1.ServicePort{{
				Name: "http",
				Port: port,
			}},
		},
	}, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	return func() {
		err := kubeClient.AppsV1().Deployments(namespace).Delete(ctx, appName, metav1.DeleteOptions{GracePeriodSeconds: pointerToInt64(0)})
		ExpectWithOffset(1, err).NotTo(HaveOccurred())
		err = kubeClient.CoreV1().Services(namespace).Delete(ctx, appName, metav1.DeleteOptions{GracePeriodSeconds: pointerToInt64(0)})
		ExpectWithOffset(1, err).NotTo(HaveOccurred())

		EventuallyWithOffset(1, func() bool {
			_, err := kubeClient.AppsV1().Deployments(namespace).Get(ctx, appName, metav1.GetOptions{})
			return isNotFound(err)
		}, "10s", "0.5s").Should(BeTrue())

		EventuallyWithOffset(1, func() bool {
			_, err := kubeClient.CoreV1().Services(namespace).Get(ctx, appName, metav1.GetOptions{})
			return isNotFound(err)
		}, "10s", "0.5s").Should(BeTrue())
	}, nil
}

func expectedResponse(appName string) string {
	return fmt.Sprintf("Hello from %s!", appName)
}

func pointerToInt64(value int64) *int64 {
	return &value
}

func isNotFound(err error) bool {
	return err != nil && kubeerrors.IsNotFound(err)
}

func buildBasicAuthConfig(name, namespace string, users map[string]*extauthapi.BasicAuth_Apr_SaltedHashedPassword) *extauthapi.AuthConfig {
	return &extauthapi.AuthConfig{
		Metadata: &core.Metadata{
			Name:      name,
			Namespace: namespace,
		},
		Configs: []*extauthapi.AuthConfig_Config{{
			AuthConfig: &extauthapi.AuthConfig_Config_BasicAuth{
				BasicAuth: &extauthapi.BasicAuth{
					Realm: "gloo",
					Apr: &extauthapi.BasicAuth_Apr{
						Users: users,
					},
				},
			},
		}},
	}
}

func getComplexAuthConfig(name, namespace string) *extauthapi.AuthConfig {
	return &extauthapi.AuthConfig{
		Metadata: &core.Metadata{
			Name:      name,
			Namespace: namespace,
		},
		Configs: []*extauthapi.AuthConfig_Config{
			{
				Name: &wrappers.StringValue{Value: "jwt"},
				AuthConfig: &extauthapi.AuthConfig_Config_Jwt{
					Jwt: &empty.Empty{},
				},
			},
			{
				Name: &wrappers.StringValue{Value: "oauth"},
				AuthConfig: &extauthapi.AuthConfig_Config_Oauth2{
					Oauth2: &extauthapi.OAuth2{
						OauthType: &extauthapi.OAuth2_AccessTokenValidation{
							AccessTokenValidation: &extauthapi.AccessTokenValidation{
								ValidationType: &extauthapi.AccessTokenValidation_Introspection{
									Introspection: &extauthapi.IntrospectionValidation{
										IntrospectionUrl: fmt.Sprintf("fake-url-%s", name),
									},
								},
								UserinfoUrl: fmt.Sprintf("fake-url-%s", name),
							},
						},
					},
				},
			},
			{
				Name: &wrappers.StringValue{Value: "passthrough"},
				AuthConfig: &extauthapi.AuthConfig_Config_PassThroughAuth{
					PassThroughAuth: &extauthapi.PassThroughAuth{
						FailureModeAllow: true,
						Protocol: &extauthapi.PassThroughAuth_Grpc{
							Grpc: &extauthapi.PassThroughGrpc{
								Address: fmt.Sprintf("passthrough-%s", name),
							},
						},
					},
				},
			},
		},
	}
}

type Lifecycle interface {
	Start(ctx context.Context)
	Stop()
}

type pollingRunner struct {
	producer Lifecycle
	consumer Lifecycle

	workQueue chan int
}

func newPollingRunner(pollingFunc func(), pollingInterval time.Duration, concurrencyFactor int) *pollingRunner {
	workQueue := make(chan int, 50)

	return &pollingRunner{
		producer:  newIntervalProducer(workQueue, pollingInterval),
		consumer:  newPooledConsumers(workQueue, pollingFunc, concurrencyFactor),
		workQueue: workQueue,
	}
}

func (r *pollingRunner) StartPolling(ctx context.Context) {
	r.producer.Start(ctx)
	r.consumer.Start(ctx)
}

func (r *pollingRunner) StopPolling() {
	// Close the channel used to pass work between producers and consumer
	defer close(r.workQueue)

	r.producer.Stop()
	r.consumer.Stop()
}

type intervalProducer struct {
	workQueue chan int
	interval  time.Duration
	stop      chan bool
}

func newIntervalProducer(workQueue chan int, interval time.Duration) *intervalProducer {
	return &intervalProducer{
		workQueue: workQueue,
		interval:  interval,
		stop:      make(chan bool),
	}
}

func (p *intervalProducer) Start(ctx context.Context) {
	// At the polling interval, drop a new piece of work on the workQueue
	workId := 0
	go func(testContext context.Context) {
		for {
			select {
			case <-time.After(p.interval):
				workId++
				p.workQueue <- workId
			case <-p.stop:
				return
			case <-testContext.Done():
				return
			}
		}
	}(ctx)
}

func (p *intervalProducer) Stop() {
	defer close(p.stop)
	p.stop <- true
}

type pooledConsumers struct {
	workQueue     chan int
	run           func()
	poolSize      int
	stop          chan bool
	activeWorkers sync.WaitGroup
}

func newPooledConsumers(workQueue chan int, run func(), poolSize int) *pooledConsumers {
	return &pooledConsumers{
		workQueue:     workQueue,
		run:           run,
		poolSize:      poolSize,
		stop:          make(chan bool),
		activeWorkers: sync.WaitGroup{},
	}
}

func (c *pooledConsumers) Start(ctx context.Context) {
	for worker := 0; worker < c.poolSize; worker++ {
		c.startWorker(ctx, worker)
	}
}

func (c *pooledConsumers) startWorker(ctx context.Context, workerId int) {
	c.activeWorkers.Add(1)
	go func(testContext context.Context) {
		defer c.activeWorkers.Done()
		for {
			select {
			case <-c.workQueue:
				c.run() // blocking
			case <-c.stop:
				return
			case <-testContext.Done():
				return
			}
		}
	}(ctx)
}

func (c *pooledConsumers) Stop() {
	defer close(c.stop)
	for worker := 0; worker < c.poolSize; worker++ {
		c.stop <- true
	}

	c.activeWorkers.Wait()
}