package main

import(
	"fmt"
	"os"
	"github.com/slack-go/slack"
)

func main(){

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6520678253792-6835723923234-eWNxlyMdFWyfJL1y7zpRIdBT")
	os.Setenv("CHANNEL_ID","C06QGQ8V9RB")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"faizal.pdf", "khalith.pdf"}

	for i:=0; i<len(fileArr); i++{
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File: fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil{
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)
	}
}