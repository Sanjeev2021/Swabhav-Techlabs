package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func getUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	//refreshToken(w, r)
	fmt.Println("Inside getUser")
	if isValidCoockie(w, r) {
		// db, err := gorm.Open(sqlite.Open("address.db"), &gorm.Config{})
		// if err != nil {
		// 	panic("failed to connect database")
		// }
		//var usersForDisplay []User
		//usersForDisplay = alluser
		// db.Find(&usersForDisplay)
		var userforDisplay []User
		fmt.Println("Valid Token")
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				fmt.Println("Error Occured 1")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			json.NewEncoder(w).Encode("userforDisplay")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		tokenStr := cookie.Value
		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				fmt.Println("Error Occured 2")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			fmt.Println("Error Occured 3")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		fmt.Println(claims.Role)
		if claims.Role == "admin" {
			for _, u := range alluser {
				if u.Role != "admin" {
					userforDisplay = append(userforDisplay, u)
				}
			}
			json.NewEncoder(w).Encode(userforDisplay)
			w.WriteHeader(200)
		} else {
			fmt.Println("Error Occured 6")
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorised Access")
		}

		// for i, _ := range usersForDisplay {
		// 	usersForDisplay[i] = User{}
		// }

	} else {
		fmt.Println("Error Occured 5")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorised Access")
	}

}
func addUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	fmt.Println("Inside addUser")
	if isValidCoockie(w, r) {
		var newUser User
		_ = json.NewDecoder(r.Body).Decode(&newUser)
		var genrateUID = uuid.New()
		newUser.UID = genrateUID.String()
		for _, u := range alluser {
			if newUser.Username == u.Username {
				w.WriteHeader(http.StatusConflict)
				fmt.Fprintf(w, "Username Already Exist")
				break
			}
		}

		// db, err := gorm.Open(sqlite.Open("address.db"), &gorm.Config{})
		// if err != nil {
		// 	panic("failed to connect database")
		// }
		// db.AutoMigrate(&User{})
		// db.Create(&newUser)

		fmt.Println(newUser)
		alluser = append(alluser, newUser)
		json.NewEncoder(w).Encode(newUser)
		w.WriteHeader(200)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorised Access")
	}

}

func deleteUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	fmt.Println("Inside DeleteUSer")
	if isValidCoockie(w, r) {
		// db, err := gorm.Open(sqlite.Open("address.db"), &gorm.Config{})
		// if err != nil {
		// 	panic("failed to connect database")
		// }
		//var usersForDisplay []User
		//usersForDisplay = alluser
		params := mux.Vars(r)
		fmt.Println(params)
		//if getByRoll(params[("RollNo")]) {
		usertodelete := params[("id")]
		fmt.Println(usertodelete)
		for index, u := range alluser {
			if u.Username == usertodelete {
				alluser = append(alluser[:index], alluser[index+1:]...)
				break
			}
		}
		//db.Find(&usersForDisplay)
		// for _, u := range usersForDisplay {
		// 	fmt.Println(u.Username)
		// }
		//db.Where("roll_no = ?", rollToDelete).Delete(&usersForDisplay)
		// for index, _ := range usersForDisplay {
		// 	usersForDisplay[index] = User{}
		// }
		//	db.Find(&usersForDisplay)
		// _deleteUserAtUid(params[("RollNo")])
		// json.NewEncoder(w).Encode(user)
		//} else {
		//w.WriteHeader(http.StatusNotFound)
		//}
		w.WriteHeader(200)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorised Access")
	}

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	fmt.Println("Inside UpdateUser")
	// db, err := gorm.Open(sqlite.Open("address.db"), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }

	if isValidCoockie(w, r) {
		var newUser User
		_ = json.NewDecoder(r.Body).Decode(&newUser)
		params := mux.Vars(r)
		fmt.Println(params)
		var tempuser User
		//	db.Model(&usersForDisplay).Where("roll_no = ?", params[("RollNo")]).Updates(newUser)
		// _deleteUserAtUid(params[("RollNo")])
		// if flagfordelete == 0 {
		// 	genrateUID := uuid.New()
		// 	newUser.UID = genrateUID.String()
		// 	user = append(user, newUser)
		// } else if flagfordelete == 2 {
		// 	flagfordelete = 0
		// 	w.WriteHeader(http.StatusNotFound)
		// }

		for index, u := range alluser {
			if u.Username == params[("username")] {
				tempuser = u
				alluser = append(alluser[:index], alluser[index+1:]...)
				break
			}
		}
		tempuser.Fname = newUser.Fname
		tempuser.Password = newUser.Password
		tempuser.Role = newUser.Role
		tempuser.IsActive = newUser.IsActive
		alluser = append(alluser, tempuser)
		json.NewEncoder(w).Encode(tempuser)

		w.WriteHeader(200)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorised Access")
	}

}
