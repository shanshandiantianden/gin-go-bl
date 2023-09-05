package Models

import (
	"github.com/bytedance/sonic"
	uuid "github.com/satori/go.uuid"
)

// SessionUserInfo 当前用户会话信息
type SessionUserInfo struct {
	UserID   uint      `json:"user_id"`   // 用户ID
	UserName string    `json:"user_name"` // 用户名
	UUID     uuid.UUID `json:"uuid"`      // 用户名
}

// Marshal 序列化到JSON
func (user *SessionUserInfo) Marshal() (jsonRaw []byte) {
	jsonRaw, _ = sonic.Marshal(user)
	return
}
