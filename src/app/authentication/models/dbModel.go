package models

type DBConfig struct {
	Host     string `yaml:"dbHostName"`
	Port     int    `yaml:"port"`
	Username string `yaml:"userName"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbName"`
}
