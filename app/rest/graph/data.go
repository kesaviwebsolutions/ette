package graph

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/itzmeanjan/ette/app/data"
	"github.com/itzmeanjan/ette/app/rest/graph/model"
	"gorm.io/gorm"
)

var db *gorm.DB

// GetDatabaseConnection - Passing already connected database handle to this package,
// so that it can be used for handling database queries for resolving graphQL queries
func GetDatabaseConnection(conn *gorm.DB) {
	db = conn
}

// Converting block data to graphQL compatible data structure
func getGraphQLCompatibleBlock(block *data.Block) *model.Block {
	if block == nil {
		return nil
	}

	return &model.Block{
		Hash:                block.Hash,
		Number:              fmt.Sprintf("%d", block.Number),
		Time:                fmt.Sprintf("%d", block.Time),
		ParentHash:          block.ParentHash,
		Difficulty:          block.Difficulty,
		GasUsed:             fmt.Sprintf("%d", block.GasUsed),
		GasLimit:            fmt.Sprintf("%d", block.GasLimit),
		Nonce:               fmt.Sprintf("%d", block.Nonce),
		Miner:               block.Miner,
		Size:                block.Size,
		TransactionRootHash: block.TransactionRootHash,
		ReceiptRootHash:     block.ReceiptRootHash,
	}
}

// Converting block array to graphQL compatible data structure
func getGraphQLCompatibleBlocks(blocks []*data.Block) []*model.Block {
	if blocks == nil {
		return nil
	}

	_blocks := make([]*model.Block, len(blocks))

	for k, v := range blocks {
		_blocks[k] = getGraphQLCompatibleBlock(v)
	}

	return _blocks
}

// Converting transaction data to graphQL compatible data structure
func getGraphQLCompatibleTransaction(tx *data.Transaction) *model.Transaction {
	if tx == nil {
		return nil
	}

	return &model.Transaction{
		Hash:      tx.Hash,
		From:      tx.From,
		To:        tx.To,
		Contract:  tx.Contract,
		Gas:       fmt.Sprintf("%d", tx.Gas),
		GasPrice:  tx.GasPrice,
		Cost:      tx.Cost,
		Nonce:     fmt.Sprintf("%d", tx.Nonce),
		State:     fmt.Sprintf("%d", tx.State),
		BlockHash: tx.BlockHash,
	}
}

// Converting transaction array to graphQL compatible data structure
func getGraphQLCompatibleTransactions(tx []*data.Transaction) []*model.Transaction {
	if tx == nil {
		return nil
	}

	_tx := make([]*model.Transaction, len(tx))

	for k, v := range tx {
		_tx[k] = getGraphQLCompatibleTransaction(v)
	}

	return _tx
}

// Extracted from, to field of range based block query ( using block numbers/ time stamps )
// gets parsed into unsigned integers
func rangeChecker(from string, to string, limit uint64) (uint64, uint64, error) {
	_from, err := strconv.ParseUint(from, 10, 64)
	if err != nil {
		return 0, 0, errors.New("Failed to parse integer")
	}

	_to, err := strconv.ParseUint(to, 10, 64)
	if err != nil {
		return 0, 0, errors.New("Failed to parse integer")
	}

	if !(_to-_from < limit) {
		return 0, 0, errors.New("Range too long")
	}

	return _from, _to, nil
}
