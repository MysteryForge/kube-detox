package auditor

type Rule struct {
	ID            string `yaml:"id"`
	Name          string `yaml:"name"`
	Description   string `yaml:"description"`
	Target        string `yaml:"target"`
	Field         string `yaml:"field"`
	Expected      any    `yaml:"expected"`
	ExpectedRegex string `yaml:"expected_regex"`
	Severity      string `yaml:"severity"`
}

type Ruleset struct {
	Rules []Rule `yaml:"rules"`
}
