package main

import (
	"context"
	"log"

	"github.com/cchaijm/gRPC/doingSSH/proto"
)

func doSSH(c proto.DoingSSHClient) {
	log.Println("doSSH was invoked")
	r, err := c.SSHServer(context.Background(),
		&proto.SSHServerRequest{
			ServerUsername: "root",
			ServerIp:       "",
			ServerPassword: "",
			Command:        "ls -la",
		},
	)

	if err != nil {
		log.Fatalf("Could not ssh: %v\n", err)
	}

	log.Printf("sshing: %s\n", r.Result)
}
