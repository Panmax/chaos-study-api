package users

import (
	"github.com/Panmax/chaos-study-api/common"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"time"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func NewGinJWTMiddleware() (*jwt.GinJWTMiddleware, error) {
	middleware := jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     8 * time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: common.JWTIdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*UserModel); ok {
				return jwt.MapClaims{
					common.JWTIdentityKey: v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)

			userModel, err := FindOneUser(claims[common.JWTIdentityKey].(string))
			if err != nil {
				return nil
			}
			c.Set(common.UserIDKey, userModel.ID)
			return &userModel
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVal login
			if err := c.ShouldBind(&loginVal); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			username := loginVal.Username
			password := loginVal.Password

			user, err := FindOneUser(username)
			if err == nil && user.checkPassword(password) == nil {
				return &user, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*UserModel); ok {
				return true
			}

			return false
		},
	}
	return jwt.New(&middleware)
}
