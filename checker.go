package main

import (
	"fmt"
	"strings"
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
	items []Checkable
}

func NewChecker() *Checker {
	return &Checker{}
}

func (c *Checker) Check() {
	for _, v := range c.items {
		if !v.Health() {
			fmt.Println(v.GetID(), "не работает")
		}
	}
}

func (c *Checker) String() string {
	var sl []string
	for _, v := range c.items {
		sl = append(sl, v.GetID())
	}

	return strings.Join(sl, " ")
}

func (c *Checker) Add(newItems ...Checkable) {
	c.items = append(c.items, newItems...)
}
