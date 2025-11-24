package blockchain

import (
	"context"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"github.com/tonkeeper/tongo"
	"github.com/tonkeeper/tongo/liteapi"
	"github.com/tonkeeper/tongo/tlb"
	"github.com/tonkeeper/tongo/wallet"
	"log"
)

func Transfer(receiverAddress string, amount uint64, comment string) error {
	client, err := liteapi.NewClientWithDefaultMainnet()
	if err != nil {
		log.Printf("Unable to create lite client: %v\n", err)
		return err
	}
	Mnemonic := viper.GetString("ton.mnemonic")
	log.Printf(fmt.Sprintf("%s====%d====%s", comment, amount, receiverAddress))
	w, err := wallet.DefaultWalletFromSeed(Mnemonic, client)
	if err != nil {
		log.Printf("Unable to create wallet: %v\n", err)
		return err
	}

	simpleTransfer := wallet.SimpleTransfer{
		Amount:  tlb.Grams(amount),
		Address: tongo.MustParseAddress(receiverAddress).ID,
		Comment: comment,
	}

	messageID, err := w.SendV2(context.TODO(), 0, simpleTransfer)
	if err != nil {
		log.Printf("Unable to generate transfer message: %v", err)
		return err
	}
	log.Printf("转账成功，消息 ID: %v", messageID)
	return nil
}
