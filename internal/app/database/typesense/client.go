package typesense

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/typesense/typesense-go/v3/typesense"
)

type Client struct {
	client *typesense.Client
}

func (c *Client) init() (err error) {
	ok, err := c.client.Health(context.Background(), 2*time.Second)
	if err != nil {
		return
	}

	if !ok {
		err = fmt.Errorf("failed connect to typesense client")
		return
	}

	log.Println("success connect to typesense client")

	err = c.initProductSchema()
	return
}

func NewClient() (c *Client, err error) {
	c = &Client{
		client: typesense.NewClient(
			typesense.WithNodes([]string{
				"http://localhost:8108",
			}),
			typesense.WithAPIKey("xyz"),
			typesense.WithConnectionTimeout(2*time.Second),
		),
	}

	c.init()

	return
}
