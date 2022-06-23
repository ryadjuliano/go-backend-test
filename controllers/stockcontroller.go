package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"inventory/database"
	"inventory/entity"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

var (
	secretkey string = "secretkeyjwt"
)

//set error message in Error struct
func SetError(err entity.Error, message string) entity.Error {
	err.IsError = true
	err.Message = message
	return err
}

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//Generate JWT token
func GenerateJWT(email, role string) (string, error) {
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("Something went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func GetAllStock(w http.ResponseWriter, r *http.Request) {
	var inventory []entity.Inventorys
	// database.Connector.Find(&inventory)
	// errors.Is(result.Error, gorm.ErrRecordNotFound)
	result := database.Connector.Find(&inventory)
	// result.RowsAffected // returns count of records found
	// result.Error        // returns error or nil
	errors.Is(result.Error, gorm.ErrRecordNotFound)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(inventory)
}

func GetStockById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var inventory entity.Inventorys
	database.Connector.First(&inventory, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventory)
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var inventory entity.Inventorys
	json.Unmarshal(requestBody, &inventory)

	database.Connector.Create(&inventory)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(inventory)
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var user entity.Users
	err := json.NewDecoder(r.Body).Decode(&user)
	var dbuser entity.Users
	database.Connector.Where("email = ?", dbuser.Email).First(&dbuser)

	//check email is alredy registered or not
	if user.Email != "" {
		var err entity.Error
		err = SetError(err, "Email already in use")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	user.Password, err = GeneratehashPassword(user.Password)
	if err != nil {
		log.Fatalln("Error in password hashing.")
	}

	json.Unmarshal(requestBody, &user)
	database.Connector.Create(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var authDetails entity.Authentication

	err := json.NewDecoder(r.Body).Decode(&authDetails)
	if err != nil {
		var err entity.Error
		err = SetError(err, "Error in reading payload.")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var authUser entity.Users
	database.Connector.Where("email = 	?", authDetails.Email).First(requestBody, &authUser)

	if authUser.Email == "" {
		var err entity.Error
		err = SetError(err, "Username or Password is incorrect")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	check := CheckPasswordHash(authDetails.Password, authUser.Password)

	if !check {
		var err entity.Error
		err = SetError(err, "Username or Password is incorrect")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	validToken, err := GenerateJWT(authUser.Email, authUser.Role)
	if err != nil {
		var err entity.Error
		err = SetError(err, "Failed to generate token")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var token entity.Token
	token.Email = authUser.Email
	token.Role = authUser.Role
	token.TokenString = validToken
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}
