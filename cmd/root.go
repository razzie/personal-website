package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().String("o", "dist", "Distribution/output directory")
}

var rootCmd = &cobra.Command{
	Use: "personal-website",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
