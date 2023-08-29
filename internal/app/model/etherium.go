package model

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pullya/tx_parcer/restapi/operations/parser_interface"
)

type EthResponse struct {
	Id     int64     `json:"id"`
	Result string    `json:"result"`
	Error  *EthError `json:"error"`
}

type EthError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type EthLog struct {
	Removed          bool     `json:"removed"`
	LogIndex         string   `json:"logIndex"`
	TransactionIndex string   `json:"transactionIndex"`
	TransactionHash  string   `json:"transactionHash"`
	BlockHash        string   `json:"blockHash"`
	BlockNumber      string   `json:"blockNumber"`
	Address          string   `json:"address"`
	Data             string   `json:"data"`
	Topics           []string `json:"topics"`
}

type EthLogsResponse struct {
	Id     int64     `json:"id"`
	Result []EthLog  `json:"result"`
	Error  *EthError `json:"error"`
}

type EthTransaction struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	From             string `json:"from"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Hash             string `json:"hash"`
	Input            string `json:"input"`
	Nonce            string `json:"nonce"`
	To               string `json:"to"`
	TransactionIndex string `json:"transactionIndex"`
	Value            string `json:"value"`
	V                string `json:"v"`
	R                string `json:"r"`
	S                string `json:"s"`
}

type EthTransactionResponse struct {
	Id     int64          `json:"id"`
	Result EthTransaction `json:"result"`
	Error  *EthError      `json:"error"`
}

func Hex2Int(inHex string) int {
	result, err := strconv.ParseInt(strings.Replace(inHex, "0x", "", -1), 16, 64)
	if err != nil {
		fmt.Println("Failed to convert hex to int", err)
		return 0
	}
	return int(result)
}

func (rt *EthTransaction) AsResponse() *parser_interface.TxParserGetTransactionsDefaultBodyTransactionsItems0 {
	return &parser_interface.TxParserGetTransactionsDefaultBodyTransactionsItems0{
		BlockHash:        rt.BlockHash,
		BlockNumber:      rt.BlockNumber,
		From:             rt.From,
		Gas:              rt.Gas,
		GasPrice:         rt.GasPrice,
		Hash:             rt.Hash,
		Input:            rt.Input,
		Nonce:            rt.Nonce,
		R:                rt.R,
		S:                rt.S,
		To:               rt.To,
		TransactionIndex: rt.TransactionIndex,
		V:                rt.V,
		Value:            rt.Value,
	}
}
