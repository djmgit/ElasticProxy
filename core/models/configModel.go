
package models

type backendConfig struct {
	Port int64 `yaml:"port"`
}

type haproxyConfig struct {
	ConfigFile int64 `yaml:"config_file"`
	Binary string `yaml:"binary"`
}

type ConfigModel struct {
	BackendConfig backendConfig `yaml:"backend_config"`
	HaproxyConfig haproxyConfig `yaml:"haproxy_config"`
}
