package api

import (
	"context"
	"net/http"
	"time"
)

type DateGames struct {
	API struct {
		Status  int      `json:"status"`
		Message string   `json:"message"`
		Results int      `json:"results"`
		Filters []string `json:"filters"`
		Games   []struct {
			SeasonYear      string    `json:"seasonYear"`
			League          string    `json:"league"`
			GameID          string    `json:"gameId"`
			StartTimeUTC    time.Time `json:"startTimeUTC"`
			EndTimeUTC      time.Time `json:"endTimeUTC"`
			Arena           string    `json:"arena"`
			City            string    `json:"city"`
			Country         string    `json:"country"`
			Clock           string    `json:"clock"`
			GameDuration    string    `json:"gameDuration"`
			CurrentPeriod   string    `json:"currentPeriod"`
			Halftime        string    `json:"halftime"`
			EndOfPeriod     string    `json:"EndOfPeriod"`
			SeasonStage     string    `json:"seasonStage"`
			StatusShortGame string    `json:"statusShortGame"`
			StatusGame      string    `json:"statusGame"`
			VTeam           struct {
				TeamID    string `json:"teamId"`
				ShortName string `json:"shortName"`
				FullName  string `json:"fullName"`
				NickName  string `json:"nickName"`
				Logo      string `json:"logo"`
				Score     struct {
					Points string `json:"points"`
				} `json:"score"`
			} `json:"vTeam"`
			HTeam struct {
				TeamID    string `json:"teamId"`
				ShortName string `json:"shortName"`
				FullName  string `json:"fullName"`
				NickName  string `json:"nickName"`
				Logo      string `json:"logo"`
				Score     struct {
					Points string `json:"points"`
				} `json:"score"`
			} `json:"hTeam"`
		} `json:"games"`
	} `json:"api"`
}

func (c *Client) GetDateGames(ctx context.Context, date string) (*interface{}, error) {
	relativePath := "games/date/"
	games := new(DateGames)
	req, err := c.GetRequestResult(ctx, http.MethodGet, relativePath, date, games)
	if err != nil {
		return nil, err
	}

	return &req, err
}
