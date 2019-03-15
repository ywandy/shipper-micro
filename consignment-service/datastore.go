/** 
* Created by yewei_andy on 19-3-15. 
* Author : yewei_andy
* Email : 896882701yw@gmail.com
*/
package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var (
	G_Mongodb_Client   *mongo.Client
	G_Mongodb_DataBase *mongo.Database
)

func InitDBConn() (error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://192.168.10.125:27018"))
	if err != nil {
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db := client.Database("youpat_visual_bc")
	err = client.Connect(ctx)

	//等待链接
	for {
		err = client.Ping(ctx, readpref.Primary())
		if (err == nil) {
			break
		}
		time.Sleep(1 * time.Second)
	}
	fmt.Println("链接成功")
	G_Mongodb_Client = client
	G_Mongodb_DataBase = db
	return nil
}


