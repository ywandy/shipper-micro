package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"log"
	pb "shipper/consignment-service/proto/consignment"
	vesselpb "shipper/vessel-service/proto/vessel"
)

//定义微服务
type service struct {
	repo         Repository
	vesselClient vesselpb.VesselService
}

//实现shipperservicehandler接口
//service 作为grpc服务端

//创建一个托运的货物
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) (error) {
	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselpb.Specification{MaxWeight: req.Weight, Capacity: int32(len(req.Containers))})
	log.Printf("find vessel:%s \n", vesselResponse)
	consignment, err := s.repo.Create(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	resp.Created = true
	resp.Consignment = consignment
	resp = &pb.Response{Created: true, Consignment: consignment}
	return nil
}

//获取所有
func (s *service) GetConsignments(xtx context.Context, req *pb.GetRequest, resp *pb.Response) (error) {
	consignments := s.repo.GetAll()
	resp.Consignments = consignments
	return nil
}

func main() {
	m_server := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)
	//作为vessel的客户端
	vessel_client := vesselpb.NewVesselService("go.micro.srv.vessel", m_server.Client())
	//解析命令行参数
	m_server.Init()
	repo := &ConsignmentRepository{}                                                                               //初始化Repository接口
	err := pb.RegisterShippingServiceHandler(m_server.Server(), &service{repo: repo, vesselClient: vessel_client}) //服务注册
	if err != nil {
		log.Fatal(err)
	}
	if err := m_server.Run(); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
