package commands

import (
	"os"

	"github.com/skycoin/dmsg/cipher"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var unsafe = false

func init() {
	confgenCmd.Flags().BoolVar(&unsafe, "unsafe", unsafe,
		"will unsafely write config if set")

	rootCmd.AddCommand(confgenCmd)
}

var confgenCmd = &cobra.Command{
	Use:    "confgen <config.json>",
	Short:  "generates config file",
	Args:   cobra.MaximumNArgs(1),
	PreRun: prepareVariables,
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) == 0 {
			confPath = "./config.json"
		} else {
			confPath = args[0]
		}

		if _, err := os.Stat(confPath); err == nil {

			log.Info(confPath, "Already exists. Run 'dmsgpty-host --help' for usage.")

		} else {

			pk, sk := cipher.GenerateKeyPair()
			log.WithField("pubkey", pk).
				WithField("seckey", sk).
				Info("Generating key pair")

			viper.Set("sk", sk)
		}

		if unsafe {
			return viper.WriteConfigAs(confPath)
		}
		return viper.SafeWriteConfigAs(confPath)
	},
}
