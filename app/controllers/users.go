package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/omar-ozgur/gram/app/models"
	"github.com/omar-ozgur/gram/utilities"
	"io/ioutil"
	"net/http"
)

var UsersCreate = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &user)

	status, message, createdUser := models.CreateUser(user)

	JSON, _ := json.Marshal(map[string]interface{}{
		"status":  status,
		"message": message,
		"user":    createdUser,
	})
	w.Write(JSON)
})

var UsersLogin = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &user)

	status, message, loginToken := models.LoginUser(user)

	JSON, _ := json.Marshal(map[string]interface{}{
		"status":  status,
		"message": message,
		"token":   loginToken,
	})
	w.Write(JSON)
})

var UsersProfile = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims := utilities.GetClaims(r.Header.Get("Authorization")[len("Bearer "):])
	current_user_id := fmt.Sprintf("%v", claims["user_id"])

	status, message, retrievedUser := models.GetUser(current_user_id)

	JSON, _ := json.Marshal(map[string]interface{}{
		"status":  status,
		"message": message,
		"user":    retrievedUser,
	})
	w.Write(JSON)
})

var UsersIndex = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	status, message, retrievedUsers := models.GetUsers()

	JSON, _ := json.Marshal(map[string]interface{}{
		"status":  status,
		"message": message,
		"users":   retrievedUsers,
	})
	w.Write(JSON)
})

var UsersSearch = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	b, _ := ioutil.ReadAll(r.Body)
	params := make(map[string]interface{})

	json.Unmarshal(b, &params)

	status, message, retrievedUsers := models.SearchUsers(params, "AND")

	JSON, _ := json.Marshal(map[string]interface{}{
		"status":  status,
		"message": message,
		"users":   retrievedUsers,
	})
	w.Write(JSON)
})

var UsersShow = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	status, message, retrievedUser := models.GetUser(vars["id"])

	JSON, _ := json.Marshal(map[string]interface{}{
		"status":  status,
		"message": message,
		"user":    retrievedUser,
	})
	w.Write(JSON)
})

var UsersUpdate = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	vars := mux.Vars(r)
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &user)

	claims := utilities.GetClaims(r.Header.Get("Authorization")[len("Bearer "):])
	current_user_id := fmt.Sprintf("%v", claims["user_id"])

	if vars["id"] != current_user_id {
		JSON, _ := json.Marshal(map[string]interface{}{
			"status":  "error",
			"message": "You do not have permission to update this user",
			"user":    models.User{},
		})
		w.Write(JSON)
		return
	}

	status, message, updatedUser := models.UpdateUser(vars["id"], user)

	JSON, _ := json.Marshal(map[string]interface{}{
		"status":  status,
		"message": message,
		"user":    updatedUser,
	})
	w.Write(JSON)
})

var UsersDelete = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	claims := utilities.GetClaims(r.Header.Get("Authorization")[len("Bearer "):])
	current_user_id := fmt.Sprintf("%v", claims["user_id"])

	if vars["id"] != current_user_id {
		JSON, _ := json.Marshal(map[string]interface{}{
			"status":  "error",
			"message": "You do not have permission to delete this user",
		})
		w.Write(JSON)
		return
	}

	status, message := models.DeleteUser(vars["id"])

	JSON, _ := json.Marshal(map[string]interface{}{
		"status":  status,
		"message": message,
	})
	w.Write(JSON)
})
