package middleware

import (
	controller "jokibro/controller"
	"jokibro/helper/messages"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustomClaims struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
	Role            string
}

func (jwtConf *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
		ErrorHandlerWithContext: middleware.JWTErrorHandlerWithContext(func(e error, c echo.Context) error {
			return controller.NewErrorResponse(c, http.StatusForbidden, e)
		}),
	}
}

// GenerateToken jwt ...
func (jwtConf *ConfigJWT) GenerateToken(userID int, role string) string {
	claims := JwtCustomClaims{
		userID,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConf.ExpiresDuration))).Unix(),
		},
	}

	// Create token with claims
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := t.SignedString([]byte(jwtConf.SecretJWT))

	return token
}

func (jwtConf *ConfigJWT) VerifyRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		claims := &JwtCustomClaims{}

		tokenAuthHeader := ctx.Request().Header.Get("Authorization")
		if !strings.Contains(tokenAuthHeader, "Bearer") {
			return controller.NewErrorResponse(ctx, http.StatusUnauthorized, messages.ErrInvalidBearerToken)
		}

		tokenAuth := strings.Replace(tokenAuthHeader, "Bearer ", "", -1)

		_, err = jwt.ParseWithClaims(tokenAuth, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtConf.SecretJWT), nil
		})
		if err != nil {
			return controller.NewErrorResponse(ctx, http.StatusUnauthorized, err)
		}

		if claims.ExpiresAt < time.Now().Unix() {
			return controller.NewErrorResponse(ctx, http.StatusUnauthorized, messages.ErrExpiredToken)
		}

		if claims.Role != jwtConf.Role {
			return controller.NewErrorResponse(ctx, http.StatusUnauthorized, messages.ErrInvalidRole)
		}

		ctx.Set(claims.Role, claims)

		return next(ctx)
	}
}

// GetUser from jwt ...
func GetCustomer(c echo.Context) (res *JwtCustomClaims) {
	user := c.Get("customer")
	if user != nil {
		res = user.(*JwtCustomClaims)
	}
	return res
}

// GetAdmin from jwt ...
func GetAdmin(c echo.Context) (res *JwtCustomClaims) {
	user := c.Get("admin")
	if user != nil {
		res = user.(*JwtCustomClaims)
	}
	return res
}
