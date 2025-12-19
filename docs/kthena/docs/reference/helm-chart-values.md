# kthena

A Helm chart for deploying Kthena

![Version: 1.0.0](https://img.shields.io/badge/Version-1.0.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Requirements

| Repository | Name | Version |
|------------|------|---------|
|  | networking | 1.0.0 |
|  | workload | 1.0.0 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| global.certManagementMode | string | `"auto"` | Certificate Management Mode. Three mutually exclusive options for managing TLS certificates:   - auto: Webhook servers generate self-signed certificates automatically (default)   - cert-manager: Use cert-manager to generate and manage certificates (requires cert-manager installation)   - manual: Provide your own certificates via caBundle |
| global.webhook.caBundle | string | `""` | caBundle is the base64-encoded CA bundle for webhook server certificates. This is ONLY required when certManagementMode is set to "manual". You can generate it with: cat /path/to/your/ca.crt | base64 | tr -d '\n' |
| networking.enabled | bool | `true` | enabled is a flag to enable or disable the networking subchart. Default is true. |
| networking.kthenaRouter.enabled | bool | `true` | Enable Kthena Router |
| networking.kthenaRouter.fairness.enabled | bool | `false` | enabled controls whether fairness scheduling is active |
| networking.kthenaRouter.fairness.inputTokenWeight | float | `1` | inputTokenWeight is the weight multiplier for input tokens |
| networking.kthenaRouter.fairness.outputTokenWeight | float | `2` | outputTokenWeight is the weight multiplier for output tokens |
| networking.kthenaRouter.fairness.windowSize | string | `"1h"` | windowSize is the sliding window duration for token usage tracking |
| networking.kthenaRouter.gatewayAPI.enabled | bool | `false` | enabled controls whether Gateway API related features are enabled |
| networking.kthenaRouter.gatewayAPI.inferenceExtension | bool | `false` | inferenceExtension controls whether Gateway API Inference Extension features are enabled. This requires gatewayAPI.enabled to be true |
| networking.kthenaRouter.image.pullPolicy | string | `"IfNotPresent"` | Image pull policy for Kthena Router |
| networking.kthenaRouter.image.repository | string | `"ghcr.io/volcano-sh/kthena-router"` | Image repository for Kthena Router |
| networking.kthenaRouter.image.tag | string | `"latest"` | Image tag for Kthena Router |
| networking.kthenaRouter.port | int | `8080` | Port for Kthena Router |
| networking.kthenaRouter.tls.dnsName | string | `"your-domain.com"` | The DNS name to use for the certificate. |
| networking.kthenaRouter.tls.enabled | bool | `false` | Enable TLS for Kthena Router |
| networking.kthenaRouter.tls.secretName | string | `"kthena-router-tls"` | The name of the secret to store the certificate and key. |
| networking.kthenaRouter.webhook.enabled | bool | `true` | Enable webhook for Kthena Router |
| networking.kthenaRouter.webhook.port | int | `8443` | Port for Kthena Router webhook |
| networking.kthenaRouter.webhook.servicePort | int | `443` | Service port for Kthena Router webhook |
| networking.kthenaRouter.webhook.tls.certFile | string | `"/etc/tls/tls.crt"` | Certificate file path |
| networking.kthenaRouter.webhook.tls.keyFile | string | `"/etc/tls/tls.key"` | Key file path |
| networking.kthenaRouter.webhook.tls.secretName | string | `"kthena-router-webhook-certs"` | Secret name for webhook certificates |
| workload.controllerManager.downloaderImage.repository | string | `"ghcr.io/volcano-sh/downloader"` | Image repository for the downloader |
| workload.controllerManager.downloaderImage.tag | string | `"latest"` | Image tag for the downloader |
| workload.controllerManager.image.pullPolicy | string | `"IfNotPresent"` | Image pull policy for the controller manager |
| workload.controllerManager.image.repository | string | `"ghcr.io/volcano-sh/kthena-controller-manager"` | Image repository for the controller manager |
| workload.controllerManager.image.tag | string | `"latest"` | Image tag for the controller manager |
| workload.controllerManager.runtimeImage.repository | string | `"ghcr.io/volcano-sh/runtime"` | Image repository for the runtime |
| workload.controllerManager.runtimeImage.tag | string | `"latest"` | Image tag for the runtime |
| workload.controllerManager.webhook.enabled | bool | `true` | Enable webhook for controller manager |
| workload.controllerManager.webhook.tls.certSecretName | string | `"kthena-controller-manager-webhook-certs"` | Secret name for webhook certificates |
| workload.controllerManager.webhook.tls.serviceName | string | `"kthena-controller-manager-webhook"` | Service name for webhook |
| workload.enabled | bool | `true` | enabled is a flag to enable or disable the workload subchart. Default is true. |

## Notes

- Values marked as “usually set by CI” are automatically updated during the release process; manual changes are not required.
- For detailed information about each component, refer to the corresponding architecture and user guide documents.
- Always review the [values.yaml](https://github.com/volcano-sh/kthena/blob/main/charts/kthena/values.yaml) file in the repository for the latest defaults and available options.