package configs

import "github.com/spf13/viper"

var cfg *config

type config struct {
	API   APIConfig
	DB    DBConfig
	REDIS REDISConfig
	JWT   JWTConfig
}

type APIConfig struct {
	Port string
}

type JWTConfig struct {
	Secret string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type REDISConfig struct {
	Addr     string
	Password string
	Db       int
}

func init() {
	viper.SetDefault("api.port", "3001")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "3306")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}
	cfg.JWT = JWTConfig{
		Secret: viper.GetString("JWT.secret"),
	}
	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		Database: viper.GetString("database.name"),
	}
	cfg.REDIS = REDISConfig{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		Db:       viper.GetInt("redis.db"),
	}

	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}

func GetJWTSecret() string {
	return cfg.JWT.Secret
}

func GetRedis() REDISConfig {
	return cfg.REDIS
}
