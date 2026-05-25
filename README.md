# Kube Detox

Enterprise-grade audit tool for Kubernetes & OpenShift manifests. Validates YAML and Helm templates against security, reliability, and governance standards.

## Installation

```bash
go build -o auditor main.go
```

## Quick Start

### 1. Audit Helm Charts
Pipe your Helm templates into the auditor:
```bash
helm template my-chart | ./auditor scan --policy rules/k8s/ --severity block
```

### 2. Audit Local Files
Scan a specific manifest directory:
```bash
./auditor scan --policy rules/k8s/security manifests/deployment.yaml
```

## Policy Management

### Severity Levels
- `--severity block`: Report only mandatory security blockers (e.g., Privileged access).
- `--severity warn`: Report all blockers AND advisory best-practices (e.g., Resource limits).

### Modular Rules
Create custom rules in `rules/custom/<category>/` by copying an existing YAML rule. The auditor automatically discovers and applies all YAML files found under the `--policy` path.

## Example Rule (rules/k8s/security/hardening.yaml)
```yaml
rules:
  - id: sec-001
    name: "Disallow Privileged"
    description: "Containers must not be privileged."
    target: "Deployment"
    field: "spec.template.spec.containers.0.securityContext.privileged"
    expected: false
    severity: "block"
```

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Support
This project is supported by [CuteTarantula](https://cutetarantula.com).

We are a UK-based software consultancy specializing in high-performance software engineering, infrastructure, and scalable systems. We design, build, and scale solutions across industries including Blockchain, AI, Fintech, AdTech and eCommerce.
Whether you need expert guidance on Kubernetes, decentralized infrastructure, or complex software challenges, we’re here to help. Reach out to collaborate on systems that endure.
