// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"fmt"
	"hash"
	"hash/fnv"
	"log"

	github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes "github.com/solo-io/solo-kit/pkg/api/v1/resources/common/kubernetes"

	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/hashutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type DiscoverySnapshot struct {
	Upstreams      UpstreamList
	Kubenamespaces github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.KubeNamespaceList
	Secrets        SecretList
}

func (s DiscoverySnapshot) Clone() DiscoverySnapshot {
	return DiscoverySnapshot{
		Upstreams:      s.Upstreams.Clone(),
		Kubenamespaces: s.Kubenamespaces.Clone(),
		Secrets:        s.Secrets.Clone(),
	}
}

func (s DiscoverySnapshot) Hash(hasher hash.Hash64) (uint64, error) {
	if hasher == nil {
		hasher = fnv.New64()
	}
	if _, err := s.hashUpstreams(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashKubenamespaces(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashSecrets(hasher); err != nil {
		return 0, err
	}
	return hasher.Sum64(), nil
}

func (s DiscoverySnapshot) hashUpstreams(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.Upstreams.AsInterfaces()...)
}

func (s DiscoverySnapshot) hashKubenamespaces(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.Kubenamespaces.AsInterfaces()...)
}

func (s DiscoverySnapshot) hashSecrets(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.Secrets.AsInterfaces()...)
}

func (s DiscoverySnapshot) HashFields() []zap.Field {
	var fields []zap.Field
	hasher := fnv.New64()
	UpstreamsHash, err := s.hashUpstreams(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("upstreams", UpstreamsHash))
	KubenamespacesHash, err := s.hashKubenamespaces(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("kubenamespaces", KubenamespacesHash))
	SecretsHash, err := s.hashSecrets(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("secrets", SecretsHash))
	snapshotHash, err := s.Hash(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	return append(fields, zap.Uint64("snapshotHash", snapshotHash))
}

func (s *DiscoverySnapshot) GetResourcesList(resource resources.Resource) (resources.ResourceList, error) {
	switch resource.(type) {
	case *Upstream:
		return s.Upstreams.AsResources(), nil
	case *github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.KubeNamespace:
		return s.Kubenamespaces.AsResources(), nil
	case *Secret:
		return s.Secrets.AsResources(), nil
	default:
		return resources.ResourceList{}, eris.New("did not contain the input resource type returning empty list")
	}
}

func (s *DiscoverySnapshot) RemoveFromResourceList(resource resources.Resource) error {
	refKey := resource.GetMetadata().Ref().Key()
	switch resource.(type) {
	case *Upstream:

		for i, res := range s.Upstreams {
			if refKey == res.GetMetadata().Ref().Key() {
				s.Upstreams = append(s.Upstreams[:i], s.Upstreams[i+1:]...)
				break
			}
		}
		return nil
	case *github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.KubeNamespace:

		for i, res := range s.Kubenamespaces {
			if refKey == res.GetMetadata().Ref().Key() {
				s.Kubenamespaces = append(s.Kubenamespaces[:i], s.Kubenamespaces[i+1:]...)
				break
			}
		}
		return nil
	case *Secret:

		for i, res := range s.Secrets {
			if refKey == res.GetMetadata().Ref().Key() {
				s.Secrets = append(s.Secrets[:i], s.Secrets[i+1:]...)
				break
			}
		}
		return nil
	default:
		return eris.Errorf("did not remove the resource because its type does not exist [%T]", resource)
	}
}

func (s *DiscoverySnapshot) RemoveMatches(predicate core.Predicate) {
	var Upstreams UpstreamList
	for _, res := range s.Upstreams {
		if matches := predicate(res.GetMetadata()); !matches {
			Upstreams = append(Upstreams, res)
		}
	}
	s.Upstreams = Upstreams
	var Kubenamespaces github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.KubeNamespaceList
	for _, res := range s.Kubenamespaces {
		if matches := predicate(res.GetMetadata()); !matches {
			Kubenamespaces = append(Kubenamespaces, res)
		}
	}
	s.Kubenamespaces = Kubenamespaces
	var Secrets SecretList
	for _, res := range s.Secrets {
		if matches := predicate(res.GetMetadata()); !matches {
			Secrets = append(Secrets, res)
		}
	}
	s.Secrets = Secrets
}

func (s *DiscoverySnapshot) UpsertToResourceList(resource resources.Resource) error {
	refKey := resource.GetMetadata().Ref().Key()
	switch typed := resource.(type) {
	case *Upstream:
		updated := false
		for i, res := range s.Upstreams {
			if refKey == res.GetMetadata().Ref().Key() {
				s.Upstreams[i] = typed
				updated = true
			}
		}
		if !updated {
			s.Upstreams = append(s.Upstreams, typed)
		}
		s.Upstreams.Sort()
		return nil
	case *github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.KubeNamespace:
		updated := false
		for i, res := range s.Kubenamespaces {
			if refKey == res.GetMetadata().Ref().Key() {
				s.Kubenamespaces[i] = typed
				updated = true
			}
		}
		if !updated {
			s.Kubenamespaces = append(s.Kubenamespaces, typed)
		}
		s.Kubenamespaces.Sort()
		return nil
	case *Secret:
		updated := false
		for i, res := range s.Secrets {
			if refKey == res.GetMetadata().Ref().Key() {
				s.Secrets[i] = typed
				updated = true
			}
		}
		if !updated {
			s.Secrets = append(s.Secrets, typed)
		}
		s.Secrets.Sort()
		return nil
	default:
		return eris.Errorf("did not add/replace the resource type because it does not exist %T", resource)
	}
}

type DiscoverySnapshotStringer struct {
	Version        uint64
	Upstreams      []string
	Kubenamespaces []string
	Secrets        []string
}

func (ss DiscoverySnapshotStringer) String() string {
	s := fmt.Sprintf("DiscoverySnapshot %v\n", ss.Version)

	s += fmt.Sprintf("  Upstreams %v\n", len(ss.Upstreams))
	for _, name := range ss.Upstreams {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Kubenamespaces %v\n", len(ss.Kubenamespaces))
	for _, name := range ss.Kubenamespaces {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Secrets %v\n", len(ss.Secrets))
	for _, name := range ss.Secrets {
		s += fmt.Sprintf("    %v\n", name)
	}

	return s
}

func (s DiscoverySnapshot) Stringer() DiscoverySnapshotStringer {
	snapshotHash, err := s.Hash(nil)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	return DiscoverySnapshotStringer{
		Version:        snapshotHash,
		Upstreams:      s.Upstreams.NamespacesDotNames(),
		Kubenamespaces: s.Kubenamespaces.Names(),
		Secrets:        s.Secrets.NamespacesDotNames(),
	}
}

var DiscoveryGvkToHashableResource = map[schema.GroupVersionKind]func() resources.HashableResource{
	UpstreamGVK: NewUpstreamHashableResource,
	github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.KubeNamespaceGVK: github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.NewKubeNamespaceHashableResource,
	SecretGVK: NewSecretHashableResource,
}
