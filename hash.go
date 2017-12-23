package main

import (
	"crypto/md5"
)

// HashTable ...
type HashTable = [5][5]bool

// HashMail ...
func HashMail(data []byte) [md5.Size]byte {
	return md5.Sum(data)
}

// MakeHashTable ...
func MakeHashTable(hash [md5.Size]byte) (table HashTable) {
	for index, arr := range table {
		table[index] = fillHashTableLine(index, arr, hash)
	}
	return table
}

func fillHashTableLine(index int, line [5]bool, hash [md5.Size]byte) [5]bool {
	for i := 0; 2 >= i; i++ {
		line[i] = (hash[index+i] % 2) == 0
	}

	line[3] = line[1]
	line[4] = line[0]

	return line
}
