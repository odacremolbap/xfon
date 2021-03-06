package rsa

import (
	"log"
	"os"

	"github.com/odacremolbap/xfon/pkg/filesystem"

	"github.com/odacremolbap/xfon/pkg/rsa"
	"github.com/spf13/cobra"
)

var (
	bits int
	out  string

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
		Run:   newFunc,
	}
)

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func init() {
	NewCmd.Flags().IntVar(&bits, "bits", 4096, "key size")
	NewCmd.Flags().StringVar(&out, "out", "", "RSA key output file")
	NewCmd.MarkFlagRequired("out")
	RootCmd.AddCommand(NewCmd)
}

// newFunc runs the new RSA command
func newFunc(cmd *cobra.Command, args []string) {
	k, err := rsa.GenerateKey(bits)
	if err != nil {
		log.Printf("error generating RSA key: %v", err.Error())
		os.Exit(-1)
	}

	p, err := rsa.WritePEM(k)
	if err != nil {
		log.Printf("error serializing RSA key into PEM: %v", err.Error())
		os.Exit(-1)
	}

	err = filesystem.WriteContentsToFile(out, p)
	if err != nil {
		log.Printf("error writing RSA key to file: %v", err.Error())
		os.Exit(-1)
	}
}
