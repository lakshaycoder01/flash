package config

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/spf13/viper"
)

//MySQLConfig holder
type mySqlConfig struct {
	Address        string `mapstructure:"address"`
	Username       string `mapstructure:"username"`
	Password       string `mapstructure:"password"`
	Database       string `mapstructure:"database"`
	MaxConnections int    `mapstructure:"maxConnections"`
}

//ServerConfig holder
type serverConfig struct {
	Port          int     `mapstructure:"port"`
	AwsRegion     *string `mapstructure:"awsRegion"`
	TemplatesPath string  `mapstructure:"templatesPath"`
}

//CommonConfig server common settings
type commonConfig struct {
	// Domain      string `mapstructure:"domain"`
	Environment string `mapstructure:"environment"`
}

type Config struct {
	Server     *serverConfig `mapstructure:"server"`
	ReadMySQL  *mySqlConfig  `mapstructure:"readsqldb"`
	WriteMySQL *mySqlConfig  `mapstructure:"writesqldb"`
}

var appConfig *Config
var common *commonConfig

//IsDebugEnv env check
func IsDebugEnv() bool {
	return EnvironmentType(common.Environment) == EnvironmentTypeDebug
}

func configLoadFailed(err error) {
	fmt.Println(err)
	panic(err)
}

//LoadConfig from file
func Load() {
	appConfig = &Config{}

	if err := readFile("env"); err != nil {
		configLoadFailed(err)
	}

	common = &commonConfig{}
	if err := viper.Unmarshal(common); err != nil {
		configLoadFailed(err)
	}

	if err := readFile(common.Environment); err != nil {
		configLoadFailed(err)
	}

	if err := viper.Unmarshal(appConfig); err != nil {
		configLoadFailed(err)
	}

	fmt.Println("config:", *appConfig.Server.AwsRegion)
}

func readFile(name string) error {

	viper.SetConfigName(name)             // name of config file (without extension)
	viper.SetConfigType("json")           // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("config/secrets") // optionally look for config in the project directory

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Error reading config: missing env.json")
		} else {
			fmt.Println("Error reading config", err)
		}

		return err
	}

	return nil
}

func readSecret(configKey string) (*secretsmanager.GetSecretValueOutput, error) {
	//Create a Secrets Manager client
	awsConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(*appConfig.Server.AwsRegion))

	if err != nil {
		return nil, err
	}

	svc := secretsmanager.NewFromConfig(awsConfig)

	secretKey := fmt.Sprintf("%s/%s", common.Environment, configKey)

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretKey),
	}

	return svc.GetSecretValue(context.TODO(), input)
}

func readDecodedSecret(configKey string, v interface{}) {
	output, err := readSecret(configKey)

	if err != nil {
		configLoadFailed(err)
	}

	if err := json.Unmarshal([]byte(*output.SecretString), v); err != nil {
		configLoadFailed(err)
	}
}

func ListenPort() int {
	if appConfig.Server != nil {
		return appConfig.Server.Port
	}

	return 80
}
