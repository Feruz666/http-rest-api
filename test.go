package main

import (
	// "fmt"
	"fmt"
	"strings"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// "hellocat", "apple,bat,cat,goodbye,hello,yellow,why"
func ArrayChallenge(strArr []string) string {
	splittedStr := strings.Split(strArr[1], ",")
	finalArr := []string{}
	var tempWord string

	for _, v := range strArr[0] {
		tempWord += string(v)
		if contains(splittedStr, tempWord) {
			finalArr = append(finalArr, tempWord)
			tempWord = ""
		}
	}
	strArr[0] = strings.Join(finalArr, ",")



	// code goes here
	return strArr[0]

}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	// fmt.Println(ArrayChallenge(readline()))

	strarray := []string{"hellocat", "apple,bat,cat,goodbye,hello,yellow,why"}

	splittedStr := strings.Split(strarray[1], ",")
	finalArr := []string{}
	var tempWord string

	for _, v := range strarray[0] {
		tempWord += string(v)
		if contains(splittedStr, tempWord) {
			finalArr = append(finalArr, tempWord)
			tempWord = ""
		}
	}
	finalStr := strings.Join(finalArr, ",")
	fmt.Println(string(finalStr))
}
