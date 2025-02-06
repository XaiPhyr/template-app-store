package main

import (
	"app_store_api/routers"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	version := os.Getenv("VERSION")

	log.SetFlags(log.Llongfile | log.LstdFlags)

	fmt.Println()
	log.Printf("-> Version: %s", version)
	log.Printf("-> Local:   http://localhost:8200")
	fmt.Println()

	r := routers.InitRouter()

	err := http.ListenAndServe(":8200", r)
	log.Printf("Error ListenAndServe: %s", err)
	fmt.Println()
}
