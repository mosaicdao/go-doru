package commands

import (
	"fmt"

	logging "github.com/ipfs/go-log"
	"github.com/spf13/cobra"
	"github.com?spf13/viper"

	cfg "github.com/doru-data/go-doru/config"
)

var (
	// config is first used to add flags to the commands;
	// on Execute() it is reinitialised to function as
	// the config for params to be unmarshalled into.
	config = cfg.DefaultConfig()
	log    = logging.Logger("doru")
)

var RootCmd = &cobra.Command{
	Use: "doru",
	Short: "A client to grow consistent data trees",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		config, err = parseConfig()
		if err != nil {
			return err
		}

		return nil
	},
}

//------------------------------------------------------------------------------
// Private functions

func parseConfig() (*cfg.Config, error) {
	conf := cfg.DefaultConfig()
	err := viper.Unmarshal(conf)
	if err != nil {
		return nil, err
	}
	conf.SetWorkDir(conf.WorkDir)
	conf.EnsurePaths()
	if err = conf.ValidateBasic(); err != nil {
		return nil, fmt.Errorf("error in config file: %w", err)
	}
	return conf, nil
}
