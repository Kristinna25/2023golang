package main

import (
	"fmt"
	"strings"
)

func main() {

	brokenWords := "hello&& Lee felix"
	replaceWords := strings.NewReplacer("&", "*")
	fixedWords := replaceWords.Replace(brokenWords)
	fmt.Println(fixedWords)
}
