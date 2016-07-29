package main

import (
	"fmt"
	"time"
)

type Weather struct {
	Description []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

type WeatherSimple struct {
	Weather

	Main struct {
		Temp     float32 `json:"temp"`
		Pressure float32 `json:"pressure"`
		Humidity float32 `json:"humidity"`
		TempMin  float32 `json:"temp_min"`
		TempMax  float32 `json:"temp_max"`
	} `json:"main"`

	Wind struct {
		Speed float32 `json:"speed"`
	} `json:"wind"`
}

type WeatherDetail struct {
	Weather

	Temp struct {
		Morn  float32 `json:"morn"`
		Day   float32 `json:"day"`
		Eve   float32 `json:"eve"`
		Night float32 `json:"night"`
	} `json:"temp"`

	Datetime int     `json:"dt"`
	Wind     float32 `json:"speed"`
	Pressure float32 `json:"pressure"`
}

type WeatherWeek struct {
	Weather []WeatherDetail `json:"list"`
}

func (w *Weather) String() string {
	return fmt.Sprintf("%s", w.Description[0].Description)
}

func (w *WeatherSimple) String() string {
	return w.Weather.String() + fmt.Sprintf("\n%d", int(w.Main.Temp))
}

func (w *WeatherDetail) String() string {
	return w.Weather.String() + fmt.Sprintf("\n%d", int(w.Temp.Day))
}

func (w *WeatherWeek) StringWeek() string {
	var result string
	for _, val := range w.Weather {
		t := time.Unix(int64(val.Datetime), 0)
		result += fmt.Sprintf("%02d.%02d\n", t.Day(), t.Month()) + val.String() + "\n\n"
	}
	return result
}

func (w *WeatherWeek) StringDetails() string {
	day := w.Weather[0]
	return day.Weather.String() + fmt.Sprintf("\nутро: %d\nдень: %d\nвечер: %d\nночь: %d\nдавление: %d", int(day.Temp.Morn), int(day.Temp.Day), int(day.Temp.Eve), int(day.Temp.Night), int(day.Pressure))
}

func (w *WeatherWeek) StringTomorow() string {
	day := w.Weather[1]
	return day.String()
}
