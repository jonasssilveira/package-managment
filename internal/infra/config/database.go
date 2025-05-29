package config

type Database struct {
	Mongo `yaml:"mongo"`
}

type Mongo struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}
