package main

import (
	"part-1/core"
	"fmt"
)

/**
*@desc: 入口1——创建区块
*@author:liuhefei
*@date: 2018/11/16 23:00
*/

func main(){
	bc := core.NewBlockchain() //初始化区块链，创建第一个区块（创世纪区块）

	bc.AddBlock("Send 1 BTC to Ivan")  //加入一个区块，发送一个比特币给伊文
	bc.AddBlock("Send 2 more BTC to Ivan")  //加入一个区块，发送更多比特币给伊文

	for _,block := range bc.Blocks{
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)  //前一个区块的哈希值
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
