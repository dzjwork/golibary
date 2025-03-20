package logrus

import (
	"bytes"
	"sync"
)

var (
	bufferPool BufferPool
)

// BufferPool 缓冲池，可读可写
type BufferPool interface {
	Put(*bytes.Buffer)
	Get() *bytes.Buffer
}

// 实现了BufferPool接口
type defaultPool struct {
	pool *sync.Pool
}

func (p *defaultPool) Put(buf *bytes.Buffer) {
	p.pool.Put(buf)
}

func (p *defaultPool) Get() *bytes.Buffer {
	return p.pool.Get().(*bytes.Buffer)
}

// SetBufferPool 设置bufferPool
func SetBufferPool(bp BufferPool) {
	bufferPool = bp
}

func init() {
	SetBufferPool(&defaultPool{
		pool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	})
}
