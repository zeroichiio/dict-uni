package data

import(
	"os"
	"time"
)


func dictUnidictBhBytes() ([]byte, error) {
	return _dictUnidictBh, nil
}

func dictUnidictBh() (*asset, error) {
	bytes, err := dictUnidictBhBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.bh", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1609246449, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
