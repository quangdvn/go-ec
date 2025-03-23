package settings

type Config struct {
	Server     ServerSetting         `mapstructure:"server"`
	Mysql      MySQLSetting          `mapstructure:"mysql"`
	Redis      RedisSetting          `mapstructure:"redis"`
	Logger     LoggerSetting         `mapstructure:"logger"`
	MailServer MailServerSetting     `mapstructure:"mail_server"`
	Gmail      GmailServerSetting    `mapstructure:"gmail"`
	SendGrid   SendGridServerSetting `mapstructure:"sendgrid"`
}

type ServerSetting struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type MySQLSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Dbname          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
	ConnMaxIdleTime int    `mapstructure:"connMaxIdleTime"`
}

type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
	PoolSize int    `mapstructure:"pool_size"`
}

type LoggerSetting struct {
	LogLevel    string `mapstructure:"log_level"`
	FileLogName string `mapstructure:"file_log_name"`
	MaxSize     int    `mapstructure:"max_size"`
	MaxBackups  int    `mapstructure:"max_backups"`
	MaxAge      int    `mapstructure:"max_age"`
	Compress    bool   `mapstructure:"compress"`
}

type MailServerSetting struct {
	Provider string `mapstructure:"provider"`
}

type GmailServerSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
type SendGridServerSetting struct {
	ApiKey string `mapstructure:"api_key"`
}
