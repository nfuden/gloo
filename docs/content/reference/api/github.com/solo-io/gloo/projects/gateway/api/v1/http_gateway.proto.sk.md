
---
title: "http_gateway.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `gateway.solo.io` 
#### Types:


- [HttpGateway](#httpgateway)
- [VirtualServiceSelectorExpressions](#virtualserviceselectorexpressions)
- [Expression](#expression)
- [Operator](#operator)
  



##### Source File: [github.com/solo-io/gloo/projects/gateway/api/v1/http_gateway.proto](https://github.com/solo-io/gloo/blob/main/projects/gateway/api/v1/http_gateway.proto)





---
### HttpGateway



```yaml
"virtualServices": []core.solo.io.ResourceRef
"virtualServiceSelector": map<string, string>
"virtualServiceExpressions": .gateway.solo.io.VirtualServiceSelectorExpressions
"virtualServiceNamespaces": []string
"options": .gloo.solo.io.HttpListenerOptions

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `virtualServices` | [[]core.solo.io.ResourceRef](../../../../../../solo-kit/api/v1/ref.proto.sk/#resourceref) | Names & namespace refs of the virtual services which contain the actual routes for the gateway. If the list is empty, all virtual services in all namespaces that Gloo watches will apply, with accordance to `ssl` flag on `Gateway` above. The default namespace matching behavior can be overridden via `virtual_service_namespaces` flag below. Only one of `virtualServices`, `virtualServiceExpressions` or `virtualServiceSelector` should be provided. If more than one is provided only one will be checked with priority virtualServiceExpressions, virtualServiceSelector, virtualServices. |
| `virtualServiceSelector` | `map<string, string>` | Select virtual services by their label. If `virtual_service_namespaces` is provided below, this will apply only to virtual services in the namespaces specified. Only one of `virtualServices`, `virtualServiceExpressions` or `virtualServiceSelector` should be provided. If more than one is provided only one will be checked with priority virtualServiceExpressions, virtualServiceSelector, virtualServices. |
| `virtualServiceExpressions` | [.gateway.solo.io.VirtualServiceSelectorExpressions](../http_gateway.proto.sk/#virtualserviceselectorexpressions) | Select virtual services using expressions. If `virtual_service_namespaces` is provided below, this will apply only to virtual services in the namespaces specified. Only one of `virtualServices`, `virtualServiceExpressions` or `virtualServiceSelector` should be provided. If more than one is provided only one will be checked with priority virtualServiceExpressions, virtualServiceSelector, virtualServices. |
| `virtualServiceNamespaces` | `[]string` | Restrict the search by providing a list of valid search namespaces here. Setting '*' will search all namespaces, equivalent to omitting this value. |
| `options` | [.gloo.solo.io.HttpListenerOptions](../../../../gloo/api/v1/options.proto.sk/#httplisteneroptions) | HTTP Gateway configuration. |




---
### VirtualServiceSelectorExpressions

 
Expressions to define which virtual services to select
Example:
expressions:
   - key: domain
     operator: in
     values: example.com

```yaml
"expressions": []gateway.solo.io.VirtualServiceSelectorExpressions.Expression

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `expressions` | [[]gateway.solo.io.VirtualServiceSelectorExpressions.Expression](../http_gateway.proto.sk/#expression) | Expressions allow for more flexible virtual service label matching, such as equality-based requirements, set-based requirements, or a combination of both. https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#equality-based-requirement. |




---
### Expression



```yaml
"key": string
"operator": .gateway.solo.io.VirtualServiceSelectorExpressions.Expression.Operator
"values": []string

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `key` | `string` | Kubernetes label key, must conform to Kubernetes syntax requirements https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#syntax-and-character-set. |
| `operator` | [.gateway.solo.io.VirtualServiceSelectorExpressions.Expression.Operator](../http_gateway.proto.sk/#operator) | The operator can only be in, notin, =, ==, !=, exists, ! (DoesNotExist), gt (GreaterThan), lt (LessThan). |
| `values` | `[]string` |  |




---
### Operator

 
Virtual Service Selector expression operator, while the set-based syntax differs from Kubernetes (kubernetes: `key: !mylabel`, gloo: `key: mylabel, operator: "!"` | kubernetes: `key: mylabel`, gloo: `key: mylabel, operator: exists`), the functionality remains the same.

| Name | Description |
| ----- | ----------- | 
| `Equals` | = |
| `DoubleEquals` | == |
| `NotEquals` | != |
| `In` | in |
| `NotIn` | notin |
| `Exists` | exists |
| `DoesNotExist` | ! |
| `GreaterThan` | gt |
| `LessThan` | lt |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->