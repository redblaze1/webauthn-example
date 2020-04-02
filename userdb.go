package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

type userdb struct {
	Users []User   `json:"user"`
	Name  []string `json:"name"`
}

var db *userdb

// DB returns a userdb singleton
func DB() *userdb {

	if db == nil {
		db = &userdb{
			// 弄5個來測試, 滿了就不給存
			Users: make([]User, 5),
			Name:  make([]string, 5),
		}
	}

	return db
}

// GetUser returns a *User by the user's username
func (db *userdb) GetUser(name string) (*User, error) {

	file, err := ioutil.ReadFile("user.json")
	if err != nil {
		return &User{}, errors.New("error reading json")
	}
	err = json.Unmarshal([]byte(file), &db)
	if err != nil {
		log.Println("OAO " + err.Error())
		return &User{}, err
	}

	for i, data := range db.Users {
		if data.Name == name {
			return &db.Users[i], nil
		}
	}
	return &User{}, errors.New("error getting user")
	// user := db.users[name]
	// if !ok {
	// 	return &User{}, fmt.Errorf("error getting user '%s': does not exist", name)
	// }

	// return user, nil
}

// PutUser stores a new user by the user's username
func (db *userdb) PutUser(user *User, name string) {

	// 存記憶體裡
	for i, data := range db.Users {
		if data.Name == "" {
			db.Users[i].Name = user.Name
			db.Users[i].ID = user.ID
			db.Users[i].DisplayName = user.DisplayName
			db.Name[i] = name
			// Call 這函式的時候根本還沒把Cred拿進來= =
			// db.Users[i].Credentials = user.Credentials
			break
		}
	}
}

// Outputjson 等有cred的時候在輸出
func Outputjson() {
	// 把有Credentials狀態的db轉成json byte
	wfile, err := json.MarshalIndent(&db, "", " ")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 輸出成檔案
	err = ioutil.WriteFile("user.json", wfile, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

// GetUsername returns a name by the user's username
func (db *userdb) GetUsername(username string) (string, error) {
	for i, data := range db.Users {
		if data.Name == username {
			return db.Name[i], nil
		}
	}
	return "", fmt.Errorf("Error find name")
}
