package main

import "C"
import (
	"os"
	"os/exec"
	"strconv"
	"strings"
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

// Tested, but not in use
func (d *DvdDrive) md5CheckDisk(isoFile string) (bool, error) {
	info, err := os.Stat(isoFile)
	if err != nil {
		return false, err
	}

	blocks := int(info.Size() / 2048)
	cmd := exec.Command("/bin/bash", "-c", "dd if="+d.file()+" bs=2048 count="+strconv.Itoa(blocks)+" | md5sum")
	outDvd, err := cmd.CombinedOutput()
	if err != nil {
		return false, err
	}

	outFile, err := exec.Command("md5sum", isoFile).CombinedOutput()
	if err != nil {
		return false, err
	}

	hash := strings.Split(string(outFile), " ")[0]
	return strings.Contains(string(outDvd), hash), nil
}
