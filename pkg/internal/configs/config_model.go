/*
Copyright 2023 Nitro Agility S.r.l.

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

package configs

type Config struct {
	Version      int            `yaml:"version"`
	Settings     Settings       `yaml:"settings"`
	Datasources  Datasources    `yaml:"datasources"`
	Metrics      []Metric       `yaml:"metrics"`
	Expectations []Expectations `yaml:"expectations"`
}

type Exporter struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
	Idle int    `yaml:"idle"`
}

type Settings struct {
	Exporter Exporter `yaml:"exporter"`
}

type Database struct {
	Name     string `yaml:"name"`
	Dialect  string `yaml:"dialect"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Datasources struct {
	Databases []Database `yaml:"databases"`
}

type Metric struct {
	Name        string `yaml:"name"`
	Type        string `yaml:"type"`
	Description string `yaml:"description"`
	Labels      string `yaml:"labels"`
}

type Expectations struct {
	Name       string   `yaml:"name"`
	Datasource string   `yaml:"datasource"`
	Metric     []string `yaml:"metrics"`
	Query      string   `yaml:"query"`
}
