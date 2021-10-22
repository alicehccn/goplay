package main

import (
	"encoding/json"
	"fmt"
	games "goplay/games"
	"io/ioutil"
	"os"
)

type Testcases struct {
	FindMedian [][]int `json:"find_median"`
}

func Print(input ...interface{}) {
	fmt.Println(input...)
}

func main() {
	Print("Let's play!")

	jsonFile, err := os.Open("testcases.json")
	if err != nil {
		Print(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var testcases Testcases
	json.Unmarshal(byteValue, &testcases)

	for i := 0; i < len(testcases.FindMedian); i++ {
		testcase := testcases.FindMedian[i]
		result := games.FindMedian(testcase)
		Print(result)
	}
}
