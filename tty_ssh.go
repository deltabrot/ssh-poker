package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/gdamore/tcell/v2"
	"golang.org/x/crypto/ssh"
)

type SshTty struct {
	channel ssh.Channel
	l       sync.Mutex
	cb      func()
	f       *os.File
	w       int
	h       int
}

func NewSshTty(s ssh.Channel, f *os.File) tcell.Tty {
	return &SshTty{
		channel: s,
		f:       f,
	}
}

func (tty *SshTty) Read(b []byte) (int, error) {
	fmt.Println("Read")
	return tty.channel.Read(b)
}

func (tty *SshTty) Write(b []byte) (int, error) {
	fmt.Println("Write")
	return tty.channel.Write(b)
}

func (tty *SshTty) NotifyResize(cb func()) {
	log.Println("NotifyResize")
	tty.l.Lock()
	fmt.Println(cb)
	tty.cb = cb
	tty.l.Unlock()
}

func (tty *SshTty) WindowSize() (int, int, error) {
	log.Println("WindowSize")
	/* w, h, err := term.GetSize(tty.fd)
	if err != nil {
		return 0, 0, err
	}
	if w == 0 {
		w, _ = strconv.Atoi(os.Getenv("COLUMNS"))
	}
	if w == 0 {
		w = 80 // default
	}
	if h == 0 {
		h, _ = strconv.Atoi(os.Getenv("LINES"))
	}
	if h == 0 {
		h = 25 // default
	}
	return w, h, nil */
	// w, h := parseDims(req.Payload[termLen+4:])
	// SetWinsize(f.Fd(), w, h)
	fmt.Println(tty.w, tty.h)
	return tty.w, tty.h, nil
}

func (tty *SshTty) SetWinSize(w int, h int) {
	tty.w = w
	tty.h = h
}

// noops
func (tty *SshTty) Start() error {
	fmt.Println("Start() :-")
	return nil
}
func (tty *SshTty) Stop() error {
	fmt.Println("Stop() :-")
	return nil
}
func (tty *SshTty) Drain() error {
	fmt.Println("Drain() :-")
	tty.channel.Close()
	return nil
}
func (tty *SshTty) Close() error {
	fmt.Println("Close() :-")
	return nil
}
