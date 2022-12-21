package main

import (
	"fmt"
	s "strings"
	"unicode"
)

var f = fmt.Printf

func main() {
	upper := s.ToUpper("Hello there!")
	f("To Upper: %s\n", upper)
	// To Upper: HELLO THERE!
	f("To Lower: %s\n", s.ToLower("Hello THERE!"))
	// To Lower: hello there!
	f("%s\n", s.Title("tHis wiLL be A title!"))
	// THis WiLL Be A Title!
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAlis"))
	// EqualFold: true
	f("EqualFold: %V\n", s.EqualFold("Mihalis", "MIHAli"))
	// EqualFold: %!V(bool=false)

	f("Prefix: %v\n", s.HasPrefix("Mihalis", "Mi"))
	// Prefix: true
	f("Prefix: %v\n", s.HasPrefix("Mihalis", "mi"))
	// Prefix: false
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "is"))
	// Suffix: true
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "IS"))
	// Suffix: false

	f("Index: %v\n", s.Index("Mihalis", "ha"))
	// Index: 2
	f("Index: %v\n", s.Index("Mihalis", "Ha"))
	// Index: -1
	f("Count: %v\n", s.Count("Mihalis", "i"))
	// Count: 2
	f("Count: %v\n", s.Count("Mihalis", "I"))
	// Count: 0
	f("Repeat: %s\n", s.Repeat("ab", 5))
	// Repeat: ababababab

	f("TrimSpace: %s\n", s.TrimSpace(" \tThis is a line. \n"))
	// TrimSpace: This is a line.
	f("TrimLeft: %s", s.TrimLeft(" \tThis is a\t line. \n", "\n\t "))
	// TrimLeft: This is a     line.
	f("TrimRight: %s\n", s.TrimRight(" \tThis is a\t line. \n", "\n\t "))
	// TrimRight:     This is a     line.

	f("Compare: %v\n", s.Compare("Mihalis", "MIHALIS"))
	// Compare: 1
	f("Compare: %v\n", s.Compare("Mihalis", "Mihalis"))
	// Compare: 0
	f("Compare: %v\n", s.Compare("Mihalis", "MIHalis"))
	// Compare: 1

	f("Fields: %v\n", s.Fields("This is a string!"))
	// Fields: [This is a string!]
	f("Fields: %v\n", s.Fields("Thisis\na\tstring!"))
	// Fields: [Thisis a string]

	f("%s\n", s.Split("abcd efg", ""))
	// [a b c d   e f g]

	f("%s\n", s.Replace("abcd efg", "", "_", -1))
	// _a_b_c_d_ _e_f_g_
	f("%s\n", s.Replace("abcd efg", "", "_", 4))
	// _a_b_c_d efg
	f("%s\n", s.Replace("abcd efg", "", "_", 2))
	// _a_bcd efg

	lines := []string{"Line 1", "Line 2", "Line 3"}
	f("Join: %s\n", s.Join(lines, "+++"))
	// Join: Line 1+++Line 2+++Line 3

	f("SplitAfter: %s\n", s.SplitAfter("123++432++", "++"))
	// SplitAfter: [123++ 432++ ]

	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	f("TrimFunc: %s\n", s.TrimFunc("123 abc ABC \t .", trimFunction))
	// TrimFunc: abc ABC
}
