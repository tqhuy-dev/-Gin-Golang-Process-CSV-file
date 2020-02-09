package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	fmt.Print("hello world")
	r := gin.Default()
	r.POST("/ping", func(c *gin.Context) {
		file, header, err := c.Request.FormFile("upload")
		filename := header.Filename
		filename = "data"
		out, err := os.Create("./tmp/" + filename + ".csv")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}
		dataCSV := GetDataCSV("./tmp/data.csv")
		c.JSON(200, gin.H{
			"message": "pong",
			"data":    dataCSV,
		})
	})
	r.Run(":3000")
	// customer := Customer{name: "Huy", phone: "0946515846"}
	//store := Store{name: "Store 1", address: "DVN Street"}

	// measure(customer)
	// measure(store)
	// WaitGroupSync()
	// ChannelSync()
}

func GetDataCSV(file string) []string {
	csvfile, err := os.Open(file)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))
	dataCSV := make([]string, 0)
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		dataCSV = append(dataCSV, record[0])
	}

	return dataCSV
}

type Crud interface {
	create() (interface{}, error)
	retrieve() (interface{}, error)
}

type Customer struct {
	name  string
	phone string
}

type Store struct {
	name    string
	address string
}

func (c Customer) create() (interface{}, error) {
	data := map[string]interface{}{}
	data["msg"] = "Create API customer"
	return data, nil
}

func (c Customer) retrieve() (interface{}, error) {
	data := map[string]interface{}{}
	data["msg"] = "Get API customer"
	return data, nil
}

func (s Store) create() (interface{}, error) {
	data := map[string]interface{}{}
	data["msg"] = "Create API store"
	return data, nil
}

func (s Store) retrieve() (interface{}, error) {
	data := map[string]interface{}{}
	data["msg"] = "Get API store"
	return data, nil
}

func measure(crud Crud) {
	fmt.Print(crud.create())
	fmt.Print(crud.retrieve())
}

func WaitGroupSync() {
	var waitGroup sync.WaitGroup
	arrNum1 := []int{1, 3, 5, 7, 9}
	arrNum2 := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	waitGroup.Add(1)
	go Todo(arrNum1, &waitGroup)
	waitGroup.Add(1)
	go Todo(arrNum2, &waitGroup)

	waitGroup.Wait()
	fmt.Print("End")
}

func ChannelSync() {
	c1 := make(chan int)
	c2 := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		c1 <- 5
	}()

	go func() {
		time.Sleep(4 * time.Second)
		c2 <- 4
	}()

	msg := <-c1
	msg2 := <-c2
	total := (msg + msg2)
	fmt.Println(total)
}

func Todo(arrNum []int, waitGroup *sync.WaitGroup) int {
	count := 0
	for _, v := range arrNum {
		count += v
		fmt.Println(v)
		time.Sleep(1 * time.Second)
	}
	waitGroup.Done()
	return count
}
