package store

import "MuchUp/app/internal/domain/repository"

// TODO: usecase 層で利用する store の公開境界をここで整理していく
type RoomUserStore = repository.RoomUserStore

// TODO: メッセージ配信まわりの store 抽象をここに集約していく
type MessageStreamStore = repository.MessageStreamStore
