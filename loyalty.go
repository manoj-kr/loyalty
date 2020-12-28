package loyalty

import (
	"fmt"

	"github.com/jinzhu/gorm/dialects/postgres"
)

type Loyalty struct {
	Channel string         `json:"channel"`
	Name    string         `json:"name"`
	Meta    postgres.Jsonb `json:"meta"`
	Data    postgres.Jsonb `json:"data"`
}

func init() {
	fmt.Println("Init")
	//TODO: Database tables creation
}

func (l *Loyalty) Publish() {
	fmt.Println("Publish")
}
