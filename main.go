package main

import (
	"encoding/json"
	games "goplay/games"
	"io/ioutil"
	"os"
)

type Testcases struct {
	FindMedian    [][]int         `json:"find_median"`
	RemoveElement []RemoveElement `json:"remove_element"`
}

type RemoveElement struct {
	Nums []int `json:"nums"`
	Val  int   `json:"val"`
}

func GetTestcasesFromFile(path string) Testcases {
	jsonFile, err := os.Open(path)
	if err != nil {
		games.Print(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var testcases Testcases
	json.Unmarshal(byteValue, &testcases)
	return testcases
}

func main() {
	games.Print("Let's play!")
	testcases := GetTestcasesFromFile("testcases.json")

	// find median
	for i := 0; i < len(testcases.FindMedian); i++ {
		testcase := testcases.FindMedian[i]
		result := games.FindMedian(testcase)
		games.Print(result)
	}

	// remove val from slice
	for i := 0; i < len(testcases.RemoveElement); i++ {
		testcase := testcases.RemoveElement[i]
		result := games.RemoveElement(testcase.Nums, testcase.Val)
		games.Print(result)
	}

	digits := "276"
	games.LetterCombinations(digits)

}
