// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package util

import (
	"errors"

	"github.com/33cn/chain33/common"
	log "github.com/33cn/chain33/common/log/log15"
	"github.com/33cn/chain33/queue"
	"github.com/33cn/chain33/types"
)

//CheckBlock : To check the block's validaty
func CheckBlock(client queue.Client, block *types.BlockDetail) error {
	req := block
	msg := client.NewMessage("consensus", types.EventCheckBlock, req)
	client.Send(msg, true)
	resp, err := client.Wait(msg)
	if err != nil {
		return err
	}
	reply := resp.GetData().(*types.Reply)
	if reply.IsOk {
		return nil
	}
	return errors.New(string(reply.GetMsg()))
}

//ExecTx : To send lists of txs within a block to exector for execution
func ExecTx(client queue.Client, prevStateRoot []byte, block *types.Block) *types.Receipts {
	list := &types.ExecTxList{
		StateHash:  prevStateRoot,
		Txs:        block.Txs,
		BlockTime:  block.BlockTime,
		Height:     block.Height,
		Difficulty: uint64(block.Difficulty),
		IsMempool:  false,
	}
	msg := client.NewMessage("execs", types.EventExecTxList, list)
	client.Send(msg, true)
	resp, err := client.Wait(msg)
	if err != nil {
		panic(err)
	}
	receipts := resp.GetData().(*types.Receipts)
	return receipts
}

//ExecKVMemSet : send kv values to memory store and set it in db
func ExecKVMemSet(client queue.Client, prevStateRoot []byte, height int64, kvset []*types.KeyValue, sync bool) []byte {
	set := &types.StoreSet{StateHash: prevStateRoot, KV: kvset, Height: height}
	setwithsync := &types.StoreSetWithSync{Storeset: set, Sync: sync}

	msg := client.NewMessage("store", types.EventStoreMemSet, setwithsync)
	client.Send(msg, true)
	resp, err := client.Wait(msg)
	if err != nil {
		panic(err)
	}
	hash := resp.GetData().(*types.ReplyHash)
	return hash.GetHash()
}

//ExecKVSetCommit : commit the data set opetation to db
func ExecKVSetCommit(client queue.Client, hash []byte) error {
	req := &types.ReqHash{Hash: hash}
	msg := client.NewMessage("store", types.EventStoreCommit, req)
	client.Send(msg, true)
	msg, err := client.Wait(msg)
	if err != nil {
		return err
	}
	hash = msg.GetData().(*types.ReplyHash).GetHash()
	_ = hash
	return nil
}

//ExecKVSetRollback : do the db's roll back operation
func ExecKVSetRollback(client queue.Client, hash []byte) error {
	req := &types.ReqHash{Hash: hash}
	msg := client.NewMessage("store", types.EventStoreRollback, req)
	client.Send(msg, true)
	msg, err := client.Wait(msg)
	if err != nil {
		return err
	}
	hash = msg.GetData().(*types.ReplyHash).GetHash()
	_ = hash
	return nil
}

func checkTxDupInner(txs []*types.TransactionCache) (ret []*types.TransactionCache) {
	dupMap := make(map[string]bool)
	for _, tx := range txs {
		hash := string(tx.Hash())
		if _, ok := dupMap[hash]; ok {
			continue
		}
		dupMap[hash] = true
		ret = append(ret, tx)
	}
	return ret
}

//CheckDupTx : check use txs []*types.Transaction and not []*types.TransactionCache
func CheckDupTx(client queue.Client, txs []*types.Transaction, height int64) (transactions []*types.Transaction, err error) {
	txcache := make([]*types.TransactionCache, len(txs))
	for i := 0; i < len(txcache); i++ {
		txcache[i] = &types.TransactionCache{Transaction: txs[i]}
	}
	cache, err := CheckTxDup(client, txcache, height)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(cache); i++ {
		transactions = append(transactions, cache[i].Transaction)
	}
	return transactions, nil
}

//CheckTxDup : check whether the tx is duplicated within the while chain
func CheckTxDup(client queue.Client, txs []*types.TransactionCache, height int64) (transactions []*types.TransactionCache, err error) {
	var checkHashList types.TxHashList
	if types.IsFork(height, "ForkCheckTxDup") {
		txs = checkTxDupInner(txs)
	}
	for _, tx := range txs {
		checkHashList.Hashes = append(checkHashList.Hashes, tx.Hash())
		checkHashList.Expire = append(checkHashList.Expire, tx.GetExpire())
	}
	checkHashList.Count = height
	hashList := client.NewMessage("blockchain", types.EventTxHashList, &checkHashList)
	client.Send(hashList, true)
	dupTxList, err := client.Wait(hashList)
	if err != nil {
		return nil, err
	}
	dupTxs := dupTxList.GetData().(*types.TxHashList).Hashes
	dupMap := make(map[string]bool)
	for _, hash := range dupTxs {
		dupMap[string(hash)] = true
		log.Debug("CheckTxDup", "TxDuphash", common.ToHex(hash))
	}
	for _, tx := range txs {
		hash := tx.Hash()
		if dupMap[string(hash)] {
			continue
		}
		transactions = append(transactions, tx)
	}
	return transactions, nil
}

//ReportErrEventToFront 上报指定错误信息到指定模块，目前只支持从store，blockchain，wallet写数据库失败时上报错误信息到wallet模块，
//然后由钱包模块前端定时调用显示给客户
func ReportErrEventToFront(logger log.Logger, client queue.Client, frommodule string, tomodule string, err error) {
	if client == nil || len(tomodule) == 0 || len(frommodule) == 0 || err == nil {
		logger.Error("SendErrEventToFront  input para err .")
		return
	}

	logger.Debug("SendErrEventToFront", "frommodule", frommodule, "tomodule", tomodule, "err", err)

	var reportErrEvent types.ReportErrEvent
	reportErrEvent.Frommodule = frommodule
	reportErrEvent.Tomodule = tomodule
	reportErrEvent.Error = err.Error()
	msg := client.NewMessage(tomodule, types.EventErrToFront, &reportErrEvent)
	client.Send(msg, false)
}
