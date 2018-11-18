package core

import (
	"time"
	"strconv"
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

/**
*@desc: 区块
*@author:liuhefei
*@date: 2018/11/16 22:33
*/

type Block struct {
	Timestamp     int64   //区块创建时间戳
	Data          []byte  //区块包含的数据
	PrevBlockHash []byte  //前一个区块的哈希值
	Hash          []byte  //区块自身的哈希值，用于校验区块数据有效
	Nonce         int     //证明工作量
}

//将区块转化为字节数组
func (b *Block) Serialize() []byte{
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)  //编码

	err := encoder.Encode(b)  //转换
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()  //返回字节数组
}

/**
* 初始化 创建一个新区块
* data: 数据
* prevBlockHash：前一个区块的哈希值
*/
func NewBlock(data string, prevBlockHash []byte) *Block{
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}

	//main1.go
	//block.SetHash()  //设置哈希值

	//main2.go
	//创建工作量证明
	pow := NewProofofWork(block)
	nonce, hash := pow.Run()  //挖矿
	block.Hash = hash[:]
	block.Nonce = nonce


	return block
}

/**
设置hash值
 */
 func (b *Block) SetHash(){
 	//将时间戳转化为字节数组
 	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
 	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
 	hash := sha256.Sum256(headers)   //校验
 	b.Hash = hash[:]   //得到hash值，hash值用于防止数据被篡改
 }

 /**
 创世区块，区块链的第一个区块
 创世区块的第一个hash值为空
  */
  func NewGenesisBlock() *Block{
  	return NewBlock("Genesis Block", []byte{})
  }

  func DeserializeBlock(d []byte) *Block {
  	var block Block

  	decoder := gob.NewDecoder(bytes.NewReader(d))
  	err := decoder.Decode(&block)
  	if err != nil {
  		log.Panic(err)
	}

	return  &block
  }
