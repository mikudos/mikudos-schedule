package clients

import (
	"fmt"
	"log"

	"github.com/mikudos/mikudos-schedule/config"
	aipb "github.com/mikudos/mikudos-schedule/proto/ai"
	"google.golang.org/grpc"
)

var (
	conns   = make(map[string]*grpc.ClientConn)
	clients = make(map[string]interface{})
)

func init() {
	log.Println("Init all grpc client: ai, learn, users, messages")
	setUpClientConn("ai")
}

func setUpClientConn(connName string) (err error) {
	confLoc := fmt.Sprintf("grpcClients.%s", connName)
	grpcAddr := config.RuntimeViper.GetString(confLoc)
	if grpcAddr == "" {
		log.Fatalln("address for " + confLoc + "must be set")
	}
	// Set up a connection to the server.
	conns[connName], err = grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	var client interface{}
	switch connName {
	case "ai":
		client = aipb.NewAiServiceClient(conns[connName])
		break
	}
	clients[connName] = client
	return err
}
