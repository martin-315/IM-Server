package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

var kafkaIp = "127.0.0.1:9092"

func Test(ctx *gin.Context) {
	//读取
	ctx.JSON(200, gin.H{
		"data": "product",
	})
}

func main() {

	//启动消息者
	go InitConsumer()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/send", SendMessage) //http://localhost:8082/send

	r.Run("0.0.0.0:8082") // 监听并在 0.0.0.0:8080 上启动服务

}

//发消息到kakfa
func SendMessage(ctx *gin.Context) {
	fmt.Println("SendMessage")
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{kafkaIp}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer client.Close()
	//例子一发单个消息
	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = Topic
	content := "this is a test log"
	sendTokafka(client, msg, content)

	//例子二发多个消息
	for _, word := range []string{"Welcome11", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		sendTokafka(client, msg, word)
	}
}

//发消息函数
func sendTokafka(client sarama.SyncProducer, msg *sarama.ProducerMessage, content string) {
	msg.Value = sarama.StringEncoder(content)

	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)

}

func InitConsumer() {
	time.Sleep(time.Second * 3)
	fmt.Println("init Counsumer success")

	var wg sync.WaitGroup
	consumer, err := sarama.NewConsumer([]string{kafkaIp}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions(Topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitionList)

	for partition := range partitionList { // 遍历所有的分区
		wg.Add(1)
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(Topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
			}
			defer pc.AsyncClose()
			wg.Done()
		}(pc)
	}
	//select{} //阻塞进程
	wg.Wait()
	consumer.Close()
}
