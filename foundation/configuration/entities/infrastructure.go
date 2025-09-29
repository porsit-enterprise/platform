package entities

type Infrastructure struct {
	PostgreSQL PostgreSQL `yaml:"postgresql"`
	Valkey     Valkey     `yaml:"valkey"`
	Ollama     Ollama     `yaml:"ollama"`
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

type PostgreSQL struct {
	Connection        string `yaml:"connection"`
	ConnectionTimeout int64  `yaml:"connection-timeout"`
}

type Valkey struct {
	Connection        string `yaml:"connection"`
	Username          string `yaml:"username"`
	Password          string `yaml:"password"`
	ConnectionTimeout int64  `yaml:"connection-timeout"`
}

type Ollama struct {
	Connection        string `yaml:"connection"`
	ConnectionTimeout int64  `yaml:"connection-timeout"`
	OperationTimeout  int64  `yaml:"operation-timeout"`
}
