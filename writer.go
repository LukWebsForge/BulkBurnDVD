package main

import (
	"log"
	"os"
	"sync"
	"time"
)

const IsoFile string = "/home/lukas/Documents/dvd.iso"

// https://wiki.debian.org/CDDVD
func main() {
	// https://www.cnet.com/products/sony-optiarc-ad-7740h-dvdrw-r-dl-dvd-ram-drive-serial-ata-series
	driveZero := DvdDrive{
		id:    "ata-Optiarc_DVD_RW_AD-7740H",
		speed: 8, // 8
	}
	// https://images-eu.ssl-images-amazon.com/images/I/71duhj7GTNS.pdf
	driveOne := DvdDrive{
		id:    "usb-TSSTcorp_CDDVDW_SE-208DB_R8X76GAC902VPR-0:0",
		speed: 8, // 8
	}

	go writeLoop(&driveZero)
	go writeLoop(&driveOne)

	wait := sync.WaitGroup{}
	wait.Add(1)
	wait.Wait()
}

func writeLoop(drive *DvdDrive) {
	l := log.New(os.Stderr, "["+drive.file()+"] ", log.Ltime)
	for true {
		for true {
			open, err := drive.isTrayOpen()
			if err != nil {
				l.Printf("Error while checking tray status: %v\n", err)
			}

			if !open {
				break
			}

			time.Sleep(time.Duration(200 * time.Millisecond))
		}

		l.Println("Tray closed. Starting write...")

		time.Sleep(time.Duration(30 * time.Second))

		open, _ := drive.isTrayOpen()
		if open {
			l.Println("Aborting write, because tray is open")
			continue
		}

		writeLog, err := drive.write(IsoFile)
		if err != nil {
			l.Printf("Write failed: %v\n%v", err, writeLog)
			_ = drive.openTray()
		} else {
			l.Println("Write successful. Waiting for next disk...")
		}

		// Without the delay the tray could be still closed
		time.Sleep(time.Duration(time.Second))
	}
}
