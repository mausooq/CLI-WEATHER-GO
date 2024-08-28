package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/fatih/color"
)


type Weather struct {
	Location struct{
		City string `json:"name"`
		Country string `json:"country"`
		}`json:"location"`
	Current struct {
		TempC float64 `json:"temp_c"`
		Condition struct{
			Text string `json:"text"`
		}`json:"condition"`
	}`json:"current"`
	Forecast struct {
		Forecastday []struct{
			Hour []struct {
			TimeEpoch int64 `json:"time_epoch"`
			TempC float64 `json:"temp_c"`
			ChanceOfRain float64 `json:"chance_of_rain"`
			Condition struct{
				Text string `json:"text"`
			}`json:"condition"`
			}`json:"hour"`
		}`json:"forecastday"`
	}`json:"forecast"`

}


func main(){
	fmt.Println("ready to go")
	 
	res , err := http.Get("https://api.weatherapi.com/v1/forecast.json?key=791a3220ed4d4099a74161535242207&q=Mangaluru&aqi=no")

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Api Not Available")
	}
	

	Body , err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var weather Weather
	err = json.Unmarshal(Body , &weather)
	if err != nil {
		panic(err)
	}
	location ,current, hours := weather.Location, weather.Current,weather.Forecast.Forecastday[0].Hour
	fmt.Printf("%s ,%s : %.0fC ,%s\n",
	location.City,
	location.Country,
	current.TempC,
	current.Condition.Text)

	for _,hour := range hours{
		date := time.Unix(hour.TimeEpoch,0)

		if date.Before(time.Now()){
			continue
		}
		message :=fmt.Sprintf("%s - %.0fC , %.0f%%, %s\n",
	date.Format("15:04"),
	hour.TempC,
	hour.ChanceOfRain,
	hour.Condition.Text)

	if hour.ChanceOfRain < 50{
		color.Green(message)
	}else{
		color.Red(message)
	}
	}

}