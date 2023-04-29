package model

type Hourly struct {
  Time []string `json:"time"`
  Cloudcover []int8 `json:"cloudcover"`
}

type DataPoint struct {
  Time string
  CloudCover int8
}
