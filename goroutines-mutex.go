package main

import (
	"fmt"
	"sync"
	"time"
)

type name struct{}

type data interface {
	first()
	second()
}

func main() {
	var m sync.Mutex
	var w sync.WaitGroup

	go func() {
		for i := 0; i < 4; i++ {
			w.Add(1)
			m.Lock()
			d := name{}
			d.first()
			m.Unlock()
			w.Done()
		}
	}()
	go func() {
		for i := 0; i < 4; i++ {
			w.Add(1)
			m.Lock()
			d := name{}
			d.second()
			m.Unlock()
			w.Done()
		}
	}()
	time.Sleep(time.Second * 2)
	w.Wait()
}

func (d *name) first() {
	fmt.Println("data1")
}

func (d *name) second() {
	fmt.Println("data2")
}
