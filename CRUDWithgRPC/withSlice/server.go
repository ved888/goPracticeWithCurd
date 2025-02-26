package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"withSlice/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	port = flag.Int("port", 50051, "gRPC server port")
)

type server struct {
	pb.UnimplementedUserServiceServer
}

var lastID int = 0

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var Users []User

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	fmt.Println("create user")
	user := req.GetUser()
	lastID++
	data := User{
		Id:    fmt.Sprintf("%d", lastID), // Convert to string
		Name:  user.Name,
		Email: user.Email,
	}
	Users = append(Users, data)
	return &pb.CreateUserResponse{
		User: &pb.User{
			Id:    data.Id,
			Name:  data.Name,
			Email: data.Email,
		},
	}, nil
}

func (s *server) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	fmt.Println("Read User by id :", req.GetId())

	for _, user := range Users {
		if user.Id == req.Id {
			return &pb.UserResponse{
				User: &pb.User{
					Id:    user.Id,
					Name:  user.Name,
					Email: user.Email,
				},
			}, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "User with ID %s not found", req.GetId())
}

func (s *server) GetUsers(ctx context.Context, req *pb.UsersRequest) (*pb.UsersResponse, error) {
	fmt.Println("Read Users:")

	var userList []*pb.User // Initialize an empty slice

	// Convert internal `User` struct to `pb.User`
	for _, user := range Users {
		userList = append(userList, &pb.User{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return &pb.UsersResponse{
		Users: userList,
	}, nil
}

func (s *server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	fmt.Println("user updated by id", req.GetId())
	userReq := req.GetUser()

	for i, user := range Users {
		if user.Id == req.GetId() {
			Users[i].Name = userReq.Name
			Users[i].Email = userReq.Email

			return &pb.UpdateUserResponse{
				User: &pb.User{
					Id:    Users[i].Id,
					Name:  Users[i].Name,
					Email: Users[i].Email,
				},
			}, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "User with ID %s not found", req.GetId())
}

func (s *server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	fmt.Println("user deleted by id", req.GetId())
	var newUser []User

	for _, user := range Users {
		if user.Id != req.GetId() {
			newUser = append(newUser, user)
		}
	}
	Users = newUser
	return &pb.DeleteUserResponse{
		Success: true,
	}, nil
}

func main() {

	fmt.Println("gRPC server running ...")

	// Listen on the specified port for incoming TCP connections.
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server instance.
	s := grpc.NewServer()

	// Register the movie service implementation with the gRPC server.
	pb.RegisterUserServiceServer(s, &server{})

	// Print the address the server is listening on.
	log.Printf("Server listening at %v", lis.Addr())

	// Start serving incoming gRPC requests.
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
