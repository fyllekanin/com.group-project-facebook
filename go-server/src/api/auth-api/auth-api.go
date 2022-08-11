package auth_api

import (
	"encoding/json"
	"github.com/fyllekanin/go-server/src/app"
	"github.com/fyllekanin/go-server/src/common/error-interface"
	"github.com/fyllekanin/go-server/src/middlewares/authorization-middleware"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"time"
)

type UserApi struct {
	application *app.Application
}

func (api *UserApi) GetWhoAmI(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("rawr")
}

func (api *UserApi) PostLogin(w http.ResponseWriter, r *http.Request) {
	var payload LoginPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(error_interface.RestError{
			Message: "Something went wrong",
		})
		return
	}

	if payload.Username != "test" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(error_interface.RestError{
			Message: "Incorrect username or password",
		})
		return
	}

	if payload.Password != "test" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(error_interface.RestError{
			Message: "Incorrect username or password",
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
	return api
}
