package configs

import (
	Logger "github.com/ewinjuman/go-lib/logger"
)

type Configuration struct {
	Apps          Apps
	Logger        Logger.Options
	Database      Database
	Redis         Redis
	Qontak        Qontak
	ForwardOption ForwardOption
}

type LoggerOptions struct {
	FileLocation       string `json:"fileLocation" envconfig:"LOGGER_FILE_LOCATION"`
	FileName           string `json:"fileName" envconfig:"LOGGER_FILE_NAME"`
	FileTdrLocation    string `json:"fileTdrLocation" envconfig:"LOGGER_FILE_TDRLOCATION"`
	FileMaxAge         int    `json:"fileMaxAge" envconfig:"LOGGER_FILE_MAXAGE"`
	Stdout             bool   `json:"stdout" envconfig:"LOGGER_STDOUT"`
	MaskingLogJsonPath string `json:"maskingLogJsonPath" envconfig:"LOGGER_MASKING_JSONPATH"`
}

type Apps struct {
	Name                   string `json:"name" envconfig:"APPS_NAME" default:"MyApp"`
	HttpPort               int    `json:"httpPort" envconfig:"APPS_PORT" default:"MyApp"`
	Mode                   string `json:"mode" envconfig:"APPS_MODE" default:"dev"`
	JwtSecretKey           string `json:"jwt_secret_key" envconfig:"APPS_JWT_KEY" default:"secret"`
	TokenExpiration        int    `json:"token_expiration" envconfig:"APPS_TOKEN_EXPIRATION" default:"3600"`
	JwtRefreshSecretKey    string
	RefreshTokenExpiration int
}

type Database struct {
	DbType      string `json:"dbType" envconfig:"DB_TYPE" default:"postgres"`
	Username    string `json:"username" envconfig:"DB_USERNAME" default:"root"`
	Password    string `json:"password" envconfig:"DB_PASSWORD" default:"password"`
	Schema      string `json:"schema" envconfig:"DB_SCHEMA" default:"postgres"`
	Host        string `json:"host" envconfig:"DB_HOST" default:"localhost"`
	Port        int    `json:"port" envconfig:"DB_PORT" default:"5432"`
	MaxIdleConn int    `json:"maxIdleConn" envconfig:"DB_MAX_IDLE_CONN" default:"0"`
	MaxOpenConn int    `json:"maxOpenConn" envconfig:"DB_MAX_OPEN_CONN" default:"0"`
	LogMode     bool   `json:"logMode" envconfig:"DB_LOG_MODE" default:"false"`
}

type Redis struct {
	Address  string `json:"address" envconfig:"REDIS_ADDRESS" default:"localhost:6379"`
	Password string `json:"password" envconfig:"REDIS_PASSWORD" default:"password"`
	Database int    `json:"database" envconfig:"REDIS_DATABASE" default:"redis"`
}

type Qontak struct {
	Option struct {
		Timeout   int  `json:"timeout" json:"timeout" envconfig:"QONTAK_OPTION_TIMEOUT" default:""`
		DebugMode bool `json:"debugMode" envconfig:"QONTAK_OPTION_DEBUGMODE" default:""`
		SkipTLS   bool `json:"skipTLS" envconfig:"QONTAK_OPTION_SKIPTLS" default:""`
	}
	Host string `json:"host" envconfig:"QONTAK_HOST" default:"localhost"`
	Path struct {
		WaTemplateList      string `json:"templateList" envconfig:"QONTAK_PATH_WA_TEMPLATE_LIST" default:"localhost"`
		WaSendMessage       string `json:"createMessage" envconfig:"QONTAK_PATH_WA_SEND_MESSAGE" default:""`
		WaBroadcastDirect   string `json:"broadcastDirect" envconfig:"QONTAK_PATH_WA_BROADCAST_DIRECT" default:""`
		TelegramSendMessage string `json:"createTelegramMessage" envconfig:"QONTAK_PATH_TELEGRAM_SEND_MESSAGE" default:""`
		Resolved            string `json:"resolved" envconfig:"QONTAK_PATH_RESOLVED" default:""`
	}
}

type ForwardOption struct {
	Option struct {
		Timeout   int  `json:"timeout" json:"timeout" envconfig:"FORWARD_OPTION_TIMEOUT" default:""`
		DebugMode bool `json:"debugMode" envconfig:"FORWARD_OPTION_DEBUGMODE" default:""`
		SkipTLS   bool `json:"skipTLS" envconfig:"FORWARD_OPTION_SKIPTLS" default:""`
	}
}
