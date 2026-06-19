package postgres

import (
	"context"

	"MuchUp/app/internal/domain/entity"
	"MuchUp/app/internal/domain/repository"
	"MuchUp/app/sqlc"
)

type messageRepository struct {
	q *sqlc.Queries
}

func NewMessageRepository(db sqlc.DBTX) repository.MessageRepository {
	return &messageRepository{q: sqlc.New(db)}
}

func (repo *messageRepository) Insert(ctx context.Context, message *entity.Message) error {
	param := sqlc.InsertMessageParams{}

	if err := repo.q.InsertMessage(ctx, param); err != nil {
		return err
	}

	return nil
}

func (repo *messageRepository) InsertText(ctx context.Context, message *entity.Message) error {
	param := sqlc.InsertTextMessageParams{}

	if err := repo.q.InsertTextMessage(ctx, param); err != nil {
		return err
	}

	return nil
}

func (repo *messageRepository) FindByID(ctx context.Context, id entity.MessageID) (*entity.Message, error) {
	uuid, err := toPGUUID(id)
	if err != nil {
		return nil, err
	}

	result, err := repo.q.FindMessageByID(ctx, uuid)
	if err != nil {
		return nil, err
	}
	_ = result

	// TODO: Reconstructメソッドを用意する
	return &entity.Message{}, nil
}

func (repo *messageRepository) FindByUserID(ctx context.Context, userID entity.UserID) ([]*entity.Message, error) {
	param := sqlc.FindAllMessagesByUserIDParams{}

	result, err := repo.q.FindAllMessagesByUserID(ctx, param)
	if err != nil {
		return nil, err
	}
	_ = result

	// TODO: Reconstructメソッドを用意する
	return make([]*entity.Message, 0), nil
}

func (repo *messageRepository) Update(ctx context.Context, message *entity.Message) error {
	param := sqlc.UpdateMessageParams{}

	if err := repo.q.UpdateMessage(ctx, param); err != nil {
		return err
	}

	return nil
}

func (repo *messageRepository) Delete(ctx context.Context, id entity.MessageID) error {
	uuid, err := toPGUUID(id)
	if err != nil {
		return err
	}

	if err := repo.q.DeleteMessage(ctx, uuid); err != nil {
		return err
	}

	return nil
}

func (repo *messageRepository) FindByRoomID(ctx context.Context, roomID entity.RoomID, offset, limit int) ([]*entity.Message, error) {
	param := sqlc.FindAllMessagesByRoomParams{}

	result, err := repo.q.FindAllMessagesByRoom(ctx, param)
	if err != nil {
		return nil, err
	}
	_ = result

	// TODO: Reconstructメソッドを用意する
	return make([]*entity.Message, 0), nil

}
