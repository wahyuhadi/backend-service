package postgres

import (
	"backend-services/handlers/models/users"
	"context"
	"errors"

	"github.com/gin-gonic/gin"
)

// -- Get users
func (p *Postgres) GetUsersDetails(ctx context.Context, c *gin.Context) (user *users.Users, err error) {
	user_id := c.GetInt64("user_id")
	tx := p.db.WithContext(ctx).
		Select("users.id, username, email, role_id, phone, uuid, users.created_at, users.updated_at").
		Joins("Role").
		Where("users.id = ?", user_id).
		Find(&user)
	if tx.Error != nil { // -- if error
		return nil, tx.Error
	}

	return user, nil
}

// -- check user by username phone and  email
// -- if username , phone and email is exist with count >= 1
func (p *Postgres) CheckUsers(ctx context.Context, username, phone string, email *string) (bool, error) {
	var count int64
	err := p.db.WithContext(ctx).
		Model(&users.Users{}).
		Where("LOWER(username) = LOWER(?) OR LOWER(phone) = LOWER(?) OR LOWER(email) = LOWER(?)", username, phone, email).
		Count(&count).Error
	if err != nil {
		return true, err
	}
	// -- if count != 0 cannot create users
	if count != 0 {
		return true, err
	}
	return false, nil
}

// -- add users
func (p *Postgres) CreateUsers(ctx context.Context, users users.Users) (user users.Users, err error) {
	tx := p.db.WithContext(ctx).Create(&users)
	if tx.Error != nil {
		return user, errors.New("Error when add users.")
	}
	user.ID = users.ID
	return user, nil
}

// -- for check existing user
// -- by email, phone,  and username whis include in variable username
func (p *Postgres) GetUserToLogin(ctx context.Context, username string) (user users.Users, err error) {
	tx := p.db.WithContext(ctx).
		Where("email = ? OR phone = ? OR username = ?", username, username, username).
		First(&user)
	if tx.Error != nil {
		return user, errors.New("Error when get data  users.")
	}
	return user, nil
}
