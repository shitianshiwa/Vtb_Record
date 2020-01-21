package worker

import (
	"Vtb_Record/src/plugins/structUtils"
	"Vtb_Record/src/utils"
	"encoding/json"
	"log"
)

type UploadDict struct {
	Title       string
	Filename    string
	Date        string
	Path        string
	User        string
	OriginTitle string `json:"Origin_Title"`
	ASS         string
	Txt         string
}

func UploadVideo(video *structUtils.VideoInfo) {
	u := UploadDict{
		Title:       video.Title,
		Filename:    video.FileName,
		Date:        video.Date,
		Path:        video.FilePath,
		User:        video.UsersConfig.Name,
		OriginTitle: video.Title,
		ASS:         "",
		Txt:         "",
	}
	data, _ := json.Marshal(u)
	log.Println(string(data))
	utils.Publish(data, "upload")
}