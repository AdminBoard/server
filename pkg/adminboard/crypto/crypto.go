package crypto

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/speps/go-hashids"
)

func MD5(s string) string {
	hash := md5.New()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}

func EncodeIDWithSalt(id int, salt string) string {
	data := hashids.NewData()
	data.Salt = salt
	data.MinLength = 256
	hash, _ := hashids.NewWithData(data)
	newID, _ := hash.Encode([]int{id})
	return newID
}

func DecodeIDWithSalt(session string, salt string) (int, error) {
	data := hashids.NewData()
	data.Salt = salt
	data.MinLength = 256
	hash, _ := hashids.NewWithData(data)
	ids, e := hash.DecodeWithError(session)
	if e != nil {
		return 0, e
	}
	return ids[0], nil
}
