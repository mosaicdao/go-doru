package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/doru-doru/go-doru/cmd"
	"github.com/doru-doru/go-doru/v0/core"
	"google.golang.org/grpc"

	// "github.com/mitchellh/go-homedir"
	logging "github.com/ipfs/go-log/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/textileio/go-threads/core/did"
	tclient "github.com/textileio/go-threads/net/api/client"
	"github.com/textileio/go-threads/util"
)

const daemonName = "dorud"

var (
	log = logging.Logger("dorud")

	config = &cmd.Config{
		Viper:     viper.New(),
		Directory: "." + daemonName,
		Name:      "config",
		Flags: map[string]cmd.Flag{
			"debug": {
				Key:          "log.debug",
				DefaultValue: false,
			},

			// Addresses
			"addressApi": {
				Key:          "address.api",
				DefaultValue: "127.0.0.1:1414",
			},

			// Datastore
			"datastoreType": {
				Key:          "datastore.type",
				DefaultValue: "badger",
			},
			"datastoreBadgerRepo": {
				Key:     "datastore.badger.repo",
				DefaultValue: "${HOME}/." + daemonName + "/repo",
			},

			// Threads
			"threadsAddress": {
				Key: "threads.address",
				DefaultValue: "127.0.0.1:4000",
			},

			// IPFS
			"ipfsMultiaddress": {
				Key: "ipfs.multiaddr",
				DefaultValue: "/ip4/127.0.0.1/tcp/5001",
			},
		},
		EnvPrefix: "DORU",
	}
)

var rootCmd = &cobra.Command{
	Use: daemonName,
	Short: "Doru daemon",
	Long:  "Doru daemon to grow consistent data tries.",
	PersistentPreRun: func(c *cobra.Command, args []string) {
		config.Viper.SetConfigType("yaml")
		//TODO: cmd.ExpandConfigVars(config.Viper, config.Flags)

		if config.Viper.GetBool("log.debug") {
			err := util.SetLogLevels(map[string]logging.LogLevel{
				daemonName:    logging.LevelDebug,
			})
			cmd.ErrCheck(err)
		}
	},
	Run: func(c *cobra.Command, args []string) {
		settings, err := json.MarshalIndent(config.Viper.AllSettings(), "", "  ")
		cmd.ErrCheck(err)
		log.Debug("loaded config: %s", string(settings))

		debug := config.Viper.GetBool("log.debug")
		// TODO: setup log file

		addressApi := config.Viper.GetString("address.api")
		datastoreType := config.Viper.GetString("datastore.type")
		datastoreBadgerRepo := config.Viper.GetString("datastore.badger.repo")
		threadsAddress := config.Viper.GetString("threads.addr")
		ipfsMultiaddress := config.Viper.GetString("ipfs.multiaddr")

		// net, err := tclient.NewClient(threadsAddress, getClientRPCOpts(threadsAddress)...)
		var opts []core.Option
		opts = append(opts, core.WithBadgerThreadsPersistance(
			config.Viper.GetString("datastore.badger.repo")))

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		doru, err := core.NewDoru(ctx, core.Config{
			Debug: debug,

			AddressApi: addressApi,
			AddressThreadsHost: threadsAddress,
			AddressIpfsHost:    ipfsMultiaddress,
		})
	},
}

func init() {
	cobra.OnInitialize(cmd.InitConfig(config))

	rootCmd.PersistentFlags().StringVar(
		&config.Path,
		"path",
		"",
		"Working directory for Doru")

	rootCmd.PersistentFlags().StringVar(
		&config.File,
		"config",
		"",
		"Config file")

	rootCmd.PersistentFlags().BoolP(
		"debug",
		"d",
		config.Flags["debug"].DefaultValue.(bool),
		"Enable debug logging")

	rootCmd.PersistentFlags().String(
		"addressApi",
		config.Flags["addressApi"].DefaultValue.(string),
		"API listen address")

	rootCmd.PersistentFlags().String(
		"datastoreType",
		config.Flags["datastoreType"].DefaultValue.(string),
		"Datastore type (only badger atm)")

	rootCmd.PersistentFlags().String(
		"datastoreBadgerRepo",
		config.Flags["datastoreBadgerRepo"].DefaultValue.(string),
		"Path to badger repository")

	rootCmd.PersistentFlags().String(
		"threadsAddress",
		config.Flags["threadsAddress"].DefaultValue.(string),
		"Threads API address")

	rootCmd.PersistentFlags().String(
		"ipfsMultiaddress",
		config.Flags["ipfsMultiaddress"].DefaultValue.(string),
		"IPFS API multiaddress")

	err := cmd.BindFlags(config.Viper, rootCmd, config.Flags)
	cmd.ErrCheck(err)
}

func main() {
	cmd.ErrCheck(rootCmd.Execute())
}


// TODO: taken from go-buckets with threads refactor towards DID;
// once stable, also move threads to DID
// see https://github.com/textileio/go-buckets/pull/2 and related docs
//
// func getClientRPCOpts(target string) (opts []grpc.DialOption) {
// 	creds := did.RPCCredentials{}
// 	if strings.Contains(target, "443") {
// 		tcreds := credentials.NewTLS(&tls.Config{})
// 		opts = append(opts, grpc.WithTransportCredentials(tcreds))
// 		creds.Secure = true
// 	} else {
// 		opts = append(opts, grpc.WithInsecure())
// 	}
// 	opts = append(opts, grpc.WithPerRPCCredentials(creds))
// 	return opts
// }
