package main

import (
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
)

func main() {
	consumer, err := sarama.NewConsumerGroup([]string, nil)
	if err != nil {
		panic(err)
	}
	log.Println("consumer created")
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	log.Println("commence consuming")
	partitionConsumer, err := consumer.Consume(*topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumed := 0
ConsumerLoop:
	for {
		log.Println("in the for")
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Consumed message offset %d\n", msg.Offset)
			if *logMsg {
				log.Printf("KEY: %s VALUE: %s", msg.Key, msg.Value)
			}
			consumed++
		case <-signals:
			break ConsumerLoop
		}
	}

	log.Printf("Consumed: %d\n", consumed)

}
