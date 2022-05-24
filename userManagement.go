package main

import (
	"context"
	"flag"
	"fmt"
	pb "github/userManagement/userManagement"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net"
)

type User struct {
	gorm.Model
	FullName string `gorm:"size:100"`
	Email    string `gorm:"size:256"`
	Password string `gorm:"size:256"`
}

var (
	port = flag.Int("port", 50052, "The server port")
)

type server struct {
	pb.UnimplementedUserManagementServer
}

func (s *server) CreateUser(ctx context.Context, userInfo *pb.User) (*pb.Response, error) {
	dsn := "root2:Zz123456789!@@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database")
	}
	user := User{FullName: userInfo.GetFullname(), Email: userInfo.GetEmail(), Password: userInfo.GetPassword()}

	result := db.Create(&user)
	if result.Error == nil {
		return &pb.Response{Message: "user created", Status: "200", Data: "user"}, nil
	}
	return &pb.Response{Message: "problem", Status: "500", Data: "error"}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
