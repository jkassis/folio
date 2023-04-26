package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/fatih/color"
	"jkassis.com/tests/v2/golang/shell/data"
)

var Blue = color.New(color.FgBlue)
var Yellow = color.New(color.FgYellow)
var Green = color.New(color.FgGreen)
var White = color.New(color.FgWhite)
var Spaces = regexp.MustCompile(`\s+`)

type Event struct {
	when time.Time
	what string
}

type Shell struct {
	cwd             string
	stdinBuf        *bufio.Reader
	stdin           io.Reader
	stdout          io.Writer
	stderr          io.Writer
	schedule        *data.Heap[*Event]
	scheduleTicker  *time.Ticker
	scheduleStopper chan struct{}
}

// Run reads input infinitely (one line at a time)
func (s *Shell) Run() {
	// setup stdin, stdout, stderr
	s.stdin = os.Stdin
	s.stdinBuf = bufio.NewReader(os.Stdin)
	s.stdout = os.Stdout // This is overly verbose, but shows how to access stdout / stderr
	s.stderr = os.Stderr

	// get cwd
	path, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error during init: %s\n", err.Error())
		os.Exit(1)
	}
	s.cwd = path

	// intercept interrupts
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for range signalChan {
			fmt.Fprintln(s.stderr, "\ntype 'exit' to cancel the shell")
			s.Prompt()
		}
	}()

	// init the schedule
	s.schedule = data.NewHeap(func(a, b *Event) bool { return a.when.Before(b.when) })

	// welcome!
	Green.Fprintln(s.stdout, "welcome to jsh!")
	fmt.Fprintln(s.stdout, "schedule commands with...")
	Yellow.Fprintln(s.stdout, "doin <duration> <cmd>")
	s.TickEach("1s")

	// loop forever
	for {
		s.Prompt()

		// read input (the goofy way)
		var cmdLine = ""
		var prefix []byte
		var hasMore = true
		for hasMore {
			prefix, hasMore, err = s.stdinBuf.ReadLine()
			if err != nil {
				fmt.Fprintf(s.stderr, "failed to read input: %v\n", err)
				os.Exit(1)
			}
			cmdLine += string(prefix)
		}

		s.Do(cmdLine)
	}
}

// TickEach changes the frequency of Ticks
func (s *Shell) TickEach(args ...string) {
	if len(args) == 0 {
		fmt.Fprintf(s.stderr, "not enough args for TickEach\n")
	}

	d, err := time.ParseDuration(args[0])
	if err != nil {
		fmt.Fprintf(s.stderr, "could not parse tick duration %s: %v\n", args[0], err)
		return
	}

	if s.scheduleTicker != nil {
		s.scheduleTicker.Stop()
		s.scheduleStopper <- struct{}{}
	}

	fmt.Fprintf(s.stdout, "%s ticker started\n", d.String())
	s.scheduleTicker = time.NewTicker(d)
	s.scheduleStopper = make(chan struct{})

	stopper := s.scheduleStopper
	go func() {
		for {
			select {
			case <-s.scheduleTicker.C:
				s.Tick()
			case <-stopper:
				fmt.Fprintf(s.stdout, "%s ticker stopped\n", d.String())
				return
			}
		}
	}()
}

// Tick checks for Events on the schedule to do
func (s *Shell) Tick() {
	if s.schedule.Len() == 0 {
		return
	}
	event := s.schedule.Pop()
	if time.Now().Before(event.when) {
		s.schedule.Push(event)
	} else {
		s.Do(event.what)
	}
}

// Do does one command
// Runs built-ins first...
// https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html
// https://fishshell.com/docs/current/commands.html
func (s *Shell) Do(cmdLine string) {
	args := Spaces.Split(cmdLine, -1)
	cmd := strings.ToLower(args[0])
	switch cmd {
	case "alias":
		//alias [-p] [name[=value] …]
		return
	case "bind":
		// bind [-m keymap] [-lpsvPSVX]
		// bind [-m keymap] [-q function] [-u function] [-r keyseq]
		// bind [-m keymap] -f filename
		// bind [-m keymap] -x keyseq:shell-command
		// bind [-m keymap] keyseq:function-name
		// bind [-m keymap] keyseq:readline-command
		// bind readline-command-line
	case "builtin":
		// builtin [shell-builtin [args]]
	case "caller":
		// caller [expr]
	case "command":
		// command [-pVv] command [arguments …]
	case "declare":
		// declare [-aAfFgiIlnrtux] [-p] [name[=value] …]
	case "echo":
		// echo [-neE] [arg …]
	case "enable":
		// enable [-a] [-dnps] [-f filename] [name …]
	case "help":
		// help [-dms] [pattern]
	case "let":
		// let expression [expression …]
	case "local":
		// local [option] name[=value] …
	case "logout":
		// logout [n]
	case "mapfile":
		// mapfile [-d delim] [-n count] [-O origin] [-s count]
		// [-t] [-u fd] [-C callback] [-c quantum] [array]
	case "printf":
		// printf [-v var] format [arguments]
	case "read":
		// read [-ers] [-a aname] [-d delim] [-i text] [-n nchars]
		// [-N nchars] [-p prompt] [-t timeout] [-u fd] [name …]
	case "readarray":
		// readarray [-d delim] [-n count] [-O origin] [-s count]
		// [-t] [-u fd] [-C callback] [-c quantum] [array]
	case "source":
		// source filename
	case "type":
		// type [-afptP] [name …]
	case "typeset":
		// typeset [-afFgrxilnrtux] [-p] [name[=value] …]
	case "ulimit":
		// ulimit [-HS] -a
		// ulimit [-HS] [-bcdefiklmnpqrstuvxPRT] [limit]
	case "unalias":
		// unalias [-a] [name … ]
	case "exit":
		os.Exit(0) // successful exit
	case "pwd":
		s.Pwd()
	case "cd":
		s.Cd(args[1])
	case "ls":
		s.Ls(args[1:]...)
	case "tickeach":
		s.TickEach(args[1:]...)
	case "doin":
		s.Doin(args[1:]...)
	default:
		s.Exec(args[0], args[1:]...)
	}
}

