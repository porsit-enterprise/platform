package entities

//──────────────────────────────────────────────────────────────────────────────────────────────────

type Provider struct {
	AI  ProviderAI  `yaml:"ai"`
	SMS ProviderSMS `yaml:"sms"`
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

type ProviderAI struct {
	Connection     string `yaml:"connection"`
	RequestTimeout int64  `yaml:"request-timeout"`
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

type ProviderSMS struct {
	Kavenegar Kavenegar `yaml:"kavenegar"`
}

type Kavenegar struct {
	ApiKey      string `yaml:"apikey"`
	OtpTemplate string `yaml:"otp-template"`
}
