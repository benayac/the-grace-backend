package middleware

import (
	"encoding/json"
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strings"
	"thegrace/pkg"
	"time"
)

const (
	iss       = "the.grace"
	audUser   = "user"
	audAdmin  = "admin"
	KeyClient = "client"
)

func GetJWTUser(client string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims[KeyClient] = client
	claims["aud"] = audUser
	claims["iss"] = iss
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString([]byte(pkg.Conf.SigningKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GetJWTAdmin(client string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims[KeyClient] = client
	claims["aud"] = audAdmin
	claims["iss"] = iss
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString([]byte(pkg.Conf.SigningKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func IsAuthorizedUser(endpoint func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("[AUTHENTICATION USER][REQUEST]")
		if r.Header["Authorization"] != nil {
			valid, err := verifyAuth(r.Header, audUser)
			if err != nil {
				log.Println("[AUTHENTICATION USER][ERROR] PARSING HEADER: ", err.Error())
				res := DefaultResponse{Status: false, Error: err.Error()}
				json.NewEncoder(w).Encode(&res)
			}
			if valid {
				endpoint(w, r)
			}
		} else {
			res := DefaultResponse{Status: false, Error: "Invalid authorization"}
			json.NewEncoder(w).Encode(&res)
		}
	}
}

func IsAuthorizedAdmin(endpoint func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("[AUTHENTICATION ADMIN][REQUEST]")
		if r.Header["Authorization"] != nil {
			valid, err := verifyAuth(r.Header, audAdmin)
			if err != nil {
				log.Println("[AUTHENTICATION ADMIN][ERROR] PARSING HEADER: ", err.Error())
				res := DefaultResponse{Status: false, Error: err.Error()}
				json.NewEncoder(w).Encode(&res)
			}
			if valid {
				endpoint(w, r)
			}
		} else {
			res := DefaultResponse{Status: false, Error: "Invalid authorization"}
			json.NewEncoder(w).Encode(&res)
		}
	}
}

func ParseAuth(header map[string][]string, key string) (string, error) {
	bearer := strings.Split(header["Authorization"][0], "Bearer ")[1]
	bearer = strings.TrimSpace(bearer)
	token, err := jwt.Parse(bearer, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid Signing Method")
		}
		if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
			return nil, errors.New("expired Token")
		}
		checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(audUser, false)
		if !checkAudience {
			return nil, errors.New("invalid aud")
		}
		iss := "the.grace"
		checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
		if !checkIss {
			return nil, errors.New("invalid iss")
		}
		return []byte(pkg.Conf.SigningKey), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		data, ok := claims[key].(string)
		if !ok {
			return "", err
		}
		return data, nil
	}
	return "", err
}

func verifyAuth(header map[string][]string, aud string) (bool, error) {
	bearer := strings.Split(header["Authorization"][0], "Bearer ")[1]
	bearer = strings.TrimSpace(bearer)
	token, err := jwt.Parse(bearer, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid Signing Method")
		}
		if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
			return nil, errors.New("expired Token")
		}
		checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
		if !checkAudience {
			return nil, errors.New("invalid aud")
		}
		iss := "the.grace"
		checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
		if !checkIss {
			return nil, errors.New("invalid iss")
		}
		return []byte(pkg.Conf.SigningKey), nil
	})
	if err != nil {
		return false, err
	}
	if !token.Valid {
		return false, nil
	}
	return true, nil
}
