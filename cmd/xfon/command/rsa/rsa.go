package rsa

import (
	"log"
	"os"

	"github.com/odacremolbap/xfon/pkg/cert"
	"github.com/spf13/cobra"
)

var (
	bits int

	// RootCmd manages private keys
	RootCmd = &cobra.Command{
		Use:   "rsa",
		Short: "manages private keys",
		Run:   runHelp,
	}

	// NewCmd creates new RSA key
	NewCmd = &cobra.Command{
		Use:   "new",
		Short: "creates new RSA key",
		Run: func(cmd *cobra.Command, args []string) {
			_, err := cert.GenRSAKey(bits)
			if err != nil {
				log.Printf("error generating RSA key: %v", err.Error())
				os.Exit(-1)
			}
		},
	}
)

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func init() {
	NewCmd.PersistentFlags().IntVar(&bits, "bits", 4096, "key size")
	RootCmd.AddCommand(NewCmd)
}
