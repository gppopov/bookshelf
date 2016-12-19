package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/redis.v5"
)

var redisClient *redis.Client

func main() {
	confFilePath := "./config.json"
	conf := readConfig(confFilePath)
	connectToRedis(conf.RedisAddr)

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	// set routes
	http.HandleFunc("/book", addBookHandler)
	http.HandleFunc("/books", listBooksHandler)
	http.HandleFunc("/login", loginHandler)
	err := http.ListenAndServeTLS(conf.AppPort, conf.PublicKey, conf.PrivateKey, nil)
	logErr(err)
}

func connectToRedis(addr string) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	pong, redisErr := redisClient.Ping().Result()
	logErr(redisErr)
	fmt.Println("Redis server accessible ", pong)

	bookIDKey := "book:id"
	bookID, err := redisClient.Get(bookIDKey).Result()
	if err == redis.Nil {
		setErr := redisClient.Set(bookIDKey, 1, 0).Err()
		checkErr(setErr)
	} else {
		fmt.Println(bookID)
	}
}

func readConfig(path string) Config {
	confFile, e := ioutil.ReadFile(path)
	checkErr(e)
	var config Config
	json.Unmarshal(confFile, &config)
	fmt.Println(config.RedisAddr)
	return config
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func logErr(err error) {
	if err != nil {
		log.Fatal("Server error: ", err)
	}
}

// Book type holds books info.
type Book struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Picurl string `json:"picurl"`
}

// Config holds app's config details as raed from the config json file.
type Config struct {
	RedisAddr  string `json:"redis_addr"`
	PrivateKey string `json:"server_private_key"`
	PublicKey  string `json:"server_public_key"`
	AppPort    string `json:"app_port"`
}
