package core

import (
	"math"
	"math/big"
	"bytes"
	"fmt"
	"crypto/sha256"
)

/**
*@desc: 工作量证明
*@author:liuhefei
*@date: 2018/11/16 23:48
*/

var (
	maxNonce = math.MaxInt64
)

const targetBits = 20  //目标位(20-5个0，4-1个0)，决定计算的难易程度

//工作量证明
type ProofofWork struct {
	block *Block    //区块
	target *big.Int //目标
}

/**
创建工作量证明
 */
func NewProofofWork(b *Block) *ProofofWork{
	target := big.NewInt(1)
	//移位操作,左移，前20个字节为0
	target.Lsh(target, uint(256 - targetBits))

	pow := &ProofofWork{b, target}

	return pow
}

/**
拼接数据
 */
func (pow *ProofofWork) prepareData(nonce int) []byte{
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

/**
挖矿
 */
func (pow *ProofofWork) Run() (int, []byte){
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	fmt.Printf("Mining the block containing \"%s\"\n" , pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)

		hash = sha256.Sum256(data)  //对数据进行hash计算
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])  //将hash值转化为hash整数

		if hashInt.Cmp(pow.target) == -1{   //hash整数与target做对比
			break  //对比成功就退出循环
		}else{
			nonce ++  //nonce从0开始自增到int64为的最大值，每次自增都要调用prepareData方法拼接数据
		}
	}

	fmt.Print("\n\n")

	return nonce, hash[:]
}

/**
校验区块算法
 */
func (pow *ProofofWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}