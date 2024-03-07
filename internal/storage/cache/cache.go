package cache

import (
	"fmt"
	"sync"
	"woonbeaj/L0/internal/jsonStruct"
)

type Loader interface {
	GetAll() (*jsonStruct.AllRows, error)
}

type Cache struct {
	Mut 	sync.RWMutex
	CacheMap map[string][]byte
}

func MustLoad(srg Loader) (*Cache, error) {
	const op = "cache.cache.MustLoad"
	var cch Cache
	cch.Mut = sync.RWMutex{}
	cch.CacheMap = make(map[string][]byte)
	data, err := srg.GetAll()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	cch.Mut.Lock()
	for i, j := range data.OrdersUIDs {
		cch.CacheMap[j] = data.OrderINFOs[i]
	}
	cch.Mut.Unlock()

	return &cch, nil
}