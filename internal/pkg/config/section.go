package config

type ServerConfig struct {
	Runmode    string `mapstructure:"runmode"`
	Name       string `mapstructure:"name"`
	HttpPort   string `mapstructure:"http_port"`
	HttpsOpen  bool   `mapstructure:"https_open"`
	HttpsPort  string `mapstructure:"https_port"`
	TlsCert    string `mapstructure:"tls_cert"`
	TlsKey     string `mapstructure:"tls_key"`
	PingUrl    string `mapstructure:"ping_url"`
	PingMaxNum int    `mapstructure:"ping_max_num"`
}

type SafetyConfig struct {
	JwtSecret string `mapstructure:"jwt_secret"`
}

type LogConfig struct {
	LoggerFileInfo   string `mapstructure:"logger_file_info"`
	LoggerFileError  string `mapstructure:"logger_file_error"`
	LoggerMaxSize    int    `mapstructure:"logger_max_size"`
	LoggerMaxBackups int    `mapstructure:"logger_max_backups"`
	LoggerMaxAge     int    `mapstructure:"logger_max_age"`
}

type DbConfig struct {
	DbType       string `mapstructure:"db_type"`
	Name         string `mapstructure:"name"`
	Addr         string `mapstructure:"addr"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
}
