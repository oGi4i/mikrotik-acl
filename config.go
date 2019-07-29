package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

type YamlConfig struct {
	Access      ConfigAccess      `yaml:"access"`
	Database    ConfigDatabase    `yaml:"database"`
	Application ConfigApplication `yaml:"application"`
}

type ConfigAccess struct {
	Users []*User `yaml:"users"`
}

type ConfigDatabase struct {
	DSN        string        `yaml:"dsn"`
	Name       string        `yaml:"name"`
	Collection string        `yaml:"collection"`
	Timeout    time.Duration `yaml:"timeout"`
}

type ConfigApplication struct {
	Templates []*Template `yaml:"templates"`
}

type Template struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

func (cfg *YamlConfig) initConfig() error {
	configFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(configFile, &cfg)
	if err != nil {
		return err
	}

	users = cfg.Access.Users

	templates.Delims("#(", ")#")
	_, err = templates.ParseFiles(cfg.getTemplatesPaths()...)
	if err != nil {
		return err
	}

	return nil
}

func (cfg *YamlConfig) getTemplatesPaths() []string {
	result := make([]string, 0)
	for _, t := range cfg.Application.Templates {
		result = append(result, t.Path)
	}
	return result
}
