package client


type MemoryStorage struct {
	data map[string]string
}

func (storage *MemoryStorage) GetItem(key string) string {
	return storage.data[key]
}

func (storage *MemoryStorage) SetItem(key, value string) {
	storage.data[key] = value
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[string]string),
	}
}