package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	authPb "github.com/wisaitas/grpc-chat-system/internal/server/protogen/auth"
	db "github.com/wisaitas/grpc-chat-system/internal/server/sqlc"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authService struct {
	authPb.UnimplementedAuthServiceServer
	Queries *db.Queries
}

func NewAuthService(Queries *db.Queries) authPb.AuthServiceServer {
	return &authService{
		Queries: Queries,
	}
}

func (s *authService) Register(ctx context.Context, req *authPb.RegisterRequest) (*authPb.RegisterResponse, error) {
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
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, status.Error(codes.Internal, err.Error())
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

	return &authPb.RegisterResponse{
		Id:    uuid.UUID(user.ID.Bytes).String(),
		Email: user.Email,
	}, nil
}

func (s *authService) Login(ctx context.Context, req *authPb.LoginRequest) (*authPb.LoginResponse, error) {
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
	return &authPb.LoginResponse{
		AccessToken:  "access_token_placeholder",
		RefreshToken: "refresh_token_placeholder",
	}, nil
}
