package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"quick-tricks/modules/redirect"
	"quick-tricks/utils/colors"
)

// redirectCmd represents the redirect command.
var redirectCmd = &cobra.Command{
	Use:   "redirect",
	Short: "Module 'redirect' checks endpoints vulnerable to Open Redirect.",
	Run: func(cmd *cobra.Command, args []string) {
		target, _ := cmd.Flags().GetString("url")

		redirectUrls, err := redirect.Detect(target)
		if err != nil {
			fmt.Println(err)
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
}
