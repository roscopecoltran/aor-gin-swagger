package config

type ApiConfig struct {
	Enabled bool         `env:"SERVER_API_DISABLED" default:"true" json:"enabled" yaml:"enabled" toml:"enabled"`
	Port    uint         `env:"SERVER_API_PORT" default:"7000" json:"port" yaml:"port" toml:"port"`
	CORS    CorsConfig   `json:"cors" yaml:"cors" toml:"cors"`
	RPC     RpcConfig    `json:"rpc" yaml:"rpc" toml:"rpc"`
	JWT     JwtConfig    `json:"jwt" yaml:"jwt" toml:"jwt"`
	OAuth2  OAuth2Config `json:"oauth2" yaml:"oauth2" toml:"oauth2"`
	Debug   bool         `env:"SERVER_API_DEBUG" default:"false" json:"debug" yaml:"debug" toml:"debug"`
}

type CorsConfig struct {
	Enabled                   bool   `env:"SERVER_API_CORS_DISABLED" default:"true" json:"enabled" yaml:"enabled" toml:"enabled"`
	ContentType               string `env:"SERVER_API_CORS_CONTENT_TYPE" default:"application/json" json:"content-type" yaml:"content-type" toml:"content-type"`
	AccessControlAllowOrigin  string `env:"SERVER_API_CORS_ALLOW_ORIGIN" default:"*" json:"access-control-allow-origin" yaml:"access-control-allow-origin" toml:"access-control-allow-origin"`
	AccessControlAllowHeaders string `env:"SERVER_API_CORS_ALLOW_HEADERS" default:"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization" json:"access-control-allow-headers" yaml:"access-control-allow-headers" toml:"access-control-allow-headers"`
	Debug                     bool   `env:"SERVER_API_DEBUG" default:"false" json:"debug" yaml:"debug" toml:"debug"`
}

type OAuth2Config struct {
	Enabled   bool                    `env:"SERVER_API_OAUTH2_DISABLED" default:"true" json:"enabled" yaml:"enabled" toml:"enabled"`
	Providers []*OAuth2ProviderConfig `json:"providers" yaml:"providers" toml:"providers"`
	Debug     bool                    `env:"SERVER_API_OAUTH2_DEBUG" default:"false" json:"debug" yaml:"debug" toml:"debug"`
}

type OAuth2ProviderConfig struct {
	Enabled     bool   `json:"enabled" yaml:"enabled" toml:"enabled"`
	Name        string `json:"name" yaml:"name" toml:"name"`
	Website     string `json:"website" yaml:"website" toml:"website"`
	ClientId    string `json:"client_id" yaml:"client_id" toml:"client_id"`
	SecretKey   string `json:"secret_key" yaml:"secret_key" toml:"secret_key"`
	CallbackUrl string `json:"callback_url" yaml:"callback_url" toml:"callback_url"`
	Debug       bool   `json:"debug" yaml:"debug" toml:"debug"`
}

type JwtConfig struct {
	Enabled   bool   `env:"SERVER_API_JWT_DISABLED" default:"true" json:"enabled" yaml:"enabled" toml:"enabled"`
	PublicKey string `env:"SERVER_API_JWT_PUBLIC_KEY" json:"public_key" yaml:"public_key" toml:"public_key"`
	SecretKey string `env:"SERVER_API_JWT_SECRET_KEY" description:"Key used to encryption cookies" default:"randombitsreplacedlkjsa" json:"secret_key" yaml:"secret_key" toml:"secret_key"`
	Debug     bool   `env:"SERVER_API_JWT_DEBUG" default:"false" json:"debug" yaml:"debug" toml:"debug"`
}

type RpcConfig struct {
	Enabled bool   `env:"SERVER_API_RPC_DISABLED" default:"true" json:"enabled" yaml:"enabled" toml:"enabled"`
	Address string `env:"SERVER_API_RPC_ADDRESS" description:"Rpc Server Address" default:"localhost:8901" json:"address" yaml:"address" toml:"address"`
	Debug   bool   `env:"SERVER_API_RPC_DEBUG" default:"false" json:"debug" yaml:"debug" toml:"debug"`
}

type WsConfig struct {
	Enabled bool   `env:"SERVER_API_WS_DISABLED" default:"true" json:"enabled" yaml:"enabled" toml:"enabled"`
	Address string `env:"SERVER_API_WS_ADDRESS" description:"Rpc Server Address" default:"localhost:8901" json:"address" yaml:"address" toml:"address"`
	SSL     bool   `env:"SERVER_API_WS_SSL" default:"false" json:"ssl_mode" yaml:"ssl_mode" toml:"ssl_mode"`
	Debug   bool   `env:"SERVER_API_WS_DEBUG" default:"false" json:"debug" yaml:"debug" toml:"debug"`
}

type SSLConfig struct {
	Enabled bool `env:"SERVER_API_SSL_DISABLED" default:"true" json:"enabled" yaml:"enabled" toml:"enabled"`
	Debug   bool `env:"SERVER_API_SSL_DEBUG" default:"false" json:"debug" yaml:"debug" toml:"debug"`
}
