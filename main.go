package main

import (
	"fmt"
	"time"
)

type name struct{}

type data interface {
	first()
	second()
}

func main() {
	for i := 0; i < 8; i++ {
		d := name{}
		if i < 4 {
			go d.first()
			continue
		}
		go d.second()
	}
	time.Sleep(time.Second * 2)
}

func (d *name) first() {
	fmt.Println("data1")
}

func (d *name) second() {
	fmt.Println("data2")
}
