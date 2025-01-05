package main

import (
    "fmt"
    "log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
    db, err := connectToDB()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    router := gin.Default()

    router.GET("/userinfo", func(c *gin.Context) {
        getUserInfo(c, db)
    })

	router.GET("/userinfo/:steamAuth", func(c *gin.Context) {
        getUserInfo(c, db)
    })

    router.Run(":8080") // Run on port 8080
    fmt.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}