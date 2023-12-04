package reports

// Report 报告

type Validator struct {
	Check   string      `json:"check" yaml:"check"` // get value with jmespath
	Assert  string      `json:"assert" yaml:"assert"`
	Expect  interface{} `json:"expect" yaml:"expect"`
	Message string      `json:"msg,omitempty" yaml:"msg,omitempty"` // optional
}

type ValidationResult struct {
	Validator
	CheckValue  interface{} `json:"check_value" yaml:"check_value"`
	CheckResult string      `json:"check_result" yaml:"check_result"`
}
