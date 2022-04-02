package data

import(
	"os"
	"time"
)


func dictUnidictCoBytes() ([]byte, error) {
	return _dictUnidictCo, nil
}

func dictUnidictCo() (*asset, error) {
	bytes, err := dictUnidictCoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.co", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1609246449, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
