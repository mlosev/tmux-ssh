# tmux-ssh

Small helper to use with tmux and manage ssh connections to multiple servers. Mostly proof of concept, but still useful in daily routine.\
Creates new tmux window (or selects already existing one) for given server and executes `ssh SERVER` command.\
Recommended usage is in conjunction with key bindings in tmux config.

## Installation

Compiled binary can be grabbed from [releases page](https://github.com/mlosev/tmux-ssh/releases)

Alternatively you can build it from source.\
Requires golang 1.9+ installed and [kingpin](https://github.com/alecthomas/kingpin) library of version `v2`.

Install options:
- Quick install with `go get`:\
  `go get -u github.com/mlosev/tmux-ssh`

- Reqular install with `go install`:
  - Clone the repo:\
    `git clone https://github.com/mlosev/tmux-ssh`
  - Make sure you have `dep` installed (see https://github.com/golang/dep for instructions)
  - Populate vendor folder by running:\
    `dep ensure -v`
  - Finally, build and install binary to `$GOPATH/bin`:\
    `go install -v`

## Usage

Add to tmux config key binding of your choice for 'run-shell' tmux command.\
Example:

```
bind -n M-s command-prompt -p '(ssh)' 'run-shell "tmux-ssh %%"'
```

To override `ssh` command use `--ssh-cmd` flag, e.g. `tmux-ssh --ssh-cmd=mosh TARGET`

To override default prefix of tmux window name use flag `--name-prefix`, e.g. `tmux-ssh --name-prefix=mydomain- --ssh-cmd=mosh TARGET`

## Feedback

Feel free to leave any appropriate feedback)

## Contribution

Contributions are welcome)
