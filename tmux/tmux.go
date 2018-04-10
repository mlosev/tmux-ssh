package tmux

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
)

// Tmux describes tmux client
type Tmux struct{}

// Window describes tmux window
// Borrowed from https://github.com/bpowers/go-tmux project
type Window struct {
	Index       int
	Name        string
	SessionName string
	NPanes      int
	Width       int
	Height      int
	Active      bool
}

// NewTmux creates instance of tmux client
func NewTmux() *Tmux {
	return &Tmux{}
}

// ListWindows returns list of tmux windows
// Mostly borrowed from https://github.com/bpowers/go-tmux project
func (t *Tmux) ListWindows() ([]Window, error) {
	jsonFmt := `{
		"Index": #{window_index},
		"Name": "#{window_name}",
		"SessionName": "#{session_name}",
		"Active": #{?window_active,true,false},
		"NPanes": #{window_panes},
		"Width": #{window_width},
		"Height": #{window_height}
	}`
	out, err := NewCommand("list-windows", "-a", "-F", strings.Replace(jsonFmt, "\n", "", -1)).Run()
	if err != nil {
		return nil, err
	}
	windowBufs := bytes.Split(out, []byte{'\n'})
	windows := make([]Window, len(windowBufs))
	for i, buf := range windowBufs {
		err := json.Unmarshal(buf, &windows[i])
		if err != nil {
			return nil, err
		}
	}
	return windows, nil
}

// CreateWindowFromCommand creates new window with given name and command
func (t *Tmux) CreateWindowFromCommand(name string, cmd []string) {
	err := NewCommand(append([]string{"new-window", "-n", name}, cmd...)...).Exec()
	if err != nil {
		panic(err)
	}

}

// SelectWindow selects window by its index and performs switch
func (t *Tmux) SelectWindow(idx int) {
	err := NewCommand("select-window", "-t", strconv.FormatInt(int64(idx), 10)).Exec()
	if err != nil {
		panic(err)
	}
}
