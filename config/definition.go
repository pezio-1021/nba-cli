package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Load() error {
	viper.SetConfigName("config")  // 設定ファイル名を指定
	viper.SetConfigType("yaml")    // 設定ファイルの形式を指定
	viper.AddConfigPath("config/") // ファイルのpathを指定

	err := viper.ReadInConfig() // 設定ファイルを探索して読み取る
	if err != nil {
		return fmt.Errorf("設定ファイル読み込みエラー: %s", err)
	}

	return nil
}
