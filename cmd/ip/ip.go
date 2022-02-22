package ip

import (
	"fmt"
	"io"
	"logdna/logging"
	"net/http"

	"github.com/spf13/cobra"
)

var sublog logging.Logger

// ipCmd represents the ip command
var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		sublog.Print("testing logger")
		resp, err := http.Get("http://ip.jsontest.com")
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		fmt.Println(string(body))
		return nil
	},
}

// Attach the main ip command to a parent cobra command
func Attach(cmd *cobra.Command, logger logging.Logger) {
	sublog = logger.Child("ip")
	cmd.AddCommand(ipCmd)
}
