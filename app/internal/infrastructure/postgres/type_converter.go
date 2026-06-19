package postgres

import (
	"time"

	"MuchUp/app/internal/domain/entity"

	"github.com/jackc/pgx/v5/pgtype"
)

type uuidID interface {
	entity.UserID | entity.RoomID | entity.MessageID
}

func toPGUUID[T uuidID](id T) (pgtype.UUID, error) {
	var uuid pgtype.UUID
	if err := uuid.Scan(string(id)); err != nil {
		return pgtype.UUID{}, err
	}

	return uuid, nil
}

func toNullablePGUUID[T uuidID](id *T) (pgtype.UUID, error) {
	if id == nil {
		return pgtype.UUID{}, nil
	}

	return toPGUUID(*id)
}

func toPGTimestamptz(t *time.Time) pgtype.Timestamptz {
	if t == nil {
		return pgtype.Timestamptz{}
	}

	return pgtype.Timestamptz{Time: *t, Valid: true}
}
