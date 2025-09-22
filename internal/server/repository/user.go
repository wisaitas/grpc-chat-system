package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wisaitas/grpc-chat-system/internal/database/convert"
	db "github.com/wisaitas/grpc-chat-system/internal/database/sqlc"
	"github.com/wisaitas/grpc-chat-system/internal/server/model"
)

type Repository interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error)
	ListUsers(ctx context.Context) ([]*model.User, error)
}

type repository struct {
	db      *pgxpool.Pool
	queries *db.Queries
}

func NewRepository(dbPool *pgxpool.Pool) Repository {
	return &repository{
		db:      dbPool,
		queries: db.New(dbPool),
	}
}

func (r *repository) CreateUser(ctx context.Context, user *model.User) error {
	_, err := r.queries.CreateUser(ctx, db.CreateUserParams{
		ID:        convert.UUIDToPgtype(user.ID),
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: convert.TimeToPgtype(user.CreatedAt),
		UpdatedAt: convert.TimeToPgtype(user.UpdatedAt),
	})
	return err
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	dbUser, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &model.User{
		ID:        convert.PgtypeToUUID(dbUser.ID),
		Email:     dbUser.Email,
		Password:  dbUser.Password,
		CreatedAt: convert.PgtypeToTime(dbUser.CreatedAt),
		UpdatedAt: convert.PgtypeToTime(dbUser.UpdatedAt),
	}, nil
}

func (r *repository) GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	dbUser, err := r.queries.GetUserByID(ctx, convert.UUIDToPgtype(id))
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &model.User{
		ID:        convert.PgtypeToUUID(dbUser.ID),
		Email:     dbUser.Email,
		Password:  dbUser.Password,
		CreatedAt: convert.PgtypeToTime(dbUser.CreatedAt),
		UpdatedAt: convert.PgtypeToTime(dbUser.UpdatedAt),
	}, nil
}

func (r *repository) ListUsers(ctx context.Context) ([]*model.User, error) {
	dbUsers, err := r.queries.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	users := make([]*model.User, len(dbUsers))
	for i, dbUser := range dbUsers {
		users[i] = &model.User{
			ID:        convert.PgtypeToUUID(dbUser.ID),
			Email:     dbUser.Email,
			Password:  dbUser.Password,
			CreatedAt: convert.PgtypeToTime(dbUser.CreatedAt),
			UpdatedAt: convert.PgtypeToTime(dbUser.UpdatedAt),
		}
	}

	return users, nil
}
