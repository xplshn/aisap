package aisap

import (
	"os"
	"os/exec"
	"strconv"
	"errors"
	"path"
	"path/filepath"
)

// mount mounts the requested AppImage (src) to the destination
// directory (dest)
// Quick, hacky implementation, ideally this should be redone using the
// squashfuse library
func mount(src string, dest string, offset int) error {
	var squashfuse string

	e, _ := os.Executable()

	// Make sure squashfuse exists
	if squashfuse, err = exec.LookPath("squashfuse"); err != nil {
		squashfuse, err = exec.LookPath(filepath.Join(path.Dir(e), "squashfuse"))
		if err != nil {
			return errors.New("failed to find squashfuse binary! Cannot mount AppImage")
		}
	}

	// Convert the offset to a string and mount using squashfuse
	o := strconv.Itoa(offset)
	mnt = exec.Command(squashfuse, "-o", "offset=" + o, src, dest)

	return mnt.Run()
}

// Unmounts an AppImage
func Unmount(ai *AppImage) error {
	if (ai == nil) {
		return errors.New("AppImage is nil")
	} else if ai.Path == "" {
		return errors.New("AppImage contains no path")
	}

	err = unmountDir(ai.MountDir())
	if err != nil { return err }

	// Clean up
	err = os.RemoveAll(ai.TempDir())

	return err
}

// Unmounts a directory (lazily in case the process is finishing up)
func unmountDir(mntPt string) error {
	var umount *exec.Cmd

	if _, err := exec.LookPath("fusermount"); err == nil {
		umount = exec.Command("fusermount", "-uz", mntPt)
	} else {
		umount = exec.Command("umount", "-l", mntPt)
	}

	// Run unmount command, returning the stdout if failed
	out, err := umount.CombinedOutput()
	if err != nil {
		err = errors.New(string(out))
	}

	return err
}
