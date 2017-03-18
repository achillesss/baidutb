package client

import (
	"time"

	"github.com/achillesss/baidutb/config"
	"github.com/achillesss/log"
)

// Start cycles a signing day by day
func Start(path string) {
	log.Parse()
	go autoReply(path)
	for {
		log.Infofln("Signing begins at %v.", time.Now())
		c := new(config.C)
		c.Decode(path)
		bdussChan := make(chan string)
		go transBdussChan(bdussChan, c.BdussList)
		countMap := signByBDUSS(bdussChan, c)
		signingCount(countMap)
		countDown()
		broadcast()
		now := time.Now()
		sleepTime := time.NewTimer(tomorrow(now).Sub(now)).C
		log.Infofln("tomorrow's signing begins at %v.", tomorrow(now))
		<-sleepTime
	}
}

func autoReply(path string) {
	for {
		c := new(config.C)
		c.Decode(path)
		if *debug {
			log.Infofln("config: %#v\n", c)
		}
		for _, bduss := range c.BdussList {
			getTopicList(bduss, c)
		}
		time.Sleep(time.Hour * 2)
	}
}
