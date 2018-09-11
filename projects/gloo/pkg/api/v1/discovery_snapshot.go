package v1

import (
	"github.com/mitchellh/hashstructure"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

type DiscoverySnapshot struct {
	Secrets   SecretsByNamespace
	Upstreams UpstreamsByNamespace
}

func (s DiscoverySnapshot) Clone() DiscoverySnapshot {
	return DiscoverySnapshot{
		Secrets:   s.Secrets.Clone(),
		Upstreams: s.Upstreams.Clone(),
	}
}

func (s DiscoverySnapshot) Hash() uint64 {
	snapshotForHashing := s.Clone()
	for _, secret := range snapshotForHashing.Secrets.List() {
		resources.UpdateMetadata(secret, func(meta *core.Metadata) {
			meta.ResourceVersion = ""
		})
	}
	for _, upstream := range snapshotForHashing.Upstreams.List() {
		resources.UpdateMetadata(upstream, func(meta *core.Metadata) {
			meta.ResourceVersion = ""
		})
		upstream.SetStatus(core.Status{})
	}
	h, err := hashstructure.Hash(snapshotForHashing, nil)
	if err != nil {
		panic(err)
	}
	return h
}
