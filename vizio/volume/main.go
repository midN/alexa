package volume

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/midn/alexa/vizio/helpers"
)

type setVolumeData struct {
	Request string `json:"REQUEST"`
	HashVal int    `json:"HASHVAL"`
	Value   int    `json:"VALUE"`
}

type audioItem struct {
	HashVal int    `json:"HASHVAL"`
	CName   string `json:"CNAME"`
}

type getAudioData struct {
	Items []audioItem
}

func SetVolume(vol int) {
	hVal := getVolumeHashval()
	setVolume(hVal, vol)
}

func setVolume(hVal int, vol int) {
	volumeData := setVolumeData{Request: "MODIFY", HashVal: hVal, Value: vol}
	path := "/menu_native/dynamic/tv_settings/audio/volume"
	helpers.GenerateRequest("PUT", path, volumeData, false)
}

func getURL() string {
	return fmt.Sprintf("https://%s:%s/key_command/", os.Getenv("HOST"), os.Getenv("PORT"))
}

func getAUTH() string {
	return os.Getenv("AUTH")
}

func getVolumeHashval() int {
	_, body := helpers.GenerateRequest("GET", "menu_native/dynamic/tv_settings/audio", nil, true)

	getAudio := getAudioData{}
	jsonErr := json.Unmarshal(body, &getAudio)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	volumeHashVal := 0
	for _, item := range getAudio.Items {
		if item.CName == "volume" {
			volumeHashVal = item.HashVal
		}
	}

	return volumeHashVal
}
