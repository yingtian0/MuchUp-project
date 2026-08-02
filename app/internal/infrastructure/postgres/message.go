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
	param, err := toInsertMessageParams(message)
	if err != nil {
		return err
	}

	if err := repo.q.InsertMessage(ctx, param); err != nil {
		return err
	}

	return nil
}

func (repo *messageRepository) InsertText(ctx context.Context, message *entity.Message) error {
	param, err := toInsertTextMessageParams(message)
	if err != nil {
		return err
	}

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

	row, err := repo.q.FindMessageByID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return toMessageEntity(row), nil
}

func (repo *messageRepository) FindByUserID(ctx context.Context, userID entity.UserID, limit, offset uint) ([]*entity.Message, error) {
	senderID, err := toPGUUID(userID)
	if err != nil {
		return nil, err
	}

	param := sqlc.FindAllMessagesByUserIDParams{
		SenderID: senderID,
		Limit:    toPGInt32(limit),
		Offset:   toPGInt32(offset),
	}

	rows, err := repo.q.FindAllMessagesByUserID(ctx, param)
	if err != nil {
		return nil, err
	}

	messages := make([]*entity.Message, len(rows))
	for i, row := range rows {
		messages[i] = toMessageEntity(row)
	}

	return messages, nil
}

func (repo *messageRepository) Update(ctx context.Context, message *entity.Message) error {
	param, err := toUpdateMessageParams(message)
	if err != nil {
		return err
	}

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

func (repo *messageRepository) FindByRoomID(ctx context.Context, roomID entity.RoomID, limit, offset uint) ([]*entity.Message, error) {
	uuid, err := toPGUUID(roomID)
	if err != nil {
		return nil, err
	}

	param := sqlc.FindAllMessagesByRoomParams{
		RoomID: uuid,
		Limit:  toPGInt32(limit),
		Offset: toPGInt32(offset),
	}

	rows, err := repo.q.FindAllMessagesByRoom(ctx, param)
	if err != nil {
		return nil, err
	}

	messages := make([]*entity.Message, len(rows))
	for i, row := range rows {
		messages[i] = toMessageEntity(row)
	}

	return messages, nil
}

func toInsertMessageParams(message *entity.Message) (sqlc.InsertMessageParams, error) {
	roomID, err := toPGUUID(message.RoomID)
	if err != nil {
		return sqlc.InsertMessageParams{}, err
	}

	senderID, err := toPGUUID(message.SenderID)
	if err != nil {
		return sqlc.InsertMessageParams{}, err
	}

	return sqlc.InsertMessageParams{
		RoomID:     roomID,
		SenderID:   senderID,
		SenderType: string(message.SenderType),
		Kind:       string(message.Kind),
		Status:     string(message.Status),
		Text:       toPGText(message.Text),
		MediaUrl:   toPGText(message.MediaURL),
		StickerID:  toPGText(message.StickerID),
		StreamID:   toPGText(message.StreamID),
		Sequence:   toPGSequence(message.Sequence),
	}, nil
}

func toInsertTextMessageParams(message *entity.Message) (sqlc.InsertTextMessageParams, error) {
	roomID, err := toPGUUID(message.RoomID)
	if err != nil {
		return sqlc.InsertTextMessageParams{}, err
	}

	senderID, err := toPGUUID(message.SenderID)
	if err != nil {
		return sqlc.InsertTextMessageParams{}, err
	}

	return sqlc.InsertTextMessageParams{
		RoomID:     roomID,
		SenderID:   senderID,
		SenderType: string(message.SenderType),
		Text:       toPGText(message.Text),
		StreamID:   toPGText(message.StreamID),
		Sequence:   toPGSequence(message.Sequence),
	}, nil
}

func toUpdateMessageParams(message *entity.Message) (sqlc.UpdateMessageParams, error) {
	id, err := toPGUUID(message.ID)
	if err != nil {
		return sqlc.UpdateMessageParams{}, err
	}

	return sqlc.UpdateMessageParams{
		ID:        id,
		Status:    string(message.Status),
		Text:      toPGText(message.Text),
		MediaUrl:  toPGText(message.MediaURL),
		StickerID: toPGText(message.StickerID),
		StreamID:  toPGText(message.StreamID),
		Sequence:  toPGSequence(message.Sequence),
	}, nil
}

func toMessageEntity(row sqlc.Message) *entity.Message {
	return entity.ReconstructMessage(
		fromPGUUID[entity.MessageID](row.ID),
		nil,
		fromPGUUID[entity.UserID](row.SenderID),
		fromPGUUID[entity.RoomID](row.RoomID),
		fromPGText(row.Text),
		fromPGText(row.MediaUrl),
		fromPGText(row.StickerID),
		row.CreatedAt.Time,
		row.UpdatedAt.Time,
		row.DeletedAt.Time,
		entity.SenderType(row.SenderType),
		entity.MessageKind(row.Kind),
		entity.MessageStatus(row.Status),
		nil,
		fromPGText(row.StreamID),
		row.Sequence.Int64,
	)
}
