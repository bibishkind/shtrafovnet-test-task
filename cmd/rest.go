package main

import (
	"context"
	"flag"
	"github.com/bibishkind/shtrafovnet-test-task/pkg/service"
	"github.com/bibishkind/shtrafovnet-test-task/pkg/transport/grpc"
	"github.com/bibishkind/shtrafovnet-test-task/pkg/transport/rest"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	port := flag.String("port", "80", "REST server port")
	flag.Parse()

	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.SetReportCaller(true)

	srvc := service.NewService()
	grpcServer := grpc.NewServer(srvc)
	restServer := rest.NewServer(grpcServer)

	go func() {
		logrus.Infof("starting rest server...")
		if err := restServer.ListenAndServe(*port); err != http.ErrServerClosed {
			logrus.Fatalf("error starting rest server %v", err)
		}
	}()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)

	<-sigint

	logrus.Infof("stopping rest server...")
	if err := restServer.Stop(context.Background()); err != nil {
		logrus.Fatalf("error stopping rest server: %v", err)
	}
}
