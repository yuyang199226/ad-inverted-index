package index

type Indexer interface {
}

type IndexService struct {
	forwardIndex  *ForwardIndex
	invertedIndex *InvertedIndex
}

func NewIndexService() *IndexService {
	return nil
}


func (isvc *IndexService) Load() {
	
}


