package service

import (
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/pullya/tx_parcer/internal/app/client"
	"github.com/pullya/tx_parcer/internal/app/model"
)

var counter int64

type Parser interface {
	// last parsed block
	GetCurrentBlock() int

	// add address to observer
	Subscribe(address string) bool

	// list of inbound or outbound transactions for an address
	GetTransactions(address string) []model.EthTransaction
}

type ITxParserRepository interface {
	AddSubscription(address string) (int64, error)
	GetSubsctiptionByAddress(addr string) (int64, error)
	DeleteSubscriptionByAddress(addr string) error
}

type TxParserService struct {
	TxParseRepository ITxParserRepository
}

func NewTxParserService(repo ITxParserRepository) TxParserService {
	return TxParserService{
		TxParseRepository: repo,
	}
}

func (tps TxParserService) GetCurrentBlock() int {

	payload := fmt.Sprintf("{\"jsonrpc\":\"2.0\",\"method\":\"eth_blockNumber\",\"params\":[],\"id\":%d}", generateUniqueID())

	resp, err := client.CloudflareCall(payload)
	if err != nil {
		fmt.Println("GetCurrentBlock eth_blockNumber: failed to get current block", err)
		return 0
	}

	var response model.EthResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		fmt.Println("GetCurrentBlock: failed to unmarshal response", err)
		return 0
	}
	if response.Error != nil {
		errr := *response.Error
		fmt.Printf("GetCurrentBlock: error received. code=%d message=%s\n", errr.Code, errr.Message)
		return 0
	}

	return model.Hex2Int(response.Result)
}

func (tps *TxParserService) Subscribe(address string) bool {

	fmt.Println("received request, address =", address)
	id, err := tps.TxParseRepository.AddSubscription(address)
	if err != nil {
		fmt.Println("Subscribe: failed to save subscription info", err)
		return false
	}

	fmt.Println("New subscription added with id =", id)

	return true
}

func (tps *TxParserService) GetTransactions(address string) []model.EthTransaction {

	payload := fmt.Sprintf("{\"jsonrpc\":\"2.0\",\"method\":\"eth_getLogs\",\"params\": [{\"address\": \"%s\"}],\"id\":%d}", address, generateUniqueID())
	resp, err := client.CloudflareCall(payload)
	if err != nil {
		fmt.Println("GetTransactions eth_getLogs: failed to get logs by address", err)
		return nil
	}

	var response model.EthLogsResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		fmt.Println("GetTransactions: failed to unmarshal logs response", err)
		return nil
	}

	if response.Error != nil {
		errr := *response.Error
		fmt.Printf("Subscribe: error received. code=%d message=%s\n", errr.Code, errr.Message)
		return nil
	}

	transactions := make([]model.EthTransaction, 0, len(response.Result))
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for _, log := range response.Result {
		wg.Add(1)
		log := log
		go func() {
			defer wg.Done()

			payload := fmt.Sprintf("{\"jsonrpc\":\"2.0\",\"method\":\"eth_getTransactionByHash\",\"params\": [\"%s\"],\"id\":%d}", log.TransactionHash, generateUniqueID())
			resp, err := client.CloudflareCall(payload)
			if err != nil {
				fmt.Println("GetTransactions eth_getTransactionByHash: failed to get transaction by hash", err)
				return
			}

			var tran model.EthTransactionResponse
			err = json.Unmarshal(resp, &tran)
			if err != nil {
				fmt.Println("GetTransactions: failed to unmarshal transaction response", err)
				return
			}
			if tran.Error != nil {
				errr := *tran.Error
				fmt.Printf("GetTransactions: error received. code=%d message=%s\n", errr.Code, errr.Message)
				return
			}

			mu.Lock()
			transactions = append(transactions, tran.Result)
			mu.Unlock()
		}()
	}
	wg.Wait()

	return transactions
}

func generateUniqueID() int64 {
	return atomic.AddInt64(&counter, 1)
}
