package config

import "github.com/caarlos0/env/v6"

type Config struct {
	HttpPort   int    `env:"HTTP_PORT" envDefault:"8080"`
	DBHost     string `env:"DB_HOST" envDefault:"127.0.0.1"`
	DBPort     int    `env:"DB_PORT" envDefault:"3306"`
	DBUser     string `env:"DB_USER" envDefault:"taro"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"pass"`
	DBName     string `env:"DB_NAME" envDefault:"fcoin-balances-db"`
}

/*
環境変数を読み込む
環境変数の値が設定されていない場合は、envDefaultの値が代入される
ローカルでは、envDefaultの値を使用し、
CIやdev/stg/prod各環境では、別途適切な値を設定するようにする
*/
func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
