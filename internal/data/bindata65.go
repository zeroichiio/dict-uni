package data

import(
	"os"
	"time"
)


func dictUnidictCnBytes() ([]byte, error) {
	return _dictUnidictCn, nil
}

func dictUnidictCn() (*asset, error) {
	bytes, err := dictUnidictCnBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.cn", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1609246449, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
