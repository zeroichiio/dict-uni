package data

import(
	"os"
	"time"
)


func dictUnidictCqBytes() ([]byte, error) {
	return _dictUnidictCq, nil
}

func dictUnidictCq() (*asset, error) {
	bytes, err := dictUnidictCqBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.cq", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1609246449, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
