package users

import (
	"backend-services/handlers"
	"backend-services/handlers/models/users"
	"backend-services/middleware"
	"backend-services/services"
	"backend-services/services/constant"
	"backend-services/services/wrapper"
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

type server struct {
	handlers.Handler
}

func (s *server) RegisUsers(c *gin.Context) {
	ctx := context.Background()
	var obj users.Users
	obj.RoleID = 2 // -- default role from users
	err := c.ShouldBind(&obj)
	// -- validate handling
	if err != nil {
		s.Log.Warnln(err.Error())
		c.JSON(wrapper.StatusBadRequest.New(constant.ApiInvalidJson, nil))
		return
	}
	// -- username cannot be empty
	if obj.Username == "" {
		c.JSON(wrapper.StatusBadRequest.New("Username cannot be empty", nil))
		return
	}

	// -- Phone cannot be empty
	if obj.Phone == "" {
		c.JSON(wrapper.StatusBadRequest.New("Phone cannot be empty", nil))
		return
	}

	// -- password cannot be empty
	if obj.Password == "" {
		c.JSON(wrapper.StatusBadRequest.New("Password cannot be empty", nil))
		return
	}

	// - Validate if useres is exist
	IsUserExist, err := s.Postgres.CheckUsers(ctx, obj.Username, obj.Phone, obj.Email)
	// -- sometime we dealing with error query
	if err != nil {
		c.JSON(wrapper.StatusInternalServerError.New("Error when validating users ", nil))
		return
	}
	// -- if users exist which mean cannot create users
	if IsUserExist {
		c.JSON(wrapper.StatusConflict.New("Username Or Phone Or Email  is already use.", nil))
		return
	}

	// -- use salt password
	password, salt := services.HashAndSalt(obj.Password)
	obj.Password = password
	obj.Salt = salt
	obj.CreatedAt = time.Now()
	obj.UpdatedAt = time.Now()

	users, err := s.Postgres.CreateUsers(ctx, obj)
	if err != nil {
		c.JSON(wrapper.StatusInternalServerError.New("Error when add users", err))
		return
	}

	obj.ID = users.ID
	c.JSON(wrapper.StatusOK.New(constant.ApiSuccess, obj))
	return
}

// - get func data userdetails
func (s *server) GetUsers(c *gin.Context) {
	ctx := context.Background()
	users, err := s.Postgres.GetUsersDetails(ctx, c)
	if err != nil {
		c.JSON(wrapper.StatusInternalServerError.New(constant.ErrorGetData, nil))
		return
	}
	c.JSON(wrapper.StatusOK.New(constant.ApiSuccess, users))
	return
}

func (s *server) DoLogin(c *gin.Context) {
	ctx := context.Background()
	var obj users.Login
	err := c.ShouldBind(&obj)
	// -- validate handling
	if err != nil {
		s.Log.Warnln(err.Error())
		c.JSON(wrapper.StatusBadRequest.New(constant.ApiInvalidJson, nil))
		return
	}
	// -- user cannot be empty
	if obj.User == "" {
		c.JSON(wrapper.StatusBadRequest.New("User cannot be empty", nil))
		return
	}
	// -- Password cannot be empty
	if obj.Password == "" {
		c.JSON(wrapper.StatusBadRequest.New("Password cannot be empty", nil))
		return
	}

	// -- Get users from databases
	usersData, err := s.Postgres.GetUserToLogin(ctx, obj.User)
	if err != nil {
		c.JSON(wrapper.StatusUnauthorized.New("Opps user not found", nil))
		return
	}
	// -- still validate users , if user dont have password will return users not found
	if usersData.Password == "" {
		c.JSON(wrapper.StatusUnauthorized.New("Opps user not found", nil))
		return
	}

	// -- check password is match
	isPassMatch := services.IsPasswordMatch(obj.Password, usersData.Password, usersData.Salt)
	// -- if password not macth
	if !isPassMatch {
		c.JSON(wrapper.StatusUnauthorized.New("Opps your password not match", nil))
		return
	}

	// -- validate role by ID
	role := "Admin"
	if usersData.RoleID == 2 {
		role = "User"
	}

	jwt := services.UserRoleJwt{
		UserID: int64(usersData.ID),
		RoleID: int64(usersData.RoleID),
	}

	token, exp := jwt.CreateToken(s.Env.JwtSecretKey)
	usrToken := users.UserToken{
		ID:      int64(usersData.ID),
		Role:    role,
		Token:   token,
		Expired: exp,
	}
	c.JSON(wrapper.StatusOK.New(constant.ApiSuccess, usrToken))
}

func Router(handler handlers.Handler, group string) {
	ref := &server{Handler: handler}
	authenticated := ref.Gin.Group(group)

	users := ref.Gin.Group(group)
	users.POST("/users/regis", ref.RegisUsers)

	login := ref.Gin.Group(group)
	login.POST("/login", ref.DoLogin)

	authenticated.Use(middleware.UserAuthenticationRequired(handler))
	{
		authenticated.GET("/users", ref.GetUsers)
	}
}
