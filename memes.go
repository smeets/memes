package memes

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
)

type Meme struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

type Data struct {
	Memes []Meme `json:"memes"`
}

type Response struct {
	Success bool `json:"success"`
	Data Data `json:"data"`
}

func GetMemes() ([]Meme, error) {
	url := "https://api.imgflip.com/get_memes"

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := Response{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return response.Data.Memes, nil
}