package main

import (
	"context"
	"fmt"
	"time"
)

type paramKey struct{}

func main() {
	c := context.WithValue(context.Background(), paramKey{}, "abc")
	c, cancel := context.WithTimeout(c, 5*time.Second)

	go mainTask(c)

	var cmd string
	for {
		fmt.Scan(&cmd)
		if cmd == "c" {
			cancel()
		}
	}
}

func mainTask(c context.Context) {
	fmt.Printf("main task started with param %q\n", c.Value(paramKey{}))

	go func() {
		c1, cancel := context.WithTimeout(c, time.Second*2)
		defer cancel()
		smallTask(c1, "task1", 4*time.Second)
	}()
	smallTask(c, "task2", 2*time.Second)
	go smallTask(context.Background(), "task3", 10*time.Second)
}

func smallTask(c context.Context, name string, timeout time.Duration) {
	fmt.Println("%s started with param %v %q\n", name, timeout, c.Value(paramKey{}))
	select {
	case <-time.After(timeout * time.Second):
		fmt.Printf("%s done\n", name)
	case <-c.Done():
		fmt.Printf("%s cancelled\n", name)
	}
}
