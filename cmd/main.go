package main

import (
	"OfflineSearchEngine/Internals/SearchEngines/builder"
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strings"
)

func main() {
	wordSearch := ""
	ret := builder.NewSearchEngine("LinearFastAddEngine", 10000)
	file, err := os.Open("C:\\Users\\zariii\\Documents\\myFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	ret.AddData(scanner, 10)
	data2 := "  Golang: the( Golang( is is Golang golang memory golang Golang safety Golang;"
	scanner = bufio.NewScanner(strings.NewReader(data2))
	scanner.Split(bufio.ScanWords)
	ret.AddData(scanner, 2)
	fmt.Println("add:", ret)
	fmt.Print("Please Enter word : ")
	fmt.Scan(&wordSearch)
	//ret1, err := ret.Search(wordSearch)
	//if err != nil {
	//	fmt.Println("Search:", err)
	//} else {
	//	fmt.Println("Search:", ret1)
	//}
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": ret,
		})
	})
	router.Run()

}
