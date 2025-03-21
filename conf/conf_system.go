package conf

import "fmt"

type System struct {
	IP      string `yaml:"ip"`
	Port    string `yaml:"port"`
	GinMode string `yaml:"gin_mode"`
	Env     string `yaml:"env"`
}

func (s *System) Addr() string {
	return fmt.Sprintf("%s:%s", s.IP, s.Port)
}
