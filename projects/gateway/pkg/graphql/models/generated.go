// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	fmt "fmt"
	io "io"
	strconv "strconv"

	customtypes "github.com/solo-io/solo-kit/projects/gateway/pkg/graphql/customtypes"
)

type AwsDestinationSpec struct {
	LogicalName     string                   `json:"logicalName"`
	InvocationStyle AwsLambdaInvocationStyle `json:"invocationStyle"`
}
type AwsLambdaFunction struct {
	LogicalName  string `json:"logicalName"`
	FunctionName string `json:"functionName"`
	Qualifier    string `json:"qualifier"`
}
type AwsUpstreamSpec struct {
	Region    string               `json:"region"`
	SecretRef string               `json:"secretRef"`
	Functions []*AwsLambdaFunction `json:"functions"`
}
type AzureDestinationSpec struct {
	FunctionName string `json:"functionName"`
}
type AzureFunction struct {
	FunctionName string `json:"functionName"`
	AuthLevel    string `json:"authLevel"`
}
type AzureUpstreamSpec struct {
	FunctionAppName string           `json:"functionAppName"`
	SecretRef       *string          `json:"secretRef"`
	Functions       []*AzureFunction `json:"functions"`
}
type Destination interface{}
type DestinationSpec interface{}
type GRPCServiceSpec struct {
	Empty *string `json:"empty"`
}
type InputAwsDestinationSpec struct {
	LogicalName     string                   `json:"logicalName"`
	InvocationStyle AwsLambdaInvocationStyle `json:"invocationStyle"`
}
type InputAwsLambdaFunction struct {
	FunctionName string `json:"functionName"`
	Qualifier    string `json:"qualifier"`
}
type InputAwsUpstreamSpec struct {
	Region    string                    `json:"region"`
	SecretRef string                    `json:"secretRef"`
	Functions []*InputAwsLambdaFunction `json:"functions"`
}
type InputAzureDestinationSpec struct {
	FunctionName string `json:"functionName"`
}
type InputAzureFunction struct {
	FunctionName string `json:"functionName"`
	AuthLevel    string `json:"authLevel"`
}
type InputAzureUpstreamSpec struct {
	FunctionAppName string                `json:"functionAppName"`
	SecretRef       *string               `json:"secretRef"`
	Functions       []*InputAzureFunction `json:"functions"`
}
type InputDestination struct {
	SingleDestination *InputSingleDestination `json:"singleDestination"`
	MultiDestination  *InputMultiDestination  `json:"multiDestination"`
}
type InputDestinationSpec struct {
	Aws   *InputAwsDestinationSpec   `json:"aws"`
	Azure *InputAzureDestinationSpec `json:"azure"`
}
type InputGRPCServiceSpec struct {
	Empty *string `json:"empty"`
}
type InputKeyValueMatcher struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	IsRegex bool   `json:"isRegex"`
}
type InputKubeUpstreamSpec struct {
	ServiceName      string                       `json:"serviceName"`
	ServiceNamespace string                       `json:"serviceNamespace"`
	ServicePort      int                          `json:"servicePort"`
	Selector         *customtypes.MapStringString `json:"selector"`
}
type InputMatcher struct {
	PathMatch       string                  `json:"pathMatch"`
	PathMatchType   PathMatchType           `json:"pathMatchType"`
	Headers         []*InputKeyValueMatcher `json:"headers"`
	QueryParameters []*InputKeyValueMatcher `json:"queryParameters"`
	Methods         []*string               `json:"methods"`
}
type InputMetadata struct {
	Name            string                       `json:"name"`
	Namespace       string                       `json:"namespace"`
	ResourceVersion string                       `json:"resourceVersion"`
	Labels          *customtypes.MapStringString `json:"labels"`
	Annotations     *customtypes.MapStringString `json:"annotations"`
}
type InputMultiDestination struct {
	Destinations []*InputWeightedDestination `json:"destinations"`
}
type InputRoute struct {
	Matcher     InputMatcher     `json:"matcher"`
	Destination InputDestination `json:"destination"`
}
type InputServiceSpec struct {
	Swagger *InputSwaggerServiceSpec `json:"swagger"`
	Grpc    *InputGRPCServiceSpec    `json:"grpc"`
}
type InputSingleDestination struct {
	UpstreamName    string                `json:"upstreamName"`
	DestinationSpec *InputDestinationSpec `json:"destinationSpec"`
}
type InputSslConfig struct {
	SecretRef string `json:"secretRef"`
}
type InputStatus struct {
	State  State   `json:"state"`
	Reason *string `json:"reason"`
}
type InputSwaggerServiceSpec struct {
	Empty *string `json:"empty"`
}
type InputUpstream struct {
	Spec     InputUpstreamSpec `json:"spec"`
	Metadata InputMetadata     `json:"metadata"`
}
type InputUpstreamSpec struct {
	Aws   *InputAwsUpstreamSpec   `json:"aws"`
	Azure *InputAzureUpstreamSpec `json:"azure"`
	Kube  *InputKubeUpstreamSpec  `json:"kube"`
}
type InputVirtualService struct {
	Domains   []*string                   `json:"domains"`
	Routes    []*InputRoute               `json:"routes"`
	SslConfig *InputSslConfig             `json:"sslConfig"`
	Plugins   *InputVirtualServicePlugins `json:"plugins"`
	Metadata  *InputMetadata              `json:"metadata"`
}
type InputVirtualServicePlugins struct {
	Empty *string `json:"empty"`
}
type InputWeightedDestination struct {
	Destination InputDestination `json:"destination"`
	Weight      int              `json:"weight"`
}
type KeyValueMatcher struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	IsRegex bool   `json:"isRegex"`
}
type KubeUpstreamSpec struct {
	ServiceName      string                       `json:"serviceName"`
	ServiceNamespace string                       `json:"serviceNamespace"`
	ServicePort      int                          `json:"servicePort"`
	Selector         *customtypes.MapStringString `json:"selector"`
}
type Matcher struct {
	PathMatch       string             `json:"pathMatch"`
	PathMatchType   PathMatchType      `json:"pathMatchType"`
	Headers         []*KeyValueMatcher `json:"headers"`
	QueryParameters []*KeyValueMatcher `json:"queryParameters"`
	Methods         []*string          `json:"methods"`
}
type Metadata struct {
	Name            string                       `json:"name"`
	Namespace       string                       `json:"namespace"`
	ResourceVersion string                       `json:"resourceVersion"`
	Labels          *customtypes.MapStringString `json:"labels"`
	Annotations     *customtypes.MapStringString `json:"annotations"`
}
type MultiDestination struct {
	Destinations []*WeightedDestination `json:"destinations"`
}
type Route struct {
	Matcher     Matcher     `json:"matcher"`
	Destination Destination `json:"destination"`
}
type ServiceSpec interface{}
type SingleDestination struct {
	UpstreamName    string          `json:"upstreamName"`
	DestinationSpec DestinationSpec `json:"destinationSpec"`
}
type SslConfig struct {
	SecretRef string `json:"secretRef"`
}
type Status struct {
	State  State   `json:"state"`
	Reason *string `json:"reason"`
}
type SwaggerServiceSpec struct {
	Empty *string `json:"empty"`
}
type Upstream struct {
	Spec     UpstreamSpec `json:"spec"`
	Metadata Metadata     `json:"metadata"`
	Status   Status       `json:"status"`
}
type UpstreamMutation struct {
	Create *Upstream `json:"create"`
	Update *Upstream `json:"update"`
	Delete *Upstream `json:"delete"`
}
type UpstreamQuery struct {
	List []*Upstream `json:"list"`
	Get  *Upstream   `json:"get"`
}
type UpstreamSpec interface{}
type VirtualService struct {
	Domains   []*string              `json:"domains"`
	Routes    []*Route               `json:"routes"`
	SslConfig *SslConfig             `json:"sslConfig"`
	Plugins   *VirtualServicePlugins `json:"plugins"`
	Metadata  Metadata               `json:"metadata"`
	Status    Status                 `json:"status"`
}
type VirtualServiceMutation struct {
	Create *VirtualService `json:"create"`
	Update *VirtualService `json:"update"`
	Delete *VirtualService `json:"delete"`
}
type VirtualServicePlugins struct {
	Empty *string `json:"empty"`
}
type VirtualServiceQuery struct {
	List []*VirtualService `json:"list"`
	Get  *VirtualService   `json:"get"`
}
type WeightedDestination struct {
	Destination SingleDestination `json:"destination"`
	Weight      int               `json:"weight"`
}

