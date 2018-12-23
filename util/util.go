package util

import (
	"math/rand"
	"time"
	"math"
)

func Rand_generator() chan int{
	out:=make(chan int)
	go func(){
		for{
			rand.Seed(time.Now().Unix())
			out <- rand.Intn(math.MaxInt8)

		}
	}()
	return out
}

