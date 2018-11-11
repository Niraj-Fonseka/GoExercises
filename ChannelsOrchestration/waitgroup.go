package main

import "sync"

type WorkerGroup struct {
	WaitGroup   sync.WaitGroup
	RespnseChan chan string
}

func NewWorkerGroup() *WorkerGroup {

	worker := WorkerGroup{
		WaitGroup: sync.WaitGroup,
		RespnseChan : make(chan string)
	}

	return &worker
}
