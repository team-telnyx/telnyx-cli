package config

import (
	"errors"
	"os"
	"path"

	"github.com/spf13/cobra"
)

const ConfigFolder = ".telnyx-cli"
const ConfigFile = "config.yaml"
const CacheFolder = "cache"

type Config struct {
	ConsulDevUrl  string `yaml:"consul_dev_url"`
	ConsulProdUrl string `yaml:"consul_prod_url"`
}

func DefaultConfig() *Config {
	return &Config{
		ConsulDevUrl:  "https://consuldev.internal.telnyx.com",
		ConsulProdUrl: "https://consul.internal.telnyx.com",
	}
}

func DefaultConfigFolder() string {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	return path.Join(home, ConfigFolder)
}

func DefaultConfigPath() string {
	return path.Join(DefaultConfigFolder(), ConfigFile)
}

func ConfigExists() (bool, error) {
	_, err := os.Stat(DefaultConfigPath())
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}
