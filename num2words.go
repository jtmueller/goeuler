package main

import "math"

// how many digit's groups to process
const groupsNumber int = 4

var _smallNumbers = []int{
	//"zero", "one", "two", "three", "four",
	4, 3, 3, 5, 4,
	//"five", "six", "seven", "eight", "nine",
	4, 3, 5, 5, 4,
	//"ten", "eleven", "twelve", "thirteen", "fourteen",
	3, 6, 6, 8, 8,
	//"fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
	7, 7, 9, 8, 8,
}
var _tens = []int{
	//"", "", "twenty", "thirty", "forty", "fifty",
	0, 0, 6, 6, 5, 5,
	//"sixty", "seventy", "eighty", "ninety",
	5, 7, 6, 6,
}
var _scaleNumbers = []int{
	//"", "thousand", "million", "billion",
	0, 8, 7, 7,
}

type digitGroup int

// WordLetterCount converts number into the words representation and counts the letters (ignoring spaces and punctuation)
func WordLetterCount(number int) int {
	// Zero rule
	if number == 0 {
		return _smallNumbers[0]
	}

	// Divide into three-digits group
	var groups [groupsNumber]digitGroup
	positive := math.Abs(float64(number))

	// Form three-digit groups
	for i := 0; i < groupsNumber; i++ {
		groups[i] = digitGroup(math.Mod(positive, 1000))
		positive /= 1000
	}

	var textGroup [groupsNumber]int
	for i := 0; i < groupsNumber; i++ {
		textGroup[i] = digitGroup2Text(groups[i])
	}
	combined := textGroup[0]
	appendAnd := groups[0] > 0 && groups[0] < 100

	for i := 1; i < groupsNumber; i++ {
		if groups[i] != 0 {
			prefix := textGroup[i] + _scaleNumbers[i]

			if combined != 0 {
				if appendAnd {
					prefix += 3 //" and "
				}
			}

			appendAnd = false

			combined = prefix + combined
		}
	}

	if number < 0 {
		combined += 5 //"minus " + combined
	}

	return combined
}

func intMod(x, y int) int {
	return int(math.Mod(float64(x), float64(y)))
}

func digitGroup2Text(group digitGroup) (ret int) {
	hundreds := group / 100
	tensUnits := intMod(int(group), 100)

	if hundreds != 0 {
		ret += _smallNumbers[hundreds] + 7 //" hundred"

		if tensUnits != 0 {
			ret += 3 //" and "
		}
	}

	tens := tensUnits / 10
	units := intMod(tensUnits, 10)

	if tens >= 2 {
		ret += _tens[tens]

		if units != 0 {
			ret += _smallNumbers[units]
		}
	} else if tensUnits != 0 {
		ret += _smallNumbers[tensUnits]
	}

	return
}
