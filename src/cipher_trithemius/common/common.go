package common

var Alphabet = map[string]int {
	"a": 0,
	"b": 1,
	"c": 2,
	"d": 3,
	"e": 4,
	"f": 5,
	"g": 6,
	"h": 7,
	"i": 8,
	"j": 9,
	"k": 10,
	"l": 11,
	"m": 12,
	"n": 13,
	"o": 14,
	"p": 15,
	"q": 16,
	"r": 17,
	"s": 18,
	"t": 19,
	"u": 20,
	"v": 21,
	"w": 22,
	"x": 23,
	"y": 24,
	"z": 25,
	".": 26,
	" ": 27,
	"1": 28,
	"2": 29,
	"3": 30,
	"4": 31,
	"5": 32,
	"6": 33,
	"7": 34,
	"8": 35,
	"9": 36,
	"0": 37,
	",": 38,
	"-": 39,
}

func ReverseMap(m map[string]int) map[int]string {
    n := make(map[int]string)
    for k, v := range m {
        n[v] = k
    }
    return n
}


func GetOffsetStep(pos int) int {
	return 2 * pos + 15
}

