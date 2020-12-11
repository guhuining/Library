package config

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Sql struct {
	UserName string `yaml:"user_name"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"db_name"`
}
