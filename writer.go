package main

import (
	"log"
	"os"
	"sync"
	"time"
)

// TODO: Set your own path the ISO file you want to burn
const IsoFile string = "/home/lukas/Documents/dvd.iso"

// https://wiki.debian.org/CDDVD
func main() {

	// TODO: Customize the settings of your DvdDrive
	// You can add more drives by copying the following lines and changing the properties
	go writeLoop(&DvdDrive{
		// You can find out the id of your drive via 'ls /dev/disk/by-id/'.
		// If you've got no idea how your drive could be called
		// use a burn program with a GUI and check the available drives.
		id: "ata-Optiarc_DVD_RW_AD-7740H",
		// You can find out the write speed of your drive via the device documentation.
		// Just google the you've got from 'ls /dev/disk/by-id/'
		// DIRTY HACK: Use a burn program with a GUI like XfBurn
		// and look for the maximum speed you can select
		speed: 8,
	})

	// You're done. Run the program with 'go run ./..'

	log.Println("Ready. Put your DVDs into the configured drives and let them burn!")

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
