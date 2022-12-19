package service

import (
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func FindTx(block *types.Block, backupSync bool, contractAddress string, chainId int, client *ethclient.Client) {
	openFlag := false
	for _, tx := range block.Transactions() {
		if tx.To() == nil {
			continue
		}

		// if tx.To().String() == contractAddress {
		fmt.Println("matched------------------")

		openFlag = true
		SaveTx(tx.Hash().String(), contractAddress, uint(block.Number().Uint64()), chainId)
		// }
	}

	SaveLastConfirmed(int(block.Number().Int64()), chainId, backupSync)

	fmt.Println("Block parsed : ", block.Number().Uint64())
	if openFlag {
		// go OpenTx(client, chainId)
	}
	// signatureService();// 10000

}

func blockConf() {

}

// 1000

// 995
// signed

// func signatureService(){

// if tx status = true {

// }
// else{
// 	delete tx fromn table;
// }
// select * from g_txburninfo where chainid=? AND completed=? AND blocknum < (currentblock-5) AND signed = false;

// for(on tx list){
// 	getsign(tx)
// 	save in db;
// }

// }
