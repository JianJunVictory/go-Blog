package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-Blog/db"

	"github.com/go-Blog/model"
	"github.com/gorilla/mux"
)

// ResponseWithJSON return json data
func ResponseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// ValidateTokenMiddleware check token value
func ValidateTokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte("secret"), nil
				})
				if error != nil {
					ResponseWithJSON(w, http.StatusOK, model.Response{
						Code:    -2001,
						Message: error.Error(),
					})
					return
				}
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					idStr := fmt.Sprintf("%v", claims["id"])
					// req.Header.Set("uId", idStr)
					ctx := context.WithValue(req.Context(), "uId", idStr)
					req = req.WithContext(ctx)
					next(w, req)
				} else {
					ResponseWithJSON(w, http.StatusOK, model.Response{
						Code:    -2001,
						Message: "Invalid authorization token",
					})
				}

			} else {
				ResponseWithJSON(w, http.StatusOK, model.Response{
					Code:    -2001,
					Message: "Invalid authorization token",
				})
			}
		} else {
			ResponseWithJSON(w, http.StatusOK, model.Response{
				Code:    -2001,
				Message: "token not exist",
			})
		}
	}
}

// GetName test get name handler
func GetName(w http.ResponseWriter, r *http.Request) {
	var godusers []model.Goduser
	goduser := model.Goduser{}

	stmt, _ := db.DB.Prepare(`SELECT * FROM  customer WHERE AGE <?`)
	defer stmt.Close()
	rows, err := stmt.Query(33)
	if err != nil {
		log.Printf("select data error: %v\n", err.Error())
	}
	for rows.Next() {
		rows.Scan(&goduser.CUST_ID, &goduser.NAME, &goduser.AGE)
		godusers = append(godusers, goduser)
	}
	ResponseWithJSON(w, http.StatusOK, model.Response{
		Code:    0,
		Message: "ok",
		Data:    godusers,
	})
}

// GetAge test get age handler
func GetAge(w http.ResponseWriter, r *http.Request) {
	age := mux.Vars(r)["age"]
	// w.Header().Set("Content-Type","application/json;charset=UTF-8")
	w.Write([]byte("age is " + age))
}
