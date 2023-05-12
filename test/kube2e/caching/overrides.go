package cachinggrpc

const overrideYaml = `
gloo:
  gatewayProxies:
    gatewayProxy:
      healthyPanicThreshold: 0
  gateway:
    persistProxySpec: true
  rbac:
    namespaced: true
    nameSuffix: e2e-test-rbac-suffix
settings:
  singleNamespace: true
  create: true
global:
  extensions:
    caching:
      enabled: true
gloo-fed:
  enabled: false
  glooFedApiserver:
    enable: false
redis:
  deployment:
    replicas: 1
`