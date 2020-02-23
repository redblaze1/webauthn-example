package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

	// 註解這幾行, 因為讀了會出錯,
	// 出錯原因: Credentials無法輸出, 所以讀了就會讀到空的Credentials
	// 所以讀記憶裡的資料就不會出錯
	// 然後name還沒放進去,先搞定這功能在說
	// file, err := ioutil.ReadFile("user.json")
	// if err != nil {
	// 	return &User{}, errors.New("error reading json")
	// }
	// err = json.Unmarshal([]byte(file), &db)
	// if err != nil {
	// 	log.Println("OAO " + err.Error())
	// 	return &User{}, err
	// }

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
func (db *userdb) PutUser(user *User) {

	// 存記憶體裡
	for i, data := range db.Users {
		if data.Name == "" {
			db.Users[i].Name = user.Name
			db.Users[i].ID = user.ID
			db.Users[i].DisplayName = user.DisplayName
			db.Users[i].Credentials = user.Credentials
			break
		}
	}
	// 存完在存成json
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
