package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/juandserrano/astro-cover/model"
	"github.com/juandserrano/astro-cover/net"
)

var ccLimit int8 = 30

var sleepTime time.Duration = time.Second * 60
var loc *time.Location

func main()  {
  var lastRun time.Time
  loc, _ = time.LoadLocation("America/Toronto")

  email := os.Getenv("GMAIL_ADDRESS")
  appPass := os.Getenv("GMAIL_ASTRO_COVER_APP_PASS")
  if email == "" {
    log.Printf("GMAIL_ADDRESS env must be set")
    log.Printf("Exiting...")
    os.Exit(1)
  }
  if appPass == "" {
    log.Printf("GMAIL_ASTRO_COVER_APP_PASS env must be set")
    log.Printf("Exiting...")
    os.Exit(1)
  }

  for {
    now := time.Now().In(loc)
      log.Printf("Checking on time %s", now)

    if now.Hour() == 19 && now.Sub(lastRun).Minutes() >= 100 {
      run()
      lastRun = time.Now().In(loc)
    }
    time.Sleep(sleepTime)
  }
}

func run() {
  cover := net.FetchData()
  if cover.Cloudcover != nil {
    var notification model.Notification
    notification.Day = time.Now().In(loc).Format(time.RFC850)
    if checkCloudCoverAtNight(&cover, &notification) {
      notification.Result = "ASTRO-COVER: Hoy es un gran dia!"
      sendNotification("OK", &notification)
    } else {
      notification.Result = "ASTRO-COVER: No es un buen dia para astrophotography :("
      sendNotification("Not OK", &notification)
    }
  }
}

func checkCloudCoverAtNight(cc *model.Hourly, notification *model.Notification) bool {
  arr := []int8{21, 22, 23, 24, 25, 26, 27}
  isGood := false
  for _, v := range arr {
    time := cc.Time[v]
    percentage := cc.Cloudcover[v]
    dataPoint := model.DataPoint {
      Time: time,
      CloudCover: percentage,
    }
    notification.Data = append(notification.Data, dataPoint)
    if percentage <= ccLimit {
      isGood = true
    } 
  }
  return isGood
}

func sendNotification(s string, n *model.Notification)  {
  if s == "OK" {
    fmt.Printf("Day: %s,\nResult: %s,\n%+v\n\n\n", n.Day, n.Result, n.Data)
    net.SendEmailNotification(n)
    

  } else {
    fmt.Printf("Day: %s,\nResult: %s,\n%+v\n\n\n", n.Day, n.Result, n.Data)
    net.SendEmailNotification(n)
  }
}
