package main

import (
	"flag"
	"github.com/bibishkind/shtrafovnet-test-task/pkg/service"
	"github.com/bibishkind/shtrafovnet-test-task/pkg/transport/grpc"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

func main() {
	port := flag.String("port", "50051", "GRPC server port")
	flag.Parse()

	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.SetReportCaller(true)

	srvc := service.NewService()
	grpcServer := grpc.NewServer(srvc)

	go func() {
		logrus.Infof("starting grpc server...")
		if err := grpcServer.ListenAndServe(*port); err != nil {
			logrus.Fatalf("error starting grpc server: %v", err)
		}
	}()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)

	<-sigint

	logrus.Infof("stopping grpc server...")
	grpcServer.Stop()
}
