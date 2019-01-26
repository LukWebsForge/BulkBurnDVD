package main

import "C"
import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

type DvdDrive struct {
	id    string
	speed int
}

func (d *DvdDrive) file() string {
	return "/dev/disk/by-id/" + d.id
}

func (d *DvdDrive) isTrayOpen() (bool, error) {
	open, err := isNativeTrayOpen(d.file())
	if err != nil {
		return false, err
	}
	return open, nil
}

func (d *DvdDrive) openTray() error {
	cmd := exec.Command("eject", d.file())
	return cmd.Run()
}

func (d *DvdDrive) write(isoFile string) (string, error) {
	cmd := exec.Command("growisofs", "-Z"+d.file()+"="+isoFile, "-dvd-compat", "-speed="+strconv.Itoa(d.speed))
	output, err := cmd.CombinedOutput()

	if err != nil {
		outStr := ""
		if output != nil {
			outStr = string(output)
		}
		return outStr, err
	}

	return string(output), nil
}

func (d *DvdDrive) md5CheckDisk(isoFile string) (bool, error) {
	info, err := os.Stat(isoFile)
	if err != nil {
		return false, err
	}

	blocks := int64(info.Size() / 2048)
	cmd := exec.Command("bash", "-c 'dd if="+d.file()+" bs=2048 count="+string(blocks)+" | md5sum'")
	outDvd, err := cmd.CombinedOutput()
	if err != nil {
		return false, err
	}

	outFile, err := exec.Command("md5sum " + isoFile).CombinedOutput()
	if err != nil {
		return false, err
	}

	fmt.Printf("DVD Hash: %v\nFile Hash: %v", string(outDvd), string(outFile))

	return true, nil
}
