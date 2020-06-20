# pwn-mongodb

Tool for finding mongodb databases without a password.

```zsh
# install to $HOME/.local/bin
$ make build && PREFIX=~/.local make install 
```

```zsh
$ pwn-mongodb -i ~/path/to/hosts.txt -o ./pwnd.txt -c 10000 -l 10000 -t 15s
```
