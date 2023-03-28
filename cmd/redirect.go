package cmd

import (
	"fmt"
	"github.com/indigo-sadland/quick-tricks/modules/redirect"
	"github.com/indigo-sadland/quick-tricks/utils/colors"
	"github.com/spf13/cobra"
)

// redirectCmd represents the redirect command.
var redirectCmd = &cobra.Command{
	Use:   "redirect",
	Short: "Module 'redirect' checks endpoints vulnerable to Open Redirect.",
	Run: func(cmd *cobra.Command, args []string) {
		target, _ := cmd.Flags().GetString("url")
		proxy, _ := cmd.Flags().GetString("proxy")

		redirectUrls, err := redirect.Detect(target, proxy)
		if err != nil {
			colors.BAD.Println(err)
			return
		}
		if len(redirectUrls) != 0 {
			colors.OK.Println("Links vulnerable to redirect:")
			for _, v := range redirectUrls {
				fmt.Println(v)
			}
		} else {
			colors.BAD.Println("No links vulnerable to redirect.")
		}
	},
}

func init() {
	rootCmd.AddCommand(redirectCmd)

	redirectCmd.Flags().StringP("url", "u", "", "Target Bitrix site")
	redirectCmd.MarkFlagRequired("url")
	redirectCmd.Flags().String("proxy", "", "http/socks5 proxy to use. Example: socks5://IP:PORT")
}
