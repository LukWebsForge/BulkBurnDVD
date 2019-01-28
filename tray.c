#include <stdio.h>
#include <stdlib.h>
#include <sys/ioctl.h>
#include <fcntl.h>
#include <linux/cdrom.h>
#include <unistd.h>

// https://askubuntu.com/a/483721
int trayOpen(char *device) {
    int cdrom;

    if (!device) {
        return 2;
    }

    if ((cdrom = open(device, O_RDONLY | O_NONBLOCK)) < 0) {
        return 2;
    }

    /* Check CD tray status */
    if (ioctl(cdrom, CDROM_DRIVE_STATUS) == CDS_TRAY_OPEN) {
        close(cdrom);
        return 0;
    }

    close(cdrom);
    return 1;
}