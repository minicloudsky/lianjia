package biz

import (
	"math/rand"
	"time"
)

type UserAgent string

var userAgents = [10]UserAgent{
	// mobile phone
	"Mozilla/5.0 (iPhone; CPU iPhone OS 15_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 Safari/604.1",
	"Mozilla/5.0 (Android 12; Mobile; rv:100.0) Gecko/100.0 Firefox/100.0",
	"Mozilla/5.0 (Windows Phone 10; Android 4.2; Mobile) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Mobile Safari/537.36",
	// iPad
	"Mozilla/5.0 (iPad; CPU OS 15_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 Safari/604.1",
	"Mozilla/5.0 (iPad; CPU OS 15_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.0 Mobile/15E148 Safari/604.1",
	// computer
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Firefox/100.0",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36",
	// laptop
	"Mozilla/5.0 (Android 12; Tablet; rv:100.0) Gecko/100.0 Firefox/100.0",
	"Mozilla/5.0 (Android 12; Tablet; rv:100.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36",
}

func RandomUA() string {
	rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	randIndex := rand.Intn(len(userAgents))
	return string(userAgents[randIndex])
}
