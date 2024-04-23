package models

type DbDetails struct {
	DBUsername string `yaml:"dbusername"`
	DBPassword string `yaml:"dbpassword"`
	DBName     string `yaml:"dbname"`
	DBHost     string `yaml:"dbhost"`
	DBPort     string `yaml:"dbport"`
}
