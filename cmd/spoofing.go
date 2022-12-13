package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"quick-tricks/modules/spoofing"
	"quick-tricks/utils/colors"
)

// spoofingCmd represents the spoofing command.
var spoofingCmd = &cobra.Command{
	Use:   "spoofing",
	Short: "The 'spoofing' module tests target for possibility of Content Spoofing attack.",

	Run: func(cmd *cobra.Command, args []string) {
		target, _ := cmd.Flags().GetString("url")
		if target == "" {
			cmd.Help()
			colors.BAD.Println("Target must be specified.")
			os.Exit(-1)
		}
		spoofingUrl, err := spoofing.Detect(target)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
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
}
