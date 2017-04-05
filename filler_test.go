package filler

import (
	"testing"
)

var demoFiller = Filler{
	Tag: "demoFiller1",
	Fn: func(obj interface{}) (interface{}, error) {
		return "hello", nil
	},
}

type demoStruct struct {
	Name string `fill:"demoFiller1"`
	Val  string `fill:"demoFiller2"`
}

// RegFiller - register new filler into []fillers
func TestRegFiller(t *testing.T) {
	RegFiller(demoFiller)
	v1, err1 := fillers[0].Fn("hello")
	v2, err2 := demoFiller.Fn("hello")
	if fillers[0].Tag != demoFiller.Tag || v1 != v2 || err1 != err2 {
		t.FailNow()
	}
}

// Fill - fill the object with all the current fillers
func TestFill(t *testing.T) {
	RegFiller(demoFiller)
	m := demoStruct{
		Name: "nameVal",
		Val:  "valVal",
	}
	// check non ptr - should panic
	func() {
		defer func() {
			if err := recover(); err != nil {
				return
			}
		}()
		Fill(m)
		t.FailNow()
	}()
	// check if got filled
	Fill(&m)
	// should be filled
	if m.Name != "hello" || m.Val != "valVal" {
		t.FailNow()
	}
}