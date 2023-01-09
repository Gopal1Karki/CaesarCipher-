package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func encryption() {
	clearScreen()
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a sentence you want to encrypt: ")
	sentence, _ := reader.ReadString('\n')
	var key int
	fmt.Printf("Enter the key: ")
	fmt.Scanln(&key)
	var text int

	var b = make([]string, len(sentence))
	for i := 0; i < len(sentence); i++ {
		if int(sentence[i]) >= 97 && int(sentence[i]) <= 122 {
			text = int(sentence[i]) + key
			if text > 122 {
				a := text - 122
				text = 96 + a
				b[i] = string(text)
			}
			b[i] = string(text)
		} else {
			text = int(sentence[i])
			b[i] = string(text)
		}
	}
	fmt.Println("The Encrypted key text: ")
	fmt.Println(b)
A:
	var ex string
	fmt.Printf("Press 'e' to exit or 'c' to return to mainmenu!! ")
	fmt.Scanln(&ex)

	if ex == "e" {
		os.Exit(1)
	} else if ex == "c" {
		clearScreen()
		main()
	} else {
		fmt.Println("Invalid Input!! Enter the valid input!")
		goto A
	}
}

func decryption() {
	clearScreen()
	clearScreen()
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a text you want to decrypt: ")
	sentence, _ := reader.ReadString('\n')
	var key int
	fmt.Printf("Enter the key")
	fmt.Scanln(&key)
	var text int

	var b = make([]string, len(sentence))
	for i := 0; i < len(sentence); i++ {
		if int(sentence[i]) >= 97 && int(sentence[i]) <= 122 {
			text = int(sentence[i]) - key
			if text < 97 {
				a := 97 - text
				text = 123 - a
				b[i] = string(text)
			}
			b[i] = string(text)
		} else {
			text = int(sentence[i])
			b[i] = string(text)
		}
	}
	fmt.Println("The Decrypted key text: ")
	fmt.Println(b)
A:
	var ex string
	fmt.Printf("Press 'e' to exit or 'c' to return to mainmenu!! ")
	fmt.Scanln(&ex)

	if ex == "e" {
		os.Exit(1)
	} else if ex == "c" {
		clearScreen()
		main()
	} else {
		fmt.Println("Invalid Input!! Enter the valid input!")
		goto A
	}

}
func main() {
	clearScreen()
A:

	fmt.Println("         C A E S A R  C I P H E R           ")

	fmt.Printf("\n\n")
	fmt.Println("Press 1 for data Encryption ")
	fmt.Println("Press 2 for data Decryption ")
	fmt.Println("Press 3 to exit! ")

	fmt.Printf("\n\n")

	var inp int
	fmt.Printf("Enter your choice: ")
	fmt.Scanln(&inp)
	if inp == 1 {
		encryption()
	} else if inp == 2 {
		decryption()
	} else if inp == 3 {
		os.Exit(1)
	} else {
		fmt.Println("Invalid Input ! Enter the valid Input !!")
		time.Sleep(1 * time.Second)
		goto A
	}

}
