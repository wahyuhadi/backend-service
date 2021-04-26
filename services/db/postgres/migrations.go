package postgres

import (
	"backend-services/handlers/models/users"
	"backend-services/services"
)

func (p *Postgres) ModelMigrate() error {
	err := p.db.AutoMigrate(
		&users.Role{},
		&users.Users{},
	)
	return err
}

func (p *Postgres) Extensi() error {
	p.db.Raw(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Rows()
	return nil
}

func (p *Postgres) roleInit() error {
	roles := []users.Role{
		{ID: 1, Role: "Admin"},
		{ID: 2, Role: "User"},
	}

	for _, item := range roles {
		var count int64
		err := p.db.Model(&users.Role{}).Where("LOWER(role) = LOWER(?)", item.Role).Count(&count).Error
		if err != nil {
			return err
		}

		if count != 0 {
			continue
		}

		err = p.db.Save(&item).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Postgres) userInit() error {
	password, salt := services.HashAndSalt("wahyu")
	roles := []users.Users{
		{ID: 1, Username: "wahyu", Phone: "6285205039835", RoleID: 1, Password: password, Salt: salt},
		{ID: 2, Username: "user", Phone: "6285205039834", RoleID: 2, Password: password, Salt: salt},
	}
	for _, item := range roles {
		var count int64
		err := p.db.Model(&users.Users{}).Where("LOWER(username) = LOWER(?)", item.Username).Count(&count).Error
		if err != nil {
			return err
		}
		if count != 0 {
			continue
		}
		err = p.db.Save(&item).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Postgres) ModelInit() error {

	// -- user role seeders
	if err := p.roleInit(); err != nil {
		return err
	}

	// -- user admin seeders
	if err := p.userInit(); err != nil {
		return err
	}

	return nil
}
