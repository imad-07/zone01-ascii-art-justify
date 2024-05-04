package funcs

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"unsafe"
)

func SplitNl(str string) []string {
	word := ""
	splitedword := []string{}
	skip := false
	for i := 0; i < len(str); i++ {
		if skip {
			skip = false
			continue
		}
		if i != len(str)-1 && str[i] == '\\' && str[i+1] == 'n' {
			if word != "" {
				splitedword = append(splitedword, word)
			}
			word = ""
			skip = true
			splitedword = append(splitedword, "\n")
			continue
		}
		word = word + string(str[i])
	}
	splitedword = append(splitedword, word)
	return splitedword
}

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

func CheckArgs(sstr []string) {
	kinds := []string{"standard", "shadow", "thinkertoy"}
	alignments := []string{"--align=justify", "--align=right", "--align=left", "--align=center"}
	if len(sstr) < 3 {
		log.Fatal("please provide me with 3 arguments")
		return
	}

	for _, arg := range sstr {
		for _, char := range arg {
			if char < 32 || char > 126 {
				log.Fatal("I only accept chars between 32 and 126 on ASCII.")
			}
		}
	}

	if !contains(alignments, sstr[0]) {
		log.Fatal("the alignment should be either one of these: --align=justify, --align=right, --align=left, --align=center")
	}

	if !contains(kinds, sstr[2]) {
		log.Fatal("the type should be either one of these: standard, shadow, thinkertoy")
	}
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func TerminalSize() (width int, err error) {
	// Get the file descriptor for stdout
	fd := int(os.Stdout.Fd())

	// Create a new termios struct
	var dimensions [4]uint16

	// Use the TIOCGWINSZ ioctl to get terminal dimensions
	if _, _, err := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(fd),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&dimensions)),
	); err != 0 {
		return 0, err
	}

	// Extract width and height from dimensions array
	width = int(dimensions[1])

	return width, nil
}

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

func Printfinal(str []string, lettres [][]string, padding int) {
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
					fmt.Printf("%*s", padding+len(lettres[str[l][j]-32][i])/2, lettres[str[l][j]-32][i])
				} else {
					fmt.Printf(lettres[str[l][j]-32][i])
				}
			}
			fmt.Print("\n")
		}
		bl = true
	}
}

func CalculPadding(str string, wwidth, twidth int) int {
	var padding int
	if str == "--align=left" {
		padding = 0
	} else if str == "--align=center" {
		padding = (twidth - wwidth) / 2
	} else if str == "--align=right" {
		padding = (twidth - wwidth) + 2
	}
	return padding
}
