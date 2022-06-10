package syncc

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// TestWaitGroup 测试的是自带的WaitGroup
func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		// 每次done计数都会-1
		defer wg.Done()
		time.Sleep(50 * time.Millisecond)
		fmt.Println("go1执行结束")
	}()
	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond)
		fmt.Println("go2执行结束")
	}()
	// 等待执行结束
	wg.Wait()
	fmt.Println("全部执行结束")
}
