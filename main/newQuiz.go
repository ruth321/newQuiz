package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

type questionAnswer struct {
	Question string
	Answer   string
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var fileName string
	fmt.Print("Enter file name: ")
	_, _ = fmt.Scan(&fileName)
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("File does not exist")
		return
	}
	var quiz []questionAnswer
	_ = json.Unmarshal(file, &quiz)
	var answer string
	var count int
	var lim int
	fmt.Printf("Enter number of questions (max %d)\n", len(quiz))
	fmt.Print("->")
	_, _ = fmt.Scan(&lim)
	a := rand.Perm(len(quiz))
	a = a[:lim]
	//c := make(chan string)
	//go timeLim(c)
	for i := 0; i < len(a); i++ {
		fmt.Print("Question: ")
		fmt.Println(quiz[a[i]].Question)
		fmt.Print("Answer: ")
		_, _ = fmt.Scan(&answer)

		if answer == quiz[a[i]].Answer {
			count++
		}
	}
	fmt.Printf("Right answers: %d out of %d\n", count, len(a))
}
