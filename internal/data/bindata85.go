package data

import(
	"os"
	"time"
)


func dictUnidictDhBytes() ([]byte, error) {
	return _dictUnidictDh, nil
}

func dictUnidictDh() (*asset, error) {
	bytes, err := dictUnidictDhBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.dh", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1609246449, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
