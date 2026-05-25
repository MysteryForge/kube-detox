# Kube Detox

Enterprise-grade audit tool for Kubernetes manifests. Validates YAML and Helm templates against security, reliability, and governance standards.

## Installation

### Download Pre-built Binaries
Download the latest binaries for Linux and macOS (amd64/arm64) from the [Releases page](https://github.com/MysteryForge/kube-detox/releases).

### Build from Source
```bash
go build -o kube-detox .
```

## Quick Start

### 1. Audit Helm Charts
Pipe your Helm templates into the binary:
```bash
helm template my-chart | ./kube-detox scan --policy rules/k8s/
```

### 2. Audit Local Files
Scan a specific manifest directory:
```bash
./kube-detox scan --policy rules/k8s/security manifests/deployment.yaml
```

## Policy Management

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
    
```

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Support
This project is supported by [CuteTarantula](https://cutetarantula.com).

We are a UK-based software consultancy specializing in high-performance software engineering, infrastructure, and scalable systems. We design, build, and scale solutions across industries including Blockchain, AI, Fintech, AdTech and eCommerce.
Whether you need expert guidance on Kubernetes, decentralized infrastructure, or complex software challenges, we’re here to help. Reach out to collaborate on systems that endure.
