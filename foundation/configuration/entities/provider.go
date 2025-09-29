package entities

//──────────────────────────────────────────────────────────────────────────────────────────────────

type Provider struct {
	AI ProviderAI `yaml:"ai"`
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

type ProviderAI struct {
	Connection     string `yaml:"connection"`
	RequestTimeout int64  `yaml:"request-timeout"`
}