type AwsLambdaInvocationStyle string

const (
	AwsLambdaInvocationStyleSync  AwsLambdaInvocationStyle = "SYNC"
	AwsLambdaInvocationStyleAsync AwsLambdaInvocationStyle = "ASYNC"
)

func (e AwsLambdaInvocationStyle) IsValid() bool {
	switch e {
	case AwsLambdaInvocationStyleSync, AwsLambdaInvocationStyleAsync:
		return true
	}
	return false
}

func (e AwsLambdaInvocationStyle) String() string {
	return string(e)
}

func (e *AwsLambdaInvocationStyle) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AwsLambdaInvocationStyle(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AwsLambdaInvocationStyle", str)
	}
	return nil
}

func (e AwsLambdaInvocationStyle) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PathMatchType string

const (
	PathMatchTypePrefix PathMatchType = "PREFIX"
	PathMatchTypeRegex  PathMatchType = "REGEX"
	PathMatchTypeExact  PathMatchType = "EXACT"
)

func (e PathMatchType) IsValid() bool {
	switch e {
	case PathMatchTypePrefix, PathMatchTypeRegex, PathMatchTypeExact:
		return true
	}
	return false
}

func (e PathMatchType) String() string {
	return string(e)
}

func (e *PathMatchType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PathMatchType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PathMatchType", str)
	}
	return nil
}

func (e PathMatchType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type State string

const (
	StatePending  State = "PENDING"
	StateAccepted State = "ACCEPTED"
	StateRejected State = "REJECTED"
)

func (e State) IsValid() bool {
	switch e {
	case StatePending, StateAccepted, StateRejected:
		return true
	}
	return false
}

func (e State) String() string {
	return string(e)
}

func (e *State) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = State(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid State", str)
	}
	return nil
}

func (e State) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
