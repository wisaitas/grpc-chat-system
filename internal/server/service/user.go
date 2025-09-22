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
	if err != nil {
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
	// Validate input
	if req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}
	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	// Get user by email
	user, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get user")
	}
	if user == nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}

	// TODO: Generate JWT tokens
	return &userPb.LoginResponse{
		AccessToken:  "access_token_placeholder",
		RefreshToken: "refresh_token_placeholder",
	}, nil
}

func (s *UserService) ListUsers(ctx context.Context, req *userPb.Empty) (*userPb.ListUsersResponse, error) {
	users, err := s.userRepo.ListUsers(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to list users")
	}

	pbUsers := make([]*userPb.User, len(users))
	for i, user := range users {
		pbUsers[i] = &userPb.User{
			Id:        user.ID.String(),
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		}
	}

	return &userPb.ListUsersResponse{
		Users: pbUsers,
	}, nil
}
