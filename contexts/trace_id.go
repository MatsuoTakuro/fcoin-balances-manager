package contexts

import (
	"context"
	"log"
	"sync"
)

var (
	logNo int = 1
	mu    sync.Mutex
)

// contextの使用例としてtraceIDを実装
type traceIDKey struct{}

func NewTraceID() int {
	var no int

	mu.Lock()
	no = logNo
	// 本来は、in-memory db（RDS等）で採番（インクリメント）・保存すべき？
	logNo += 1
	mu.Unlock()

	return no
}

func GetTracdID(ctx context.Context) int {
	v := ctx.Value(traceIDKey{})

	if id, ok := v.(int); ok {
		return id
	}

	log.Printf("not found trace_id: %+v", ctx)
	return 0
}

func SetTraceID(ctx context.Context, tracdID int) context.Context {
	return context.WithValue(ctx, traceIDKey{}, tracdID)
}
