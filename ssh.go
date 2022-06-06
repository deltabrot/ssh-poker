package main

import (
	"io/ioutil"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

var (
	DEFAULT_SHELL string = "sh"
)

func startSsh() {
	// An SSH server is represented by a ServerConfig, which holds
	// certificate details and handles authentication of ServerConns.
	sshConfig := &ssh.ServerConfig{
		NoClientAuth: true,
	}

	// Define server private key.
	privateBytes, err := ioutil.ReadFile("./keys/server_key")
	if err != nil {
		log.Fatal("Failed to load private key (./keys/server_key)")
	}

	private, err := ssh.ParsePrivateKey(privateBytes)
	if err != nil {
		log.Fatal("Failed to parse private key")
	}

	sshConfig.AddHostKey(private)

	// Once a ServerConfig has been configured, connections can be accepted.
	listener, err := net.Listen("tcp4", ":2222")
	if err != nil {
		log.Fatalf("failed to listen on *:2222")
	}

	// Accept all connections
	log.Printf("listening on %s", ":2222")
	for {
		tcpConn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept incoming connection (%s)", err)
			continue
		}
		// Before use, a handshake must be performed on the incoming net.Conn.
		sshConn, chans, reqs, err := ssh.NewServerConn(tcpConn, sshConfig)
		if err != nil {
			log.Printf("failed to handshake (%s)", err)
			continue
		}

		// Check remote address
		log.Printf("new ssh connection from %s (%s)", sshConn.RemoteAddr(), sshConn.ClientVersion())

		// Print incoming out-of-band Requests
		go handleRequests(reqs)
		// Accept all channels
		go handleChannels(chans)
	}
}

func handleRequests(reqs <-chan *ssh.Request) {
	for req := range reqs {
		log.Printf("recieved out-of-band request: %+v", req)
	}
}
