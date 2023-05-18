// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-projects/projects/apiserver/api/fed.rpc/v1/federated_gateway_resources.proto

package v1

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"

	safe_hasher "github.com/solo-io/protoc-gen-ext/pkg/hasher"
	"github.com/solo-io/protoc-gen-ext/pkg/hasher/hashstructure"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = new(hash.Hash64)
	_ = fnv.New64
	_ = hashstructure.Hash
	_ = new(safe_hasher.SafeHasher)
)

// Hash function
func (m *FederatedGateway) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-projects/projects/apiserver/pkg/api/fed.rpc/v1.FederatedGateway")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMetadata()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Metadata")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMetadata(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Metadata")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetSpec()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Spec")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSpec(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Spec")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetStatus()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Status")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetStatus(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Status")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *FederatedMatchableHttpGateway) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-projects/projects/apiserver/pkg/api/fed.rpc/v1.FederatedMatchableHttpGateway")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMetadata()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Metadata")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMetadata(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Metadata")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetSpec()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Spec")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSpec(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Spec")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetStatus()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Status")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetStatus(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Status")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *FederatedMatchableTcpGateway) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-projects/projects/apiserver/pkg/api/fed.rpc/v1.FederatedMatchableTcpGateway")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMetadata()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Metadata")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMetadata(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Metadata")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetSpec()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Spec")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSpec(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Spec")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetStatus()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Status")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetStatus(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Status")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *FederatedVirtualService) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-projects/projects/apiserver/pkg/api/fed.rpc/v1.FederatedVirtualService")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMetadata()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Metadata")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMetadata(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Metadata")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetSpec()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Spec")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSpec(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Spec")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetStatus()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Status")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetStatus(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Status")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *FederatedRouteTable) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-projects/projects/apiserver/pkg/api/fed.rpc/v1.FederatedRouteTable")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMetadata()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Metadata")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMetadata(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Metadata")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetSpec()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Spec")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSpec(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Spec")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetStatus()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Status")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetStatus(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Status")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ListFederatedGatewaysRequest) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-projects/projects/apiserver/pkg/api/fed.rpc/v1.ListFederatedGatewaysRequest")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ListFederatedGatewaysResponse) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-projects/projects/apiserver/pkg/api/fed.rpc/v1.ListFederatedGatewaysResponse")); err != nil {
		return 0, err
	}

	for _, v := range m.GetFederatedGateways() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GetFederatedGatewayYamlRequest) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-projects/projects/apiserver/pkg/api/fed.rpc/v1.GetFederatedGatewayYamlRequest")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetFederatedGatewayRef()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("FederatedGatewayRef")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetFederatedGatewayRef(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("FederatedGatewayRef")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GetFederatedGatewayYamlResponse) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-projects/projects/apiserver/pkg/api/fed.rpc/v1.GetFederatedGatewayYamlResponse")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetYamlData()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("YamlData")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetYamlData(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("YamlData")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ListFederatedMatchableHttpGatewaysRequest) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-projects/projects/apiserver/pkg/api/fed.rpc/v1.ListFederatedMatchableHttpGatewaysRequest")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ListFederatedMatchableHttpGatewaysResponse) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-projects/projects/apiserver/pkg/api/fed.rpc/v1.ListFederatedMatchableHttpGatewaysResponse")); err != nil {
		return 0, err
	}

	for _, v := range m.GetFederatedMatchableHttpGateways() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GetFederatedMatchableHttpGatewayYamlRequest) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-projects/projects/apiserver/pkg/api/fed.rpc/v1.GetFederatedMatchableHttpGatewayYamlRequest")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetFederatedMatchableHttpGatewayRef()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("FederatedMatchableHttpGatewayRef")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetFederatedMatchableHttpGatewayRef(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("FederatedMatchableHttpGatewayRef")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GetFederatedMatchableHttpGatewayYamlResponse) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-projects/projects/apiserver/pkg/api/fed.rpc/v1.GetFederatedMatchableHttpGatewayYamlResponse")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetYamlData()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("YamlData")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetYamlData(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("YamlData")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ListFederatedMatchableTcpGatewaysRequest) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-projects/projects/apiserver/pkg/api/fed.rpc/v1.ListFederatedMatchableTcpGatewaysRequest")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ListFederatedMatchableTcpGatewaysResponse) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-projects/projects/apiserver/pkg/api/fed.rpc/v1.ListFederatedMatchableTcpGatewaysResponse")); err != nil {
		return 0, err
	}

	for _, v := range m.GetFederatedMatchableTcpGateways() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GetFederatedMatchableTcpGatewayYamlRequest) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-projects/projects/apiserver/pkg/api/fed.rpc/v1.GetFederatedMatchableTcpGatewayYamlRequest")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetFederatedMatchableTcpGatewayRef()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("FederatedMatchableTcpGatewayRef")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetFederatedMatchableTcpGatewayRef(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("FederatedMatchableTcpGatewayRef")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GetFederatedMatchableTcpGatewayYamlResponse) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-projects/projects/apiserver/pkg/api/fed.rpc/v1.GetFederatedMatchableTcpGatewayYamlResponse")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetYamlData()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("YamlData")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetYamlData(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("YamlData")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ListFederatedVirtualServicesRequest) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-projects/projects/apiserver/pkg/api/fed.rpc/v1.ListFederatedVirtualServicesRequest")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ListFederatedVirtualServicesResponse) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-projects/projects/apiserver/pkg/api/fed.rpc/v1.ListFederatedVirtualServicesResponse")); err != nil {
		return 0, err
	}

	for _, v := range m.GetFederatedVirtualServices() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GetFederatedVirtualServiceYamlRequest) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-projects/projects/apiserver/pkg/api/fed.rpc/v1.GetFederatedVirtualServiceYamlRequest")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetFederatedVirtualServiceRef()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("FederatedVirtualServiceRef")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetFederatedVirtualServiceRef(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("FederatedVirtualServiceRef")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GetFederatedVirtualServiceYamlResponse) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-projects/projects/apiserver/pkg/api/fed.rpc/v1.GetFederatedVirtualServiceYamlResponse")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetYamlData()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("YamlData")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetYamlData(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("YamlData")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ListFederatedRouteTablesRequest) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-projects/projects/apiserver/pkg/api/fed.rpc/v1.ListFederatedRouteTablesRequest")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ListFederatedRouteTablesResponse) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-projects/projects/apiserver/pkg/api/fed.rpc/v1.ListFederatedRouteTablesResponse")); err != nil {
		return 0, err
	}

	for _, v := range m.GetFederatedRouteTables() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GetFederatedRouteTableYamlRequest) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-projects/projects/apiserver/pkg/api/fed.rpc/v1.GetFederatedRouteTableYamlRequest")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetFederatedRouteTableRef()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("FederatedRouteTableRef")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetFederatedRouteTableRef(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("FederatedRouteTableRef")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GetFederatedRouteTableYamlResponse) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-projects/projects/apiserver/pkg/api/fed.rpc/v1.GetFederatedRouteTableYamlResponse")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetYamlData()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("YamlData")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetYamlData(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("YamlData")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}
