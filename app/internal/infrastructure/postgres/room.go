package postgres

import (
	"context"

	"MuchUp/app/internal/domain/entity"
	"MuchUp/app/internal/domain/repository"
	"MuchUp/app/sqlc"
)

type roomRepository struct {
	q *sqlc.Queries
}

func NewRoomRepository(db sqlc.DBTX) repository.RoomRepository {
	return &roomRepository{q: sqlc.New(db)}
}

func (repo *roomRepository) Insert(ctx context.Context, room *entity.Room) error {
	param, err := toInsertRoomParams(room)
	if err != nil {
		return err
	}

	if err := repo.q.InsertRoom(ctx, param); err != nil {
		return err
	}

	return nil
}

func (repo *roomRepository) FindByID(ctx context.Context, id entity.RoomID) (*entity.Room, error) {
	uuid, err := toPGUUID(id)
	if err != nil {
		return nil, err
	}

	row, err := repo.q.FindRoomByID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return toRoomEntity(row), nil
}

func (repo *roomRepository) FindByUserID(ctx context.Context, userID entity.UserID) ([]*entity.Room, error) {
	uuid, err := toPGUUID(userID)
	if err != nil {
		return nil, err
	}

	rows, err := repo.q.FindAllRoomsByUserID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	rooms := make([]*entity.Room, len(rows))
	for i, row := range rows {
		rooms[i] = toRoomEntity(row)
	}

	return rooms, nil
}

func (repo *roomRepository) Update(ctx context.Context, room *entity.Room) error {
	param, err := toUpdateRoomParams(room)
	if err != nil {
		return err
	}

	if err := repo.q.UpdateRoom(ctx, param); err != nil {
		return err
	}

	return nil
}

func (repo *roomRepository) Delete(ctx context.Context, id entity.RoomID) error {
	uuid, err := toPGUUID(id)
	if err != nil {
		return err
	}

	if err := repo.q.DeleteRoom(ctx, uuid); err != nil {
		return err
	}

	return nil
}

func toInsertRoomParams(room *entity.Room) (sqlc.InsertRoomParams, error) {
	createdBy, err := toNullablePGUUID(room.CreatedBy)
	if err != nil {
		return sqlc.InsertRoomParams{}, err
	}

	return sqlc.InsertRoomParams{
		Type:      string(room.Type),
		Status:    string(room.Status),
		Capacity:  toPGInt32(room.Capacity),
		CreatedBy: createdBy,
	}, nil
}

func toUpdateRoomParams(room *entity.Room) (sqlc.UpdateRoomParams, error) {
	id, err := toPGUUID(room.ID)
	if err != nil {
		return sqlc.UpdateRoomParams{}, err
	}

	return sqlc.UpdateRoomParams{
		ID:                 id,
		Status:             string(room.Status),
		Type:               string(room.Type),
		Capacity:           toPGInt32(room.Capacity),
		ActivatedAt:        toPGTimestamptz(room.ActivatedAt),
		ClosedAt:           toPGTimestamptz(room.ClosedAt),
		LastMessageAt:      toPGTimestamptz(room.LastMessageAt),
		LastAiIntervenedAt: toPGTimestamptz(room.LastAIIntervenedAt),
	}, nil
}

func toRoomEntity(row sqlc.Room) *entity.Room {
	return entity.ReconstructRoom(
		fromPGUUID[entity.RoomID](row.ID),
		entity.RoomType(row.Type),
		entity.RoomStatus(row.Status),
		fromPGUint32(row.Capacity),
		fromNullablePGUUID[entity.UserID](row.CreatedBy),
		row.CreatedAt.Time,
		fromPGTimestamptz(row.ActivatedAt),
		fromPGTimestamptz(row.ClosedAt),
		fromPGTimestamptz(row.LastMessageAt),
		fromPGTimestamptz(row.LastAiIntervenedAt),
	)
}
