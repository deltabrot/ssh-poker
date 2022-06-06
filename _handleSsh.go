import (
	"fmt"
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/gliderlabs/ssh"
)

func handleSsh() {
	ssh.Handle(func(s ssh.Session) {
		defer s.Close()

		// alert server a new connection has been made
		fmt.Printf("\033[0;33m'%s' has connected\033[0m\n", s.User())

		tty := NewSshTty(s)
		screen, err := tcell.NewTerminfoScreenFromTty(tty)
		if err != nil {
			log.Println(err)
			return
		}

		if err := screen.Init(); err != nil {
			panic(err)
		}

		/* defStyle := tcell.StyleDefault.
			Background(tcell.ColorBlack).
			Foreground(tcell.ColorWhite)
		screen.SetStyle(defStyle) */

		displayHelloWorld(screen)

		/* term := terminal.NewTerminal(s, "$ ")
		go func() {
			for {
				line, err := term.ReadLine()
				if err != nil {
					fmt.Printf(
						"\033[0;33m'%s' has disconnected\033[0m\n", s.User(),
					)
					s.Exit(1)
					break
				}

				// ping/pong
				if line == "ping" {
					io.WriteString(s, "pong")
					continue
				}

				time.Sleep(100 * time.Millisecond)
			}
		}() */

		for {
			switch ev := screen.PollEvent().(type) {
			case *tcell.EventResize:
				screen.Sync()
				displayHelloWorld(screen)
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape {
					screen.Fini()
					os.Exit(0)
				}
			}
		}

		// greet new user
		// io.WriteString(s, fmt.Sprintf("Hello %s\n", s.User()))

		// bring this back potentially
		/* term := terminal.NewTerminal(s, "$ ")
		go func() {
			for {
				line, err := term.ReadLine()
				if err != nil {
					fmt.Printf(
						"\033[0;33m'%s' has disconnected\033[0m\n", s.User(),
					)
					s.Exit(1)
					break
				}

				// ping/pong
				if line == "ping" {
					io.WriteString(s, "pong")
					continue
				}

				time.Sleep(100 * time.Millisecond)
			}
		}() */

		// loop (probably deprecated because of tcell)
		/* for {
			time.Sleep(100 * time.Millisecond)
		} */
	})

}
