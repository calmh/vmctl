// +build solaris

package vmadm

import "os/exec"

func vmadmList() ([]byte, error) {
	return exec.Command("vmadm", "list", "-p").CombinedOutput()
}

func vmadmGet(uuid string) ([]byte, error) {
	return exec.Command("vmadm", "get", uuid).CombinedOutput()
}
