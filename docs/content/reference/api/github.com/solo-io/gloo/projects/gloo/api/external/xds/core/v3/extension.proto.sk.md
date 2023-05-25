
---
title: "extension.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `xds.core.v3` 
#### Types:


- [TypedExtensionConfig](#typedextensionconfig)
  



##### Source File: [github.com/solo-io/gloo/projects/gloo/api/external/xds/core/v3/extension.proto](https://github.com/solo-io/gloo/blob/master/projects/gloo/api/external/xds/core/v3/extension.proto)





---
### TypedExtensionConfig

 
Message type for extension configuration.

```yaml
"name": string
"typedConfig": .google.protobuf.Any

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `name` | `string` | The name of an extension. This is not used to select the extension, instead it serves the role of an opaque identifier. |
| `typedConfig` | [.google.protobuf.Any](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/any) | The typed config for the extension. The type URL will be used to identify the extension. In the case that the type URL is *xds.type.v3.TypedStruct* (or, for historical reasons, *udpa.type.v1.TypedStruct*), the inner type URL of *TypedStruct* will be utilized. See the :ref:`extension configuration overview <config_overview_extension_configuration>` for further details. |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->