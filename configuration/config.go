package configuration

import (
	"fmt"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"strings"
)

type Provider interface {
	Config() (*Config, error)
}

type DefaultProvider struct{}

func (*DefaultProvider) Config() (*Config, error) {
	var out Config
	var k = koanf.New(".")

	if err := k.Load(file.Provider("barkconf.yaml"), yaml.Parser()); err != nil {
		return nil, fmt.Errorf("parsing configuration from yaml: %w", err)
	}
	err := k.Load(env.Provider("BARKDOG_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "BARKDOG_")), "_", ".", -1)
	}), nil)
	if err != nil {
		return nil, fmt.Errorf("parsing configuration from env: %w", err)
	}
	err = k.UnmarshalWithConf("", &out, koanf.UnmarshalConf{Tag: "koanf"})
	return &out, err
}

type Config struct {
	Site     Site     `koanf:"site"`
	Database Database `koanf:"database"`
	Env      Env      `koanf:"env"`
}
