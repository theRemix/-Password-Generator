package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/boltdb/bolt"
)

type PasswordResponse struct {
	Password string `json:"password"`
}

type UserInfo struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Timestamp string `json:"timestamp"`
}

var randomWords = [10]string{"random!", "password", "alsdfasdf", "poonicorn", "SSH", "ohhai", "psdoifu", "1232142341", "hello", "1312"}

func TestApiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	test := make(map[string]bool)
	test["success"] = true

	j, _ := json.Marshal(test)
	w.Write(j)
}

func CheckApiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	randomPassword := getRandomPassword()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	var userInfo *UserInfo
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		fmt.Println(err)
	}

	userInfo.Password = randomPassword

	response := &PasswordResponse{
		Password: randomPassword,
	}

	j, _ := json.Marshal(response)
	w.Write(j)
	postData(userInfo)
}

func postData(userInfo *UserInfo) {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, openDbErr := bolt.Open("users.db", 0600, nil)
	if openDbErr != nil {
		fmt.Println("ERR: ", openDbErr)
	}
	defer db.Close()

	// Store the user model in the user bucket using the username as the key.
	updateErr := db.Update(func(tx *bolt.Tx) error {
		b, insertErr := tx.CreateBucketIfNotExists([]byte("Users"))
		if insertErr != nil {
			return insertErr
		}

		encoded, encodeErr := json.Marshal(userInfo)
		if encodeErr != nil {
			return encodeErr
		}

		return b.Put([]byte(userInfo.Username), encoded)
	})

	if updateErr != nil {
		fmt.Println("ERR:", updateErr)
	}
}

func getRandomPassword() string {
	password := randomWords[rand.Intn(11)] + randomWords[rand.Intn(11)] + randomWords[rand.Intn(11)]
	fmt.Println(randomWords)
	return password
}
