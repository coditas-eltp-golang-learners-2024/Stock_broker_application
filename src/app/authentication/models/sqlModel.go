package models

type SQLConfig struct {
    Username string `yaml:"username"`
    Password string `yaml:"password"`
    Host     string `yaml:"host"`
    Port     string `yaml:"port"`
    DBName   string `yaml:"dbname"`
}
