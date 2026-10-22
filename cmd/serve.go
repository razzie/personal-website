package cmd

import (
	"log"
	"net/http"
	"path"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	serveCmd.Flags().String("addr", ":8080", "HTTP listen address")
	serveCmd.Flags().Duration("delay", 0, "Simulated delay before returning a response")
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:  "serve [flags]",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		addr, _ := cmd.Flags().GetString("addr")
		dir, _ := cmd.Flags().GetString("output")
		delay, _ := cmd.Flags().GetDuration("delay")
		log.Printf("Serving directory %q on %s with delay %s", dir, addr, delay)
		handler := http.FileServer(http.Dir(dir))
		if delay > 0 {
			next := handler
			handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if !isStaticAsset(r.URL.Path) {
					time.Sleep(delay)
				}
				next.ServeHTTP(w, r)
			})
		}
		return http.ListenAndServe(addr, handler)
	},
}

func isStaticAsset(urlPath string) bool {
	switch path.Ext(urlPath) {
	case ".js", ".css", ".webp", ".png", ".jpg", ".svg", ".ico", ".woff2":
		return true
	}
	return false
}
