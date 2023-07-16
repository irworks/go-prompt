package strings

import (
	"unicode/utf8"

	"github.com/mattn/go-runewidth"
)

// Get the length of the string in bytes.
func Len(s string) ByteNumber {
	return ByteNumber(len(s))
}

// Get the length of the string in runes.
func RuneCount(s string) RuneNumber {
	return RuneNumber(utf8.RuneCountInString(s))
}

// Get the width of the string (how many columns it takes upt in the terminal).
func GetWidth(s string) Width {
	return Width(runewidth.StringWidth(s))
}

// Get the width of the rune (how many columns it takes upt in the terminal).
func GetRuneWidth(char rune) Width {
	return Width(runewidth.RuneWidth(char))
}

// IndexNotByte is similar with strings.IndexByte but showing the opposite behavior.
func IndexNotByte(s string, c byte) ByteNumber {
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] != c {
			return ByteNumber(i)
		}
	}
	return -1
}

// LastIndexNotByte is similar with strings.LastIndexByte but showing the opposite behavior.
func LastIndexNotByte(s string, c byte) ByteNumber {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != c {
			return ByteNumber(i)
		}
	}
	return -1
}

type asciiSet [8]uint32

func (as *asciiSet) notContains(c byte) bool {
	return (as[c>>5] & (1 << uint(c&31))) == 0
}

func makeASCIISet(chars string) (as asciiSet, ok bool) {
	for i := 0; i < len(chars); i++ {
		c := chars[i]
		if c >= utf8.RuneSelf {
			return as, false
		}
		as[c>>5] |= 1 << uint(c&31)
	}
	return as, true
}

// IndexNotAny is similar with strings.IndexAny but showing the opposite behavior.
func IndexNotAny(s, chars string) ByteNumber {
	if len(chars) > 0 {
		if len(s) > 8 {
			if as, isASCII := makeASCIISet(chars); isASCII {
				for i := 0; i < len(s); i++ {
					if as.notContains(s[i]) {
						return ByteNumber(i)
					}
				}
				return -1
			}
		}

	LabelFirstLoop:
		for i, c := range s {
			for j, m := range chars {
				if c != m && j == len(chars)-1 {
					return ByteNumber(i)
				} else if c != m {
					continue
				} else {
					continue LabelFirstLoop
				}
			}
		}
	}
	return -1
}

// LastIndexNotAny is similar with strings.LastIndexAny but showing the opposite behavior.
func LastIndexNotAny(s, chars string) ByteNumber {
	if len(chars) > 0 {
		if len(s) > 8 {
			if as, isASCII := makeASCIISet(chars); isASCII {
				for i := len(s) - 1; i >= 0; i-- {
					if as.notContains(s[i]) {
						return ByteNumber(i)
					}
				}
				return -1
			}
		}
	LabelFirstLoop:
		for i := len(s); i > 0; {
			r, size := utf8.DecodeLastRuneInString(s[:i])
			i -= size
			for j, m := range chars {
				if r != m && j == len(chars)-1 {
					return ByteNumber(i)
				} else if r != m {
					continue
				} else {
					continue LabelFirstLoop
				}
			}
		}
	}
	return -1
}
