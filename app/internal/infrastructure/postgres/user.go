package postgres

import (
	"context"

	"MuchUp/app/internal/domain/entity"
	"MuchUp/app/internal/domain/repository"
	"MuchUp/app/sqlc"
)

type userRepository struct {
	q *sqlc.Queries
}

func NewUserRepository(db sqlc.DBTX) repository.UserRepository {
	return &userRepository{q: sqlc.New(db)}
}
func (repo *userRepository) Insert(ctx context.Context, _ *entity.UserProfile) error {
	param := sqlc.InsertUserParams{}

	if err := repo.q.InsertUser(ctx, param); err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) FindByID(ctx context.Context, id entity.UserID) (*entity.UserProfile, error) {
	uuid, err := toPGUUID(id)
	if err != nil {
		return nil, err
	}

	row, err := repo.q.FindUserByID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return toUserEntity(row), nil
}

func (repo *userRepository) Update(ctx context.Context, user *entity.UserProfile) error {
	param := sqlc.UpdateUserParams{}

	if err := repo.q.UpdateUser(ctx, param); err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) Delete(ctx context.Context, id entity.UserID) error {
	uuid, err := toPGUUID(id)
	if err != nil {
		return err
	}

	if err := repo.q.DeleteUser(ctx, uuid); err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) FindAll(ctx context.Context, limit, offset int) ([]*entity.UserProfile, error) {
	param := sqlc.FindAllUsersParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	rows, err := repo.q.FindAllUsers(ctx, param)
	if err != nil {
		return nil, err
	}

	users := make([]*entity.UserProfile, len(rows))
	for i, row := range rows {
		users[i] = toUserEntity(row)
	}

	return users, nil
}

func (repo *userRepository) FindByEmail(ctx context.Context, email string) (*entity.UserProfile, error) {
	return nil, ErrNotImplemented
}

func (repo *userRepository) FindByRoom(ctx context.Context, roomID entity.RoomID) ([]*entity.UserProfile, error) {
	return nil, ErrNotImplemented
}

func toUserEntity(row sqlc.User) *entity.UserProfile {
	return entity.ReconstructUserProfile(
		fromPGUUID[entity.UserID](row.ID),
		row.Nickname,
		fromPGText(row.AvatarUrl),
		row.UsagePurpose.String,
		entity.UserStatus(row.Status),
		row.CreatedAt.Time,
		row.UpdatedAt.Time,
	)
}
