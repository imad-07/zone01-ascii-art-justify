package main

import (
	"fmt"
	"funcs/funcs"
	"os"
)

func main() { //{anass,"","imad"}
	args := os.Args[1:]
	funcs.CheckArgs(args)
	// alignement := args[0]
	alignement := args[0]
	victim := args[1]
	kind := args[2] + ".txt"
	terlen, _ := funcs.TerminalSize()
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
		padding := funcs.CalculPadding(alignement, linlen, terlen)
		funcs.Printfinal(word, lettres, padding)
	} else if alignement == "--align=justify" {
		x:=funcs.JustifyText(victim,terlen)
		word := funcs.SplitNl(x)
		funcs.Printfinal(word,lettres,0)

	}
	for i := 0; i < terlen; i++ {
		fmt.Printf("*")
	}
}
