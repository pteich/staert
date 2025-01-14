package staert

import (
	"context"
	"errors"
	"strings"

	"github.com/kvtools/valkeyrie/store"
)

// Extremely limited mock store so we can test initialization
type Mock struct {
	Error           bool
	KVPairs         []*store.KVPair
	WatchTreeMethod func() <-chan []*store.KVPair
	ListError       error
	GetError        error
}

func (s *Mock) Put(ctx context.Context, key string, value []byte, opts *store.WriteOptions) error {
	s.KVPairs = append(s.KVPairs, &store.KVPair{Key: key, Value: value, LastIndex: 0})

	return nil
}

func (s *Mock) Get(ctx context.Context, key string, options *store.ReadOptions) (*store.KVPair, error) {
	if s.Error {
		return nil, errors.New("error")
	}

	if s.GetError != nil {
		return nil, s.GetError
	}

	for _, kvPair := range s.KVPairs {
		if kvPair.Key == key {
			return kvPair, nil
		}
	}
	return nil, nil
}

func (s *Mock) Delete(ctx context.Context, key string) error {
	return errors.New("delete not supported")
}

// Exists mock
func (s *Mock) Exists(ctx context.Context, key string, options *store.ReadOptions) (bool, error) {
	return false, errors.New("exists not supported")
}

// Watch mock
func (s *Mock) Watch(ctx context.Context, key string, options *store.ReadOptions) (<-chan *store.KVPair, error) {
	return nil, errors.New("watch not supported")
}

// WatchTree mock
func (s *Mock) WatchTree(ctx context.Context, prefix string, options *store.ReadOptions) (<-chan []*store.KVPair, error) {
	return s.WatchTreeMethod(), nil
}

// NewLock mock
func (s *Mock) NewLock(ctx context.Context, key string, options *store.LockOptions) (store.Locker, error) {
	return nil, errors.New("NewLock not supported")
}

// List mock
func (s *Mock) List(ctx context.Context, prefix string, options *store.ReadOptions) ([]*store.KVPair, error) {
	if s.Error {
		return nil, errors.New("error")
	}

	var kv []*store.KVPair
	for _, kvPair := range s.KVPairs {
		if s.ListError != nil {
			return nil, s.ListError
		}
		if strings.HasPrefix(kvPair.Key, prefix) && kvPair.Key != prefix {
			kv = append(kv, kvPair)
		}
	}

	return kv, nil
}

// DeleteTree mock
func (s *Mock) DeleteTree(ctx context.Context, prefix string) error {
	return errors.New("DeleteTree not supported")
}

// AtomicPut mock
func (s *Mock) AtomicPut(ctx context.Context, key string, value []byte, previous *store.KVPair, opts *store.WriteOptions) (bool, *store.KVPair, error) {
	return false, nil, errors.New("AtomicPut not supported")
}

// AtomicDelete mock
func (s *Mock) AtomicDelete(ctx context.Context, key string, previous *store.KVPair) (bool, error) {
	return false, errors.New("AtomicDelete not supported")
}

// Close mock
func (s *Mock) Close() error { return nil }
