package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	answer := rand.Intn(100) + 1

	fmt.Println("guess game ( 1 ~ 100) : ")
	fmt.Println(answer)

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 10; i++ {
		fmt.Println("chance : ", 10-i)
		fmt.Print("guess number : ")
		inputNumberString, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		inputNumberString = strings.TrimSpace(inputNumberString)
		inputNumber, err := strconv.Atoi(inputNumberString)
		if err != nil {
			log.Fatal(err)
		}
		if inputNumber == answer {
			fmt.Println("you got it right.congrats!!")
			break
		} else if inputNumber < answer {
			fmt.Println("your guess number is lower than the answer")
		} else if inputNumber > answer {
			fmt.Println("your guess number is greater than the answer")
		}

	}

}
