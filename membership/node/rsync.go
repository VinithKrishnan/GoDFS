package node

import (
	"fmt"
	"os/exec"
	"time"
	"os/user"
	"strings"
)

// RSyncSend sends a file somewhere else
func (node *ThisNode) RSyncSend(src, SDFSName string, dst OtherNode, ackChan chan bool) {

	user, err := user.Current()
	if err != nil {
		fmt.Println("Couldn't get current user")
	}

	addr := user.Username + "@" + dst.Hostname + ":/shared/" + SDFSName

	fmt.Printf("Sending file %s to be named %s to node %s\n", src, strings.ReplaceAll(SDFSName, "^", "/"), dst.Hostname)

	_, err = exec.Command("rsync", "-v", "-e", "ssh -o StrictHostKeyChecking=no", src, addr).Output()
	if err != nil {
		fmt.Println("Error running rsync send command to: " + addr + " " + err.Error())
		ackChan <- false
	} else {
		// signal success ACK
		ackChan <- true
	}

}

// RSyncFetch fetches a file
func (node *ThisNode) RSyncFetch(filename, localPath, src string) {

	user, err := user.Current()
	if err != nil {
		fmt.Println("Couldn't get current user")
	}
	addr := user.Username + "@" + src + ":/shared/" + filename

	fmt.Printf("Getting file: %s from node %s\n", strings.ReplaceAll(filename, "^", "/"), src)

	_, err = exec.Command("rsync", "-v", "-e", "ssh -o StrictHostKeyChecking=no", addr, localPath).Output()
	if err != nil {
		fmt.Println("Error running rsync fetch command to: " + addr + " " + err.Error())
	} else {
		node.Logger.Println("Done fetching " + filename + " to " + localPath)
		fmt.Println("Done fetching " + strings.ReplaceAll(filename, "^", "/") + " to " + localPath)
		fmt.Println(time.Now().Format("15:04:05.000"))
	}
}
