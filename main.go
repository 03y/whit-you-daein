package main

import (
    "os"
    "fmt"
    "os/exec"
    "net/http"
    "github.com/gin-gonic/gin"
)

const VERSION           string = "0.0.1-Alpha"

func uptime(c *gin.Context) {
    app := "uptime"

    cmd := exec.Command(app)
    stdout, err := cmd.Output()

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    var output string = string(stdout)
    c.IndentedJSON(http.StatusOK, output[1:len(output)-2])
}

func free(c *gin.Context) {
    app := "free"
    arg0 := "-hL"

    cmd := exec.Command(app, arg0)
    stdout, err := cmd.Output()

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    var output string = string(stdout)
    c.IndentedJSON(http.StatusOK, output[1:len(output)-2])
}

func main() {
    fmt.Printf("Whit you daein %s\n", VERSION) 
   
    if len(os.Args) < 2 {
        fmt.Println("Usage: ./whit <URL>:<PORT>")
        return
    }


    var listen string = os.Args[1]

    router := gin.Default()
    router.GET("/uptime", uptime)
    router.GET("/free", free)
    fmt.Printf("\nListening on %s...\n", listen)
    router.Run(listen)
}

