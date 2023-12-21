package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func transferETH(client *ethclient.Client,fromPrivKey *ecdsa.PrivateKey, to common.Address, amount *big.Int) error{
	ctx := context.Background()
	publicKey := fromPrivKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok{
		return fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return err
	}
	var gasLimit = uint64(21000)

	chainid, err := client.ChainID(ctx)
	if err != nil {
		log.Fatal(err)
	}


	tipCap, _ := client.SuggestGasTipCap(ctx)
	feeCap, _ := client.SuggestGasPrice(ctx)

		tx := types.NewTx(
		&types.DynamicFeeTx{
			ChainID:   chainid,
			Nonce:     nonce,
			GasTipCap: tipCap,
			GasFeeCap: feeCap,
			Gas:       gasLimit,
			To:        &to,
			Value:     amount,
			Data:      nil,
		})

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainid), fromPrivKey)
	if err !=nil{
		return err
	}
	return client.SendTransaction(ctx, signedTx)
}