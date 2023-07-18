package sanitize_cluster_header_test

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/wrappers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/extauth"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	. "github.com/solo-io/solo-projects/projects/gloo/pkg/plugins/sanitize_cluster_header"
)

var _ = Describe("Plugin", func() {
	var (
		plugin plugins.HttpFilterPlugin
	)

	BeforeEach(func() {
		plugin = NewPlugin()
	})

	Context("with cluster header sanitation enabled", func() {

		It("should only sanitize one header name even through multiple Routes with identical Destinations", func() {

			listener := &v1.HttpListener{
				Options: &v1.HttpListenerOptions{
					SanitizeClusterHeader: &wrappers.BoolValue{Value: true},
				},
				VirtualHosts: []*v1.VirtualHost{
					{
						Routes: []*v1.Route{
							{
								Action: &v1.Route_RouteAction{
									RouteAction: &v1.RouteAction{
										Destination: &v1.RouteAction_ClusterHeader{
											ClusterHeader: "test-cluster-header",
										},
									},
								},
							},
							{
								Action: &v1.Route_RouteAction{
									RouteAction: &v1.RouteAction{
										Destination: &v1.RouteAction_ClusterHeader{
											ClusterHeader: "test-cluster-header",
										},
									},
								},
							},
						},
					},
				},
			}

			filters, err := plugin.HttpFilters(plugins.Params{}, listener)
			Expect(err).NotTo(HaveOccurred())
			Expect(filters).To(HaveLen(1))
			goTypedConfig := filters[0].HttpFilter.GetTypedConfig()
			var filter = extauth.Sanitize{}
			err = ptypes.UnmarshalAny(goTypedConfig, &filter)
			Expect(err).NotTo(HaveOccurred())
			Expect(filter.HeadersToRemove).To(Equal([]string{"test-cluster-header"}))
		})

		It("should sanitize each cluster header name available", func() {
			listener := &v1.HttpListener{
				Options: &v1.HttpListenerOptions{
					SanitizeClusterHeader: &wrappers.BoolValue{Value: true},
				},
				VirtualHosts: []*v1.VirtualHost{
					{
						Routes: []*v1.Route{
							{
								Action: &v1.Route_RouteAction{
									RouteAction: &v1.RouteAction{
										Destination: &v1.RouteAction_ClusterHeader{
											ClusterHeader: "test-cluster-header",
										},
									},
								},
							},
							{
								Action: &v1.Route_RouteAction{
									RouteAction: &v1.RouteAction{
										Destination: &v1.RouteAction_ClusterHeader{
											ClusterHeader: "another-cluster-header",
										},
									},
								},
							},
						},
					},
					{
						Routes: []*v1.Route{
							{
								Action: &v1.Route_RouteAction{
									RouteAction: &v1.RouteAction{
										Destination: &v1.RouteAction_ClusterHeader{
											ClusterHeader: "last-cluster-header",
										},
									},
								},
							},
						},
					},
				},
			}

			filters, err := plugin.HttpFilters(plugins.Params{}, listener)
			Expect(err).NotTo(HaveOccurred())
			goTypedConfig := filters[0].HttpFilter.GetTypedConfig()
			var filter = extauth.Sanitize{}
			err = ptypes.UnmarshalAny(goTypedConfig, &filter)
			Expect(err).NotTo(HaveOccurred())
			Expect(filter.HeadersToRemove).To(HaveLen(3))
			Expect(filter.HeadersToRemove).To(ConsistOf("last-cluster-header", "test-cluster-header", "another-cluster-header"))
		})
	})

	Context("with cluster header sanitation disabled", func() {

		It("should not provide any sanitize filters", func() {
			listener := &v1.HttpListener{
				Options: &v1.HttpListenerOptions{
					SanitizeClusterHeader: nil,
				},
				VirtualHosts: []*v1.VirtualHost{
					{
						Routes: []*v1.Route{
							{
								Action: &v1.Route_RouteAction{
									RouteAction: &v1.RouteAction{
										Destination: &v1.RouteAction_ClusterHeader{
											ClusterHeader: "test-cluster-header",
										},
									},
								},
							},
						},
					},
				},
			}

			filters, err := plugin.HttpFilters(plugins.Params{}, listener)
			Expect(err).NotTo(HaveOccurred())
			Expect(filters).To(BeEmpty())

		})
	})

})