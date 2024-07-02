// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/ai/ai.proto

package ai

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
func (m *UpstreamSpec) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.UpstreamSpec")); err != nil {
		return 0, err
	}

	switch m.AuthToken.(type) {

	case *UpstreamSpec_InlineAuthToken:

		if _, err = hasher.Write([]byte(m.GetInlineAuthToken())); err != nil {
			return 0, err
		}

	case *UpstreamSpec_AuthTokenRef:

		if _, err = hasher.Write([]byte(m.GetAuthTokenRef())); err != nil {
			return 0, err
		}

	}

	switch m.Llm.(type) {

	case *UpstreamSpec_Openai:

		if h, ok := interface{}(m.GetOpenai()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Openai")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetOpenai(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Openai")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *UpstreamSpec_Mistral_:

		if h, ok := interface{}(m.GetMistral()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Mistral")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetMistral(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Mistral")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *UpstreamSpec_Anthropic_:

		if h, ok := interface{}(m.GetAnthropic()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Anthropic")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetAnthropic(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Anthropic")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *UpstreamSpec_Custom_:

		if h, ok := interface{}(m.GetCustom()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Custom")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetCustom(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Custom")); err != nil {
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
func (m *RouteSettings) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.RouteSettings")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetPromptEnrichment()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("PromptEnrichment")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetPromptEnrichment(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("PromptEnrichment")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetPromptGuard()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("PromptGuard")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetPromptGuard(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("PromptGuard")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetRateLimiting()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RateLimiting")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRateLimiting(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RateLimiting")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetRag()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Rag")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRag(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Rag")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetSemanticCaching()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("SemanticCaching")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSemanticCaching(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("SemanticCaching")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	for _, v := range m.GetBackupModels() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	for _, v := range m.GetDefaults() {

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
func (m *FieldDefault) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.FieldDefault")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetField())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetValue()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Value")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetValue(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Value")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetOverride())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *Postgres) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.Postgres")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetConnectionString())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetCollectionName())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *Redis) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.Redis")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetConnectionString())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *Embedding) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.Embedding")); err != nil {
		return 0, err
	}

	switch m.Embedding.(type) {

	case *Embedding_Openai:

		if h, ok := interface{}(m.GetOpenai()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Openai")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetOpenai(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Openai")); err != nil {
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
func (m *SemanticCaching) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.SemanticCaching")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetDatastore()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Datastore")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDatastore(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Datastore")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetEmbedding()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Embedding")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetEmbedding(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Embedding")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetTtl())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *RAG) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.RAG")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetDatastore()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Datastore")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDatastore(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Datastore")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetEmbedding()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Embedding")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetEmbedding(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Embedding")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte(m.GetPromptTemplate())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *RateLimiting) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.RateLimiting")); err != nil {
		return 0, err
	}

	for _, v := range m.GetRateLimitConfigs() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *AIPromptEnrichment) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.AIPromptEnrichment")); err != nil {
		return 0, err
	}

	for _, v := range m.GetPrepend() {

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

	for _, v := range m.GetAppend() {

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
func (m *AIPromptGaurd) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.AIPromptGaurd")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetRequest()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Request")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRequest(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Request")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetResponse()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Response")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetResponse(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Response")); err != nil {
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
func (m *UpstreamSpec_OpenAI) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.UpstreamSpec_OpenAI")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *UpstreamSpec_Mistral) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.UpstreamSpec_Mistral")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *UpstreamSpec_Anthropic) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.UpstreamSpec_Anthropic")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetVersion())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *UpstreamSpec_Custom) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.UpstreamSpec_Custom")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetHost())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetPort())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *Embedding_OpenAI) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.Embedding_OpenAI")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *SemanticCaching_DataStore) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.SemanticCaching_DataStore")); err != nil {
		return 0, err
	}

	switch m.Datastore.(type) {

	case *SemanticCaching_DataStore_Redis:

		if h, ok := interface{}(m.GetRedis()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Redis")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetRedis(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Redis")); err != nil {
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
func (m *RAG_DataStore) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.RAG_DataStore")); err != nil {
		return 0, err
	}

	switch m.Datastore.(type) {

	case *RAG_DataStore_Postgres:

		if h, ok := interface{}(m.GetPostgres()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Postgres")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetPostgres(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Postgres")); err != nil {
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
func (m *AIPromptEnrichment_Message) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.AIPromptEnrichment_Message")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetRole())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetContent())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetRoleOverride())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *AIPromptGaurd_Request) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.AIPromptGaurd_Request")); err != nil {
		return 0, err
	}

	for _, v := range m.GetMatches() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	for _, v := range m.GetBuiltins() {

		err = binary.Write(hasher, binary.LittleEndian, v)
		if err != nil {
			return 0, err
		}

	}

	if _, err = hasher.Write([]byte(m.GetCustomResponseMessage())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *AIPromptGaurd_Response) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("ai.options.gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/ai.AIPromptGaurd_Response")); err != nil {
		return 0, err
	}

	for _, v := range m.GetMatches() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	for _, v := range m.GetBuiltins() {

		err = binary.Write(hasher, binary.LittleEndian, v)
		if err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}
