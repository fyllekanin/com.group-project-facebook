package auth_api

import (
	"encoding/json"
	"github.com/fyllekanin/go-server/src/app"
	"github.com/fyllekanin/go-server/src/common/error-interface"
	"github.com/fyllekanin/go-server/src/middlewares/authorization-middleware"
	user_repository "github.com/fyllekanin/go-server/src/repositories/user-repository"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strings"
	"time"
)

type UserApi struct {
	application *app.Application
}

func (userApi *UserApi) GetWhoAmI(w http.ResponseWriter, r *http.Request) {
	hashedTwo, _ := bcrypt.GenerateFromPassword([]byte("password"), 10)
	json.NewEncoder(w).Encode(string(hashedTwo))
}

func (userApi *UserApi) PostRegister(w http.ResponseWriter, r *http.Request) {
	var payload RegisterPayload
	var repository = user_repository.NewUserRepository(userApi.application.Db)
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(error_interface.RestError{
			Message: "Something went wrong",
		})
		return
	}

	isUsernameTaken, err := repository.DoUsernameExist(strings.ToLower(payload.Username))
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(error_interface.RestError{
			Message: "Something went wrong",
		})
		return
	}

	if isUsernameTaken {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(error_interface.RestError{
			Message: "Username is already taken",
		})
		return
	}

	json.NewEncoder(w).Encode(error_interface.RestError{
		Message: "Yeay",
	})
}

func (userApi *UserApi) PostLogin(w http.ResponseWriter, r *http.Request) {
	var payload LoginPayload
	var repository = user_repository.NewUserRepository(userApi.application.Db)

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(error_interface.RestError{
			Message: "Something went wrong",
		})
		return
	}

	user, err := repository.GetUserByUsername(strings.ToLower(payload.Username))
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(error_interface.RestError{
			Message: "Incorrect username and/or password",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(error_interface.RestError{
			Message: "Incorrect username and/or password",
		})
		return
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &authorization_middleware.AuthClaims{
		1,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Second)),
		},
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &authorization_middleware.AuthClaims{
		1,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(48 * time.Hour)),
		},
	})

	accessTokenString, err := accessToken.SignedString([]byte("AllYourBase"))
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(error_interface.RestError{
			Message: "Something went wrong",
		})
		return
	}

	refreshTokenString, err := refreshToken.SignedString([]byte("AllYourBase"))
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(error_interface.RestError{
			Message: "Something went wrong",
		})
		return
	}

	json.NewEncoder(w).Encode(LoginResponse{
		Username:     "test",
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	})
}

func GetApi(application *app.Application) *UserApi {
	var api = &UserApi{
		application: application,
	}
	var subRouter = application.Router.PathPrefix("/auth").Subrouter()
	var userRouter = application.Router.PathPrefix("/auth/user").Subrouter()
	userRouter.Use(authorization_middleware.Middleware)

	userRouter.HandleFunc("/whoami", api.GetWhoAmI).Methods("GET")
	subRouter.HandleFunc("/login", api.PostLogin).Methods("POST")
	subRouter.HandleFunc("/register", api.PostRegister).Methods("POST")
	return api
}
