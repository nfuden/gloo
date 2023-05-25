
---
title: "transformation.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `envoy.config.filter.http.transformation_ee.v2` 
#### Types:


- [FilterTransformations](#filtertransformations)
- [TransformationRule](#transformationrule)
- [RouteTransformations](#routetransformations)
- [Transformation](#transformation)
- [DlpTransformation](#dlptransformation)
- [Action](#action)
- [RegexMatcher](#regexmatcher)
- [KeyValueMatcher](#keyvaluematcher)
- [DlpMatcher](#dlpmatcher)
- [RegexAction](#regexaction)
  



##### Source File: [github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/transformation_ee/transformation.proto](https://github.com/solo-io/gloo/blob/master/projects/gloo/api/external/envoy/extensions/transformation_ee/transformation.proto)





---
### FilterTransformations



```yaml
"transformations": []envoy.config.filter.http.transformation_ee.v2.TransformationRule

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `transformations` | [[]envoy.config.filter.http.transformation_ee.v2.TransformationRule](../transformation.proto.sk/#transformationrule) | Specifies transformations based on the route matches. The first matched transformation will be applied. If there are overlapped match conditions, please put the most specific match first. |




---
### TransformationRule



```yaml
"match": .solo.io.envoy.api.v2.route.RouteMatch
"matchV3": .solo.io.envoy.config.route.v3.RouteMatch
"routeTransformations": .envoy.config.filter.http.transformation_ee.v2.RouteTransformations

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `match` | [.solo.io.envoy.api.v2.route.RouteMatch](../../../../../../../../../../../envoy/api/v2/route/route.proto.sk/#routematch) | The route matching parameter. Only when the match is satisfied, the "requires" field will apply. For example: following match will match all requests. .. code-block:: yaml match: prefix: /. |
| `matchV3` | [.solo.io.envoy.config.route.v3.RouteMatch](../../../config/route/v3/route_components.proto.sk/#routematch) |  |
| `routeTransformations` | [.envoy.config.filter.http.transformation_ee.v2.RouteTransformations](../transformation.proto.sk/#routetransformations) | transformation to perform. |




---
### RouteTransformations



```yaml
"requestTransformation": .envoy.config.filter.http.transformation_ee.v2.Transformation
"clearRouteCache": bool
"responseTransformation": .envoy.config.filter.http.transformation_ee.v2.Transformation
"onStreamCompletionTransformation": .envoy.config.filter.http.transformation_ee.v2.Transformation

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `requestTransformation` | [.envoy.config.filter.http.transformation_ee.v2.Transformation](../transformation.proto.sk/#transformation) |  |
| `clearRouteCache` | `bool` | clear the route cache if the request transformation was applied. |
| `responseTransformation` | [.envoy.config.filter.http.transformation_ee.v2.Transformation](../transformation.proto.sk/#transformation) |  |
| `onStreamCompletionTransformation` | [.envoy.config.filter.http.transformation_ee.v2.Transformation](../transformation.proto.sk/#transformation) | Apply a transformation in the onStreamComplete callback (for modifying headers and dynamic metadata for access logs). |




---
### Transformation



```yaml
"dlpTransformation": .envoy.config.filter.http.transformation_ee.v2.DlpTransformation

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `dlpTransformation` | [.envoy.config.filter.http.transformation_ee.v2.DlpTransformation](../transformation.proto.sk/#dlptransformation) |  |




---
### DlpTransformation



```yaml
"actions": []envoy.config.filter.http.transformation_ee.v2.Action
"enableHeaderTransformation": bool
"enableDynamicMetadataTransformation": bool

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `actions` | [[]envoy.config.filter.http.transformation_ee.v2.Action](../transformation.proto.sk/#action) | list of actions to apply. |
| `enableHeaderTransformation` | `bool` | If true, headers will be transformed. Should only be true for the on_stream_complete_transformation route transformation type. |
| `enableDynamicMetadataTransformation` | `bool` | If true, dynamic metadata will be transformed. Should only be used for the on_stream_complete_transformation route transformation type. |




---
### Action



```yaml
"name": string
"regex": []string
"regexActions": []envoy.config.filter.http.transformation_ee.v2.RegexAction
"shadow": bool
"percent": .solo.io.envoy.type.Percent
"maskChar": string
"matcher": .envoy.config.filter.http.transformation_ee.v2.Action.DlpMatcher

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `name` | `string` | Identifier for this action. Used mostly to help ID specific actions in logs. If left null will default to unknown. |
| `regex` | `[]string` | Deprecated in favor of DlpMatcher List of regexes to apply to the response body to match data which should be masked They will be applied iteratively in the order which they are specified. |
| `regexActions` | [[]envoy.config.filter.http.transformation_ee.v2.RegexAction](../transformation.proto.sk/#regexaction) | Deprecated in favor of DlpMatcher List of regexes to apply to the response body to match data which should be masked. They will be applied iteratively in the order which they are specified. If this field and `regex` are both provided, all the regexes will be applied iteratively in the order provided, starting with the ones from `regex`. |
| `shadow` | `bool` | If specified, this rule will not actually be applied, but only logged. |
| `percent` | [.solo.io.envoy.type.Percent](../../../../../../../../../solo-kit/api/external/envoy/type/percent.proto.sk/#percent) | The percent of the string which should be masked. If not set, defaults to 75%. |
| `maskChar` | `string` | The character which should overwrite the masked data If left empty, defaults to "X". |
| `matcher` | [.envoy.config.filter.http.transformation_ee.v2.Action.DlpMatcher](../transformation.proto.sk/#dlpmatcher) | The matcher used to determine which values will be masked by this action. |




---
### RegexMatcher

 
List of regexes to apply to the response body to match data which should be
masked. They will be applied iteratively in the order which they are
specified.

```yaml
"regexActions": []envoy.config.filter.http.transformation_ee.v2.RegexAction

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `regexActions` | [[]envoy.config.filter.http.transformation_ee.v2.RegexAction](../transformation.proto.sk/#regexaction) |  |




---
### KeyValueMatcher

 
List of headers for which associated values will be masked.
Note that enable_header_transformation must be set for this to take effect.
Note that if enable_dynamic_metadata_transformation is set, proto struct dynamic metadata
(i.e., the values matching any JSON keys specified in `keys`; primarily for json-formatted WAF audit logs) will also be masked accordingly.

```yaml
"keys": []string

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `keys` | `[]string` |  |




---
### DlpMatcher



```yaml
"regexMatcher": .envoy.config.filter.http.transformation_ee.v2.Action.RegexMatcher
"keyValueMatcher": .envoy.config.filter.http.transformation_ee.v2.Action.KeyValueMatcher

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `regexMatcher` | [.envoy.config.filter.http.transformation_ee.v2.Action.RegexMatcher](../transformation.proto.sk/#regexmatcher) |  Only one of `regexMatcher` or `keyValueMatcher` can be set. |
| `keyValueMatcher` | [.envoy.config.filter.http.transformation_ee.v2.Action.KeyValueMatcher](../transformation.proto.sk/#keyvaluematcher) |  Only one of `keyValueMatcher` or `regexMatcher` can be set. |




---
### RegexAction



```yaml
"regex": string
"subgroup": int

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `regex` | `string` | The regex to match for masking. |
| `subgroup` | `int` | If provided and not 0, only this specific subgroup of the regex will be masked. |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->