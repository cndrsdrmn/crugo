package config

var App *AppConfig

type AppConfig struct {
	Name  string `yaml:"name"`
	Port  int    `yaml:"port"`
	Debug bool   `yaml:"debug"`
}
