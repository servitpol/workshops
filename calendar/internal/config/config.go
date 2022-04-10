package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

// Application holds application configuration values
type Application struct {
	DB      *Database `yaml:"db"`
	Lsn     *Listen   `yaml:"listen"`
	Jwt     *Jwt      `yaml:"jwt"`
	IsDebug *bool     `yaml:"is_debug"`
}

type Database struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	User   string `yaml:"user"`
	Pass   string `yaml:"pass"`
	DbName string `yaml:"db_name"`
}

type Listen struct {
	BindIP string `yaml:"bind_ip"`
	Port   string `yaml:"port"`
}

type Jwt struct {
	Secret  string `yaml:"secret_key"`
	Issuer  string `yaml:"issuer"`
	ExHours int64  `yaml:"ex_hours"`
}

var instance *Application
var once sync.Once

func GetConfig() *Application {
	once.Do(func() {
		instance = &Application{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Println(help)
			log.Fatal(err)
		}
	})
	return instance
}
