package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	userPb "github.com/wisaitas/grpc-chat-system/internal/server/protogen/user"
	db "github.com/wisaitas/grpc-chat-system/internal/server/sqlc"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	userPb.UnimplementedUserServiceServer
	Queries *db.Queries
}

func NewUserService(Queries *db.Queries) *UserService {
	return &UserService{
		Queries: Queries,
	}
}

func (s *UserService) Register(ctx context.Context, req *userPb.RegisterRequest) (*userPb.RegisterResponse, error) {
	if req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}
	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}
	if req.Password != req.ConfirmPassword {
		return nil, status.Error(codes.InvalidArgument, "passwords do not match")
	}

	existingUser, err := s.Queries.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to check existing user")
	}
	if existingUser != nil && existingUser.Email != "" {
		return nil, status.Error(codes.AlreadyExists, "user already exists with this email")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to hash password")
	}

	user := &db.TblUser{
		ID:        pgtype.UUID{Bytes: uuid.New(), Valid: true},
		Email:     req.Email,
		Password:  string(hashedPassword),
		CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
		UpdatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	if _, err := s.Queries.CreateUser(ctx, db.CreateUserParams{
		ID:        user.ID,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}); err != nil {
		return nil, status.Error(codes.Internal, "failed to create user")
	}

	return &userPb.RegisterResponse{
		Id:    uuid.UUID(user.ID.Bytes).String(),
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
	user, err := s.Queries.GetUserByEmail(ctx, req.Email)
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
	users, err := s.Queries.ListUsers(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to list users")
	}

	pbUsers := make([]*userPb.User, len(users))
	for i, user := range users {
		pbUsers[i] = &userPb.User{
			Id:        uuid.UUID(user.ID.Bytes).String(),
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Time.Format("2006-01-02T15:04:05Z07:00"),
		}
	}

	return &userPb.ListUsersResponse{
		Users: pbUsers,
	}, nil
}
