package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

type book struct {
	title string
	path  string
}

type histogram struct {
	chars map[rune]int
	mu    sync.Mutex
}

var books = []*book{
	&book{title: "The Iliad", path: "../data/the-iliad.txt"},
	&book{title: "The Underground Railroad", path: "../data/the-underground-railroad.txt"},
	&book{title: "Pride and Prejudice", path: "../data/pride-and-prejudice.txt"},
	&book{title: "The Republic", path: "../data/the-republic.txt"},
	&book{title: "My Bondage and My Freedom", path: "../data/my-bondage-and-my-freedom.txt"},
	&book{title: "War and Peace", path: "../data/war-and-peace.txt"},
	&book{title: "Moby Dick", path: "../data/moby-dick.txt"},
	&book{title: "Meditations", path: "../data/meditations.txt"},
}

func main() {
	log.SetFlags(log.Ltime | log.Lmicroseconds)
	log.Println("starting...")

	h := histogram{chars: make(map[rune]int)}

	var wg sync.WaitGroup
	wg.Add(len(books))

	for _, b := range books {
		go func(b *book) {
			defer wg.Done()

			log.Printf("Processing %s...", b.title)

			file, err := os.Open(b.path)
			if err != nil {
				log.Println(err)
				//continue
				return
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			//		scanner.Split(bufio.ScanLines)
			for scanner.Scan() {
				for _, c := range scanner.Text() {
					h.mu.Lock()
					h.chars[c]++
					h.mu.Unlock()
				}
			}

			log.Printf("Done with %s", b.title)
		}(b)
	}
	wg.Wait()
	printHist(&h)
}

func printHist(h *histogram) {
	for k, v := range h.chars {
		fmt.Printf("%q=%d, ", k, v)

	}
}
