package auditor

import (
	"fmt"
	"regexp"

	"github.com/tidwall/gjson"
	"sigs.k8s.io/yaml"
)

func Validate(manifest []byte, rule Rule) (bool, string) {
	jsonData, err := yaml.YAMLToJSON(manifest)
	if err != nil {
		return false, fmt.Sprintf("Error parsing YAML: %v", err)
	}

	// Check if target matches kind
	kind := gjson.GetBytes(jsonData, "kind").String()
	if rule.Target != "" && kind != rule.Target {
		return true, fmt.Sprintf("Rule %s skipped (Target %s != Kind %s)", rule.ID, rule.Target, kind)
	}

	result := gjson.GetBytes(jsonData, rule.Field)

	// Logic for different expected types
	if rule.Expected == "exists" {
		return result.Exists(), rule.Description
	}
	if rule.Expected == "not_null" {
		return result.Exists(), rule.Description
	}

	// For slices like capabilities: drop: ["ALL"]
	// Result.String() might be "[\"ALL\"]"
	if result.IsArray() {
		for _, item := range result.Array() {
			if item.String() == fmt.Sprintf("%v", rule.Expected) {
				return true, rule.Description
			}
		}
		return false, rule.Description
	}

	if rule.ExpectedRegex != "" {
		return regexp.MustCompile(rule.ExpectedRegex).MatchString(result.String()), rule.Description
	}

	// Default: direct value comparison
	return result.String() == fmt.Sprintf("%v", rule.Expected), rule.Description
}
