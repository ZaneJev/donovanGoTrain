package main

/*
В данном решение добавлены в вывод имена файлов,
в которых первоначально встречается повторяющаяся строка
*/

import (
	"bufio"
	"fmt"
	"os"
)

type duplicates struct {
	Filenames []string
	Counter   int
}

func main() {
	counts := make(map[string]duplicates)
	files := os.Args[1:]
	if len(files) == 0 {
		countLinesWithFilename(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLinesWithFilename(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n.Counter > 1 {
			fmt.Printf("%d times repeats string: %s\n", n.Counter, line)
		}
	}
}

func countLinesWithFilename(f *os.File, counts map[string]duplicates) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		count, ok := counts[input.Text()]
		if !ok {
			counts[input.Text()] = duplicates{
				Filenames: []string{f.Name()},
				Counter:   1,
			}
		} else {
			count.Filenames = append(count.Filenames, f.Name())
			count.Counter++

			counts[input.Text()] = count
		}

		if counts[input.Text()].Counter > 1 {
			fmt.Printf("%d files with repeated strings: %v\n", len(counts[input.Text()].Filenames), counts[input.Text()].Filenames)
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}
