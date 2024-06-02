package main

import (
	"context"
	"fmt"
	natsServer "github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type config struct {
	ServerName string
	Host       string
	Port       int
	Subject    string
}

func createServer(cfg config) {

	// options to create nats server
	opts := &natsServer.Options{
		ServerName:     cfg.ServerName,
		Host:           cfg.Host,
		Port:           cfg.Port,
		NoLog:          false,
		NoSigs:         false,
		MaxControlLine: 4096,
		MaxPayload:     65536,
	}

	// initialize the server struct
	server, err := natsServer.NewServer(opts)
	if err != nil {
		log.Fatal(err)
	}

	// run the nats server based on the options
	err = natsServer.Run(server)
	if err != nil {
		log.Fatal("Failed to start NATS server:", err)
	}

	log.Println("NATS server started")
}

func producer(ctx context.Context, cfg config) {
	natsURL := fmt.Sprintf("nats://%s:%d", cfg.Host, cfg.Port)
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal("Failed to connect to NATS server:", err)
	}
	// close the connection on function exit
	defer nc.Close()

	// listen for messaages on this subject
	subject := cfg.Subject

	i := 0

	for {
		select {
		case <-ctx.Done():
			log.Println("exiting from producer")
			return
		default:
			i += 1
			message := fmt.Sprintf("message %v", i)
			time.Sleep(1 * time.Second)
			// Publish the message to the nats server
			err = nc.Publish(subject, []byte(message))
			if err != nil {
				log.Println("Failed to publish message:", err)
			} else {
				log.Println(message)
			}
		}
	}
}

func consumer(ctx context.Context, cfg config) {
	natsURL := fmt.Sprintf("nats://%s:%d", cfg.Host, cfg.Port)
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal("Failed to connect to NATS server:", err)
	}
	defer nc.Close()

	fmt.Printf("Connected to NATS server on port %d\n", cfg.Port)

	messages := make(chan *nats.Msg, 1000)

	subject := cfg.Subject

	// we're subscribing to the subject
	// and assigning our channel as reference to receive messages there
	subscription, err := nc.ChanSubscribe(subject, messages)
	if err != nil {
		log.Fatal("Failed to subscribe to subject:", err)
	}

	defer func() {
		subscription.Unsubscribe()
		close(messages)
	}()

	log.Println("*****Subscribed to", subject)

	for {
		select {
		case <-ctx.Done():
			log.Println("exiting from consumer")
			return
		case msg := <-messages:
			log.Println("received", string(msg.Data))
		}
	}
}
func main() {
	cfg := config{
		ServerName: "local_nats_server",
		Host:       "localhost",
		Port:       4222,
		Subject:    "example.logs",
	}

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		// listen for interrupts to exit gracefully
		sigChannel := make(chan os.Signal, 1)
		signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)
		<-sigChannel
		close(sigChannel)
		cancel()
	}()

	// create the local server
	createServer(cfg)

	// register the consumer
	go consumer(ctx, cfg)

	// register the producer
	go producer(ctx, cfg)
	<-ctx.Done()

	log.Println("server shutdown completed")
	log.Println("exiting gracefully")
}
