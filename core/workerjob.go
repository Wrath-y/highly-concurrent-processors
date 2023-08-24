package core

import "fmt"

var (
	MaxWorker = 100
	MaxQueue  = 100
)

// Job 表示要运行的作业
type Job struct {
	Payload Payload
}

// JobQueue 在 Job 队列上发送工作请求的缓冲通道。
var JobQueue chan Job

// Worker 代表执行作业的 Worker。
type Worker struct {
	WorkerPool chan chan Job
	// 无缓冲 Job 通道
	JobChan chan Job
	quit    chan struct{}
}

func NewWorker(workerPool chan chan Job) *Worker {
	return &Worker{
		WorkerPool: workerPool,
		JobChan:    make(chan Job),
		quit:       make(chan struct{})}
}

// Start 方法为 Worker 启动循环监听。
func (w *Worker) Start() {
	go func() {
		for {
			// 将当前 woker 注册到工作队列中。
			w.WorkerPool <- w.JobChan

			select {
			case job := <-w.JobChan:
				// 接收 work 请求。
				if err := job.Payload.Handler(); err != nil {
					fmt.Printf("Playlod Handler err: %s", err.Error())
				}

			case <-w.quit:
				// 接收一个退出的信号。
				return
			}
		}
	}()
}

// Stop 将退出信号传递给 Worker 进程以停止处理清理。
func (w *Worker) Stop() {
	go func() {
		w.quit <- struct{}{}
	}()
}
