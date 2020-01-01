package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	aipb "github.com/mikudos/mikudos-schedule/proto/ai"
	pb "github.com/mikudos/mikudos-schedule/proto/schedule"
	"github.com/mikudos/mikudos-schedule/schedule"
	"github.com/robfig/cron/v3"
)

// AiService AiService
type AiService struct {
	baseService
	jobID cron.EntryID
}

type baseService struct {
	HelloRequest *aipb.HelloRequest
}

// ClientFunc ClientFunc
func (ai *AiService) ClientFunc(call *pb.GrpcCall) {
	serviceName := "ai"
	state := conns[serviceName].GetState()
	if state.String() != "READY" {
		conns[serviceName].Close()
		setUpClientConn(serviceName)
	}
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var err error
	switch call.GetMethodName() {
	case "SayHello":
		json.Unmarshal([]byte(call.PayloadStr), &ai.HelloRequest)
		_, err = clients[serviceName].(aipb.AiServiceClient).SayHello(ctx, ai.HelloRequest)
		break
	case "SayHi":
		json.Unmarshal([]byte(call.PayloadStr), &ai.HelloRequest)
		break
	default:
		fmt.Printf("set cron task fail: %s", "没有对应grpc method")
		break
	}
	log.Printf("%s called", call.GetMethodName())
	if err != nil {
		log.Printf("could not call method on %s: %v", serviceName, err)
	}
	ai.checkCancelList()
}

func (ai *AiService) checkCancelList() {
	if _, ok := schedule.OneTimeJobs[ai.jobID]; ok {
		schedule.Cron.Remove(ai.jobID)
		delete(schedule.OneTimeJobs, ai.jobID)
	}
}
