/*
Copyright 2021 The Pixiu Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	yamlConfig = "yaml"
)

type Config struct {
	configFile string
	configType string

	data []byte
}

func New() *Config {
	return &Config{}
}

func (c *Config) SetConfigFile(configFile string) {
	c.configFile = configFile
}

func (c *Config) SetConfigType(in string) {
	c.configType = in
}

func (c *Config) readInConfig() error {
	var err error
	c.data, err = ioutil.ReadFile(c.configFile)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) Binding(out interface{}) error {
	if err := c.readInConfig(); err != nil {
		return err
	}
	switch c.configType {
	case yamlConfig:
		if err := yaml.Unmarshal(c.data, out); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported config type %s", c.configType)
	}

	return nil
}
