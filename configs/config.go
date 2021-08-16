package configs

import (
	"encoding/json"
	"flag"
	"log"
	"strings"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const ServiceName = "api"

var options = []option{
	{"config", "string", "", "config file"},

	{"server.http.port", "int", 8080, "server http port"},
	{"server.http.timeout_sec", "int", 86400, "server http connection timeout"},
	{"server.grpc.port", "int", 9090, "server grpc port"},
	{"server.grpc.timeout_sec", "int", 86400, "server grpc connection timeout"},
	{"server.grpc.tls.enabled", "bool", false, "Enable or disable TLS"},

	{"postgres.master.host", "string", "localhost", "postgres master host"},
	{"postgres.master.port", "int", 5432, "postgres master port"},
	{"postgres.master.user", "string", "postgres", "postgres master user"},
	{"postgres.master.password", "string", "postgres", "postgres master password"},
	{"postgres.master.database_name", "string", "example", "postgres master database name"},
	{"postgres.master.secure", "string", "disable", "postgres master SSL support"},
	{"postgres.master.max_conns_pool", "int", 150, "max number of connections pool postgres"},

	{"postgres.replica.host", "string", "localhost", "postgres replica host"},
	{"postgres.replica.port", "int", 5432, "postgres replica port"},
	{"postgres.replica.user", "string", "postgres", "postgres replica user"},
	{"postgres.replica.password", "string", "postgres", "postgres replica password"},
	{"postgres.replica.database_name", "string", "example", "postgres replica database name"},
	{"postgres.replica.secure", "string", "disable", "postgres replica SSL support"},
	{"postgres.replica.max_conns_pool", "int", 150, "max number of connections pool postgres"},

	{"nats.host", "string", "127.0.0.1", "The nats host"},
	{"nats.port", "int", 4222, "The nats port"},
	{"nats.username", "string", "", "The nats user login"},
	{"nats.password", "string", "", "The nats user password"},
	{"nats.request_timeout_msec", "int", 500000, "The nats connection timeout in msec"},
	{"nats.retry_limit", "int", 5, "Reconnection limit to the nats"},
	{"nats.reconnect_time_wait_msec", "int", 500, "Reconnect time wait to the nats in msec"},

	{"sentry.enabled", "bool", false, "Enables or disables sentry"},
	{"sentry.dsn", "string", "https://7e67a2b5fd034e9dbb7cdc7d4cd1bccd@sentry.lenvendo.ru//11", "Data source name. Sentry addr"},
	{"sentry.environment", "string", "local", "The environment to be sent with events."},

	{"tracer.host", "string", "127.0.0.1", "The tracer host"},
	{"tracer.port", "int", 5775, "The tracer port"},
	{"tracer.enabled", "bool", false, "Enables or disables tracing"},
	{"tracer.name", "string", "github.com/lenvendo/ig-absolut-api", "The tracer name"},

	{"metrics.enabled", "bool", false, "Enables or disables metric"},
	{"metrics.port", "int", 9153, "server http port"},

	{"logger.level", "string", "emerg",
		"Level of logging. A string that correspond to the following levels: emerg, alert, crit, err, warning, notice, info, debug"},
	{"logger.time_format", "string", "2006-01-02T15:04:05.999999999", "Date format in logs"},
}

type Config struct {
	Server   Server
	Postgres struct {
		Master  Database
		Replica Database
	}
	Tracer struct {
		Enabled bool
		Host    string
		Port    int
		Name    string
	}
	Metrics struct {
		Enabled bool
		Port    int
	}

	Logger Logger
	Sentry struct {
		Enabled     bool
		Dsn         string
		Environment string
	}
	Nats struct {
		Host           string
		Port           int
		UserName       string
		Password       string
		RequestTimeOut time.Duration `mapstructure:"request_timeout_msec"`
		RetryLimit     int           `mapstructure:"retry_limit"`
		WaitLimit      int           `mapstructure:"reconnect_time_wait_msec"`
	}
	Limiter struct {
		Enabled bool
		Limit   float64
	}
}

type Database struct {
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string `mapstructure:"database_name"`
	Secure       string
	MaxConnsPool int `mapstructure:"max_conns_pool"`
}

type Server struct {
	GRPC struct {
		Port       int
		TimeoutSec int `mapstructure:"timeout_sec"`
		TLS        TLS
	}
	HTTP struct {
		Port       int
		TimeoutSec int `mapstructure:"timeout_sec"`
	}
}

// TLS - tls config
type TLS struct {
	Enabled bool
}

type GRPC struct {
	Host string
	Port int
	TLS  struct {
		InsecureSkipVerify bool `mapstructure:"insecure_skip_verify"`
		Enabled            bool
	}
}

type Logger struct {
	Level      string
	TimeFormat string `mapstructure:"time_format"`
}

type option struct {
	name        string
	typing      string
	value       interface{}
	description string
}

func NewConfig() *Config {
	return &Config{}
}

// Read read parameters for config.
// Read from environment variables, flags or file.
func (c *Config) Read() error {

	for _, o := range options {
		switch o.typing {
		case "string":
			flag.String(o.name, o.value.(string), o.description)
		case "int":
			flag.Int(o.name, o.value.(int), o.description)
		default:
			viper.SetDefault(o.name, o.value)
		}
	}
	viper.SetEnvPrefix(ServiceName)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		panic("Read config error: " + err.Error())
	}

	if fileName := viper.GetString("config"); fileName != "" {
		viper.SetConfigFile(fileName)
		viper.SetConfigType("toml")

		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	}

	if err := viper.Unmarshal(c); err != nil {
		return err
	}

	return nil
}

// Print print config structure
func (c *Config) Print() error {
	b, err := json.Marshal(c)
	if err != nil {
		return err
	}
	log.Println(string(b))
	return nil
}
