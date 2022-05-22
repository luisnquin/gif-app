package auth

import (
	"crypto/rsa"
	"database/sql"
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/luisnquin/meow-app/src/server/config"
	"github.com/luisnquin/meow-app/src/server/log"
	"github.com/luisnquin/meow-app/src/server/models"
	userProxy "github.com/luisnquin/meow-app/src/server/provider/user"
	"github.com/luisnquin/meow-app/src/server/utils"
	"golang.org/x/crypto/bcrypt"
)

var (
	Config middleware.JWTConfig

	privateKey *rsa.PrivateKey
)

const (
	AdminRole       string = "PRIVILEGED"
	UserDefaultRole string = "USER"
)

func init() {
	privCont, err := ioutil.ReadFile("./private.rsa.key")
	if err != nil {
		panic(err)
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privCont)
	if err != nil {
		panic(err)
	}

	pubCont, err := ioutil.ReadFile("./public.rsa.key")
	if err != nil {
		panic(err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(pubCont)
	if err != nil {
		panic(err)
	}

	Config = middleware.JWTConfig{
		Claims:        &models.Claims{},
		SigningKey:    privateKey,
		SigningMethod: jwt.SigningMethodRS256.Alg(),
		KeyFunc: func(token *jwt.Token) (interface{}, error) {
			return publicKey, nil
		},
	}
}

func LoginHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userReq models.User

		err := c.Bind(&userReq)
		if err != nil {
			return echo.ErrUnauthorized
		}

		user, err := userProxy.GetByEmailOrUsername(userReq.Username, userReq.Email)
		if err != nil {
			log.Error(err)

			if errors.Is(err, sql.ErrNoRows) {
				return echo.ErrUnauthorized
			}

			return echo.ErrInternalServerError
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userReq.Password))
		if err != nil {
			log.Error(err)

			return echo.ErrUnauthorized
		}

		user.Password = ""

		token, err := genSignedJWTToken(user)
		if err != nil {
			log.Error(err)

			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, models.TokenResponse{
			Token: token,
		})
	}
}

func RegisterHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var user models.User

		err := c.Bind(&user)
		if err != nil {
			return echo.ErrBadRequest
		}

		if user.Username == "" || user.Firstname == "" || user.Lastname == "" || user.Email == "" || user.Password == "" {
			return echo.ErrBadRequest
		}

		if !emailIsValid(user.Email) {
			return echo.ErrBadRequest
		}

		exists, err := userProxy.UsernameOrEmailExists(user.Username, user.Email)
		if err != nil {
			return echo.ErrInternalServerError
		}

		if exists {
			return echo.ErrBadRequest
		}

		password, err := utils.GenHashedPassword(user.Password)
		if err != nil {
			return echo.ErrInternalServerError
		}

		user.Role = UserDefaultRole
		user.Password = password
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()

		// Add email verification.

		err = userProxy.Save(user)
		if err != nil {
			log.Error(err)

			return echo.ErrInternalServerError
		}

		return c.NoContent(http.StatusOK)
	}
}

func emailIsValid(email string) bool {
	rexp := regexp.MustCompile(config.Server.Internal.EmailRegex)
	if !rexp.MatchString(email) {
		return false
	}

	return true
}
