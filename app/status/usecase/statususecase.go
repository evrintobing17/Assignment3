package usecase

import (
	"Assignment3/app/model"
	"Assignment3/app/status"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

var startTime time.Time

type UC struct{}

func NewStatusUsecase() status.StatusUsecase {
	return &UC{}
}

func (u *UC) UpdateStatus() (model.StatusBahaya, error) {
	var data model.Data
	var statusData model.StatusBahaya
	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("failed to read json file, error: %v", err)
		return statusData, err
	}

	_ = json.Unmarshal(jsonData, &data)

	if uptime() > 15*time.Second {
		data.Data.Water = rand.Intn(100)
		data.Data.Wind = rand.Intn(100)
		jsonString, _ := json.Marshal(data)
		_ = os.WriteFile("data.json", jsonString, os.ModePerm)
		startTime = time.Now()
	}

	statusData.Water = data.Data.Water
	statusData.Wind = data.Data.Wind

	if data.Data.Water <= 5 {
		statusData.StatusAir = "Aman"
	} else if data.Data.Water > 5 && data.Data.Water <= 8 {
		statusData.StatusAir = "Siaga"
	} else {
		statusData.StatusAir = "Bahaya"
	}

	if data.Data.Wind <= 6 {
		statusData.StatusAngin = "Aman"
	} else if data.Data.Wind > 6 && data.Data.Wind <= 15 {
		statusData.StatusAngin = "Siaga"
	} else {
		statusData.StatusAngin = "Bahaya"
	}

	fmt.Println(statusData)

	return statusData, nil

}

func uptime() time.Duration {
	return time.Since(startTime)
}

func init() {
	startTime = time.Now()
}
