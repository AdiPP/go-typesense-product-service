package typesense

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/AdiPP/go-typesense-product-service/internal/app/config"
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

	log.Println("Success connect to typesense client")

	err = c.initProductSchema()
	return
}

func NewClient(cfg config.Typesense) (c *Client, err error) {
	c = &Client{
		client: typesense.NewClient(
			typesense.WithNodes([]string{
				fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
			}),
			typesense.WithAPIKey(cfg.APIKey),
			typesense.WithConnectionTimeout(2*time.Second),
		),
	}

	err = c.init()
	return
}
