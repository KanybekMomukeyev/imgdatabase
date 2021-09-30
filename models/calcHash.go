package models

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"math"
	"time"
)

//MakeTimestamp lalal
func MakeTimestamp() int64 {
	time.Sleep(time.Millisecond)
	return time.Now().UnixNano() / int64(time.Millisecond)
}

//HmacMD5 lalal
func HmacMD5(key, data string) string {
	hmac := hmac.New(md5.New, []byte(key))
	hmac.Write([]byte(data))
	return hex.EncodeToString(hmac.Sum([]byte("")))
}

//GetFloatSwitchOnly lalal
func GetFloatSwitchOnly(unk interface{}) (float64, error) {
	switch i := unk.(type) {
	case float64:
		return i, nil
	case float32:
		return float64(i), nil
	case int64:
		return float64(i), nil
	case int32:
		return float64(i), nil
	case int:
		return float64(i), nil
	case uint64:
		return float64(i), nil
	case uint32:
		return float64(i), nil
	case uint:
		return float64(i), nil
	default:
		return math.NaN(), errors.New("getFloat: unknown value is of incompatible type")
	}
}
