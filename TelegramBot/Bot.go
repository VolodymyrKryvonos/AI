package main

import (
	"bytes"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/vova616/screenshot"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func DownloadFile(url string, filepath string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func makeScreenShot() {
	img, err := screenshot.CaptureScreen()
	if err != nil {
		panic(err)
	}
	f, err := os.Create("G:/Programming/Go/Projects/AI/TelegramBot/IMGs/f1.jpg")
	if err != nil {
		panic(err)
	}
	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
	f.Close()
}

func main() {

	bot, err := tgbotapi.NewBotAPI("902103535:AAE72ua0d9xmaOiNSLM9VTfOe9CFvOniNEk")
	if err != nil {
		log.Panic(err)
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.Photo != nil {
			fileID := (*update.Message.Photo)[1].FileID
			s, e := bot.GetFileDirectURL(fileID)
			if e == nil {
				DownloadFile(s, "G:/Programming/Go/Projects/AI/TelegramBot/IMGs/f.jpg")
			}
			makeScreenShot()
			photo, err := ioutil.ReadFile("G:/Programming/Go/Projects/AI/TelegramBot/IMGs/f1.jpg")
			if err == nil {
				reder := tgbotapi.FileReader{
					Name:   "f1.jpg",
					Reader: bytes.NewReader(photo),
					Size:   -1,
				}
				msg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, reder)
				bot.Send(msg)
			}
		}
	}
}
