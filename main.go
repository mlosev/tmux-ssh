package main

import (
	"fmt"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/mlosev/tmux-ssh/tmux"
)

func findWindow(t *tmux.Tmux, target string) int {
	notFound := -1
	ww, err := t.ListWindows()
	if err != nil {
		return notFound
	}
	for _, w := range ww {
		if w.Name == target {
			return w.Index
		}
	}
	return notFound
}

var (
	sshTarget  = kingpin.Arg("target", "Target to ssh to").Required().String()
	sshCommand = kingpin.Flag("ssh-cmd", "SSH command to call").Default("ssh").String()
	namePrefix = kingpin.Flag("name-prefix", "Prefix to tmux window name").Default("ssh-").String()
)

func main() {
	kingpin.Parse()

	t := tmux.NewTmux()

	target := *sshTarget
	name := fmt.Sprintf("%s%s", *namePrefix, strings.SplitN(target, ".", 2)[0])
	idx := findWindow(t, name)

	switch idx {
	case -1:
		t.CreateWindowFromCommand(name, append(strings.Fields(*sshCommand), target))
	default:
		t.SelectWindow(idx)
	}
}
