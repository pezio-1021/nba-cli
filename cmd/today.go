package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
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
	todayGames, err := client.GetDateGames(ctx, n.Format("2006-01-02"))
	if err != nil {
		fmt.Println(err)
		fmt.Println("here")
		os.Exit(1)
	}

	for _, v := range todayGames.API.Games {

		// 試合がまだ開催されていない場合はそのメッセージを表示する
		if v.StatusGame == "Scheduled" {
			fmt.Println(v.VTeam.ShortName + " vs " + v.HTeam.ShortName)
			fmt.Println("Game is Scheduled")
			continue
		}

		winT := map[string]string{"name": v.VTeam.ShortName, "score": v.VTeam.Score.Points}
		loseT := map[string]string{"name": v.HTeam.ShortName, "score": v.HTeam.Score.Points}
		if v.VTeam.Score.Points < v.HTeam.Score.Points {
			winT = map[string]string{"name": v.HTeam.ShortName, "score": v.HTeam.Score.Points}
			loseT = map[string]string{"name": v.VTeam.ShortName, "score": v.VTeam.Score.Points}
		}

		data := [][]string{
			[]string{"Victory team", winT["name"]},
			[]string{"Losing team", loseT["name"]},
			[]string{"Score", winT["score"] + "-" + loseT["score"]},
		}

		table := tablewriter.NewWriter(os.Stdout)

		for _, v := range data {
			table.Append(v)
		}
		table.Render()

		b, err := ioutil.ReadFile("img/" + winT["name"] + ".txt")
		if err != nil {
			fmt.Println("Sorry, No Image")
			continue
		}
		lines := string(b)
		color.Blue(lines)
	}

}
