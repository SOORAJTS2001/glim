package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func main() {
	yellow := "\033[33m"
	cyan := "\033[36m"
	reset := "\033[0m"
	bold := "\033[1m"

	// Dynamic Info
	now := time.Now().Format("Mon Jan 2 15:04:05 2006")
	goVer := runtime.Version()

	// ASCII Gopher (credit: https://github.com/ashleymcnamara/gophers)
	gopher := `
        ,_---~~~~~----._
  _,,_,*^____      _____''*g*\"*,
 / __/ /'     ^.  /      \ ^@q   f
[  @f | @))    |  | @))   l  0 _/
 \~~  \~______/ __ \_____/    \
  |           |/  \|          |
  |   GLIM    ||  ||  SHELL   |
  |___________\\__//__________|
`
	banner := fmt.Sprintf(`
%s%s🌟 Welcome to GLIM Shell%s
%s🐹 Built with Go%s
%s📦 Go Version:%s %s%s
%s🕒 Launched at:%s %s%s
%s💡 Type 'help' to get started
`,
		yellow, bold, yellow, // 🌟 Welcome...
		cyan, yellow, // Built with Go
		cyan, yellow, goVer, yellow, // Go Version: <yellow goVer>
		cyan, yellow, now, yellow, // Launched at: <yellow now>
		reset, // reset everything
	)

	fmt.Println(banner)
	fmt.Println(cyan + gopher + reset)
	time.Sleep(1 * time.Second)

	var reader = bufio.NewReader(os.Stdin)
	for {

		fmt.Printf("\033[1mglim>\033[0m")
		var input, e = reader.ReadString('\n')
		if e != nil {
			fmt.Printf("Error...")
			continue
		}
		input = strings.TrimSpace(input)
		if input == "help" {
			fmt.Println(`
        ╭────────────────────────────────────────────╮
        │              🆘 GLIM HELP MENU             │
        ╰────────────────────────────────────────────╯

        Glim is a lightweight shell wrapper built in Go.
        You can use it to run any system command as usual.

        ▶ Supported Examples:

          ls -la           → List files in long format
          pwd              → Show current working directory
          echo "hi"        → Print a message
          whoami           → Show current user
          date             → Show current system time
          ping google.com  → Ping an address

        ▶ Shell Shortcuts:

          help             → Show this help menu
          exit / quit      → Exit the Glim shell

        ▶ Notes:
          - Glim sends your command to the system using subprocess
          - Any command that works in your regular terminal should work here

        Happy hacking! 🚀
        `)
			continue
		}
		if input == "exit" || input == "quit" {
			fmt.Println("Ohk Bie!")
			break
		}
		var input_array = strings.Split(input, " ")
		switch input_array[0] {
		case "cd":
			if len(input_array) < 2 {
				fmt.Println("Could not change directory, since path is not given")
				continue
			}
			os.Chdir(input_array[1])
			continue
		}

		var command = input_array[0]
		var args = input_array[1:]
		var exe = exec.Command(command, args...)
		exe.Stdout = os.Stdout
		exe.Stderr = os.Stderr
		exe.Run()

	}

}
