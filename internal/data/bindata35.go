package data

import(
	"os"
	"time"
)


func dictUnidictBjBytes() ([]byte, error) {
	return _dictUnidictBj, nil
}

func dictUnidictBj() (*asset, error) {
	bytes, err := dictUnidictBjBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.bj", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1609246449, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
