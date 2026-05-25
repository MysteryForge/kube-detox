package auditor

import (
	"os"
	"testing"
)

func TestValidate(t *testing.T) {
	manifest, err := os.ReadFile("testdata/deployment.yaml")
	if err != nil {
		t.Fatalf("failed to read testdata: %v", err)
	}

	tests := []struct {
		name     string
		rule     Rule
		wantPass bool
	}{
		{
			name: "Privileged check - PASS",
			rule: Rule{
				ID:       "test-001",
				Field:    "spec.template.spec.containers.0.securityContext.privileged",
				Expected: false,
			},
			wantPass: true,
		},
		{
			name: "Read-Only FS check - PASS",
			rule: Rule{
				ID:       "test-002",
				Field:    "spec.template.spec.containers.0.securityContext.readOnlyRootFilesystem",
				Expected: true,
			},
			wantPass: true,
		},
		{
			name: "Privileged check - FAIL",
			rule: Rule{
				ID:       "test-003",
				Field:    "spec.template.spec.containers.0.securityContext.privileged",
				Expected: true,
			},
			wantPass: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPass, _ := Validate(manifest, tt.rule)
			if gotPass != tt.wantPass {
				t.Errorf("Validate() gotPass = %v, want %v", gotPass, tt.wantPass)
			}
		})
	}
}
