package utils

import (
	"sync"
	"time"
)

// Snowflake 结构体
type Snowflake struct {
	mu         sync.Mutex // 互斥锁
	timestamp  int64      // 时间戳
	sequenceID int64      // 序列号
	machineID  int64      // 机器ID
}

// 饿汉式创建唯一的 Snowflake 实例
var snowflake = &Snowflake{
	timestamp:  0,
	sequenceID: 1,
	machineID:  0,
}

// NewSnowflake 函数，返回一个 Snowflake 实例
func NewSnowflake() *Snowflake {
	return snowflake
}

// Generate 生成唯一 ID
func (s *Snowflake) Generate() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().UnixNano() / int64(time.Millisecond) // 获取当前时间戳(毫秒)
	if s.timestamp == now {                                // 如果当前时间戳和上次生成时间戳相同，则序列号增 1
		s.sequenceID = (s.sequenceID + 1) & 4095 // 确保序列号不溢出(4095 = 2^12 - 1表明序列号占12位)
		if s.sequenceID == 0 {
			// 序列号溢出，等待下一毫秒
			for now <= s.timestamp {
				now = time.Now().UnixNano() / int64(time.Millisecond)
			}
		}
	} else {
		s.sequenceID = 0 // 否则，重置序列号为 0
	}
	s.timestamp = now // 更新时间戳为当前时间

	ID := (now << 22) | (s.machineID << 10) | s.sequenceID // 生成 ID
	return ID
}
