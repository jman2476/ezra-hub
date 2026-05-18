package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func MenuRepl(items []string, index int) (string, int, error) {
	if index >= len(items) {
		return "", 0, fmt.Errorf("Index %d is out of range of slice", index)
	}
	scrollReader := bufio.NewReader(os.Stdin)
	var buff []byte

	PrintMenu(items, index)
	for {
		b, err := scrollReader.ReadByte()
		if err != nil {
			fmt.Printf("error reading input: %v", err)
		}
		if int(b) == 13 {
			return items[index], index, nil
		}
		buff = append(buff, b)
		if len(buff) < 3 || buff == nil {
			continue
		}
		if int(buff[len(buff)-3]) == 27 && int(buff[len(buff)-2]) == 91 {
			switch last := rune(buff[len(buff)-1]); last {
			case 'A':
				newIndex := index - 1
				if newIndex >= 0 {
					index = newIndex
				}
				ClearMenu(len(items))
				PrintMenu(items, index)
			case 'B':
				newIndex := index + 1
				if newIndex < len(items) {
					index = newIndex
				}
				ClearMenu(len(items))
				PrintMenu(items, index)
			}
		}
	}
}

func PrintMenu(items []string, index int) {
	fmt.Println("[\u2191]/[\u2193] to navigate, [Enter] to select\r")
	for idx, item := range items {
		margin := "  "
		if idx == index {
			margin = "> "
		}
		fmt.Printf("%s%s\n\r", margin, item)
	}
}

func ClearMenu(num int) {
	var wipe strings.Builder
	for range num + 1 {
		wipe.Write([]byte("\r\033[K\033[A"))
	}

	fmt.Print(wipe.String())
}

func ClearWindow() error {
	_, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return err
	}

	for range height {
		fmt.Print("\r\033[K\033[A")
	}
	return nil
}

func SelectYesNo() (selection bool) {
	var stateYes = ">Yes  No"
	var stateNo = " Yes >No"
	// var clearLine = "\r        \r"
	var current = stateYes

	boolReader := bufio.NewReader(os.Stdin)
	var buff []byte
	fmt.Print(current)
	for {
		selection = stateYes == current
		b, err := boolReader.ReadByte()
		if err != nil {
			fmt.Printf("error reading input: %v", err)
		}
		if int(b) == 13 {
			return
		}
		buff = append(buff, b)
		if len(buff) < 3 || buff == nil {
			continue
		}
		if int(buff[len(buff)-3]) == 27 && int(buff[len(buff)-2]) == 91 {
			switch last := rune(buff[len(buff)-1]); last {
			case 'C':
				fallthrough
			case 'D':
				if selection {
					current = stateNo
				} else {
					current = stateYes
				}
				EraseCharacters(len(current))
				fmt.Print(current)
			default:
				continue
			}
		}
	}
}

func EraseCharacters(chars int) {
	var clear strings.Builder
	for range chars {
		clear.Write([]byte("\033[D"))
	}

	fmt.Print(clear.String())
}
