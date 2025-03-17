package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kgo"
)

// Recommending golang package for kafka :- https://github.com/twmb/franz-go

// Kafka UI Visualization is running on port :- localhost:8080

const topicName string = "test-kafka"

var brokers = []string{"localhost:29092", "localhost:29093", "localhost:29094"}

// Admin Client is responsible for creating the partitions, insertion in the topics or delete operations so all this done will be handle by this special client not by the normal client.
var adminClient *kadm.Client // kadm -->> kafka Admin

// One client can both produce and consume!
// Consuming can either be direct (no consumer group), or through a group. Below, we use a group.
func getAdminClient() *kgo.Client {
	// Data will distribute into the muliple partitions by round robbin
	balancer := kgo.RoundRobinBalancer()
	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.Balancers(balancer),
	)
	if err != nil {
		panic(err)
	}

	// defer client.Close()

	adminClient = kadm.NewClient(client)
	_, err = adminClient.CreateTopic(context.Background(), 5, -1, nil, topicName)
	if err != nil {
		if !strings.Contains(err.Error(), "TOPIC_ALREADY_EXISTS") {
			panic(err)
		}
	}

	return client
}

// func getSimpleKafkaClient() *kgo.Client {
// 	client, err := kgo.NewClient(
// 		kgo.SeedBrokers(brokers...),
// 	)
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	return client
// }

func main() {
	// Init Admin Client
	simpleClient := getAdminClient()

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		KafkaProducer(i, &wg, simpleClient)
	}

	wg.Wait()

	// CreateTopic Parameter Passed :-
	// ctx context.Context,
	// partitions 5,
	// replicationFactor int16 = you can leave as default by putting -1,
	// configs - In configuration we can set up the retention time and some other properties or u can set the nil
	// topicName
	// _, err := adminClient.CreateTopic(context.Background(), 5, -1, nil, topicName)
	// if err != nil {
	// 	log.Panic(err)
	// }

}

func KafkaProducer(i int, wg *sync.WaitGroup, simpleClient *kgo.Client) {
	kafkaKey := strconv.Itoa(i)
	wg.Add(1)
	// prepare record to produce over kafka
	record := &kgo.Record{Topic: topicName, Key: []byte("kafka_" + kafkaKey), Value: []byte(fmt.Sprintf("Our %s hello to kafka", kafkaKey))}
	// produce
	simpleClient.Produce(context.Background(), record, func(_ *kgo.Record, err error) {
		defer wg.Done()
		if err != nil {
			fmt.Printf("record had a produce error: %v\n", err)
		}

	})
}

func KafkaProducerClientClose(client *kadm.Client) {
	client.Close()
}
