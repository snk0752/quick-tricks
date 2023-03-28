package cmd

import (
	"github.com/spf13/cobra"
)

// quickCmd represents the quick command
var quickCmd = &cobra.Command{
	Use:   "quick",
	Short: "Run all quick modules ('recon', 'lfi', 'redirect', 'spoofing' and 'xss')",
	Run: func(cmd *cobra.Command, args []string) {
		//lfiCmd.Run(cmd, args)
		reconCmd.Run(cmd, args)
		redirectCmd.Run(cmd, args)
		spoofingCmd.Run(cmd, args)
		xssCmd.Run(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(quickCmd)
	quickCmd.Flags().StringP("url", "u", "", "Target Bitrix site")
	quickCmd.MarkFlagRequired("url")
	quickCmd.Flags().String("proxy", "", "http/socks5 proxy to use. Example: socks5://IP:PORT")

}
