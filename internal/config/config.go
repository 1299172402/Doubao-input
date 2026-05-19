package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var globalConfig *Config

type Config struct {
	Port     string `mapstructure:"port"`
	AutoType bool   `mapstructure:"auto_type"`
}

func setDefaults() {
	viper.SetDefault("port", "2828")     // 默认端口
	viper.SetDefault("auto_type", false) // 默认不启用自动输入
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path) // 如 "config.yml"
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()               // 允许环境变量覆盖
	viper.SetEnvPrefix("DOUBAO_INPUT") // 环境变量前缀

	// 先设置默认值
	setDefaults()

	if err := viper.ReadInConfig(); err != nil {
		err := viper.WriteConfigAs(path)
		if err != nil {
			return nil, fmt.Errorf("创建默认配置文件失败: %w", err)
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func InitConfig() {
	cfg, err := LoadConfig("./doubao-input-config.yml")
	if err != nil {
		log.Fatal("加载配置失败: ", err)
	}
	globalConfig = cfg
}

func GetConfig() *Config {
	return globalConfig
}

func SaveConfig(cfg *Config) error {
	// 将结构体的值写回 viper
	viper.Set("port", cfg.Port)
	viper.Set("auto_type", cfg.AutoType)

	// 写入配置文件（会覆盖原文件）
	if err := viper.WriteConfig(); err != nil {
		// 如果 WriteConfig 失败（比如文件不存在），可以尝试 WriteConfigAs
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return viper.WriteConfigAs(viper.ConfigFileUsed())
		}
		return err
	}
	return nil
}
