package broker

import (
	"fmt"
	"log"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/mikudos/mikudos-schedule/config"
)

// BrokerInstance faaa
var BrokerInstance Broker

// Broker aaa
type Broker struct {
	config   *sarama.Config
	producer sarama.AsyncProducer
	Client   sarama.ConsumerGroup
}

// Msg aaa
type Msg struct {
	Topic   string
	Key     string
	Message string
}

var (
	brokers = config.RuntimeViper.GetString("brokers.endPoints")
	version = config.RuntimeViper.GetString("brokers.version")
	group   = config.RuntimeViper.GetString("brokers.group")
	topics  = config.RuntimeViper.GetString("brokers.topics")
)

// Send 发送消息
func (b *Broker) Send(m Msg) {
	msg := &sarama.ProducerMessage{
		Topic: m.Topic,
		Key:   sarama.StringEncoder(m.Key),
	}

	msg.Value = sarama.ByteEncoder(m.Message)
	fmt.Printf("input [%s]\n", m.Message)

	// // send to chain
	BrokerInstance.producer.Input() <- msg

	select {
	case suc := <-BrokerInstance.producer.Successes():
		fmt.Printf("offset: %d,  timestamp: %s", suc.Offset, suc.Timestamp.String())
		break
	case fail := <-BrokerInstance.producer.Errors():
		fmt.Printf("err: %s\n", fail.Err.Error())
		defer BrokerInstance.producer.Close()
		generateProducer()
		break
	}
}

func init() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V0_11_0_2
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	BrokerInstance = Broker{config: config}
	var err error
	err = generateProducer()
	if err != nil {
		log.Printf("producer_test create producer error :%s\n", err.Error())
		return
	}
	err = generateClient()
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}
}

func generateProducer() (err error) {
	BrokerInstance.producer, err = sarama.NewAsyncProducer(strings.Split(brokers, ","), BrokerInstance.config)
	return err
}

func generateClient() (err error) {
	BrokerInstance.Client, err = sarama.NewConsumerGroup(strings.Split(brokers, ","), group, BrokerInstance.config)
	return err
}
