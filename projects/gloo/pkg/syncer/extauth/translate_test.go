package extauth_test

import (
	"context"
	"reflect"
	"time"

	extauthsyncer "github.com/solo-io/solo-projects/projects/gloo/pkg/syncer/extauth"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/golang/protobuf/ptypes/wrappers"

	"github.com/golang/protobuf/ptypes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/core/matchers"

	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	extauth "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1"
	v1snap "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/gloosnapshot"
	static_plugin_gloo "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/static"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

var _ = Describe("Translate", func() {

	var (
		params            plugins.Params
		virtualHost       *v1.VirtualHost
		upstream          *v1.Upstream
		secret            *v1.Secret
		cipherSecret      *v1.Secret
		route             *v1.Route
		authConfig        *extauth.AuthConfig
		authConfigRef     *core.ResourceRef
		extAuthExtension  *extauth.ExtAuthExtension
		clientSecret      *extauth.OauthSecret
		apiKey            *extauth.ApiKey
		credentialsSecret *v1.AccountCredentialsSecret
		ldapSecret        *v1.Secret
		encryptionKey     string
	)

	BeforeEach(func() {
		encryptionKey = "an example encryption key1234567"
		upstream = &v1.Upstream{
			Metadata: &core.Metadata{
				Name:      "extauth",
				Namespace: "default",
			},
			UpstreamType: &v1.Upstream_Static{
				Static: &static_plugin_gloo.UpstreamSpec{
					Hosts: []*static_plugin_gloo.Host{{
						Addr: "test",
						Port: 1234,
					}},
				},
			},
		}
		route = &v1.Route{
			Matchers: []*matchers.Matcher{{
				PathSpecifier: &matchers.Matcher_Prefix{
					Prefix: "/",
				},
			}},
			Action: &v1.Route_RouteAction{
				RouteAction: &v1.RouteAction{
					Destination: &v1.RouteAction_Single{
						Single: &v1.Destination{
							DestinationType: &v1.Destination_Upstream{
								Upstream: upstream.Metadata.Ref(),
							},
						},
					},
				},
			},
		}

		apiKey = &extauth.ApiKey{
			ApiKey: "apiKey1",
		}

		credentialsSecret = &v1.AccountCredentialsSecret{
			Username: "user",
			Password: "pass",
		}
		ldapSecret = &v1.Secret{
			Metadata: &core.Metadata{
				Name:      "ldapSecret",
				Namespace: "default",
			},
			Kind: &v1.Secret_Credentials{
				Credentials: credentialsSecret,
			},
		}
		clientSecret = &extauth.OauthSecret{
			ClientSecret: "1234",
		}

		secret = &v1.Secret{
			Metadata: &core.Metadata{
				Name:      "secret",
				Namespace: "default",
			},
			Kind: &v1.Secret_Oauth{
				Oauth: clientSecret,
			},
		}
		secretRef := secret.Metadata.Ref()
		authConfig = &extauth.AuthConfig{
			Metadata: &core.Metadata{
				Name:      "oauth",
				Namespace: "gloo-system",
			},
			Configs: []*extauth.AuthConfig_Config{{
				AuthConfig: &extauth.AuthConfig_Config_Oauth2{
					Oauth2: &extauth.OAuth2{
						OauthType: &extauth.OAuth2_OidcAuthorizationCode{

							OidcAuthorizationCode: &extauth.OidcAuthorizationCode{
								ClientSecretRef:          secretRef,
								ClientId:                 "ClientId",
								IssuerUrl:                "IssuerUrl",
								AuthEndpointQueryParams:  map[string]string{"test": "additional_auth_query_params"},
								TokenEndpointQueryParams: map[string]string{"test": "additional_token_query_params"},
								AppUrl:                   "AppUrl",
								CallbackPath:             "CallbackPath",
								AutoMapFromMetadata: &extauth.AutoMapFromMetadata{
									Namespace: "test_namespace",
								},
								Session: &extauth.UserSession{
									FailOnFetchFailure: true,
									CookieOptions: &extauth.UserSession_CookieOptions{
										MaxAge: &wrapperspb.UInt32Value{Value: 20},
									},
									Session: &extauth.UserSession_Cookie{
										Cookie: &extauth.UserSession_InternalSession{
											AllowRefreshing: &wrapperspb.BoolValue{Value: true},
											KeyPrefix:       "prefix",
										},
									},
									CipherConfig: &extauth.UserSession_CipherConfig{
										Key: &extauth.UserSession_CipherConfig_KeyRef{
											KeyRef: &core.ResourceRef{
												Name:      "cipher-key-name",
												Namespace: "cipher-key-namespace",
											},
										},
									},
								},
							},
						},
					},
				},
			}},
		}

		authConfigRef = authConfig.Metadata.Ref()
		extAuthExtension = &extauth.ExtAuthExtension{
			Spec: &extauth.ExtAuthExtension_ConfigRef{
				ConfigRef: authConfigRef,
			},
		}

		params.Snapshot = &v1snap.ApiSnapshot{
			Upstreams:   v1.UpstreamList{upstream},
			AuthConfigs: extauth.AuthConfigList{authConfig},
			Secrets:     v1.SecretList{cipherSecret},
		}
	})

	JustBeforeEach(func() {
		virtualHost = &v1.VirtualHost{
			Name:    "virt1",
			Domains: []string{"*"},
			Options: &v1.VirtualHostOptions{
				Extauth: extAuthExtension,
			},
			Routes: []*v1.Route{route},
		}

		proxy := &v1.Proxy{
			Metadata: &core.Metadata{
				Name:      "secret",
				Namespace: "default",
			},
			Listeners: []*v1.Listener{{
				Name: "default",
				ListenerType: &v1.Listener_HttpListener{
					HttpListener: &v1.HttpListener{
						VirtualHosts: []*v1.VirtualHost{virtualHost},
					},
				},
			}},
		}

		cipherSecret = &gloov1.Secret{
			Metadata: &core.Metadata{
				Name:      "cipher-key-name",
				Namespace: "cipher-key-namespace",
			},
			Kind: &gloov1.Secret_Encryption{
				Encryption: &gloov1.EncryptionKeySecret{
					Key: encryptionKey,
				},
			},
		}

		params.Snapshot.Proxies = v1.ProxyList{proxy}
		params.Snapshot.Secrets = v1.SecretList{secret, ldapSecret, cipherSecret}
	})

	It("should translate oauth config for extauth server", func() {
		translated, err := extauthsyncer.TranslateExtAuthConfig(context.TODO(), params.Snapshot, authConfigRef)
		Expect(err).NotTo(HaveOccurred())
		Expect(translated.AuthConfigRefName).To(Equal(authConfigRef.Key()))
		Expect(translated.Configs).To(HaveLen(1))
		actual := translated.Configs[0].GetOauth2()
		expected := authConfig.Configs[0].GetOauth2()
		actualOidc := actual.GetOidcAuthorizationCode()
		expectedOidc := expected.GetOidcAuthorizationCode()
		Expect(actualOidc.IssuerUrl).To(Equal(expectedOidc.IssuerUrl))
		Expect(actualOidc.AuthEndpointQueryParams).To(Equal(expectedOidc.AuthEndpointQueryParams))
		Expect(actualOidc.TokenEndpointQueryParams).To(Equal(expectedOidc.TokenEndpointQueryParams))
		Expect(actualOidc.ClientId).To(Equal(expectedOidc.ClientId))
		Expect(actualOidc.ClientSecret).To(Equal(clientSecret.ClientSecret))
		Expect(actualOidc.AppUrl).To(Equal(expectedOidc.AppUrl))
		Expect(actualOidc.CallbackPath).To(Equal(expectedOidc.CallbackPath))
		Expect(actualOidc.AutoMapFromMetadata.Namespace).To(Equal("test_namespace"))
		// verify translation of the User Session
		//lint:ignore SA1019 testing for upgrades
		Expect(actualOidc.Session).To(BeNil())
		Expect(actualOidc.UserSession.FailOnFetchFailure).To(Equal(expectedOidc.Session.FailOnFetchFailure))
		Expect(actualOidc.UserSession.CookieOptions).To(Equal(expectedOidc.Session.CookieOptions))
		Expect(actualOidc.UserSession.CipherConfig.Key).To(Equal(cipherSecret.GetEncryption().GetKey()))
		Expect(actualOidc.UserSession.GetCookie()).To(Equal(expectedOidc.Session.GetCookie()))
	})
	It("should translate session when cipher is not included", func() {
		// set the cipher to nil
		params.Snapshot.AuthConfigs[0].GetConfigs()[0].GetOauth2().GetOidcAuthorizationCode().GetSession().CipherConfig = nil
		translated, err := extauthsyncer.TranslateExtAuthConfig(context.TODO(), params.Snapshot, authConfigRef)
		Expect(err).NotTo(HaveOccurred())
		Expect(translated.AuthConfigRefName).To(Equal(authConfigRef.Key()))
		Expect(translated.Configs).To(HaveLen(1))
		actual := translated.Configs[0].GetOauth2()
		expected := authConfig.Configs[0].GetOauth2()
		actualOidc := actual.GetOidcAuthorizationCode()
		expectedOidc := expected.GetOidcAuthorizationCode()
		Expect(actualOidc.IssuerUrl).To(Equal(expectedOidc.IssuerUrl))
		Expect(actualOidc.AuthEndpointQueryParams).To(Equal(expectedOidc.AuthEndpointQueryParams))
		Expect(actualOidc.TokenEndpointQueryParams).To(Equal(expectedOidc.TokenEndpointQueryParams))
		Expect(actualOidc.ClientId).To(Equal(expectedOidc.ClientId))
		Expect(actualOidc.ClientSecret).To(Equal(clientSecret.ClientSecret))
		Expect(actualOidc.AppUrl).To(Equal(expectedOidc.AppUrl))
		Expect(actualOidc.CallbackPath).To(Equal(expectedOidc.CallbackPath))
		Expect(actualOidc.AutoMapFromMetadata.Namespace).To(Equal("test_namespace"))
		// verify translation of the User Session
		//lint:ignore SA1019 testing for upgrades
		Expect(actualOidc.Session.GetCookie()).To(Equal(expectedOidc.Session.GetCookie()))
		//lint:ignore SA1019 testing for upgrades
		Expect(actualOidc.Session.FailOnFetchFailure).To(Equal(expectedOidc.Session.GetFailOnFetchFailure()))
		//lint:ignore SA1019 testing for upgrades
		Expect(actualOidc.Session.CookieOptions).To(Equal(expectedOidc.Session.CookieOptions))
		//lint:ignore SA1019 testing to ensure that the Session is Nil
		Expect(actualOidc.Session.GetCipherConfig()).To(BeNil())
		Expect(actualOidc.UserSession).To(BeNil())
	})
	It("should translate oidc config without a clientSecret", func() {
		oidcConfig := authConfig.GetConfigs()[0].GetOauth2().GetOauthType().(*extauth.OAuth2_OidcAuthorizationCode).OidcAuthorizationCode
		oidcConfig.DisableClientSecret = &wrappers.BoolValue{Value: true}
		oidcConfig.ClientSecretRef = nil

		translated, err := extauthsyncer.TranslateExtAuthConfig(context.TODO(), params.Snapshot, authConfigRef)
		Expect(err).NotTo(HaveOccurred())
		Expect(translated.AuthConfigRefName).To(Equal(authConfigRef.Key()))
		Expect(translated.Configs).To(HaveLen(1))
		actual := translated.Configs[0].GetOauth2()
		expected := authConfig.Configs[0].GetOauth2()
		actualOidc := actual.GetOidcAuthorizationCode()
		expectedOidc := expected.GetOidcAuthorizationCode()
		// Verify that the client secret is empty
		Expect(actualOidc.ClientSecret).To(BeEmpty())
		// Verify the rest of the translation
		Expect(actualOidc.IssuerUrl).To(Equal(expectedOidc.IssuerUrl))
		Expect(actualOidc.AuthEndpointQueryParams).To(Equal(expectedOidc.AuthEndpointQueryParams))
		Expect(actualOidc.TokenEndpointQueryParams).To(Equal(expectedOidc.TokenEndpointQueryParams))
		Expect(actualOidc.ClientId).To(Equal(expectedOidc.ClientId))
		Expect(actualOidc.AppUrl).To(Equal(expectedOidc.AppUrl))
		Expect(actualOidc.CallbackPath).To(Equal(expectedOidc.CallbackPath))
		Expect(actualOidc.AutoMapFromMetadata.Namespace).To(Equal("test_namespace"))
	})
	Context("Encryption Key error", func() {
		BeforeEach(func() {
			encryptionKey = "an example encryption key"
		})
		It("should error because it does not meet the key length", func() {
			_, err := extauthsyncer.TranslateExtAuthConfig(context.TODO(), params.Snapshot, authConfigRef)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("32 characters in length"))
		})
	})

	It("will fail if the oidc auth proto has a new top level field", func() {
		// This test is important as it checks whether the oidc auth code proto have a new top level field.
		// This should happen very rarely, and should be used as an indication that the `translateOidcAuthorizationCode` function
		// most likely needs to change.

		Expect(reflect.TypeOf(extauth.ExtAuthConfig_OidcAuthorizationCodeConfig{}).NumField()).To(
			Equal(23),
			"wrong number of fields found",
		)
	})

	Context("with plain OAuth2 extauth", func() {
		BeforeEach(func() {
			clientSecret = &extauth.OauthSecret{
				ClientSecret: "1234",
			}

			secret = &v1.Secret{
				Metadata: &core.Metadata{
					Name:      "secret",
					Namespace: "default",
				},
				Kind: &v1.Secret_Oauth{
					Oauth: clientSecret,
				},
			}
			secretRef := secret.Metadata.Ref()

			authConfig = &extauth.AuthConfig{
				Metadata: &core.Metadata{
					Name:      "oauth",
					Namespace: "gloo-system",
				},
				Configs: []*extauth.AuthConfig_Config{{
					AuthConfig: &extauth.AuthConfig_Config_Oauth2{
						Oauth2: &extauth.OAuth2{
							OauthType: &extauth.OAuth2_Oauth2{
								Oauth2: &extauth.PlainOAuth2{
									AppUrl:             "app.url",
									CallbackPath:       "/callback",
									ClientId:           "cid",
									ClientSecretRef:    secretRef,
									Scopes:             []string{"trust"},
									AuthEndpoint:       "login.url/auth",
									TokenEndpoint:      "login.url/token",
									RevocationEndpoint: "login.url/revoke",
									Session: &extauth.UserSession{
										FailOnFetchFailure: true,
										CookieOptions: &extauth.UserSession_CookieOptions{
											MaxAge: &wrapperspb.UInt32Value{Value: 20},
										},
										Session: &extauth.UserSession_Cookie{
											Cookie: &extauth.UserSession_InternalSession{
												AllowRefreshing: &wrapperspb.BoolValue{Value: true},
												KeyPrefix:       "prefix",
											},
										},
										CipherConfig: &extauth.UserSession_CipherConfig{
											Key: &extauth.UserSession_CipherConfig_KeyRef{
												KeyRef: &core.ResourceRef{
													Name:      "cipher-key-name",
													Namespace: "cipher-key-namespace",
												},
											},
										},
									},
								},
							},
						},
					},
				}},
			}
			authConfigRef = authConfig.Metadata.Ref()
			extAuthExtension = &extauth.ExtAuthExtension{
				Spec: &extauth.ExtAuthExtension_ConfigRef{
					ConfigRef: authConfigRef,
				},
			}

			params.Snapshot = &v1snap.ApiSnapshot{
				Upstreams:   v1.UpstreamList{upstream},
				AuthConfigs: extauth.AuthConfigList{authConfig},
			}
		})

		// This test checks whether the plain oauth2 proto has a new field.
		// If so, the `translatePlainOAuth2` function most likely needs to be updated
		It("will fail if the plain oauth2 proto has a new top level field", func() {
			Expect(reflect.TypeOf(extauth.ExtAuthConfig_PlainOAuth2Config{}).NumField()).To(
				Equal(17),
				"wrong number of fields found")
		})

		It("should translate plain oauth2 config", func() {
			translated, err := extauthsyncer.TranslateExtAuthConfig(context.TODO(), params.Snapshot, authConfigRef)
			Expect(err).NotTo(HaveOccurred())
			Expect(translated.AuthConfigRefName).To(Equal(authConfigRef.Key()))
			Expect(translated.Configs).To(HaveLen(1))
			actual := translated.Configs[0].GetOauth2()
			actualPlainOAuth2 := actual.GetOauth2Config()
			expected := authConfig.Configs[0].GetOauth2()
			expectedPlainOAuth2 := expected.GetOauth2()

			Expect(actualPlainOAuth2.AppUrl).To(Equal(expectedPlainOAuth2.AppUrl))
			Expect(actualPlainOAuth2.CallbackPath).To(Equal(expectedPlainOAuth2.CallbackPath))
			Expect(actualPlainOAuth2.ClientId).To(Equal(expectedPlainOAuth2.ClientId))
			Expect(actualPlainOAuth2.ClientSecret).To(Equal(clientSecret.ClientSecret))
			Expect(actualPlainOAuth2.AuthEndpointQueryParams).To(Equal(expectedPlainOAuth2.AuthEndpointQueryParams))
			Expect(actualPlainOAuth2.TokenEndpointQueryParams).To(Equal(expectedPlainOAuth2.TokenEndpointQueryParams))
			Expect(actualPlainOAuth2.Scopes).To(Equal(expectedPlainOAuth2.Scopes))
			Expect(actualPlainOAuth2.AuthEndpoint).To(Equal(expectedPlainOAuth2.AuthEndpoint))
			Expect(actualPlainOAuth2.TokenEndpoint).To(Equal(expectedPlainOAuth2.TokenEndpoint))
			Expect(actualPlainOAuth2.RevocationEndpoint).To(Equal(expectedPlainOAuth2.RevocationEndpoint))
			// verify translation of the Session is nil, when the cipher config is set
			//lint:ignore SA1019 testing for upgrades
			Expect(actualPlainOAuth2.Session).To(BeNil())
			//lint:ignore SA1019 testing for upgrades
			Expect(actualPlainOAuth2.Session.GetCipherConfig()).To(BeNil())
			Expect(actualPlainOAuth2.UserSession.FailOnFetchFailure).To(Equal(expectedPlainOAuth2.Session.FailOnFetchFailure))
			Expect(actualPlainOAuth2.UserSession.CookieOptions).To(Equal(expectedPlainOAuth2.Session.CookieOptions))
			Expect(actualPlainOAuth2.UserSession.CipherConfig.Key).To(Equal(cipherSecret.GetEncryption().GetKey()))
			Expect(actualPlainOAuth2.UserSession.GetCookie()).To(Equal(expectedPlainOAuth2.Session.GetCookie()))
		})
		It("should translate plain oauth2 config without a secretRef", func() {
			plainOauthConfig := authConfig.Configs[0].GetOauth2().GetOauthType().(*extauth.OAuth2_Oauth2).Oauth2
			plainOauthConfig.DisableClientSecret = &wrappers.BoolValue{Value: true}
			plainOauthConfig.ClientSecretRef = nil
			translated, err := extauthsyncer.TranslateExtAuthConfig(context.TODO(), params.Snapshot, authConfigRef)
			Expect(err).NotTo(HaveOccurred())
			Expect(translated.AuthConfigRefName).To(Equal(authConfigRef.Key()))
			Expect(translated.Configs).To(HaveLen(1))
			actual := translated.Configs[0].GetOauth2()
			actualPlainOAuth2 := actual.GetOauth2Config()
			expected := authConfig.Configs[0].GetOauth2()
			expectedPlainOAuth2 := expected.GetOauth2()

			// Expect an empty client secret
			Expect(actualPlainOAuth2.ClientSecret).To(BeEmpty())
			//validate the rest of the translation for good measure
			Expect(actualPlainOAuth2.AppUrl).To(Equal(expectedPlainOAuth2.AppUrl))
			Expect(actualPlainOAuth2.CallbackPath).To(Equal(expectedPlainOAuth2.CallbackPath))
			Expect(actualPlainOAuth2.ClientId).To(Equal(expectedPlainOAuth2.ClientId))
			Expect(actualPlainOAuth2.AuthEndpointQueryParams).To(Equal(expectedPlainOAuth2.AuthEndpointQueryParams))
			Expect(actualPlainOAuth2.TokenEndpointQueryParams).To(Equal(expectedPlainOAuth2.TokenEndpointQueryParams))
			Expect(actualPlainOAuth2.Scopes).To(Equal(expectedPlainOAuth2.Scopes))
			Expect(actualPlainOAuth2.AuthEndpoint).To(Equal(expectedPlainOAuth2.AuthEndpoint))
			Expect(actualPlainOAuth2.TokenEndpoint).To(Equal(expectedPlainOAuth2.TokenEndpoint))
			Expect(actualPlainOAuth2.RevocationEndpoint).To(Equal(expectedPlainOAuth2.RevocationEndpoint))
			// verify translation of the Session is nil, when the cipher config is set
			//lint:ignore SA1019 testing for upgrades
			Expect(actualPlainOAuth2.Session).To(BeNil())
			//lint:ignore SA1019 testing for upgrades
			Expect(actualPlainOAuth2.Session.GetCipherConfig()).To(BeNil())
			Expect(actualPlainOAuth2.UserSession.FailOnFetchFailure).To(Equal(expectedPlainOAuth2.Session.FailOnFetchFailure))
			Expect(actualPlainOAuth2.UserSession.CookieOptions).To(Equal(expectedPlainOAuth2.Session.CookieOptions))
			Expect(actualPlainOAuth2.UserSession.CipherConfig.Key).To(Equal(cipherSecret.GetEncryption().GetKey()))
			Expect(actualPlainOAuth2.UserSession.GetCookie()).To(Equal(expectedPlainOAuth2.Session.GetCookie()))
		})
		It("should translate session if not using cipherConfig", func() {
			authConfigs := params.Snapshot.AuthConfigs
			// set the cipherConfig to nil for translation of the session
			authConfigs[0].Configs[0].GetOauth2().GetOauth2().Session.CipherConfig = nil
			translated, err := extauthsyncer.TranslateExtAuthConfig(context.TODO(), params.Snapshot, authConfigRef)
			Expect(err).NotTo(HaveOccurred())
			Expect(translated.AuthConfigRefName).To(Equal(authConfigRef.Key()))
			Expect(translated.Configs).To(HaveLen(1))
			actual := translated.Configs[0].GetOauth2()
			actualPlainOAuth2 := actual.GetOauth2Config()
			expected := authConfig.Configs[0].GetOauth2()
			expectedPlainOAuth2 := expected.GetOauth2()

			Expect(actualPlainOAuth2.AppUrl).To(Equal(expectedPlainOAuth2.AppUrl))
			Expect(actualPlainOAuth2.CallbackPath).To(Equal(expectedPlainOAuth2.CallbackPath))
			Expect(actualPlainOAuth2.ClientId).To(Equal(expectedPlainOAuth2.ClientId))
			Expect(actualPlainOAuth2.ClientSecret).To(Equal(clientSecret.ClientSecret))
			Expect(actualPlainOAuth2.AuthEndpointQueryParams).To(Equal(expectedPlainOAuth2.AuthEndpointQueryParams))
			Expect(actualPlainOAuth2.TokenEndpointQueryParams).To(Equal(expectedPlainOAuth2.TokenEndpointQueryParams))
			Expect(actualPlainOAuth2.Scopes).To(Equal(expectedPlainOAuth2.Scopes))
			Expect(actualPlainOAuth2.AuthEndpoint).To(Equal(expectedPlainOAuth2.AuthEndpoint))
			Expect(actualPlainOAuth2.TokenEndpoint).To(Equal(expectedPlainOAuth2.TokenEndpoint))
			Expect(actualPlainOAuth2.RevocationEndpoint).To(Equal(expectedPlainOAuth2.RevocationEndpoint))
			// verify translation of the Session, because the cipher config is nil
			//lint:ignore SA1019 testing for upgrades
			Expect(actualPlainOAuth2.Session.GetCookie()).To(Equal(expectedPlainOAuth2.Session.GetCookie()))
			//lint:ignore SA1019 testing for upgrades
			Expect(actualPlainOAuth2.Session.FailOnFetchFailure).To(Equal(expectedPlainOAuth2.Session.GetFailOnFetchFailure()))
			//lint:ignore SA1019 testing for upgrades
			Expect(actualPlainOAuth2.Session.CookieOptions).To(Equal(expectedPlainOAuth2.Session.CookieOptions))
			//lint:ignore SA1019 testing to ensure that the Session Cipher Config is Nil
			Expect(actualPlainOAuth2.Session.GetCipherConfig()).To(BeNil())
			// verify translation of the User Session, which is nil when the cipher config is set
			Expect(actualPlainOAuth2.UserSession).To(BeNil())
		})
		It("should translate session if not using cipherConfig, and only setting Cookie Options", func() {
			authConfigs := params.Snapshot.AuthConfigs
			// set the cipherConfig to nil for translation of the session
			session := authConfigs[0].Configs[0].GetOauth2().GetOauth2().Session
			session.CipherConfig = nil
			session.Session = nil
			Expect(session.CookieOptions).NotTo(BeNil())
			translated, err := extauthsyncer.TranslateExtAuthConfig(context.TODO(), params.Snapshot, authConfigRef)
			Expect(err).NotTo(HaveOccurred())
			Expect(translated.AuthConfigRefName).To(Equal(authConfigRef.Key()))
			Expect(translated.Configs).To(HaveLen(1))
			actual := translated.Configs[0].GetOauth2()
			actualPlainOAuth2 := actual.GetOauth2Config()
			expected := authConfig.Configs[0].GetOauth2()
			expectedPlainOAuth2 := expected.GetOauth2()

			Expect(actualPlainOAuth2.AppUrl).To(Equal(expectedPlainOAuth2.AppUrl))
			Expect(actualPlainOAuth2.CallbackPath).To(Equal(expectedPlainOAuth2.CallbackPath))
			Expect(actualPlainOAuth2.ClientId).To(Equal(expectedPlainOAuth2.ClientId))
			Expect(actualPlainOAuth2.ClientSecret).To(Equal(clientSecret.ClientSecret))
			Expect(actualPlainOAuth2.AuthEndpointQueryParams).To(Equal(expectedPlainOAuth2.AuthEndpointQueryParams))
			Expect(actualPlainOAuth2.TokenEndpointQueryParams).To(Equal(expectedPlainOAuth2.TokenEndpointQueryParams))
			Expect(actualPlainOAuth2.Scopes).To(Equal(expectedPlainOAuth2.Scopes))
			Expect(actualPlainOAuth2.AuthEndpoint).To(Equal(expectedPlainOAuth2.AuthEndpoint))
			Expect(actualPlainOAuth2.TokenEndpoint).To(Equal(expectedPlainOAuth2.TokenEndpoint))
			Expect(actualPlainOAuth2.RevocationEndpoint).To(Equal(expectedPlainOAuth2.RevocationEndpoint))
			// verify translation of the Session, because the cipher config is nil
			//lint:ignore SA1019 testing for upgrades
			Expect(actualPlainOAuth2.Session.GetCookie()).To(Equal(expectedPlainOAuth2.Session.GetCookie()))
			//lint:ignore SA1019 testing for upgrades
			Expect(actualPlainOAuth2.Session.FailOnFetchFailure).To(Equal(expectedPlainOAuth2.Session.GetFailOnFetchFailure()))
			//lint:ignore SA1019 testing for upgrades
			Expect(actualPlainOAuth2.Session.CookieOptions).To(Equal(expectedPlainOAuth2.Session.CookieOptions))
			//lint:ignore SA1019 testing to ensure that the Session Cipher Config is Nil
			Expect(actualPlainOAuth2.Session.GetCipherConfig()).To(BeNil())
			// verify translation of the User Session, which is nil when the cipher config is set
			Expect(actualPlainOAuth2.UserSession).To(BeNil())
		})
	})

	Context("with api key extauth", func() {
		BeforeEach(func() {

			secret = &v1.Secret{
				Metadata: &core.Metadata{
					Name:      "secretName",
					Namespace: "default",
					Labels:    map[string]string{"team": "infrastructure"},
				},
				Kind: &v1.Secret_ApiKey{
					ApiKey: apiKey,
				},
			}
			secretRef := secret.Metadata.Ref()

			authConfig = &extauth.AuthConfig{
				Metadata: &core.Metadata{
					Name:      "apikey",
					Namespace: "gloo-system",
				},
				Configs: []*extauth.AuthConfig_Config{{
					AuthConfig: &extauth.AuthConfig_Config_ApiKeyAuth{
						ApiKeyAuth: &extauth.ApiKeyAuth{
							HeaderName:       "x-api-key",
							ApiKeySecretRefs: []*core.ResourceRef{secretRef},
						},
					},
				}},
			}
			authConfigRef = authConfig.Metadata.Ref()
			extAuthExtension = &extauth.ExtAuthExtension{
				Spec: &extauth.ExtAuthExtension_ConfigRef{
					ConfigRef: authConfigRef,
				},
			}

			params.Snapshot = &v1snap.ApiSnapshot{
				AuthConfigs: extauth.AuthConfigList{authConfig},
			}
		})

		Context("secret is malformed", func() {
			It("returns expected error when secret is not of API key type", func() {
				secret.Kind = &v1.Secret_Aws{}
				_, err := extauthsyncer.TranslateExtAuthConfig(context.TODO(), params.Snapshot, authConfigRef)
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(ContainSubstring(extauthsyncer.NonApiKeySecretError(secret).Error())))
			})

			It("returns expected error when the secret does not contain an API key", func() {
				secret.GetApiKey().ApiKey = ""
				_, err := extauthsyncer.TranslateExtAuthConfig(context.TODO(), params.Snapshot, authConfigRef)
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(ContainSubstring(extauthsyncer.EmptyApiKeyError(secret).Error())))
			})
		})

		Context("with api key extauth, secret ref matching", func() {
			It("should translate api keys config for extauth server - matching secret ref", func() {
				translated, err := extauthsyncer.TranslateExtAuthConfig(context.TODO(), params.Snapshot, authConfigRef)
				Expect(err).NotTo(HaveOccurred())
				Expect(translated.AuthConfigRefName).To(Equal(authConfigRef.Key()))
				Expect(translated.Configs).To(HaveLen(1))
				actual := translated.Configs[0].GetApiKeyAuth()
				Expect(actual).To(Equal(&extauth.ExtAuthConfig_ApiKeyAuthConfig{
					HeaderName: "x-api-key",
					ValidApiKeys: map[string]*extauth.ExtAuthConfig_ApiKeyAuthConfig_KeyMetadata{
						"apiKey1": {
							Username: "secretName",
						},
					},
				}))
			})

			It("should translate api keys config for extauth server - mismatching secret ref", func() {
				secret.Metadata.Name = "mismatchName"
				_, err := extauthsyncer.TranslateExtAuthConfig(context.TODO(), params.Snapshot, authConfigRef)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("list did not find secret"))
			})
		})

		Describe("API keys with metadata", func() {

			BeforeEach(func() {
				secret = &v1.Secret{
					Metadata: &core.Metadata{
						Name:      "secretName",
						Namespace: "default",
						Labels:    map[string]string{"team": "infrastructure"},
					},
					Kind: &v1.Secret_ApiKey{
						ApiKey: &extauth.ApiKey{
							ApiKey: "apiKey1",
							Metadata: map[string]string{
								"user-id": "123",
							},
						},
					},
				}
				secretRef := secret.Metadata.Ref()
				authConfig = &extauth.AuthConfig{
					Metadata: &core.Metadata{
						Name:      "apikey",
						Namespace: "gloo-system",
					},
					Configs: []*extauth.AuthConfig_Config{{
						AuthConfig: &extauth.AuthConfig_Config_ApiKeyAuth{
							ApiKeyAuth: &extauth.ApiKeyAuth{
								HeaderName:       "x-api-key",
								ApiKeySecretRefs: []*core.ResourceRef{secretRef},
								HeadersFromMetadataEntry: map[string]*extauth.ApiKeyAuth_MetadataEntry{
									"x-user-id": {
										Name:     "user-id",
										Required: true,
									},
								},
							},
						},
					}},
				}
				authConfigRef = authConfig.Metadata.Ref()
				extAuthExtension = &extauth.ExtAuthExtension{
					Spec: &extauth.ExtAuthExtension_ConfigRef{
						ConfigRef: authConfigRef,
					},
				}

				params.Snapshot = &v1snap.ApiSnapshot{
					AuthConfigs: extauth.AuthConfigList{authConfig},
				}
			})

			It("should fail if required metadata is missing on the secret", func() {
				secret.GetApiKey().Metadata = nil

				_, err := extauthsyncer.TranslateExtAuthConfig(context.TODO(), params.Snapshot, authConfigRef)
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(ContainSubstring(extauthsyncer.MissingRequiredMetadataError("user-id", secret).Error())))
			})

			It("should include secret metadata in the API key metadata", func() {
				translated, err := extauthsyncer.TranslateExtAuthConfig(context.TODO(), params.Snapshot, authConfigRef)
				Expect(err).NotTo(HaveOccurred())
				Expect(translated.AuthConfigRefName).To(Equal(authConfigRef.Key()))
				Expect(translated.Configs).To(HaveLen(1))
				actual := translated.Configs[0].GetApiKeyAuth()
				Expect(actual).To(Equal(&extauth.ExtAuthConfig_ApiKeyAuthConfig{
					HeaderName: "x-api-key",
					ValidApiKeys: map[string]*extauth.ExtAuthConfig_ApiKeyAuthConfig_KeyMetadata{
						"apiKey1": {
							Username: "secretName",
							Metadata: map[string]string{
								"user-id": "123",
							},
						},
					},
					HeadersFromKeyMetadata: map[string]string{
						"x-user-id": "user-id",
					},
				}))
			})

			When("metadata is not required", func() {
				It("does not fail if the secret does not contain the metadata", func() {
					secret.GetApiKey().Metadata = nil
					authConfig.GetConfigs()[0].GetApiKeyAuth().GetHeadersFromMetadataEntry()["x-user-id"].Required = false

					translated, err := extauthsyncer.TranslateExtAuthConfig(context.TODO(), params.Snapshot, authConfigRef)
					Expect(err).NotTo(HaveOccurred())
					Expect(translated.AuthConfigRefName).To(Equal(authConfigRef.Key()))
					Expect(translated.Configs).To(HaveLen(1))
					actual := translated.Configs[0].GetApiKeyAuth()
					Expect(actual).To(Equal(&extauth.ExtAuthConfig_ApiKeyAuthConfig{
						HeaderName: "x-api-key",
						ValidApiKeys: map[string]*extauth.ExtAuthConfig_ApiKeyAuthConfig_KeyMetadata{
							"apiKey1": {
								Username: "secretName",
							},
						},
						HeadersFromKeyMetadata: map[string]string{
							"x-user-id": "user-id",
						},
					}))
				})
			})
		})

		Context("with api key ext auth, label matching", func() {
			BeforeEach(func() {
				authConfig = &extauth.AuthConfig{
					Metadata: &core.Metadata{
						Name:      "apikey",
						Namespace: "gloo-system",
					},
					Configs: []*extauth.AuthConfig_Config{{
						AuthConfig: &extauth.AuthConfig_Config_ApiKeyAuth{
							ApiKeyAuth: &extauth.ApiKeyAuth{
								LabelSelector: map[string]string{"team": "infrastructure"},
							},
						},
					}},
				}
				authConfigRef = authConfig.Metadata.Ref()
				extAuthExtension = &extauth.ExtAuthExtension{
					Spec: &extauth.ExtAuthExtension_ConfigRef{
						ConfigRef: authConfigRef,
					},
				}

				params.Snapshot = &v1snap.ApiSnapshot{
					AuthConfigs: extauth.AuthConfigList{authConfig},
				}
			})

			It("should translate api keys config for extauth server - matching label", func() {
				translated, err := extauthsyncer.TranslateExtAuthConfig(context.TODO(), params.Snapshot, authConfigRef)
				Expect(err).NotTo(HaveOccurred())
				Expect(translated.AuthConfigRefName).To(Equal(authConfigRef.Key()))
				Expect(translated.Configs).To(HaveLen(1))
				actual := translated.Configs[0].GetApiKeyAuth()
				Expect(actual.ValidApiKeys).To(Equal(map[string]*extauth.ExtAuthConfig_ApiKeyAuthConfig_KeyMetadata{
					"apiKey1": {
						Username: "secretName",
					},
				}))
			})

			Context("should translate apikeys config for extauth server", func() {

				It("should not error - mismatched labels", func() {
					secret.Metadata.Labels = map[string]string{"missingLabel": "missingValue"}
					_, err := extauthsyncer.TranslateExtAuthConfig(context.TODO(), params.Snapshot, authConfigRef)
					Expect(err).NotTo(HaveOccurred())
				})

				It("should not error - empty labels", func() {
					secret.Metadata.Labels = map[string]string{}
					_, err := extauthsyncer.TranslateExtAuthConfig(context.TODO(), params.Snapshot, authConfigRef)
					Expect(err).NotTo(HaveOccurred())
				})

				It("should not error - nil labels", func() {
					secret.Metadata.Labels = nil
					_, err := extauthsyncer.TranslateExtAuthConfig(context.TODO(), params.Snapshot, authConfigRef)
					Expect(err).NotTo(HaveOccurred())
				})

			})

		})
	})

	Context("with OPA extauth", func() {
		BeforeEach(func() {
			authConfig = &extauth.AuthConfig{
				Metadata: &core.Metadata{
					Name:      "oauth",
					Namespace: "gloo-system",
				},
				Configs: []*extauth.AuthConfig_Config{{
					AuthConfig: &extauth.AuthConfig_Config_OpaAuth{
						OpaAuth: &extauth.OpaAuth{
							Modules: []*core.ResourceRef{{Namespace: "namespace", Name: "name"}},
							Query:   "true",
						},
					},
				}},
			}
			authConfigRef = authConfig.Metadata.Ref()
			extAuthExtension = &extauth.ExtAuthExtension{
				Spec: &extauth.ExtAuthExtension_ConfigRef{
					ConfigRef: authConfigRef,
				},
			}

			params.Snapshot = &v1snap.ApiSnapshot{
				Upstreams:   v1.UpstreamList{upstream},
				AuthConfigs: extauth.AuthConfigList{authConfig},
			}

			params.Snapshot.Artifacts = v1.ArtifactList{
				{

					Metadata: &core.Metadata{
						Name:      "name",
						Namespace: "namespace",
					},
					Data: map[string]string{"module.rego": "package foo"},
				},
			}
		})

		It("should translate OPA config without options specified", func() {
			translated, err := extauthsyncer.TranslateExtAuthConfig(context.TODO(), params.Snapshot, authConfigRef)
			Expect(err).NotTo(HaveOccurred())
			Expect(translated.AuthConfigRefName).To(Equal(authConfigRef.Key()))
			Expect(translated.Configs).To(HaveLen(1))
			actual := translated.Configs[0].GetOpaAuth()
			expected := authConfig.Configs[0].GetOpaAuth()
			Expect(actual.Query).To(Equal(expected.Query))
			data := params.Snapshot.Artifacts[0].Data
			Expect(actual.Modules).To(Equal(data))
			Expect(actual.Options).To(Equal(expected.Options))
		})

		It("Should translate OPA config with options specified", func() {
			// Specify additional options in Opa Config.
			opaAuth := authConfig.Configs[0].GetOpaAuth()
			opaAuth.Options = &extauth.OpaAuthOptions{
				FastInputConversion: true,
			}

			translated, err := extauthsyncer.TranslateExtAuthConfig(context.TODO(), params.Snapshot, authConfigRef)
			Expect(err).NotTo(HaveOccurred())
			Expect(translated.AuthConfigRefName).To(Equal(authConfigRef.Key()))
			Expect(translated.Configs).To(HaveLen(1))
			Expect(translated.Configs[0].GetOpaAuth().GetOptions().GetFastInputConversion()).To(Equal(true))
			actual := translated.Configs[0].GetOpaAuth()
			expected := authConfig.Configs[0].GetOpaAuth()
			Expect(actual.Query).To(Equal(expected.Query))
			data := params.Snapshot.Artifacts[0].Data
			Expect(actual.Modules).To(Equal(data))
			Expect(actual.Options).To(Equal(expected.Options))
		})
	})

	Context("with AccessTokenValidation extauth", func() {
		BeforeEach(func() {
			authConfig = &extauth.AuthConfig{
				Metadata: &core.Metadata{
					Name:      "oauth",
					Namespace: "gloo-system",
				},
				Configs: []*extauth.AuthConfig_Config{{
					AuthConfig: &extauth.AuthConfig_Config_Oauth2{
						Oauth2: &extauth.OAuth2{
							OauthType: &extauth.OAuth2_AccessTokenValidation{
								AccessTokenValidation: &extauth.AccessTokenValidation{
									ValidationType: &extauth.AccessTokenValidation_Introspection{
										Introspection: &extauth.IntrospectionValidation{
											IntrospectionUrl:    "introspection-url",
											ClientId:            "client-id",
											ClientSecretRef:     secret.Metadata.Ref(),
											UserIdAttributeName: "sub",
										},
									},
									CacheTimeout: ptypes.DurationProto(time.Minute),
									UserinfoUrl:  "user-info-url",
									ScopeValidation: &extauth.AccessTokenValidation_RequiredScopes{
										RequiredScopes: &extauth.AccessTokenValidation_ScopeList{
											Scope: []string{"foo", "bar"},
										},
									},
								},
							},
						},
					},
				}},
			}
			authConfigRef = authConfig.Metadata.Ref()
			extAuthExtension = &extauth.ExtAuthExtension{
				Spec: &extauth.ExtAuthExtension_ConfigRef{
					ConfigRef: authConfigRef,
				},
			}

			params.Snapshot = &v1snap.ApiSnapshot{
				Upstreams:   v1.UpstreamList{upstream},
				AuthConfigs: extauth.AuthConfigList{authConfig},
			}

		})

		It("should succeed for IntrospectionValidation config", func() {
			translated, err := extauthsyncer.TranslateExtAuthConfig(context.TODO(), params.Snapshot, authConfigRef)
			Expect(err).NotTo(HaveOccurred())

			Expect(translated.AuthConfigRefName).To(Equal(authConfigRef.Key()))
			Expect(translated.Configs).To(HaveLen(1))

			actual := translated.Configs[0].GetOauth2().GetAccessTokenValidationConfig()
			expected := authConfig.Configs[0].GetOauth2().GetAccessTokenValidation()

			Expect(actual.GetUserinfoUrl()).To(Equal(expected.GetUserinfoUrl()))
			Expect(actual.GetCacheTimeout()).To(Equal(expected.GetCacheTimeout()))
			Expect(actual.GetIntrospection().GetIntrospectionUrl()).To(Equal(expected.GetIntrospection().GetIntrospectionUrl()))
			Expect(actual.GetIntrospection().GetClientId()).To(Equal(expected.GetIntrospection().GetClientId()))
			Expect(actual.GetIntrospection().GetClientSecret()).To(Equal(clientSecret.ClientSecret))
			Expect(actual.GetIntrospection().GetUserIdAttributeName()).To(Equal(expected.GetIntrospection().GetUserIdAttributeName()))
			Expect(actual.GetRequiredScopes().GetScope()).To(Equal(expected.GetRequiredScopes().GetScope()))
		})
	})

	Context("with Ldap extauth", func() {
		BeforeEach(func() {
			authConfig = &extauth.AuthConfig{
				Metadata: &core.Metadata{
					Name:      "ldap",
					Namespace: "gloo-system",
				},
				Configs: []*extauth.AuthConfig_Config{{
					AuthConfig: &extauth.AuthConfig_Config_Ldap{
						Ldap: &extauth.Ldap{
							Address:                 "my.server.com:389",
							UserDnTemplate:          "uid=%s,ou=people,dc=solo,dc=io",
							MembershipAttributeName: "someName",
							AllowedGroups: []string{
								"cn=managers,ou=groups,dc=solo,dc=io",
								"cn=developers,ou=groups,dc=solo,dc=io",
							},
							Pool: &extauth.Ldap_ConnectionPool{
								MaxSize: &wrappers.UInt32Value{
									Value: uint32(5),
								},
								InitialSize: &wrappers.UInt32Value{
									Value: uint32(0), // Set to 0, otherwise it will try to connect to the dummy address
								},
							},
							SearchFilter:         "(objectClass=*)",
							DisableGroupChecking: false,
							GroupLookupSettings: &extauth.LdapServiceAccount{
								CheckGroupsWithServiceAccount: true,
								CredentialsSecretRef:          ldapSecret.Metadata.Ref(),
							},
						},
					},
				}},
			}
			authConfigRef = authConfig.Metadata.Ref()
			extAuthExtension = &extauth.ExtAuthExtension{
				Spec: &extauth.ExtAuthExtension_ConfigRef{
					ConfigRef: authConfigRef,
				},
			}

			params.Snapshot = &v1snap.ApiSnapshot{
				Upstreams:   v1.UpstreamList{upstream},
				AuthConfigs: extauth.AuthConfigList{authConfig},
			}
		})
		It("translates ldap config", func() {

			translated, err := extauthsyncer.TranslateExtAuthConfig(context.TODO(), params.Snapshot, authConfigRef)
			Expect(err).NotTo(HaveOccurred())
			Expect(translated.AuthConfigRefName).To(Equal(authConfigRef.Key()))
			Expect(translated.Configs).To(HaveLen(1))
			actual := translated.Configs[0].GetLdapInternal()
			expected := authConfig.Configs[0].GetLdap()
			Expect(actual.Address).To(Equal(expected.Address))
			Expect(actual.AllowedGroups).To(BeEquivalentTo(expected.AllowedGroups))
			Expect(actual.Pool.MaxSize).To(Equal(expected.Pool.MaxSize))
			Expect(actual.Pool.InitialSize).To(Equal(expected.Pool.InitialSize))
			Expect(actual.SearchFilter).To(Equal(expected.SearchFilter))
			Expect(actual.GroupLookupSettings.Username).To(Equal("user"))
			Expect(actual.GroupLookupSettings.Password).To(Equal("pass"))
			Expect(actual.UserDnTemplate).To(Equal(expected.UserDnTemplate))
			Expect(actual.MembershipAttributeName).To(Equal(expected.MembershipAttributeName))
		})
		It("does not require credentials when service account is not required ", func() {
			authConfig.Configs[0].GetLdap().GroupLookupSettings.CheckGroupsWithServiceAccount = false
			authConfig.Configs[0].GetLdap().GroupLookupSettings.CredentialsSecretRef = nil

			_, err := extauthsyncer.TranslateExtAuthConfig(context.TODO(), params.Snapshot, authConfigRef)
			Expect(err).NotTo(HaveOccurred())
		})
		It("requires credentials when service account is required", func() {
			authConfig.Configs[0].GetLdap().GroupLookupSettings.CredentialsSecretRef = nil
			_, err := extauthsyncer.TranslateExtAuthConfig(context.TODO(), params.Snapshot, authConfigRef)
			Expect(err).To(HaveOccurred())
		})
		It("Uses the old API when settings for service account are not present", func() {
			authConfig.Configs[0].GetLdap().GroupLookupSettings = nil

			translated, err := extauthsyncer.TranslateExtAuthConfig(context.TODO(), params.Snapshot, authConfigRef)
			Expect(err).NotTo(HaveOccurred())
			final := translated.Configs[0].GetLdap()
			Expect(final).NotTo(BeNil(), "Old API should be used")
		})
	})
	Context("HMAC extauth", func() {
		BeforeEach(func() {
			authConfig = &extauth.AuthConfig{
				Metadata: &core.Metadata{
					Name:      "hmac",
					Namespace: "gloo-system",
				},
				Configs: []*extauth.AuthConfig_Config{{
					AuthConfig: &extauth.AuthConfig_Config_HmacAuth{
						HmacAuth: &extauth.HmacAuth{
							SecretStorage: &extauth.HmacAuth_SecretRefs{
								SecretRefs: &extauth.SecretRefList{
									SecretRefs: []*core.ResourceRef{
										ldapSecret.Metadata.Ref(),
									}},
							},
							ImplementationType: &extauth.HmacAuth_ParametersInHeaders{ParametersInHeaders: &extauth.HmacParametersInHeaders{}},
						},
					}},
				}}
			authConfigRef = authConfig.Metadata.Ref()
			extAuthExtension = &extauth.ExtAuthExtension{
				Spec: &extauth.ExtAuthExtension_ConfigRef{
					ConfigRef: authConfigRef,
				},
			}
			params.Snapshot = &v1snap.ApiSnapshot{
				Upstreams:   v1.UpstreamList{upstream},
				AuthConfigs: extauth.AuthConfigList{authConfig},
			}
		})
		It("Translates valid HMAC config", func() {

			translated, err := extauthsyncer.TranslateExtAuthConfig(context.TODO(), params.Snapshot, authConfigRef)
			Expect(err).NotTo(HaveOccurred())
			Expect(translated.AuthConfigRefName).To(Equal(authConfigRef.Key()))
			Expect(translated.Configs).To(HaveLen(1))
			actual := translated.Configs[0].GetHmacAuth()
			Expect(actual.GetSecretList().GetSecretList()[ldapSecret.GetCredentials().GetUsername()]).To(Equal(ldapSecret.GetCredentials().GetPassword()))
		})
		It("errors when no secrets are provided", func() {
			authConfig.Configs[0].GetHmacAuth().SecretStorage = &extauth.HmacAuth_SecretRefs{
				&extauth.SecretRefList{
					SecretRefs: []*core.ResourceRef{},
				},
			}
			_, err := extauthsyncer.TranslateExtAuthConfig(context.TODO(), params.Snapshot, authConfigRef)
			Expect(err).NotTo(BeNil())
		})
	})
})