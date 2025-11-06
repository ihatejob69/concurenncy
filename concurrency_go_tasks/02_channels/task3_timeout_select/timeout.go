package timeout

import (
	"context"
	"time"
)

// Work выполняет длительную задачу и возвращает ошибку,
// если она заняла больше 100 мс или контекст был отменён.
func Work(ctx context.Context) error {
	// TODO: реализовать через select и time.After
	//ch := make(chan error)
	select {
	case <-time.After(time.Millisecond * 100):
		return context.DeadlineExceeded

	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}
