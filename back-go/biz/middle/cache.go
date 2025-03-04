package middle

import "github.com/coocood/freecache"

var Cache *freecache.Cache

func init() {
	Cache = freecache.NewCache(3 * 1024 * 1024 * 1024)
}

func SaveFile(key string, imgDate []byte) error {
	if err := Cache.Set([]byte(key), imgDate, 0); err != nil {
		return err
	}
	return nil
}

func ReadFile(key string) ([]byte, error) {
	imgDate, err := Cache.Get([]byte(key))
	if err != nil {
		return nil, err
	}
	return imgDate, nil
}
