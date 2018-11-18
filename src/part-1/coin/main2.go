package main

import (
	"part-1/core"
	"fmt"
	"strconv"
)
/**
*@desc: 入口2——工作量证明
*@author:liuhefei
*@date: 2018/11/16 23:29
*/

func main(){
	bc := core.NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")  //加入一个区块，发送一个比特币给伊文
	bc.AddBlock("Send 2 more BTC to Ivan")  //加入一个区块，发送更多比特币给伊文

	for _,block := range bc.Blocks{
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)  //前一个区块的哈希值
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)  //本区块的hash值

		//工作量证明
		pow := core.NewProofofWork(block)
		//校验
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}

}
