package cmd

import (
	"github.com/rchaganti/dadjoke-tui/tui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dadjoke-tui",
	Short: "A terminal UI for dadjokes",
	Long:  `dadjoke-tui is a terminal UI for dadjokes. It is built using bubbletea and lipgloss.`,
	Run: func(cmd *cobra.Command, args []string) {
		p := tui.NewApp(searchTerm, limit)
		_, err := p.Run()
		if err != nil {
			panic(err)
		}
	},
}

var (
	searchTerm string
	limit      int
)

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.Flags().StringVarP(&searchTerm, "search", "s", "", "Search term for dadjokes")
	rootCmd.Flags().IntVarP(&limit, "limit", "l", 5, "Number of dadjokes to display")
}
