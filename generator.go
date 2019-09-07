package main

import (
	"math/rand"
	"time"
)

type Generator interface {
	Next() (int64, error)
}

func NewGenerator(start, end int64) Generator {
	g := &generator{
		Random: rand.New(rand.NewSource(time.Now().Unix())),
		Start:  start,
		End:    end + 1,
	}

	return g
}

type generator struct {
	Random *rand.Rand
	Start  int64
	End    int64
}

func (g *generator) Next() (int64, error) {
	return g.Random.Int63n(g.End-g.Start) + g.Start, nil
}
