package main

import "time"

var (
	systemTimeout = 500 * time.Second
)

func main() {
	getLogger().Info("Starting test run")

	loadConfig()

	mainLoop()

	finishTestRun()

}

func mainLoop() {

	// Listen for new devices on the network
	go startDeviceListener()

	time.Sleep(systemTimeout)

}

func finishTestRun() {
	getLogger().Info("Test run has finished")
}
