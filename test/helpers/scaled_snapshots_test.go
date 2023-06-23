package helpers_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/test/helpers"
)

var _ = Describe("ScaledSnapshotBuilder", func() {
	When("with endpoints", func() {
		It("generates a snapshot with the expected number of endpoints", func() {
			snap := helpers.NewScaledSnapshotBuilder().WithEndpointCount(10).Build()
			Expect(snap.Endpoints).To(HaveLen(10))
			Expect(snap.Upstreams).To(HaveLen(0))
		})
	})

	When("with upstreams", func() {
		It("generates a snapshot with the expected number of upstreams", func() {
			snap := helpers.NewScaledSnapshotBuilder().WithUpstreamCount(10).Build()
			Expect(snap.Endpoints).To(HaveLen(0))
			Expect(snap.Upstreams).To(HaveLen(10))
		})
	})

	When("with upstream builder", func() {
		When("with consistent SNI", func() {
			It("generates a snapshot with upstreams that all have the same SNI", func() {
				snap := helpers.NewScaledSnapshotBuilder().WithUpstreamCount(10).
					WithUpstreamBuilder(helpers.NewUpstreamBuilder().WithConsistentSni()).Build()
				Expect(snap.Upstreams).To(HaveLen(10))
				Expect(snap.Upstreams[0].SslConfig).NotTo(BeNil())
				firstSNI := snap.Upstreams[0].SslConfig.Sni
				for i := 1; i < len(snap.Upstreams); i++ {
					Expect(snap.Upstreams[i].SslConfig).NotTo(BeNil())
					Expect(snap.Upstreams[i].SslConfig.Sni).To(Equal(firstSNI))
				}
			})
		})

		When("with unique SNI", func() {
			It("generates a snapshot with upstreams that all have unique SNI", func() {
				snap := helpers.NewScaledSnapshotBuilder().WithUpstreamCount(10).
					WithUpstreamBuilder(helpers.NewUpstreamBuilder().WithUniqueSni()).Build()
				Expect(snap.Upstreams).To(HaveLen(10))
				foundSNI := map[string]bool{}
				for i := 0; i < len(snap.Upstreams); i++ {
					Expect(snap.Upstreams[i].SslConfig).NotTo(BeNil())
					_, ok := foundSNI[snap.Upstreams[i].SslConfig.Sni]
					Expect(ok).To(BeFalse())
					foundSNI[snap.Upstreams[i].SslConfig.Sni] = true
				}
			})
		})
	})
})
