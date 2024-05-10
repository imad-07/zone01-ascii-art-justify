package main

import (
	"fmt"
	"os"
	"strings"

	"funcs/funcs"
)

func main() {
	args := os.Args[1:]
	funcs.CheckArgs(args)
	alignement := args[0]
	victim := args[1]
	kind := args[2] + ".txt"
	terlen:= funcs.TerminalSize()
	// here we split our input with new lines while keeping each one of them in an indexed place in the array
	word := funcs.SplitNl(victim)
	fileContent, err := os.ReadFile(kind)
	if err != nil {
		fmt.Printf("error in the kind file")
		return
	}
	// here we get the standard art from the file that they gave us
	lettres := funcs.GetLettres(fileContent)
	linlen := 0
	for i := 0; i < len(victim); i++ {
		x := []rune(lettres[victim[i]-32][1])
		linlen += len(x)
	}
	if alignement != "--align=justify" {
		padding := funcs.Padding(alignement, victim, linlen, terlen)
		padword := funcs.SplitNl(padding)
		funcs.Printfinal(padword, lettres)
	} else if alignement == "--align=justify" {
		tst := strings.Fields(victim)
		if len(tst) == 1 {
			funcs.Printfinal(word, lettres)
		}
		x := funcs.JustifyText(victim, terlen, linlen)
		jusword := funcs.SplitNl(x)
		funcs.Printfinal(jusword, lettres)
	}
}
