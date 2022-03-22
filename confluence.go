package confluence

import (
	"github.com/virtomize/confluence-go-api"
	"log"
)

type confluence struct {
	client *goconfluence.API
}

func NewConfluence(domain string, username string, token string) *confluence {
	client, err := goconfluence.NewAPI(domain, username, token)
	errHandler(err)
	return &confluence{
		client: client,
	}
}

func errHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
