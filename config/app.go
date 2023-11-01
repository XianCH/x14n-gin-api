package config

type app struct {
	Desc       string `yaml:"desc"`
	Addr       string `yaml:"addr"`
	ConfigFile string `yaml:"configFile"`
	Version    string `yaml:"version"`
	Env        string `yaml:"env"`
}

type ServerConfig struct {
	App   app   `yaml:"app"`
	Mysql mysql `yaml:"mysql"`
	Log   log   `yaml:"log"`
	jwt   Jwt   `yaml:"jwt"`
	// Jwt     jwt     `yaml:"jwt"`
	// Redis   redis   `yaml:"redis"`
	// Elastic elastic `yaml:"elastic"`
}
