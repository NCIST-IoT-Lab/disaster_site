package utils

import (
	"sync"
)

// Group 是一个并发控制工具，可以方便地管理多个并发任务
type Group struct {
	wg      sync.WaitGroup
	errOnce sync.Once
	err     error
}

// Go 启动一个新的并发任务
// 该函数接收一个函数作为参数，该函数将在新的 goroutine 中执行
func (g *Group) Go(f func() error) {
	g.wg.Add(1)

	go func() {
		defer g.wg.Done()

		if err := f(); err != nil {
			g.errOnce.Do(func() {
				g.err = err
			})
		}
	}()
}

// Wait 等待所有并发任务完成
// 该函数会阻塞直到所有通过 Go 方法启动的 goroutine 都完成
// 返回遇到的第一个错误（如果有）
func (g *Group) Wait() error {
	g.wg.Wait()
	return g.err
}