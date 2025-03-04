package middle

import (
	"back/biz/model"
	reqmodel "back/biz/reqModel"
	"context"
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
)

var identityKey = "token"

func GetJwt() (*jwt.HertzJWTMiddleware, error) {
	user := model.User{}
	temp, err := jwt.New(&jwt.HertzJWTMiddleware{
		Key:         []byte("yyym-Yuyu-SCNU-SCCNU-(UESTC)(SCU)"),
		Timeout:     time.Hour * 24 * 7,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginVals reqmodel.LoginReq
			if err := c.BindJSON(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			key := loginVals.Key
			password := loginVals.Password
			password = MD5(password)
			model.DB.Model(&model.User{}).Where("username = ?", key).Find(&user)

			if password == user.Password {
				return &user, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(code, FailWithMsg(message))
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			info := make(map[string]interface{})
			info["token"] = "Bearer " + token
			info["expire"] = expire.Format(time.RFC3339)
			info["username"] = user.Username
			info["type"] = user.Type

			c.JSON(http.StatusOK, SuccessWithDate(info))
		},
	})
	return temp, err
}