// Pwd prints the working directory
func (s *Shell) Pwd() {
	fmt.Fprintln(s.stdout, s.cwd)
}

// Cd changes to the directory specified by the path
func (s *Shell) Cd(p string) {
	// resolve user home directory
	if p == "~" {
		userHome, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintf(s.stderr, "could not resolve user home: %v\n", err)
			return
		}
		p = userHome
	}

	// clean dot notation
	p = path.Clean(p)

	// enforce it's a directory
	isDir, err := s.IsDirectory(p)
	if err != nil {
		fmt.Fprintf(s.stderr, "could not test if %s is a directory: %v\n", p, err)
	}
	if !isDir {
		fmt.Fprintf(s.stderr, "%s is not a directory\n", p)
	}

	// change
	err = os.Chdir(p)
	if err != nil {
		fmt.Fprintf(s.stderr, "pwd is now %s\n", p)
	}
	s.cwd = p
}

// Ls lists the contents of the dir
func (s *Shell) Ls(args ...string) {
	// parse flags
	var recurse bool
	flagSet := flag.NewFlagSet("ls", flag.ContinueOnError)
	flagSet.BoolVar(&recurse, "r", false, "recursively list directory contents")
	err := flagSet.Parse(args)
	if err != nil {
		flagSet.Usage()
		return
	}
	args = flagSet.Args()

	if len(args) == 0 {
		args = []string{s.cwd}
	}

	for i, dirPath := range args {
		if i > 1 {
			fmt.Fprintln(s.stdout, "")
		}
		if len(dirPath) > 1 {
			fmt.Fprintf(s.stdout, "%s:\n", dirPath)
		}

		var err error
		dirPath, err = filepath.Abs(dirPath)
		if err != nil {
			fmt.Fprintf(s.stderr, "could not get abs path for %s: %v\n", dirPath, err)
			continue
		}

		if !recurse {
			files, err := os.ReadDir(dirPath)
			if err != nil {
				fmt.Fprintf(s.stderr, "could not read directory %s: %v\n", dirPath, err)
			}

			for _, file := range files {
				if file.IsDir() {
					Blue.Fprintf(s.stdout, "%s", file.Name())
					White.Fprintf(s.stdout, "/\n")
				} else {
					fmt.Fprintf(s.stdout, "%s\n", file.Name())
				}
			}
		} else {
			filepath.WalkDir(dirPath, func(filePath string, fileDirEntry fs.DirEntry, err error) error {
				if err != nil {
					fmt.Fprintf(s.stderr, "could not get dirEntry for %s: %v\n", filePath, err)
					return nil
				}
				fileRelPath, err := filepath.Rel(dirPath, filePath)
				if err != nil {
					fmt.Fprintf(s.stderr, "could not get rel path for %s: %v\n", filePath, err)
					return nil
				}

				if fileDirEntry.IsDir() {
					Blue.Fprintf(s.stdout, "%s", fileRelPath)
					White.Fprintf(s.stdout, "/\n")
				} else {
					fmt.Fprintf(s.stdout, "%s\n", fileRelPath)
				}

				return nil
			})
		}
	}

	fmt.Fprintln(s.stdout, "")
}

// Prompt returns text to display ahead of the user's input
func (s *Shell) Prompt() {
	parts := strings.Split(s.cwd, "/")
	for i, part := range parts {
		if i < len(parts)-1 && len(part) > 2 {
			parts[i] = part[:2]
		}
	}

	Blue.Fprintf(s.stdout, "%s > ", strings.Join(parts, "/"))
}

// Exec runs a command
func (s *Shell) Exec(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stderr = s.stderr
	cmd.Stdin = s.stdin
	cmd.Stdout = s.stdout
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(s.stderr, "could not run command: %v\n", err)
	}
}

// Doin runs a command after a delay
func (s *Shell) Doin(args ...string) {
	if len(args) <= 1 {
		fmt.Fprintf(s.stderr, "not enough args to doin\n")
	}

	delay, err := time.ParseDuration(args[0])
	if err != nil {
		fmt.Fprintf(s.stderr, "could not parse delay duration %s: %v\n", args[0], err)
		return
	}

	s.schedule.Push(&Event{when: time.Now().Add(delay), what: strings.Join(args[1:], " ")})
}

// isDirectory returns true if file specified by path is a directory
func (s *Shell) IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), err
}

func main() {
	s := &Shell{}
	s.Run()
}
