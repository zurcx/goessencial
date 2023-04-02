package main

import "fmt"

type Storage interface {
	download()
}

type GoogleDrive struct{}

func (g GoogleDrive) download() {
	fmt.Println("Downloading from Google Drive")
}
func Downloader(storage Storage) {
	storage.download()
}

type DropBox struct{}

func (g DropBox) download() {
	fmt.Println("Downloading from DropBox")
}

func printWithType(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("Int")

	case string:
		fmt.Println("String")

	default:
		fmt.Println("?")

	}
}

func printWithTypeGenirics[T any](value T) {
	fmt.Println(value)
}

func main() {
	//storage := GoogleDrive{}
	storage := DropBox{}
	Downloader(storage)
	printWithType(10)
	printWithType("teste")
	printWithType(1.90)
	printWithTypeGenirics(3)
}
