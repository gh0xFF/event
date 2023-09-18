package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gh0xFF/event/internal/utils"
	"github.com/gh0xFF/event/pkg/eventservice"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if err := os.RemoveAll(filepath.Join(dir, "traindata")); err != nil {
		panic(err)
	}

	if err := os.Mkdir("traindata", 0777); err != nil {
		panic(err)
	}

	var fileCounter int

	// примеры чтобы сгенирировать датасет для создания словаря. Важно создать максимально разные, чтобы достичь хорошего сжатия
	for _, devOs := range []string{"IOS 13.5.1", "IOS 13.5.2", "IOS 13.5.3", "IOS 15.5.3", "IOS 16.4.5", "Android 4.4.4", "Android 2.3.3", "Android 3.0.1"} {
		for _, event := range []string{"app_start", "on_create", "on_error", "on_pause", "on_destroy", "on_resize", "hehe", "not_hehe"} {

			sequence := rand.Uint32()
			paramInt := rand.Uint32()

			monthOffset := rand.Int() % 12
			dayOffset := rand.Int() % 60
			// 2020-12-01 23:59:00

			fmt.Println("\n", fileCounter)
			tmtt := time.Now().AddDate(0, monthOffset, dayOffset).Format("2006-01-02 15:04:05")
			fmt.Printf("time=%s\n", tmtt)

			model := eventservice.EventModel{
				ClientTime: tmtt,            // must be unique to avoid using this alias in dict
				DeviceId:   utils.NewUUID(), // must be unique to avoid using this alias in dict
				DeviceOs:   devOs,
				Session:    randStringBytesRmndr(16), // must be unique to avoid using this alias in dict
				Event:      event,
				ParamStr:   randStringBytesRmndr(80), // must be unique to avoid using this alias in dict
				Sequence:   sequence,
				ParamInt:   paramInt,
			}
			data, err := json.Marshal(&model)
			if err != nil {
				panic(err)
			}

			fmt.Println(string(data))

			file, err := os.Create(filepath.Join(dir, "traindata", strconv.Itoa(fileCounter)+".txt"))
			if err != nil {
				panic(err)
			}

			if _, err := file.WriteString(string(data)); err != nil {
				panic(err)
			}

			file.Close()

			fileCounter++
		}
	}

	fmt.Printf("created %d files\n", fileCounter)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
