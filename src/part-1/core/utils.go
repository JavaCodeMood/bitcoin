package core

import (
	"bytes"
	"encoding/binary"
	"log"
	"crypto/sha256"
)

/**
*@desc: 工具
*@author:liuhefei
*@date: 2018/11/16 23:44
*/
func IntToHex(num int64) []byte{
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil{
		log.Panic(err)
	}

	return buff.Bytes()
}

func DataToHash(data []byte) []byte{
	hash := sha256.Sum256(data)

	return hash[:]
}
