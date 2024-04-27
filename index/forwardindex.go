package index

import "sync"

type ForwardIndex struct {
	mu       sync.Mutex
	adIDToAd map[uint32]*Ad
}

func NewForwardIndex() *ForwardIndex {

	return &ForwardIndex{adIDToAd: map[uint32]*Ad{}}
}

func (idx *ForwardIndex) Get(adID uint32) *Ad {
	return idx.adIDToAd[adID]
}

func (idx *ForwardIndex) Set(ad *Ad) {
	idx.mu.Lock()
	idx.adIDToAd[ad.AdID] = ad
	idx.mu.Unlock()
}
