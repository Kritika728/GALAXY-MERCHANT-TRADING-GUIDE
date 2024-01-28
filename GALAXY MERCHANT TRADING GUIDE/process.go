package main

import (
	"fmt"
	"strconv"
	"strings"
)

// *************************************INPUT***************************************************

var romanToInteger = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}
var wordToRoman = map[string]string{}
var metalValue = map[string]float32{}

func convertRomanToInteger(roman string) int {
	var num int
	for i := 0; i < len(roman); i++ {
		//check if smaller value comes before greater
		var nextValue int
		if (i + 1) < len(roman) {
			nextValue = romanToInteger[string(roman[i+1])]
		}
		if currValue, exists := romanToInteger[string(roman[i])]; exists {
			if currValue < nextValue {
				num -= currValue
			} else {
				num += currValue
			}
		}
	}
	return num
}

func makeWordToRomanMap(statement string) {
	//split string through enter and get []string
	var ques []string
	lines := strings.Split(statement, "\n")
	for i := 0; i < len(lines); i++ {
		word := strings.Fields(lines[i])

		//********statements******
		//------------------------
		if word[len(word)-1] == "Credits" { //---------------------------------------statement
			romanString := wordToRoman[word[0]] + wordToRoman[word[1]]
			value := convertRomanToInteger(romanString)
			num, _ := strconv.Atoi(word[4])

			metalValue[word[2]] = float32(num) / float32(value)

		} else if word[len(word)-1] != "?" { //---------------------------------------statement
			wordToRoman[word[0]] = word[len(word)-1]
		} else { //-------------------------------------------------------------------question
			var romanNum string
			if strings.Contains(lines[i], "much") { //---------------------------------------how much ?
				ques := word[3 : len(word)-1]
				for i := 0; i < len(ques); i++ {
					if v, ok := wordToRoman[ques[i]]; ok {
						romanNum += v
					}

				}
				ans := convertRomanToInteger(romanNum)
				//fmt.Sprintf("", ans, romanNum)
				fmt.Printf("%s is %d\n", strings.Join(ques, " "), ans)

			} else if strings.Contains(lines[i], "many") { //---------------------------------------how many credit is ?
				var romanNum string
				ques = word[4 : len(word)-1]
				var metalRate float32
				for i := 0; i < len(ques); i++ {
					if v, ok := metalValue[ques[i]]; ok {
						metalRate = v
					} else {
						if v, ok := wordToRoman[ques[i]]; ok {
							romanNum += v
						}

					}

				}
				ans := convertRomanToInteger(romanNum)
				metalCredit := metalRate * float32(ans)
				//fmt.Println("metalCredit: ", metalCredit)
				fmt.Printf("%s is %.2f Credits\n", strings.Join(ques, " "), metalCredit)

			} else if (word[0] == "Does") && word[len(word)-1] == "?" { // ------------------------------------------- Does ?
				var metalRate float32
				var romanNum1, romanNum2 string
				divide := strings.Split(lines[i], "than")

				split1 := strings.Split(divide[0], " ")
				part1 := split1[1 : len(split1)-4]

				split2 := strings.Split(divide[1], " ")
				part2 := split2[:len(split2)-1]

				for i := 0; i < len(part1); i++ {
					if v, ok := metalValue[part1[i]]; ok {
						metalRate = v
					} else {
						if v, ok := wordToRoman[part1[i]]; ok {
							romanNum1 += v
						}

					}

				}
				number1 := convertRomanToInteger(romanNum1)
				ans1 := float32(number1) * metalRate

				//-----part2----
				for i := 0; i < len(part2); i++ {

					if v, ok := metalValue[part2[i]]; ok {
						metalRate = v
					} else {
						if v, ok := wordToRoman[part2[i]]; ok {
							romanNum2 += v
						}

					}

				}
				number2 := convertRomanToInteger(romanNum2)
				ans2 := float32(number2) * metalRate

				if ans1 > ans2 {
					fmt.Printf("%s has more Credits than %s\n", strings.Join(part1, " "), strings.Join(part2, " "))
				} else {
					fmt.Printf("%s has less Credits than %s\n", strings.Join(part1, " "), strings.Join(part2, " "))
				}
			} else if word[0] == "Is" && word[len(word)-1] == "?" {
				var metalRate float32
				var romanNum1, romanNum2 string
				divide := strings.Split(lines[i], "than")

				split1 := strings.Split(divide[0], " ")
				part1 := split1[1 : len(split1)-1]

				split2 := strings.Split(divide[1], " ")
				part2 := split2[:len(split2)-1]

				for i := 0; i < len(part1); i++ {
					if v, ok := metalValue[part1[i]]; ok {
						metalRate = v
					} else {
						if v, ok := wordToRoman[part1[i]]; ok {
							romanNum1 += v
						}

					}

				}
				number1 := convertRomanToInteger(romanNum1)
				ans1 := float32(number1) * metalRate

				//-----part2----
				for i := 0; i < len(part2); i++ {

					if v, ok := metalValue[part2[i]]; ok {
						metalRate = v
					} else {
						if v, ok := wordToRoman[part2[i]]; ok {
							romanNum2 += v
						}

					}

				}
				number2 := convertRomanToInteger(romanNum2)
				ans2 := float32(number2) * metalRate

				if ans1 > ans2 {
					fmt.Printf("%s is larger than %s\n", strings.Join(part1, " "), strings.Join(part2, " "))
				} else {
					fmt.Printf("%s is smaller than %s\n", strings.Join(part1, " "), strings.Join(part2, " "))
				}

			} else {
				fmt.Println("Requested number is in invalid format")
			}

		}
	}
}

//***********************

//******************************************************************************************

func main() {
	//roman to integer
	statement := `glob is I
	prok is V
	pish is X
	tegj is L
	glob glob Silver is 34 Credits
	glob prok Gold is 57800 Credits
	pish pish Iron is 3910 Credits
	how much is pish tegj glob glob ?
	how many Credits is glob prok Silver ?
	how many Credits is glob glob Gold ?
	how many Credits is glob glob glob glob glob glob Gold ?
	how many Credits is pish tegj glob Iron ?
	Does pish tegj glob glob Iron has more Credits than glob glob Gold ?
	Does glob glob Gold has less Credits than pish tegj glob glob Iron ?
	Is glob prok larger than pish pish ?
	Is tegj glob glob smaller than glob prok ?
	how much wood could a woodchuck chuck if a woodchuck could chuck wood ?`

	makeWordToRomanMap(statement)

}
