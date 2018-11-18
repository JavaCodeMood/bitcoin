package core

/**
*@desc: 多个区块组成一个链条
*@author:liuhefei
*@date: 2018/11/16 22:49
*/
import (
	"github.com/boltdb/bolt"   //基于文件的存储, bolt库中有多个Bucket,每个Bucket中是一个key/value集合
	"log"
	"fmt"
)

//将区块链数据存入文件进行保存
const dbFile  = "blockchain.db"
const blocksBucket = "blocks"

type Blockchain struct {
	Blocks []*Block   //将区块放入一个数组
}

/**
向区块链中添加新数据
 */
 func (bc *Blockchain) AddBlock(data string){
 	//取出数组的末尾元素，当作新的blcok的前一个block
 	prevBlock := bc.Blocks[len(bc.Blocks) - 1]
 	//通过hash值组装创建一个新的区块
 	newBlock := NewBlock(data, prevBlock.Hash)
 	//将新创建的区块添加到区块链上
 	bc.Blocks = append(bc.Blocks, newBlock)
 }

type Blockchain1 struct {
	tip []byte
	Db *bolt.DB
}

type BlockchainIterator struct {
	currentHash []byte
	Db *bolt.DB
}

//向区块链中添加新数据
func (bc1 *Blockchain1) AddBlock1(data string){
	var lastHash []byte

	err := bc1.Db.View(func(tx *bolt.Tx) error{
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("1"))

		return nil
	})

	if err != nil {
		log.Print(err)
	}

	newBlock := NewBlock(data, lastHash)

	err = bc1.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil{
			log.Print(err)
		}

		err = b.Put([]byte("1"), newBlock.Hash)
		if err != nil{
			log.Print(err)
		}

		bc1.tip = newBlock.Hash

		return nil
	})
}

/**
创建一个新的区块链
 */
 func NewBlockchain() *Blockchain{
 	return &Blockchain{[]*Block{NewGenesisBlock()}}
 }

 //创建一个区块
func NewBlockchain1() *Blockchain1{
	var tip []byte
	//打开文件
	db,err := bolt.Open(dbFile, 0600, nil)
	if err != nil{
		log.Panic(err)
	}

	//向文件提交数据
	err = db.Update(func(tx * bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			fmt.Println("No existing blockchain found, Creating a new one...")
			genesis := NewGenesisBlock()  //创建创世区块

			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Print(err)
			}

			err = b.Put(genesis.Hash, genesis.Serialize())  //以k-v的形式存储
			if err != nil {
				log.Panic(err)
			}

			err = b.Put([]byte("l"), genesis.Hash)   //创世区块的hash
			if err != nil {
				log.Panic(err)
			}

			tip = genesis.Hash
		}else {
			tip = b.Get([]byte("1"))
		}

		return  nil
	})

	if err != nil {
		log.Panic(err)
	}

	bc1 := Blockchain1{tip, db}   //创建区块链

	return &bc1
}