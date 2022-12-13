package cmd

import (
	"fmt"
	"quick-tricks/modules/lfi"
	"quick-tricks/utils/colors"

	"github.com/spf13/cobra"
)

// lfiCmd represents the spoofing command
var lfiCmd = &cobra.Command{
	Use:   "lfi",
	Short: "Module 'lfi' checks if there are endpoints vulnerable to Local File Inclusion.",
	Run: func(cmd *cobra.Command, args []string) {
		target, _ := cmd.Flags().GetString("url")

		lfiUrls, err := lfi.Detect(target)
		if err != nil {
			fmt.Println(err)
		}

		if len(lfiUrls) != 0 {
			colors.OK.Printf("Path to LFI:\n")
			for _, v := range lfiUrls {
				fmt.Println(v)
			}
		} else {
			colors.BAD.Println("There is no path to LFI.")
		}
	},
}

func init() {
	rootCmd.AddCommand(lfiCmd)
	lfiCmd.Flags().StringP("url", "u", "", "Target Bitrix site")
	lfiCmd.MarkFlagRequired("url")
}
