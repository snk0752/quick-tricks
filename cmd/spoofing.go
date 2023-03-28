package cmd

import (
	"fmt"
	"github.com/indigo-sadland/quick-tricks/modules/spoofing"
	"github.com/indigo-sadland/quick-tricks/utils/colors"
	"github.com/spf13/cobra"
)

// spoofingCmd represents the spoofing command.
var spoofingCmd = &cobra.Command{
	Use:   "spoofing",
	Short: "Module 'spoofing' tests target for possibility of Content Spoofing attack.",

	Run: func(cmd *cobra.Command, args []string) {
		target, _ := cmd.Flags().GetString("url")
		proxy, _ := cmd.Flags().GetString("proxy")

		spoofingUrl, err := spoofing.Detect(target, proxy)
		if err != nil {
			colors.BAD.Println(err)
			return
		}
		if spoofingUrl != "" {
			colors.OK.Println("Success content spoofing attack!")
			fmt.Println(spoofingUrl)
		} else {
			colors.BAD.Println("Content spoofing attack failed.")
		}
	},
}

func init() {
	rootCmd.AddCommand(spoofingCmd)

	spoofingCmd.Flags().StringP("url", "u", "", "Target Bitrix site")
	spoofingCmd.MarkFlagRequired("url")
	spoofingCmd.Flags().String("proxy", "", "http/socks5 proxy to use. Example: socks5://IP:PORT")
}
