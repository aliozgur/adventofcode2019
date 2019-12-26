package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var DEBUG bool = false


type UserOptions struct {
	Print bool
	PrintSleepMiliseconds int
}


func Println(a ...interface{}) (n int, err error) {
	if !DEBUG{
		return
	}
	return fmt.Println(a...)
}

func Max(v1 int, v2 int) int{
	if v1 > v2{
		return v1
	}
	return v2
}

func ReadIntegersFromFile(filePath string) (result []int64){
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		v,e := strconv.ParseInt(scanner.Text(),10,64)
		if e != nil {
			result = append(result, v)
		} else{
			panic(e)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return
}

func ToInt(s string) (result int){
	result,_ =strconv.Atoi(s)
	return
}
