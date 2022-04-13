package cmd

import (
	"github.com/brokeyourbike/macadmin.go/helpers"
	"github.com/spf13/cobra"
)

var feedUrl string

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Parses Apple's feed of macOS IPSWs and lets you download one.",
	Long:  `Parses Apple's feed of macOS IPSWs and lets you download one.`,
	Run: func(cmd *cobra.Command, args []string) {
		ipfs, err := helpers.FetchIpswData(feedUrl)

		if err != nil {
			panic(err)
		}
	},
}

func init() {
	getCmd.Flags().StringVarP(&feedUrl, "feed_url", "f", "https://mesu.apple.com/assets/macos/com_apple_macOSIPSW/com_apple_macOSIPSW.xml", "Url to Apple's feed of macOS IPSWs")
}
