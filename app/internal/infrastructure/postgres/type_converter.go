package postgres

import (
	"math"
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

func toPGText(s *string) pgtype.Text {
	if s == nil {
		return pgtype.Text{}
	}

	return pgtype.Text{String: *s, Valid: true}
}

// toPGInt32 はページングのlimit/offsetをint32へ変換する。
// limit/offsetはuintのため負値は型レベルで排除済みで、ここではint32の範囲を
// 超える値をint32の最大値に丸め、オーバーフローを防ぐ。
func toPGInt32(n uint) int32 {
	if n > math.MaxInt32 {
		return math.MaxInt32
	}

	return int32(n)
}

func toPGTimestamptz(t *time.Time) pgtype.Timestamptz {
	if t == nil {
		return pgtype.Timestamptz{}
	}

	return pgtype.Timestamptz{Time: *t, Valid: true}
}

// toPGSequence はSequence未設定(0)をNULLとして扱う暫定実装。
// row.Sequence.IntからのReconstruct側もNULLを0として読むため、それに合わせて0とNULLを同一視している。
// Sequenceの採番ロジック(Redisストリーム等)がまだ実装されておらず0始まりか1始まりか未確定のため、
// 採番ロジック実装時にこの前提が成立するか要議論。
// 代替案: entity.Message.Sequenceを*int64にすれば採番方式に依存せず安全になるが、
// 呼び出し側全体にnilハンドリングが波及するため今回は見送っている。
// TODO: 採番ロジック実装時に0始まり/1始まりを確定し、この前提(0=未採番)のままでよいか
// *int64への変更が必要か議論する。
func toPGSequence(sequence int64) pgtype.Int8 {
	if sequence == 0 {
		return pgtype.Int8{}
	}

	return pgtype.Int8{Int64: sequence, Valid: true}
}

// fromPGUint32 はDB由来のint32(capacityなど)を非負のuintへ変換する。
// 対象カラムにはDB側でCHECK制約(値 > 0)がある想定だが、gosecのint32→uint変換警告を
// 避けるため念のため負値は0に丸める。
func fromPGUint32(n int32) uint {
	if n < 0 {
		return 0
	}

	return uint(n)
}

func fromPGUUID[T uuidID](uuid pgtype.UUID) T {
	return T(uuid.String())
}

func fromNullablePGUUID[T uuidID](uuid pgtype.UUID) *T {
	if !uuid.Valid {
		return nil
	}

	id := fromPGUUID[T](uuid)

	return &id
}

func fromPGText(t pgtype.Text) *string {
	if !t.Valid {
		return nil
	}

	return &t.String
}

func fromPGTimestamptz(t pgtype.Timestamptz) *time.Time {
	if !t.Valid {
		return nil
	}

	return &t.Time
}
