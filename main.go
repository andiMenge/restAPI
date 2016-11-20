package main

import (
	"flag"
	"net/http"
)

var (
	apiPort   = flag.String("port", "1337", "defines the port for the webserver")
	redisAddr = flag.String("redis", "localhost:6379", "addr of redis server")
	redisPW   = flag.String("pw", "", "pw for redis server. \"\" for no pw") // no password set
	redisDB   = flag.Int("db", 0, "defines redis db. 0 for default db")      // use default DB
)

func main() {
	flag.Parse() //parse cmdline flags

	//start redis client
	err := initRedis()
	if err != nil {
		ErrLog.Fatalf("connection to redis failed: %s", err)
	}
	InfoLog.Printf("connection to redis successful")

	//start webserver
	router := NewRouter()
	InfoLog.Printf("starting webserver on port 0.0.0.0:%s", *apiPort)
	ErrLog.Fatal(http.ListenAndServe(":"+*apiPort, router))
}
