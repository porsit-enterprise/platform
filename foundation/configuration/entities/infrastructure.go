package entities

type Infrastructure struct {
	PostgreSQL PostgreSQL `yaml:"postgresql"`
	Valkey     Valkey     `yaml:"valkey"`
	Ollama     Ollama     `yaml:"ollama"`
	Centrifugo Centrifugo `yaml:"centrifugo"`
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

type PostgreSQL struct {
	Connection           string `yaml:"connection"`
	ConnectionTimeout    uint8  `yaml:"connection-timeout"`
	ConnectionRetries    uint8  `yaml:"connection-retries"`
	ConnectionRetryDelay uint8  `yaml:"connection-retry-delay"`
}

type Valkey struct {
	Connection           string `yaml:"connection"`
	Username             string `yaml:"username"`
	Password             string `yaml:"password"`
	ConnectionTimeout    uint8  `yaml:"connection-timeout"`
	ConnectionRetries    uint8  `yaml:"connection-retries"`
	ConnectionRetryDelay uint8  `yaml:"connection-retry-delay"`
}

type Ollama struct {
	Connection           string `yaml:"connection"`
	ConnectionTimeout    uint8  `yaml:"connection-timeout"`
	OperationTimeout     uint8  `yaml:"operation-timeout"`
	ConnectionRetries    uint8  `yaml:"connection-retries"`
	ConnectionRetryDelay uint8  `yaml:"connection-retry-delay"`
}

type Centrifugo struct {
	ConnectionGRPC    string `yaml:"connection-grpc"`
	ConnectionHTTP    string `yaml:"connection-http"`
	ApiKey            string `yaml:"api-key"`
	ConnectionTimeout uint8  `yaml:"connection-timeout"`
}
