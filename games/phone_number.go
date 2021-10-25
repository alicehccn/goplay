package games

func GetPhoneNums() map[string][]string {
	result := make(map[string][]string)
	result["0"] = nil
	result["1"] = nil
	result["2"] = []string{"a", "b", "c"}
	result["3"] = []string{"d", "e", "f"}
	result["4"] = []string{"g", "h", "i"}
	result["5"] = []string{"j", "k", "l"}
	result["6"] = []string{"m", "n", "o"}
	result["7"] = []string{"p", "q", "r", "s"}
	result["8"] = []string{"t", "u", "v"}
	result["9"] = []string{"w", "x", "y", "z"}
	return result
}

func LetterCombinations(digits string) []string {
	// var letters [][]string
	var result []string
	// // var maxSliceSize int
	// nums := GetPhoneNums()
	// for _, char := range digits {
	// 	d := string(char)
	// 	group := nums[d]
	// 	letters = append(letters, group)
	// 	if len(group) > maxSliceSize {
	// 		// maxSliceSize = len(group)
	// 	}
	// }

	// fmt.Println(result)
	return result
}
