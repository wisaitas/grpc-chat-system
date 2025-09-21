package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/wisaitas/grpc-chat-system/internal/server/model"
	userPb "github.com/wisaitas/grpc-chat-system/internal/server/protogen/user"
	"github.com/wisaitas/grpc-chat-system/internal/server/repository"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserService struct {
	userPb.UnimplementedUserServiceServer
	userRepo repository.Repository
}

func NewUserService(userRepo repository.Repository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Register(ctx context.Context, req *userPb.RegisterRequest) (*userPb.RegisterResponse, error) {
	// Validate input
	if req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}
	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}
	if req.Password != req.ConfirmPassword {
		return nil, status.Error(codes.InvalidArgument, "passwords do not match")
	}

	// Check if user already exists
	existingUser, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, status.Error(codes.Internal, "failed to check existing user")
	}
	if existingUser != nil && existingUser.Email != "" {
		return nil, status.Error(codes.AlreadyExists, "user already exists with this email")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to hash password")
	}

	// Create user
	user := &model.User{
		ID:        uuid.New(),
		Email:     req.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.userRepo.CreateUser(ctx, user); err != nil {
		return nil, status.Error(codes.Internal, "failed to create user")
	}

	return &userPb.RegisterResponse{
		Id:    user.ID.String(),
		Email: user.Email,
	}, nil
}

func (s *UserService) Login(ctx context.Context, req *userPb.LoginRequest) (*userPb.LoginResponse, error) {
	return nil, status.Error(codes.Unimplemented, "login not implemented yet")
}

func (s *UserService) ListUsers(ctx context.Context, req *userPb.Empty) (*userPb.ListUsersResponse, error) {
	return nil, status.Error(codes.Unimplemented, "list users not implemented yet")
}
