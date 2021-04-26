package postgres

import (
	"backend-services/services/constant"
	"errors"
	"fmt"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

type Postgres struct {
	db *gorm.DB
}

func (p *Postgres) Connect(host, port, user, password, database string) error {

	cfg := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, database)

	var err error

	if p.db, err = gorm.Open(postgres.Open(cfg), &gorm.Config{}); err != nil {
		return errors.New(constant.ComposeMessage(constant.PostgreFailedConnect, err.Error()))
	}

	return nil
}
