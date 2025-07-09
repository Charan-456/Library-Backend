package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Charan-456/funcs/middleware"
	"github.com/Charan-456/funcs/models"
	"github.com/golang-jwt/jwt/v5"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	b := r.Body
	newUser := models.User{}
	err := json.NewDecoder(b).Decode(&newUser)
	if err != nil {
		fmt.Fprintln(w, "Invalid data", http.StatusBadRequest)
		return
	}
	if newUser.Username == "" {
		fmt.Fprintln(w, "Invalid username", http.StatusForbidden)
		return
	}
	if len(newUser.Password) < 8 {
		fmt.Fprintln(w, "Password length must be greater than 8", http.StatusBadRequest)
		return
	}
	if len(newUser.Email) < 10 {
		fmt.Fprintln(w, "Invalid email ID", http.StatusBadRequest)
		return
	}
	var existingUser models.User
	if err := models.DB.First(&existingUser, "username = ? OR email = ?", newUser.Username, newUser.Email).Error; err == nil {
		fmt.Fprintln(w, "Username or Email already exists", http.StatusConflict)
		return
	}
	if err := models.DB.Create(&newUser).Error; err != nil {
		fmt.Fprintln(w, "Cannot create user", http.StatusNotImplemented)
		return
	}
	fmt.Fprintln(w, "User Created Succefully \nUse these credentials for your LOGIN", http.StatusOK)

	// if err == nil {
	// 	userName, exists := models.UserList[newUser.Username]
	// 	if exists {
	// 		fmt.Fprintln(w, "username: ", userName, " already exists, please choose a different username")
	// 		return
	// 	}
	// 	if len(newUser.Password) < 8 {
	// 		fmt.Fprintln(w, "Please choose a password that is complex and long")
	// 		return
	// 	}
	// 	if newUser.Email == "" || len(newUser.Email) < 10 {
	// 		fmt.Fprintln(w, "Enter a valid email  ID")
	// 		return
	// 	}
	// 	models.UserList[newUser.Username] = newUser
	// 	fmt.Fprintln(w, "Welcome to our world of books, We are happy you decided to learn\n You can use your credentials to log in and reserve books of your choice at your nearest store")
	// }
}

func GetAllUserNames(w http.ResponseWriter, r *http.Request) {
	log.Println("Summoning JUTSU")
	var AllUsers []models.User
	models.DB.Find(&AllUsers)
	json.NewEncoder(w).Encode(AllUsers)
}

func Login(w http.ResponseWriter, r *http.Request) {
	userCreds := models.Creds{}
	err := json.NewDecoder(r.Body).Decode(&userCreds)
	if err != nil {
		http.Error(w, "JSON unmarshalling failed", http.StatusBadRequest)
		return
	}
	// userIfExists, exists := models.UserList[userCreds.Username]
	var existingUser models.User
	err = models.DB.Where("(BINARY email = ? OR BINARY username = ?) AND BINARY password = ?", userCreds.Username, userCreds.Username, userCreds.Password).First(&existingUser).Error
	if err != nil {
		http.Error(w, "Entered username or Password is incorrect", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "Indian Library",
		"sub": userCreds.Username,
		"exp": time.Now().Add(time.Hour).Unix(),
	})
	tokenSigned, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		http.Error(w, "The token signing failed", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenSigned})

}

func Books(w http.ResponseWriter, r *http.Request) {
	var UserName = middleware.ContextKey("user_name")
	fmt.Fprintf(w, "Welocome to the store %v", r.Context().Value(UserName))
}
