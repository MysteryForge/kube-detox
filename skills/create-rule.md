---
name: create-rule
description: >
  Processes natural language requests to generate K8s compliance rules.
  Extracts resource kind, field path, and conditions to draft valid YAML
  rules placed in the rules/k8s directory structure.
---

# Skill: Rule Creation

Use this process to add new compliance rules from natural language requests:

1. **Identify the Goal**: Extract the Kubernetes resource `kind`, the field path, and desired outcome.
2. **Determine Category**: 
   - If user specifies a category (e.g., "security", "reliability"), place in `rules/k8s/<category>/`.
   - If no category is specified, prompt or default to `rules/custom/<custom_name>/`.
3. **Draft YAML**:
   - `id`: Unique identifier (e.g., `cust-001`, `pod-001`).
   - `target`: The `kind` of the resource.
   - `field`: The GJSON-compatible path.
   - `severity`: Default to `warn` unless explicitly stated as `block`.
4. **Execution**: Create the file in the determined location and confirm creation.
