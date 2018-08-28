package config

type CProgram struct {
	Path string
	Name string
}
type CWatch struct {
	Version  string              `yaml:"version"`
	Mode     string              `yaml:"mode"`
	Port     int                 `yaml:"Port"`
	Interval int                 `yaml:"interval"`
	Programs map[string]CProgram `yaml:"programs"`
}
