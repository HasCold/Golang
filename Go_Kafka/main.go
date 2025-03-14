package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kgo"
)

// Recommending golang package for kafka :- https://github.com/twmb/franz-go

// Kafka UI Visualization is running on port :- localhost:8080

var brokers = []string{"localhost:29092", "localhost:29093", "localhost:29094"}

// Admin Client is responsible for creating the partitions, insertion in the topics or delete operations so all this done will be handle by this special client not by the normal client.
var adminClient *kadm.Client // kadm -->> kafka Admin

// One client can both produce and consume!
// Consuming can either be direct (no consumer group), or through a group. Below, we use a group.
func getAdminClient() {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
	)
	if err != nil {
		panic(err)
	}

	// defer client.Close()

	adminClient = kadm.NewClient(client)
}

func getSimpleKafkaClient() *kgo.Client {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
	)
	if err != nil {
		panic(err)
	}

	return client
}

func main() {
	// Init Admin Client
	getAdminClient()

	topicName := "test-kafka"
	var wg sync.WaitGroup

	// Using waitGroup here to allow synchronization or asynchronization process.
	wg.Add(1)

	simpleClient := getSimpleKafkaClient()
	record := &kgo.Record{Topic: topicName, Value: []byte("Our second message to kafka")}
	simpleClient.Produce(context.Background(), record, func(_ *kgo.Record, err error) {
		defer wg.Done()
		if err != nil {
			fmt.Printf("record had a produce error: %v \n", err)
		}
	})

	wg.Wait()

	// CreateTopic Parameter Passed :-
	// ctx context.Context,
	// partitions 5,
	// replicationFactor int16 = you can leave as default by putting -1,
	// configs - In configuration we can set up the retention time and some other properties or u can set the nil
	// topicName
	_, err := adminClient.CreateTopic(context.Background(), 5, -1, nil, topicName)
	if err != nil {
		log.Panic(err)
	}

}

func KafkaProducerClientClose(client *kadm.Client) {
	client.Close()
}
