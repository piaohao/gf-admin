package test

import (
	"fmt"
	"github.com/gogf/gf/g/os/grpool"
	"github.com/gogf/gf/g/os/gtime"
	"github.com/gogf/gf/g/os/gtimer"
	"sync"
	"testing"
	"time"
)

func job() {
	time.Sleep(1 * time.Second)
}

func TestPool(t *testing.T) {
	pool := grpool.New(100)
	for i := 0; i < 1000; i++ {
		pool.Add(job)
	}
	fmt.Println("worker:", pool.Size())
	fmt.Println("  jobs:", pool.Jobs())
	gtimer.SetInterval(time.Second, func() {
		fmt.Println("worker:", pool.Size())
		fmt.Println("  jobs:", pool.Jobs())
		fmt.Println()
	})

	select {}
}

func TestPool2(t *testing.T) {
	start := gtime.Millisecond()
	//wg := sync.WaitGroup{}
	pool := grpool.New(1)
	for i := 0; i < 10000000; i++ {
		//wg.Add(1)
		pool.Add(func() {
			time.Sleep(time.Millisecond)
			//wg.Done()
		})
	}
	//wg.Wait()
	fmt.Println(pool.Size())
	fmt.Println("time spent:", gtime.Millisecond()-start)
}

func TestNativePool(t *testing.T) {
	start := gtime.Millisecond()
	wg := sync.WaitGroup{}
	for i := 0; i < 10000000; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(time.Millisecond)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("time spent:", gtime.Millisecond()-start)
}
