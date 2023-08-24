package dispatcher

import "highly-concurrent-processors/core"

type Dispatcher struct {
	// 通过调度器注册一个 Worker 通道池, 每个 Worker 拥有对 WorkerPool 的引用
	WorkerPool chan chan core.Job
	maxWorker  int
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	workerPool := make(chan chan core.Job, maxWorkers)
	return &Dispatcher{WorkerPool: workerPool, maxWorker: maxWorkers}
}

func (d *Dispatcher) Run() {
	// 启动指定数量的 Worker
	for i := 0; i < d.maxWorker; i++ {
		worker := core.NewWorker(d.WorkerPool)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-core.JobQueue:
			// 接收一个 job 请求
			go func(job core.Job) {
				// 尝试获取可用的 chan Job
				// 这将阻塞 worker 直到拥有空闲 chan Job
				jobChan := <-d.WorkerPool

				// 调度一个 job 到 worker job 通道
				jobChan <- job
			}(job)
		}
	}
}
