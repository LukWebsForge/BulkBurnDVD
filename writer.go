package main

import (
	"log"
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
	// https://www.lg.com/us/computer-products/pdf/h_gh22lp20_spec_sheet.pdf
	/* driveOne := DvdDrive{
		id:    "usb-HL-DT-ST_DVD-RAM_GH22LP20-0:0",
		speed: 4, // 22 -> 16
	} */
	// https://images-eu.ssl-images-amazon.com/images/I/71duhj7GTNS.pdf
	driveTwo := DvdDrive{
		id:    "usb-TSSTcorp_CDDVDW_SE-208DB_R8X76GAC902VPR-0:0",
		speed: 8, // 8
	}

	go writeLoop(&driveZero)
	// go writeLoop(&driveOne)
	go writeLoop(&driveTwo)

	driveZero.openTray()
	driveTwo.openTray()

	wait := sync.WaitGroup{}
	wait.Add(1)
	wait.Wait()
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
			drive.openTray()
		} else {
			log.Printf("[%v] Write successful. Waiting for next disk...", drive.file())
		}
	}
}
