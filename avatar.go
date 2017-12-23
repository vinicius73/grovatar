package main

import (
	"crypto/md5"
	"strings"
)

// AvatarData ...
type AvatarData struct {
	Email string
	Hash  [md5.Size]byte
	Color [3]byte
	Table HashTable
}

func (m *AvatarData) build(email string) {
	m.Email = strings.ToLower(email)
	m.Hash = HashMail([]byte(m.Email))
	m.Color = [3]byte{m.Hash[0], m.Hash[1], m.Hash[2]}
	m.Table = MakeHashTable(m.Hash)
}

// MakeAvatar ...
func MakeAvatar(email string) (AvatarData, error) {
	avatar := AvatarData{}

	avatar.build(email)

	return avatar, nil
}
