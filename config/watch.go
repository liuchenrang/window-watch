package config

type CProgram struct {
	Path string
	Name string
}
type CWatch struct {
	Version  string              `yaml:"version"`
	Port     int                 `yaml:"Port"`
	Programs map[string]CProgram `yaml:"programs"`
}
