package main

import (
	"bytes"
	"log"
	"syscall"
	"testing"
	"time"
)

func TestMain(t *testing.T) {

	t.Run("Test Main Function", func(t *testing.T) {
		var buffer bytes.Buffer
		log.SetOutput(&buffer)
		go main()
		time.Sleep(1 * time.Millisecond)
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	})
}

func TestExit(t *testing.T) {

	t.Run("Test Exit Function", func(t *testing.T) {
		var buffer bytes.Buffer
		log.SetOutput(&buffer)
		go exit()
		time.Sleep(1 * time.Millisecond)
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	})
}
