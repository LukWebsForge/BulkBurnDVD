# BulkBurnDVD

This is a simple application written in Go which helps you to burn a lot of DVDs.  
It's written in [Go](https://golang.org/) and only works on Linux.

## How to use

1. Create a ISO with the content of the DVDs.
2. Install [Go](https://golang.org/), 
[growisofs](https://packages.debian.org/de/sid/growisofs) and
[gcc](https://packages.debian.org/de/jessie/gcc)
3. Customize your settings in the `drive.go` file
4. Run the program with `go run ./..`

Now as the program runs, you just put a DVDs in one of the configured drives 
and the program will burn the ISO on the DVD. 
When it's finished it'll eject the drive and you can put a new DVD into it.

But if the drive should eject the DVD after 30s, it's very likely that a error
occurred. Report it via the GitHubs issue.

That's it. It's so simple!

### Question: I want to use multiple drives

Simple, just duplicate the following code (in `writer.go`) and customize the settings.

```go
go writeLoop(&DvdDrive{
    // You can find out the id of your drive via 'ls /dev/disk/by-id/'.
    // If you've got no idea how your drive could be called
    // use a burn program with a GUI and check the available drives.
    id:    "ata-Optiarc_DVD_RW_AD-7740H",
    // You can find out the write speed of your drive via the device documentation.
    // Just google the you've got from 'ls /dev/disk/by-id/'
    // DIRTY HACK: Use a burn program with a GUI like XfBurn
    // and look for the maximum speed you can select
    speed: 8,
})
```