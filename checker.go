package main

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

type Measurable interface {
	GetMetrics() string
}

type Checkable interface {
	Measurable
	Ping() error
	GetID() string
	Health() bool
}

type Checker struct {
	items      []Checkable
	ch         chan []Checkable
	getItems   chan []Checkable
	ctx        context.Context
	cancelFunc context.CancelFunc
	exit       chan struct{}
}

func NewChecker(ctx context.Context) *Checker {
	cancel, cancelFunc := context.WithCancel(ctx)
	return &Checker{
		items:      nil,
		ch:         make(chan []Checkable),
		getItems:   make(chan []Checkable),
		ctx:        cancel,
		cancelFunc: cancelFunc,
		exit:       make(chan struct{}),
	}
}

func (c *Checker) String() string {
	var sl []string
	for _, v := range <-c.getItems {
		sl = append(sl, v.GetID())
	}

	return strings.Join(sl, " ")
}

func (c *Checker) Add(newItems ...Checkable) {
	ch := make(chan struct{})
	defer close(ch)

	go func() {
		c.ch <- newItems
		ch <- struct{}{}
	}()

	<-ch
}

func (c *Checker) Run() {
	ticker := time.NewTicker(5 * time.Second)
	exit := make(chan struct{})
	defer func() {
		ticker.Stop()
		exit <- struct{}{}
		fmt.Println("Закончили Run")
	}()

	go func() {
		for {
			select {
			case newItem := <-c.ch:
				c.items = append(c.items, newItem...)
			case c.getItems <- c.items:
			case <-exit:
				close(exit)

				c.exit <- struct{}{}

				return
			}
		}
	}()

	func() {
		for {
			select {
			case tick := <-ticker.C:
				fmt.Println(tick)
				c.Check()
			case <-c.ctx.Done():
				fmt.Println("Процесс прерван")
				return
			}
		}
	}()
}

func (c *Checker) Stop() {
	defer close(c.getItems)
	defer close(c.ch)
	defer close(c.exit)
	c.cancelFunc()
	<-c.exit
	fmt.Println("Потушили микросервисы", runtime.NumGoroutine())
}

func (c *Checker) Check() {
	wg := sync.WaitGroup{}
	for _, v := range <-c.getItems {
		wg.Add(1)
		v := v
		go func() {
			defer func() {
				r := recover()
				if r != nil {
					err := fmt.Errorf("ошибка: %v", r)
					fmt.Println(err)
				}
			}()
			defer wg.Done()

			if !v.Health() {
				if v.GetID() == "666" {
					panic("Hello err")
				}
				fmt.Println(v.GetID(), "не работает")
			}
		}()
	}
	wg.Wait()
}
