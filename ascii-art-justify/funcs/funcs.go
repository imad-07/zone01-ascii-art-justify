package funcs

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// this func ios gonna split the input by newlines
func SplitNl(str string) []string {
	word := ""
	splitedword := []string{}

	for i := 0; i < len(str); i++ {
		if i != len(str)-1 && str[i] == '\\' && str[i+1] == 'n' {
			if word != "" {
				splitedword = append(splitedword, word)
			}
			word = ""
			i++
			splitedword = append(splitedword, "\n")
			continue
		}
		word = word + string(str[i])
	}
	splitedword = append(splitedword, word)
	return splitedword
}

// this is gonna get the letters from the file depending on the input
func GetLettres(fileContent []byte) [][]string {
	lettres := [][]string{}
	lettre := []string{}
	line := []byte{}
	s := ""
	for i := 0; i < len(fileContent); i++ {
		if fileContent[i] != 13 {
			s = s + string(fileContent[i])
		}
	}
	for i := 0; i < len(s); i++ {
		if i != len(s)-1 && s[i] == '\n' && s[i+1] == '\n' {
			lettre = append(lettre, string(line))
			lettres = append(lettres, lettre)
			lettre = nil
			line = nil
			continue
		}
		if s[i] == '\n' {
			lettre = append(lettre, string(line))
			line = nil
			continue
		}
		line = append(line, s[i])
	}
	lettres = append(lettres, lettre)
	return lettres
}

// this is gonna test the arguments given by the user to check the validity
func CheckArgs(sstr []string) {
	kinds := []string{"standard", "shadow", "thinkertoy"}
	alignments := []string{"--align=justify", "--align=right", "--align=left", "--align=center"}
	if len(sstr) < 3 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nExample: go run . --align=right something standard")
		os.Exit(0)
	}

	for _, arg := range sstr {
		for _, char := range arg {
			if char < 32 || char > 126 {
				log.Fatal("I only accept chars between 32 and 126 on ASCII.")
			}
		}
	}

	if !contains(alignments, sstr[0]) {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nExample: go run . --align=right something standard")
		os.Exit(0)
	}

	if !contains(kinds, sstr[2]) {
		log.Fatal("the type should be either one of these: standard, shadow, thinkertoy")
	}
}

// this func returns true if a string is contained in a slice of str
func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

// get the terminal size
func TerminalSize() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 0
	}
	parts := strings.Split(strings.TrimSpace(string(out)), " ")
	width, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0
	}
	return width
}

// this func returns the len of a slice
func SliceLen(str [][]string) int {
	result := 0
	for i := 0; i < len(str); i++ {
		for j := 0; j < len(str[i]); j++ {
			x := []rune(str[i][j])
			result += len(x)
		}
	}
	return result
}

// this is the func that is gonna print the final result
func Printfinal(str []string, lettres [][]string) {
	bl := false
	for l := 0; l < len(str); l++ {
		if str[l] == "" {
			continue
		}
		if str[l] == "\n" {
			if l == len(str)-1 {
				continue
			}
			if bl && str[l+1] != "\n" {
				continue
			}
			fmt.Printf("\n")
			continue
		}
		for i := 1; i < 9; i++ {
			for j := 0; j < len(str[l]); j++ {
				if j == 0 {
					fmt.Printf(lettres[str[l][j]-32][i])
				} else {
					fmt.Printf(lettres[str[l][j]-32][i])
				}
			}
			fmt.Print("\n")
		}
		bl = true
	}
}

// this is gonna calculate the needed padding to add depending on the alignement
func Padding(alignement, victim string, linlen, terlen int) string {
	if alignement == "--align=left" {
		return victim
	} else if alignement == "--align=center" {
		needed := ((terlen - linlen) / 2) / 6
		padding := ""
		for i := 0; i < needed; i++ {
			padding += " "
		}
		victim = padding + victim
	} else if alignement == "--align=right" {
		needed := (terlen - linlen) / 6
		padding := ""
		for i := 0; i < needed; i++ {
			padding += " "
		}
		victim = padding + victim
	}
	return victim
}

// this is gonna justify the input
func JustifyText(victim string, terlen, linlen int) string {
	alls := terlen - linlen
	words := strings.Fields(victim)
	numSpacesNeeded := len(words) - 1
	var justifiedText string
	if numSpacesNeeded != 0 {
		eachs := (alls / numSpacesNeeded) / 6
		extraSpaces := (alls % numSpacesNeeded) / 6
		for i, word := range words {
			justifiedText += word
			if i < len(words)-1 {
				for j := 0; j <= eachs; j++ {
					justifiedText += " "
				}
				// Add extra spaces if any
				if extraSpaces > 0 {
					justifiedText += " "
					extraSpaces--
				}
			}
		}
	}
	return justifiedText
}
