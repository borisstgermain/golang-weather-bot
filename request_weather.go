package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ReqWeather struct {
	ApiKey, ReqCurrent, ReqDetail, req string
}

func (r *ReqWeather) Init(apiKey string) {
	r.ApiKey = apiKey
	r.req = "&units=metric" + "&lang=ru" + "&appid=" + r.ApiKey
	r.ReqCurrent = "http://api.openweathermap.org/data/2.5/weather?q="
	r.ReqDetail = "http://api.openweathermap.org/data/2.5/forecast/daily?q="

}

func (r *ReqWeather) GetWeather(option []string) string {
	var req string

	if len(option) < 2 {
		req = r.getWeatherCurrent(option[0])
		w := WeatherSimple{}
		res := request(req)

		err := json.Unmarshal(res, &w)
		if err != nil {
			fmt.Println(err)
		}

		return w.String()

	} else {
		req = r.getWeatherDetail(option[0])
		w := WeatherWeek{}
		res := request(req)

		err := json.Unmarshal(res, &w)
		if err != nil {
			fmt.Println(err)
		}

		switch option[1] {
		case "неделя":
			return w.StringWeek()
		case "подробно":
			return w.StringDetails()
		case "завтра":
			return w.StringTomorow()
		default:
			return w.StringWeek()
		}
	}
}

func (r *ReqWeather) getWeatherCurrent(city string) string {
	req := r.ReqCurrent + city + r.req
	return req
}

func (r *ReqWeather) getWeatherDetail(city string) string {
	req := r.ReqDetail + city + r.req
	return req
}

func request(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
	}

	robots, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	return robots
}
