package models

type Config struct {
	BaseUrl  string `yaml:"base_url" json:"base_url"`
	Headless bool   `yaml:"headless" json:"headless"`
	Video    Video  `yaml:"video" json:"video"`
	Trace    Trace  `yaml:"trace" json:"trace"`
	Report   Report `yaml:"report" json:"report"`
	Browser  string `yaml:"browser" json:"browser"`
}

type Video struct {
	Method string `yaml:"method" json:"method"`
	Path   string `yaml:"path" json:"path"`
}

type Trace struct {
	Method string `yaml:"method" json:"method"`
	Path   string `yaml:"path" json:"path"`
}

type Report struct {
	Path string `yaml:"path" json:"path"`
}
