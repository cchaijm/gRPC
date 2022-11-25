package main

import (
	"context"
	proto "github.com/cchaijm/gRPC/doingSSH/proto"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
)

type Server struct {
	proto.DoingSSHServer
}

func (*Server) SSHServer(ctx context.Context, in *proto.SSHServerRequest) (*proto.SSHServerSSHResponse, error) {
	log.Printf("SSH was invoked with %v\n", in)

	sshConfig := &ssh.ClientConfig{
		User:            in.ServerUsername,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		// Auth: .... fill out with keys etc as normal
		Auth: []ssh.AuthMethod{
			ssh.Password(in.ServerPassword),
		},
	}

	conn, err := net.Dial("tcp", in.ServerIp)
	if err != nil {
		log.Println(err)
	}

	c, chans, reqs, err := ssh.NewClientConn(conn, in.ServerIp, sshConfig)

	client := ssh.NewClient(c, chans, reqs)

	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}

	defer session.Close()
	cmd, err := session.Output(in.Command)

	log.Println(string(cmd))

	log.Println("Hello, world!")

	return &proto.SSHServerSSHResponse{Result: "Hello " + in.ServerIp}, nil
}
