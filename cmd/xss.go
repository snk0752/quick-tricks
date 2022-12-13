package cmd

import (
	"fmt"
	"quick-tricks/modules/xss"
	"quick-tricks/utils/colors"

	"github.com/spf13/cobra"
)

// xssCmd represents the xss command.
var xssCmd = &cobra.Command{
	Use:   "xss",
	Short: "Module 'xss' checks target's endpoints that potentially can be vulnerable XSS.",
	Run: func(cmd *cobra.Command, args []string) {
		target, _ := cmd.Flags().GetString("url")

		xssUrls, err := xss.BuildPayloads(target)
		if err != nil {
			fmt.Println(err)
			return
		}
		if len(xssUrls) != 0 {
			colors.OK.Printf("Links possibly vulnerable to XSS:\n")
			for _, v := range xssUrls {
				fmt.Println(v)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(xssCmd)
	xssCmd.Flags().StringP("url", "u", "", "Target Bitrix site")
	xssCmd.MarkFlagRequired("url")
}
