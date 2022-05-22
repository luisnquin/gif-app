package utils

import (
	"os"
	"os/signal"
	"syscall"
)

func HandleExit(queueSize int) chan func() {
	queue := make(chan func(), queueSize)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGKILL)

	go func() {
		for range c {
			for task := range queue {
				task()
			}

			os.Exit(0)
		}
	}()

	return queue
}

func IsRunningInADockerContainer() bool {
	_, err := os.Stat("/.dockerenv")

	return err == nil
}
