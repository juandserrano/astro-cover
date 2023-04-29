package net

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/juandserrano/astro-cover/model"
)


const BASE_URL = "https://api.open-meteo.com/v1/forecast"

var (
  latitude string = "45.41";
  longitude string = "-75.73";
  timezone string = "America%2FNew_York";
)

func FetchData() model.Hourly {
  url := fmt.Sprintf("%s?latitude=%s&longitude=%s&hourly=cloudcover&daily=sunrise,sunset&forecast_days=3&timezone=%s", 
    BASE_URL, latitude, longitude, timezone) 
  res, err := http.Get(url)
  if err != nil {
    log.Printf("GET Error: %s", err)
    return model.Hourly{}
  }
  defer res.Body.Close()
  body, err := io.ReadAll(res.Body)
  if err != nil {
    log.Printf("io.ReadAll error: %s", err)
    return model.Hourly{}
  }

  var result map[string]any
  json.Unmarshal(body, &result)

  newjson, err := json.Marshal(result["hourly"])
  if err != nil {
    log.Printf("json.Marshal error: %s", err)
    return model.Hourly{}
  }
  var cover model.Hourly
  json.Unmarshal(newjson, &cover)

  return cover
}
