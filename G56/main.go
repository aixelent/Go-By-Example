package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type Sync struct {
	Mutex  sync.Mutex
	Writer io.Writer
}

func (s *Sync) Write(b []byte) (n int, err error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	return s.Writer.Write(b)
}

var Data = []string{
	"A",
	"B",
	"C",
	"D",
	"E",
	"F",
}

func main() {
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	write := &Sync{sync.Mutex{}, f}
	wg := sync.WaitGroup{}

	for _, val := range Data {
		wg.Add(1)
		go func(names string) {
			fmt.Fprintln(write, names)
			wg.Done()
		}(val)
	}
	wg.Wait()
}
