package data

import(
	"os"
	"time"
)


func dictUnidictAzBytes() ([]byte, error) {
	return _dictUnidictAz, nil
}

func dictUnidictAz() (*asset, error) {
	bytes, err := dictUnidictAzBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.az", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1609246449, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
