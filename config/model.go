package config

type Config struct {
	Zop   ZopConfig   `ini:"zop"`
	Mysql MysqlConfig `ini:"mysql"`
}

type ZopConfig struct {
	Host string `ini:"host"`
	Port string `ini:"port"`
}

type MysqlConfig struct {
	Username string `ini:"username"`
	Password string `ini:"password"`
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	Database string `ini:"database"`
	Dsn      string `ini:"dsn"`
}
