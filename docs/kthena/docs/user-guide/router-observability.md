# Router Observability

## Overview

Kthena provides comprehensive observability features for monitoring and debugging the router component, which serves as the data plane entry point for inference traffic. This documentation implements the comprehensive observability framework outlined in the Router Observability Proposal, providing detailed metrics, access logs, and debug interfaces for effective AI workload management.

## Metrics Specification

### Port Configuration

- **Metrics Endpoint**: Port `15000` at `/metrics`

### Request Processing Metrics

| Metric Name | Type | Description | Labels | Buckets |
|-------------|------|-------------|---------|---------|
| `kthena_router_requests_total` | Counter | Total HTTP requests processed | `model`, `path`, `status_code`, `error_type` | - |
| `kthena_router_request_duration_seconds` | Histogram | End-to-end request latency | `model`, `path`, `status_code` | [0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10, 30, 60] |
| `kthena_router_request_prefill_duration_seconds` | Histogram | Prefill phase latency (PD-disaggregated) | `model`, `path`, `status_code` | [0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10, 30, 60] |
| `kthena_router_request_decode_duration_seconds` | Histogram | Decode phase latency (PD-disaggregated) | `model`, `path`, `status_code` | [0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10, 30, 60] |
| `kthena_router_active_downstream_requests` | Gauge | Active downstream requests | `model` | - |
| `kthena_router_active_upstream_requests` | Gauge | Active upstream requests | `model_route`, `model_server` | - |

### AI-Specific Token Metrics

| Metric Name | Type | Description | Labels |
|-------------|------|-------------|---------|
| `kthena_router_tokens_total` | Counter | Total tokens processed/generated | `model`, `path`, `token_type` |

### Scheduler Plugin Metrics

| Metric Name | Type | Description | Labels | Buckets |
|-------------|------|-------------|---------|---------|
| `kthena_router_scheduler_plugin_duration_seconds` | Histogram | Plugin processing time | `model`, `plugin`, `type` | [0.001, 0.005, 0.01, 0.05, 0.1, 0.5] |

### Rate Limiting Metrics

| Metric Name | Type | Description | Labels |
|-------------|------|-------------|---------|
| `kthena_router_rate_limit_exceeded_total` | Counter | Requests rejected due to rate limiting | `model`, `limit_type`, `path` |

### Fairness Queue Metrics

| Metric Name | Type | Description | Labels | Buckets |
|-------------|------|-------------|---------|---------|
| `kthena_router_fairness_queue_size` | Gauge | Current fairness queue size | `model`, `user_id` | - |
| `kthena_router_fairness_queue_duration_seconds` | Histogram | Time in fairness queue | `model`, `user_id` | [0.001, 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5] |

## Prometheus Configuration

```yaml
global:
  scrape_interval: 15s
  evaluation_interval: 15s

scrape_configs:
  - job_name: 'kthena-router'
    static_configs:
      - targets: ['kthena-router-service:15000']
        labels:
          component: 'router'
```

## Access Log Format

### JSON Format (Default)

```json
{
  "timestamp": "2024-01-15T10:30:45.123Z",
  "method": "POST",
  "path": "/v1/chat/completions",
  "protocol": "HTTP/1.1",
  "status_code": 200,
  "model_name": "llama2-7b",
  "model_route": "default/llama2-route-v1",
  "model_server": "default/llama2-server",
  "selected_pod": "llama2-deployment-5f7b8c9d-xk2p4",
  "request_id": "550e8400-e29b-41d4-a716-446655440000",
  "input_tokens": 150,
  "output_tokens": 75,
  "duration_total": 2350,
  "duration_request_processing": 45,
  "duration_upstream_processing": 2180,
  "duration_response_processing": 5,
  "error": {
    "type": "timeout",
    "message": "Model inference timeout after 30s"
  }
}
```

### Text Format (Alternative)

`[2024-01-15T10:30:45.123Z] "POST /v1/chat/completions HTTP/1.1" 200 model_name=llama2-7b model_route=default/llama2-route-v1 model_server=default/llama2-server selected_pod=llama2-deployment-5f7b8c9d-xk2p4 request_id=550e8400-e29b-41d4-a716-446655440000 tokens=150/75 timings=2350ms(45+2180+5)`

## Debug Interface

### Debug Endpoints

| Endpoint | Description |
|-------------|----------|
| `/debug/config_dump/modelroutes` | List all ModelRoute configurations |
| `/debug/config_dump/modelservers` | List all ModelServer configurations |
| `/debug/config_dump/pods` | List all Pod information |
| `/debug/config_dump/namespaces/{namespace}/modelroutes/{name}` | Specific ModelRoute details |
| `/debug/config_dump/namespaces/{namespace}/modelservers/{name}` | Specific ModelServer details |
| `/debug/config_dump/namespaces/{namespace}/pods/{name}` | Specific Pod details |

## Configuration

### Access Log Configuration

```yaml
accessLogger:
  format: "json"  # Options: "json", "text"
  output: "stdout"  # Options: "stdout", "stderr", or file path
  enabled: true
```

Metrics Configuration

```yaml
observability:
  metrics:
    enabled: true
    port: 15000 
    path: /metrics
```

## Grafana Dashboard

Based on the proposal metrics, create dashboard panels for:

1. Request Rate by Model: rate(kthena_router_requests_total[5m])
2. Token Processing Rate: rate(kthena_router_tokens_total[5m])
3. Request Latency: histogram_quantile(0.95, kthena_router_request_duration_seconds)
4. Rate Limit Violations: rate(kthena_router_rate_limit_exceeded_total[5m])

## Alerting Rules

```yaml
groups:
- name: kthena_router_alerts
  rules:
  - alert: HighRouterErrorRate
    expr: rate(kthena_router_requests_total{status_code=~"5.."}[5m]) > 0.1
    for: 2m
    labels:
      severity: critical
    annotations:
      summary: "High error rate in Kthena router"
      description: "Router error rate is {{ $value | humanizePercentage }} for model {{ $labels.model }}"

  - alert: HighRequestLatency
    expr: histogram_quantile(0.95, kthena_router_request_duration_seconds) > 5
    for: 5m
    labels:
      severity: warning
    annotations:
      summary: "High request latency detected"
      description: "95th percentile latency is {{ $value }}s for model {{ $labels.model }}"
```

## Troubleshooting

### Check Metrics Endpoint

```bash
# Port forward to access metrics
kubectl port-forward -n kthena-system svc/kthena-router 15000:15000

# View actual metrics
curl http://localhost:15000/metrics
```

### Check Debug Endpoints

```bash
# Test debug endpoints (if implemented)
curl http://localhost:15000/debug/config_dump/modelroutes
curl http://localhost:15000/debug/config_dump/modelservers
```
