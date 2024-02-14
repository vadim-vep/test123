package main

import (
	"fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount"
	"os"
)

func main() {
	OfficialAccountApp, err := officialAccount.NewOfficialAccount(&officialAccount.UserConfig{
		AppID:     os.Getenv("AppId"), // 公众号、小程序的appid
		Secret:    os.Getenv("AppSecret"),
		Token:     os.Getenv("Token"),
		AESKey:    os.Getenv("AESKey"),
		Log:       officialAccount.Log{Level: "debug"},
		HttpDebug: true,
		Debug:     true})

	code := "TestCode"

	UserInterface, err := OfficialAccountApp.OAuth.UserFromCode(code)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(UserInterface)

}
