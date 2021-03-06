// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package executor

import (
	"github.com/33cn/chain33/types"
	pty "github.com/33cn/plugin/plugin/dapp/trade/types"
)

func (t *trade) ExecLocal_SellLimit(sell *pty.TradeForSell, tx *types.Transaction, receipt *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	return t.localAddLog(tx, receipt, index)
}

func (t *trade) ExecLocal_BuyMarket(buy *pty.TradeForBuy, tx *types.Transaction, receipt *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	return t.localAddLog(tx, receipt, index)
}

func (t *trade) ExecLocal_RevokeSell(revoke *pty.TradeForRevokeSell, tx *types.Transaction, receipt *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	return t.localAddLog(tx, receipt, index)
}

func (t *trade) ExecLocal_BuyLimit(buy *pty.TradeForBuyLimit, tx *types.Transaction, receipt *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	return t.localAddLog(tx, receipt, index)
}

func (t *trade) ExecLocal_SellMarket(sell *pty.TradeForSellMarket, tx *types.Transaction, receipt *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	return t.localAddLog(tx, receipt, index)
}

func (t *trade) ExecLocal_RevokeBuy(revoke *pty.TradeForRevokeBuy, tx *types.Transaction, receipt *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	return t.localAddLog(tx, receipt, index)
}

func (t *trade) localAddLog(tx *types.Transaction, receipt *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	var set types.LocalDBSet

	for i := 0; i < len(receipt.Logs); i++ {
		item := receipt.Logs[i]
		if item.Ty == pty.TyLogTradeSellLimit || item.Ty == pty.TyLogTradeSellRevoke {
			var receipt pty.ReceiptTradeSellLimit
			err := types.Decode(item.Log, &receipt)
			if err != nil {
				panic(err) //数据错误了，已经被修改了
			}
			kv := t.saveSell([]byte(receipt.Base.SellID), item.Ty)
			set.KV = append(set.KV, kv...)
		} else if item.Ty == pty.TyLogTradeBuyMarket {
			var receipt pty.ReceiptTradeBuyMarket
			err := types.Decode(item.Log, &receipt)
			if err != nil {
				panic(err) //数据错误了，已经被修改了
			}
			kv := t.saveBuy(receipt.Base)
			set.KV = append(set.KV, kv...)
		} else if item.Ty == pty.TyLogTradeBuyRevoke || item.Ty == pty.TyLogTradeBuyLimit {
			var receipt pty.ReceiptTradeBuyLimit
			err := types.Decode(item.Log, &receipt)
			if err != nil {
				panic(err) //数据错误了，已经被修改了
			}

			kv := t.saveBuyLimit([]byte(receipt.Base.BuyID), item.Ty)
			set.KV = append(set.KV, kv...)
		} else if item.Ty == pty.TyLogTradeSellMarket {
			var receipt pty.ReceiptSellMarket
			err := types.Decode(item.Log, &receipt)
			if err != nil {
				panic(err) //数据错误了，已经被修改了
			}
			kv := t.saveSellMarket(receipt.Base)
			//tradelog.Info("saveSellMarket", "kv", kv)
			set.KV = append(set.KV, kv...)
		}
	}

	return &set, nil
}
