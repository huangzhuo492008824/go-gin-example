package main

import (
	"log"
	"time"

	"github.com/robfig/cron"
)

func main() {
	log.Println("Starting...")

	c := cron.New()
	c.AddFunc("5-59/15 * * * * *", func() {
		log.Println("Run models.CleanAllTag... 15 minutes")
		// models.CleanAllTag()
	})
	c.AddFunc("6-59/3 * * * * *", func() {
		log.Println("Run models.CleanAllArticle... 3 minutes")
		// models.CleanAllArticle()
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
