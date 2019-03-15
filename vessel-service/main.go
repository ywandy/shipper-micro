/** 
* Created by yewei_andy on 19-3-15. 
* Author : yewei_andy
* Email : 896882701yw@gmail.com
*/
package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"github.com/pkg/errors"
	"log"
	pb "shipper/vessel-service/proto/vessel"
)

//船只接口
type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

//船只仓库
type VesselRepository struct {
	vessels []*pb.Vessel
}

//寻找合适的货船接口
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, v := range repo.vessels {
		//判断货船是否符合规格
		if v.Capacity >= spec.Capacity && v.MaxWeight >= spec.MaxWeight {
			return v, nil
		}
	}
	return nil, errors.New("no any vessel available")
}

type service struct {
	repo Repository
}

//实现接口
func (s *service) FindAvailable(ctx context.Context, spec *pb.Specification, resp *pb.Response) error {
	v, err := s.repo.FindAvailable(spec)
	if err != nil {
		return err
	}
	resp.Vessel = v
	return nil
}

func main() {
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	}
	repo := &VesselRepository{vessels: vessels}
	m_service := micro.NewService(micro.Name("go.micro.srv.vessel"), micro.Version("latest"))
	m_service.Init()
	pb.RegisterVesselServiceHandler(m_service.Server(), &service{repo: repo})
	if err := server.Run(); err != nil {
		log.Fatal("failed to serve: %v", err)
	}
}
