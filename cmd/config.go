package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Flag describes a command flag - credit textile
type Flag struct {
	Key          string
	DefaultValue interface{}
}

// Config describes a command configuration params and file info - credit textile
type Config struct {
	Viper     *viper.Viper
	File      string
	Directory string
	Path      string
	Name      string
	Flags     map[string]Flag
	EnvPrefix string
}

func InitConfig(config *Config) func() {
	return func() {
		initViperConfig(
			config.Viper,
			config.File,
			config.Path,
			config.EnvPrefix)
	}
}

// BindFlags binds the flags to the viper config values. - credit textile
func BindFlags(v *viper.Viper, root *cobra.Command, flags map[string]Flag) error {
	for n, f := range flags {
		if err := v.BindPFlag(f.Key, root.PersistentFlags().Lookup(n)); err != nil {
			return err
		}
		v.SetDefault(f.Key, f.DefaultValue)
	}
	return nil
}

func initViperConfig(
	v *viper.Viper,
	file,
	path,
	envPrefix string) {
	v.SetConfigFile(file)
	v.AddConfigPath(path)
	v.SetEnvPrefix(envPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// ignore ConfigFileNotFound error
		} else {
			// panic on error if config file is found but another error occured
			panic(err)
		}
	}
}
