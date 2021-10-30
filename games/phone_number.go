package games

func GetPhoneNums(key string) []string {
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
	return result[key]
}

func Perm(output []string, input [][]string) []string {
	var i, j, k, m int
	Print("---------")
	for k = 0; k < 3; k++ {
		for i = 0; i < 3; i++ {

			for j = 0; j < 3; j++ {
				if k >= 1 {
					m = 0
				} else {
					m = k + 1
				}
				// Print(k, k, m, i, m+1, j)
				Print(input[k][k], input[m][i], input[m+1][j])
			}
		}
	}

	Print(output)
	return output
}

func LetterCombinations(digits string) []string {
	var letters [][]string
	var numOfPerm int = 1
	// // var maxSliceSize int

	for _, char := range digits {
		dString := string(char)
		nums := GetPhoneNums(dString)
		letters = append(letters, nums)
		numOfPerm *= len(nums)
	}
	var results []string
	return Perm(results, letters)
}
