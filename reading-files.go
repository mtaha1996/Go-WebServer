package main

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// returns the slice of words
func words(file []uint8) []string {
	var words []string

	pos := 0
	for i := 0; i < len(file); {
		switch string(file[i]) {
		case "\n":
			fmt.Println(string(file[pos:i]))
			words = append(words, string(file[pos:i]))
			i++
			pos = i

		case " ":
			fmt.Println(string(file[pos:i]))
			words = append(words, string(file[pos:i]))
			i++
			pos = i

		case ",":
			fmt.Println(string(file[pos:i]))
			words = append(words, string(file[pos:i]))
			i++
			pos = i

		case ";":
			fmt.Println(string(file[pos:i]))
			words = append(words, string(file[pos:i]))
			i++
			pos = i

		case "'":
			fmt.Println(string(file[pos:i]))
			words = append(words, string(file[pos:i]))
			i++
			pos = i

		case ".":
			switch string(file[i+1]){
			case "\n":
				fmt.Println("helll yeah!")
				fmt.Println(string(file[pos:i]))
				words = append(words, string(file[pos:i-1]))
				i+=2
				pos = i

			case " ":
				fmt.Println(string(file[pos:i]))
				words = append(words, string(file[pos:i-1]))
				i+=2
				pos = i
			default:
				i++
			}

		default:
			i++
		}
	}
	words = append(words, string(file[pos:]))




	return words

}

// returns the slice of sentences
func sent(file []uint8) []string {
	var sents []string

	pos := 0
	for i := 0; i < len(file); {
		switch string(file[i]) {
		case ".":
			switch string(file[i+1]){
			case "\n":
				fmt.Println(string(file[pos:i]))
				sents = append(sents, string(file[pos:i]))
				i++
				pos = i
			case " ":
				fmt.Println(string(file[pos:i]))
				sents = append(sents, string(file[pos:i]))
				i++
				pos = i
			default:
				i++
			}

		default:

			i++
		}
	}

	return sents

}

// returns the slice of lines
func Lines(file []uint8) []string {
	var lines []string

	pos := 0
	for i := 0; i < len(file); {
		switch string(file[i]) {
		case "\n":
			fmt.Println(string(file[pos:i]))
			lines = append(lines, string(file[pos:i]))
			i++
			pos = i

		default:
			i++
		}
	}
	lines = append(lines, string(file[pos:]))
	return lines

}

func main() {
	dat, err := ioutil.ReadFile("tmp/dat")
	check(err)
	//fmt.Printf("type of dat is %v\n", dat)
	//fmt.Println(string(101))
	//
	//some := Lines(dat)
	//fmt.Println(some[2])
	fmt.Println(string(10))

	for i := range words(dat){
		fmt.Println(i)
	}

}
