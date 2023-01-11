package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/indigo-sadland/quick-tricks/modules/spoofing"
	"github.com/indigo-sadland/quick-tricks/utils/colors"
)

// spoofingCmd represents the spoofing command.
var spoofingCmd = &cobra.Command{
	Use:   "spoofing",
	Short: "Module 'spoofing' tests target for possibility of Content Spoofing attack.",

	Run: func(cmd *cobra.Command, args []string) {
		target, _ := cmd.Flags().GetString("url")
		spoofingUrl, err := spoofing.Detect(target)
		if err != nil {
			fmt.Println(err)
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
}
