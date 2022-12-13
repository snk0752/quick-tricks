package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/indigo-sadland/quick-tricks/modules/recon/license"
	"github.com/indigo-sadland/quick-tricks/modules/recon/lp"
	"github.com/indigo-sadland/quick-tricks/modules/recon/lpd"
	"github.com/indigo-sadland/quick-tricks/utils/colors"
	"strings"
)

// reconCmd represents the recon command
var reconCmd = &cobra.Command{
	Use:   "recon",
	Short: "Module 'recon' helps to find login page endpoints, local path disclosure and license key.",
	Run: func(cmd *cobra.Command, args []string) {
		var detectLoginPages, detectLicense, detectPathDisclosure bool

		detectLoginPages, _ = cmd.Flags().GetBool("lp")
		detectLicense, _ = cmd.Flags().GetBool("license")
		detectPathDisclosure, _ = cmd.Flags().GetBool("lpd")
		allChecks, _ := cmd.Flags().GetBool("all")
		if detectLoginPages || detectLicense || detectPathDisclosure {
			allChecks = false
		}
		if allChecks {
			detectLoginPages = true
			detectLicense = true
			detectPathDisclosure = true
		}

		target, _ := cmd.Flags().GetString("url")

		// Action for the lp flag.
		if detectLoginPages {
			pages, err := lp.Detect(target)
			if err != nil {
				fmt.Println(err)
				return
			}
			if len(pages) != 0 {
				fmt.Printf("Login Page Endpoints:\n")
				var c *color.Color
				for _, v := range pages {
					fmt.Printf("\r" + v[0] + " ")
					if strings.Contains(v[1], "200") {
						c = colors.OK
					} else if strings.Contains(v[1], "404") || strings.Contains(v[1], "403") ||
						strings.Contains(v[1], "500") {
						c = colors.BAD
					} else {
						c = colors.NEUTRAL
					}
					c.Printf(v[1] + "\n")
				}
			}
		}
		// Action for the license flag.
		if detectLicense {
			licenseExposed, err := license.Detect(target)
			if err != nil {
				fmt.Println(err)
				return
			}
			if licenseExposed != "" {
				colors.OK.Printf("\nLicense file is exposed!")
				fmt.Println(licenseExposed)
			} else {
				colors.BAD.Printf("\nNo license file is found.\n")
			}
		}
		// Action for the lpd flag.
		if detectPathDisclosure {
			pathDisclosure, err := lpd.Detect(target)
			if err != nil {
				fmt.Println(err)
				return
			}
			if pathDisclosure != nil {
				colors.OK.Printf("\nLinks that expose local paths:\n")
				for _, v := range pathDisclosure {
					fmt.Println(v)
				}
			} else {
				colors.BAD.Printf("\nNo local path disclosure found.\n")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(reconCmd)
	reconCmd.Flags().StringP("url", "u", "", "Target Bitrix site")
	reconCmd.MarkFlagRequired("url")
	reconCmd.Flags().Bool("lp", false, "Check possible login page endpoints")
	reconCmd.Flags().Bool("license", false, "Check if license file is exposed")
	reconCmd.Flags().Bool("lpd", false, "Check local path disclosure")
	reconCmd.Flags().BoolP("all", "a", true, "Run all checks")
}
