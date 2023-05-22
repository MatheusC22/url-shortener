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
	Port_1      string
	Port_2      string
	Port_grapql string
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
		Port_1:      viper.GetString("api.port_1"),
		Port_2:      viper.GetString("api.port_2"),
		Port_grapql: viper.GetString("api.port_graphql"),
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

func GetApi() APIConfig {
	return cfg.API
}

func GetJWTSecret() string {
	return cfg.JWT.Secret
}

func GetRedis() REDISConfig {
	return cfg.REDIS
}
