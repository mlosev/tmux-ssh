package main

import (
	"fmt"
	"os"
	"strings"

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

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s NAME\n", os.Args[0])
		os.Exit(1)
	}

	t := tmux.NewTmux()

	target := os.Args[1]
	name := fmt.Sprintf("ssh-%s", strings.SplitN(target, ".", 2)[0])
	idx := findWindow(t, name)

	switch idx {
	case -1:
		t.CreateWindowFromCommand(name, []string{"ssh", target})
	default:
		t.SelectWindow(idx)
	}
}
