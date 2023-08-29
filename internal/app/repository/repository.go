package repository

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/pullya/tx_parcer/internal/app/model"
)

type Repository struct {
	Subscriptions map[string]model.Subscription
	LastId        int64
}

func NewRepository() *Repository {
	return &Repository{
		Subscriptions: make(map[string]model.Subscription),
		LastId:        0,
	}
}

func (r *Repository) AddSubscription(address string) (int64, error) {
	var mu sync.Mutex

	mu.Lock()
	lid := atomic.AddInt64(&r.LastId, 1)
	r.Subscriptions[address] = model.Subscription{Id: lid, Address: address}
	r.LastId = lid
	mu.Unlock()

	return lid, nil
}

func (r *Repository) GetSubsctiptionByAddress(addr string) (int64, error) {
	var mu sync.Mutex

	mu.Lock()
	s, ok := r.Subscriptions[addr]
	mu.Unlock()

	if !ok {
		return 0, fmt.Errorf("GetSubscriptionById: Subscription not found by address: %s", addr)
	}

	return s.Id, nil
}

func (r *Repository) DeleteSubscriptionByAddress(addr string) error {
	var mu sync.Mutex

	mu.Lock()
	defer mu.Unlock()

	if _, ok := r.Subscriptions[addr]; !ok {
		return fmt.Errorf("DeleteSubscriptionById: Subscription not found by address: %s", addr)
	}
	delete(r.Subscriptions, addr)

	return nil
}
