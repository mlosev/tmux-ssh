# tmux-ssh

Small helper to use with tmux and manage ssh connections to multiple servers. Mostly proof of concept, but still useful in daily routine.\
Creates new tmux window (or selects already existing one) for given server and executes `ssh SERVER` command.\
Recommended usage is in conjunction with key bindings in tmux config.

## Installation

Currently only installation from source is supported.\
Requires golang 1.9+ installed.

To install: `go get -u github.com/mlosev/tmux-ssh`

## Usage

Add to tmux config key binding of your choice for 'run-shell' tmux command.\
Example:

```
bind -n M-s command-prompt -p '(ssh)' 'run-shell "tmux-ssh %%"'
```
