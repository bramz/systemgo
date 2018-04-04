package main

import( 
   "fmt"
    "github.com/gin-gonic/gin"
)

func index(context *gin.Context) {
    fmt.Println("testing1")
    fmt.Println("testingagaintest")
}

func main() {
    router := gin.Default()

    router.GET("/", index)

    router.Run(":8080")
    fmt.Println("Running test httpd")
    //test
}
