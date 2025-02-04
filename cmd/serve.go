package cmd

import (
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

func init() {
	serveCmd.Flags().String("addr", ":8080", "HTTP listen address")
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:  "serve [flags]",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		addr, _ := cmd.Flags().GetString("addr")
		dir, _ := cmd.Flags().GetString("output")
		log.Printf("Serving directory %q on %s", dir, addr)
		return http.ListenAndServe(addr, http.FileServer(http.Dir(dir)))
	},
}
