# Interactive Terminal Runner
Runs other programmes in the terminal in an interactive fashion

### Installation
```
go get github.com/rawnet/go-interactive-terminal-runner
```

### Usage

Basic
```go
    runner, err := interactive_terminal_runner.NewTerminalRunner()

	if err != nil {
		log.Fatalln(err)
	}

    if err = runner.Exec("gpg", "-c /a/file/location.txt"); err != nil {
        log.Fatalln(err)
    }
```

Custom Exec Path
```go
    runner, err := interactive_terminal_runner.NewTerminalRunner(interactive_terminal_runner.WithCustomExecPath("sh"))

	if err != nil {
		log.Fatalln(err)
	}

    if err = runner.Exec("gpg", "-c /a/file/location.txt"); err != nil {
        log.Fatalln(err)
    }
```