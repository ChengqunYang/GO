package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

// kafka消费者 实例
func main()  {
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer,err:%v\n",err)
		return
	}
	partitionList, err := consumer.Partitions("web_log")
	if err != nil {
		fmt.Printf("fail to get list of partition err:%v\n",err)
		return
	}
	fmt.Println("分区列表:",partitionList)
	for partition := range partitionList {
		// 针对每一个分区创建一个对应分区的消费者
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%\n",partition,err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v",msg.Partition,msg.Offset,msg.Key,string(msg.Value))
			}
		}(pc)
	}
	select {

	}
}
