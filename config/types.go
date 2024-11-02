package config

type secret struct {
	JwtSecret string `mapstructure:"jwt_secret"`
	Issuer    string
}

type mysql struct {
	UserName string
	Password string
	Address  string
	Database string
	Charset  string
}

type config struct {
	Secret secret
	Mysql  mysql
}
