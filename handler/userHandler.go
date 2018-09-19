package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-Blog/db"

	"github.com/go-Blog/model"
	"github.com/go-Blog/utils"
)

// Register user register
func Register(w http.ResponseWriter, r *http.Request) {

	user := new(model.User)
	json.NewDecoder(r.Body).Decode(user)
	log.Printf("request user info:%#v", user)
	// check email is nil
	if user.Email == "" || user.Password == "" {
		ResponseWithJSON(w, http.StatusOK, model.Response{
			Code:    -1,
			Message: "email and password is nil",
		})
		return
	}
	// check email format
	if !utils.CheckEmail(user.Email) {
		ResponseWithJSON(w, http.StatusOK, model.Response{
			Code:    -1,
			Message: "Email format error",
		})
		return
	}
	// select user to check user existed
	user1 := model.User{}
	stmt, _ := db.DB.Prepare(`SELECT id,email,password,status From user where email = ?`)
	defer stmt.Close()
	log.Println(user.Email)
	row := stmt.QueryRow(user.Email)
	row.Scan(&user1.ID, &user1.Email, &user1.Password, &user1.Status)

	if user1.ID != 0 && user1.Email != "" && user1.Password != "" {
		// Status 0:unactive user,1:active user
		if user1.Status != 0 {
			ResponseWithJSON(w, http.StatusOK, model.Response{
				Code:    -1,
				Message: "user existed",
			})
			return
		}
		checkToken, _ := utils.CreateTokenEndpoint(int(user1.ID))
		err := utils.SendEmail(user1.Email, checkToken)
		// send email failed
		if err != nil {
			ResponseWithJSON(w, http.StatusOK, model.Response{
				Code:    -1,
				Message: err.Error(),
			})
			return
		}
		// send email OK
		ResponseWithJSON(w, http.StatusOK, model.Response{
			Code:    0,
			Message: "resend active email,please to active",
		})
		return
	}

	// create user
	stmt, _ = db.DB.Prepare(`INSERT INTO user (email, password) Values (?,?)`)
	defer stmt.Close()

	// crypto password
	cryptoPassword := utils.CryptoPassword(user.Password)
	result, err := stmt.Exec(user.Email, cryptoPassword)
	if err != nil {
		ResponseWithJSON(w, http.StatusOK, model.Response{
			Code:    -1,
			Message: "create user fail",
		})
		return
	}
	LastInsertId, err := result.LastInsertId()
	if nil != err {
		log.Println(err)
	}
	checkToken, _ := utils.CreateTokenEndpoint(int(LastInsertId))
	err = utils.SendEmail(user.Email, checkToken)
	if err != nil {
		ResponseWithJSON(w, http.StatusOK, model.Response{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	ResponseWithJSON(w, http.StatusOK, model.Response{
		Code:    0,
		Message: "register successfully,please to active your account with email.",
	})

}

// ActiveAccount active usert account
func ActiveAccount(w http.ResponseWriter, r *http.Request) {
	token := new(model.JwtToken)
	json.NewDecoder(r.Body).Decode(token)
	log.Printf("requet data info:%#v\n", token)
	result, uid := utils.ProtectedEndpoint(token.Token)
	if result {
		//todo update user status
		stmt, _ := db.DB.Prepare(`UPDATE user SET status = 1 WHERE id = ?`)
		defer stmt.Close()
		res, err := stmt.Exec(uid)
		if err != nil {
			ResponseWithJSON(w, http.StatusOK, model.Response{
				Code:    -2001,
				Message: "update user active status failed",
			})
			return
		}
		num, _ := res.RowsAffected()
		log.Printf("update row:%v", num)
		newToken, _ := utils.CreateTokenEndpoint(uid)

		ResponseWithJSON(w, http.StatusOK, model.Response{
			Code:    0,
			Message: "OK",
			Data:    model.JwtToken{Token: newToken},
		})
	} else {
		ResponseWithJSON(w, http.StatusOK, model.Response{
			Code:    -2001,
			Message: "Token is expired",
		})
	}

}

// Login login handlerFunc
func Login(w http.ResponseWriter, r *http.Request) {
	user := new(model.User)
	json.NewDecoder(r.Body).Decode(user)
	log.Printf("request user info:%#v", user)
	if user.Email == "" || user.Password == "" {
		ResponseWithJSON(w, http.StatusOK, model.Response{
			Code:    -1,
			Message: "email and password is nil",
		})
		return
	}

	if !utils.CheckEmail(user.Email) {
		ResponseWithJSON(w, http.StatusOK, model.Response{
			Code:    -1,
			Message: "email format error",
		})
		return
	}

	user1 := model.User{}
	stmt, _ := db.DB.Prepare(`SELECT id,email,password,status FROM user WHERE email=?`)
	defer stmt.Close()

	row := stmt.QueryRow(user.Email)
	row.Scan(&user1.ID, &user1.Email, &user1.Password, &user.Status)

	if user1.ID == 0 && user1.Email == "" && user1.Password == "" {
		ResponseWithJSON(w, http.StatusOK, model.Response{
			Code:    -1,
			Message: "user not existed",
		})
		return
	}
	// TODO check user password after cryptoing
	if utils.CryptoPassword(user.Password) != user1.Password {
		ResponseWithJSON(w, http.StatusOK, model.Response{
			Code:    -1,
			Message: "password error",
		})
		return
	}
	// TODO create token and return it
	tokenString, _ := utils.CreateTokenEndpoint(user1.ID)
	ResponseWithJSON(w, http.StatusOK, model.Response{
		Code:    0,
		Message: "OK",
		Data:    model.JwtToken{Token: tokenString},
	})
}
