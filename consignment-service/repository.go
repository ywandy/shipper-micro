/** 
* Created by yewei_andy on 19-3-15. 
* Author : yewei_andy
* Email : 896882701yw@gmail.com
*/
package main

import pb "shipper/consignment-service/proto/consignment"

//仓库接口
//1.实现存放货物
//2.实现获取货物

type Repository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error) //存放
	GetAll() []*pb.Consignment                                   //获取所有
}

//存放货物的仓库，实现IRepository接口
//改变使用mongodb
type ConsignmentRepository struct {
	consignments []*pb.Consignment
}

//创建
func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

//拿所有
func (repo *ConsignmentRepository) GetAll() []*pb.Consignment {
	return repo.consignments
}
