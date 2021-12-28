package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/shohei/nba-cli/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// helloCmd represents the hello command
var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Get Today Games Results",
	Long:  `sample`,
	Run:   GetTodayGames,
}

func init() {
	rootCmd.AddCommand(todayCmd)
}

func GetTodayGames(cmd *cobra.Command, args []string) {
	client, err := api.New(viper.GetString("api.key"), nil)
	if err != nil {
		os.Exit(1)
	}
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	n := time.Now()

	req, err := client.GetDateGames(ctx, n.Format("2006-01-02"))
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(req)
}
