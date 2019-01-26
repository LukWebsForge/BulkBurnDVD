package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

const IsoFile string = "/home/lukas/Documents/dvd.iso"

func main() {
	driveZero := DvdDrive{
		number: "0",
		speed:  3,
	}
	driveOne := DvdDrive{
		number: "1",
		speed:  3,
	}
	driveTwo := DvdDrive{
		number: "2",
		speed:  6,
	}

	go writeLoop(&driveZero)
	go writeLoop(&driveOne)
	go writeLoop(&driveTwo)

	wait := sync.WaitGroup{}
	wait.Add(1)
	wait.Wait()
}

func printOpen(drive *DvdDrive) {
	open, _ := drive.isTrayOpen()
	fmt.Printf("The drive at %v: %v\n", drive.file(), open)
}

func writeLoop(drive *DvdDrive) {
	for true {
		for true {
			open, err := drive.isTrayOpen()
			if err != nil {
				log.Printf("[%v] Error while checking tray status: %v\n", drive.file(), err)
			}

			if !open {
				break
			}

			time.Sleep(time.Duration(200 * time.Millisecond))
		}

		log.Printf("[%v] Tray closed. Starting write...\n", drive.file())

		time.Sleep(time.Duration(30 * time.Second))

		writeLog, err := drive.write(IsoFile)
		if err != nil {
			log.Printf("[%v] Write failed: %v\n%v", drive.file(), err, writeLog)
		} else {
			log.Printf("[%v] Write successful. Waiting for next disk...", drive.file())
		}
	}
}
