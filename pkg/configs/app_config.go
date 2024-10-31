package configs

import (
	"fmt"
	"github.com/ewinjuman/go-lib/helper/convert"
	Logger "github.com/ewinjuman/go-lib/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var Config *Configuration

func init() {
	if len(os.Args) > 1 && strings.Contains(strings.ToLower(os.Args[1]), "test") {
		_, file, _, _ := runtime.Caller(0)
		apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, "../.."+string(filepath.Separator))))
		err := os.Chdir(apppath)
		if err != nil {
			panic(err)
		}
	}
	Config = NewENV()
}

func ReloadConfig() (err error) {
	tempConfig, err := Reload()
	if err != nil {
		return
	}
	Config = tempConfig
	return
}

func getEnvironment() string {
	if len(os.Args) > 1 && !strings.Contains(strings.ToLower(os.Args[1]), "test") {
		return os.Args[1]
	}
	return "local"
}

func getConfigFilePath(env string) string {
	return fmt.Sprintf("./resource/conf/config.%s.json", env)
}

func validateAppMode(appMode, env string) {
	if appMode != env {
		panic(errors.New(fmt.Sprintf("Please change 'apps.mode' to '%v'", env)))
	}
}

func NewENV() *Configuration {
	var cfg Configuration

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}
	var log LoggerOptions
	err = envconfig.Process("LOGGER", &log)
	if err != nil {
		fmt.Println("Error loading environment variables:", err)
	}
	libLog := Logger.Options{}
	convert.ObjectToObject(log, &libLog)

	// Load environment variables ke struct
	err = envconfig.Process("", &cfg)
	if err != nil {
		fmt.Println("Error loading environment variables:", err)
	}
	cfg.Logger = libLog
	fmt.Printf("Config: %+v\n", cfg)
	return &cfg
}

func New() *Configuration {
	env := getEnvironment()
	path := getConfigFilePath(env)

	viper.SetConfigFile(path)
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		//if strings.Contains(strings.ToLower(env), "testlog") {
		//	return nil
		//} else {
		panic(err)
		//}

	}

	defaultConfig := Configuration{}
	err := viper.Unmarshal(&defaultConfig)
	if err != nil {
		panic(err)
	}

	validateAppMode(defaultConfig.Apps.Mode, env)

	viper.OnConfigChange(func(e fsnotify.Event) {
		err := ReloadConfig()
		if err != nil {
			fmt.Println("error reload config: ", err.Error())
		} else {
			fmt.Println("Config file changed:", time.Now().Format(time.RFC1123Z))
		}
	})
	viper.WatchConfig()
	fmt.Printf("Config: %+v\n", defaultConfig)
	return &defaultConfig
}

func Reload() (*Configuration, error) {
	env := Config.Apps.Mode
	defaultConfig := Configuration{}
	path := getConfigFilePath(env)

	viper.SetConfigFile(path)
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		return &defaultConfig, err
	}

	err := viper.Unmarshal(&defaultConfig)
	if err != nil {
		return &defaultConfig, err
	}

	if env != defaultConfig.Apps.Mode {
		return &defaultConfig, errors.New("apps.mode is different from the previous configuration!")
	}

	return &defaultConfig, nil
}
