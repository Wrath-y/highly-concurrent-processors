package main

import (
	"highly-concurrent-processors/core"
	"highly-concurrent-processors/dispatcher"
	"time"
)

func payloadHandler() {
	list := make([]core.Payload, 3)
	list[0] = core.Payload{
		Data: "111",
	}
	list[1] = core.Payload{
		Data: "222",
	}
	list[2] = core.Payload{
		Data: "333",
	}

	for _, payload := range list {
		// 创建一个有效负载的job
		work := core.Job{Payload: payload}

		// 将 work push 到队列。
		core.JobQueue <- work
	}
}

func main() {
	core.JobQueue = make(chan core.Job, core.MaxQueue)
	d := dispatcher.NewDispatcher(core.MaxWorker)
	d.Run()

	payloadHandler()

	time.Sleep(time.Second * 10)
}
