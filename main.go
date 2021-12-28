/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/shohei/nba-cli/cmd"
	"github.com/shohei/nba-cli/config"
)

func main() {
	err := config.Load()
	if err != nil {
		panic(err)
	}

	cmd.Execute()
}
