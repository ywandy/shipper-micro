package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro"
	"io/ioutil"
	"log"
	"os"
	pb "shipper/consignment-service/proto/consignment"
)

const (
	defaultFilename = "consignment.json"
)

//从文件读取数据进来
func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &consignment)
	if err != nil {
		fmt.Println(err)
	}
	return consignment, err
}

func main() {
	m_server := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)
	//解析命令行参数
	//server.Init()
	//创建一个微服务客户端

	client := pb.NewShippingService("go.micro.srv.consignment", m_server.Client())
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}
	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("can not parse file:%v", err)
	}
	fmt.Println(consignment)
	r, err := client.CreateConsignment(context.TODO(), consignment)
	if err != nil {
		log.Fatalf("can not create:%v", err)
	}
	log.Printf("created:%t", r.Created)
	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("can not list consignments:%v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
