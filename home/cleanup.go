package home

import (
	"log"
	"time"

	"github.com/gibson/gophercraft/home/models"
)

func (hs *Server) cleanup() {
	for {
		deleted, err := hs.DB.Where("expiry < ?", time.Now()).Delete(new(models.WebToken))
		if err != nil {
			log.Fatal(err)
		}

		if deleted > 0 {
			log.Println("Wiped", deleted, "expired tokens")
		}

		time.Sleep(20 * time.Minute)
	}
}
