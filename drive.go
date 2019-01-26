package main

import "C"
import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

type DvdDrive struct {
	number string
	speed  int
}

func (d *DvdDrive) file() string {
	return "/dev/sr" + d.number
}

func (d *DvdDrive) isTrayOpen() (bool, error) {
	/* cmd := exec.Command("isoinfo", "-i"+d.file())
	out, err := cmd.CombinedOutput()

	if err != nil {
		if err.Error() != "exit status 1" {
			return false, err
		}
	}

	return strings.Contains(string(out), "tray open"), nil */
	return isNativeTrayOpen(d.file()), nil
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
