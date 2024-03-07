package memory

import (
	"fmt"
	"woonbeaj/L0/internal/storage/cache"
	"woonbeaj/L0/internal/storage/postgre"
)

type Cacher interface {
	Save(orderUid string, orderInfo []byte) error
	Get(orderUid string) (string, error)
}

type StorageWithCache struct {
	db 		*storage.Storage
	cache 	*cache.Cache
}

func NewWithCache(db *storage.Storage, cache *cache.Cache) *StorageWithCache {
	return &StorageWithCache{db: db, cache: cache}
}

func(s *StorageWithCache) Save(orderUid string, orderInfo []byte) error {
	const op = "storage.memory.Save" 

	s.cache.Mut.Lock()
	defer s.cache.Mut.Unlock()
	if _, ok := s.cache.CacheMap[orderUid]; ok {
		return fmt.Errorf("record already exists")
	} 
	
	err := s.db.Save(orderUid, orderInfo)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	s.cache.CacheMap[orderUid] = orderInfo
	

	return nil
}

func (s *StorageWithCache) Get(orderUid string) ([]byte, error) {
	const op = "storage.memory.Get"

	s.cache.Mut.Lock()
	defer s.cache.Mut.Unlock()

	if res, ok := s.cache.CacheMap[orderUid]; ok {
		return res, nil
	}

	res, err := s.db.Get(orderUid)
	if err != nil {
		return nil,  fmt.Errorf("%s: %w", op, err)
	}
	s.cache.CacheMap[orderUid] = res
	return res, nil
}