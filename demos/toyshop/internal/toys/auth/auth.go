package auth

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/KernelGamut32/golang-microservices/demos/toyshop/internal/users"
	"github.com/spf13/viper"
)

type AuthVerifier struct{}

var verifier *AuthVerifier

func GetAuthVerifier() *AuthVerifier {
	if verifier == nil {
		verifier = &AuthVerifier{}
	}
	return verifier
}

func getAuthConfig() string {
	dir, _ := os.Getwd()
	viper.SetConfigName("app")
	viper.AddConfigPath(dir + "/../configs")
	viper.AutomaticEnv()

	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	return viper.GetString("AUTH_ENDPOINT")
}

func (authVerifier AuthVerifier) VerifyAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authEndpoint := getAuthConfig()

		client := &http.Client{
			Timeout: time.Second * 10,
		}
		req, err := http.NewRequest("GET", authEndpoint, nil)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		req.Header.Set("x-access-token", r.Header.Get("x-access-token"))
		response, err := client.Do(req)
		if response.StatusCode != http.StatusOK || err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		defer response.Body.Close()

		usr := users.User{}
		json.NewDecoder(response.Body).Decode(&usr)

		ctx := context.WithValue(r.Context(), "user", usr)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
