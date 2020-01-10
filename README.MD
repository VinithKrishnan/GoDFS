#CS 425 Fall 2019 Distributed Systems Machine Problems


This repo contains a distributed system written in Go.

This is MP3 from CS425 at UIUC Fall 2019

## Usage:

+ To run a node, navigate to the membership directory. then run:

`go run main.go`

It is currently hard-coded to work with our class VMs. To change this, change logQuery/nodes.txt

## Tests:

+ To run unit tests go to the logQuery directory (Run regular nodes elsewhere first):

`go test -v`

These integrated tests test all aspects of our MP1 design. The first test checks that we are able to read the node list file. 
Then we check that the RPC framework is working. This starts by making sure we can create a local RPC node, launch the server on a test port, and receive connections from other nodes.
Next, we will make sure we can launch the HTTP listener. The last check runs a distributed grep query on all 10 nodes with a known output.
The test checks that all 10 nodes respond with the correct number of lines. There are 3 tests: a common query, an uncommon query, and a rare query only present on 2 VMs.

### Citations:

-golang.org
-gobyexample.org

We found the following blog example of sending and receiving in Golang helpful:
-https://www.linode.com/docs/development/go/developing-udp-and-tcp-clients-and-servers-in-go/

The following were helpful in building our deployment infrastructure scripts and testing:

https://unix.stackexchange.com/questions/22682/how-to-launch-a-set-of-program-inside-tmux-or-gnome-terminal-within-a-script
https://unix.stackexchange.com/questions/304862/creating-a-tmux-session-without-attaching-to-it
https://www.cyberciti.biz/faq/bash-for-loop/
https://stackoverflow.com/questions/9301917/can-i-ssh-somewhere-run-some-commands-and-then-leave-myself-a-prompt
https://gist.github.com/andreyvit/2921703
https://stackoverflow.com/questions/19313807/tmux-send-keys-syntax
https://stackoverflow.com/questions/16678487/wget-command-to-download-a-file-and-save-as-a-different-filename
https://stackoverflow.com/questions/4089430/how-can-i-determine-the-url-that-a-local-git-repository-was-originally-cloned-fr
https://www.jerriepelser.com/blog/execute-permissions-with-git/




