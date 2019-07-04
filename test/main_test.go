package test

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gofrs/uuid"
	"os"
	"testing"
)

func TestMd5(t *testing.T) {
	h := md5.New()
	h.Write([]byte("19911026"))
	str := hex.EncodeToString(h.Sum(nil))
	println(str)
}

func TestMd5_1(t *testing.T) {
	h := md5.New()
	h.Reset()
	bytes := []byte("8pgby")
	h.Write(bytes)
	arr := h.Sum(nil)
	println(hex.EncodeToString(arr))
}

func TestMd5_2(t *testing.T) {
	h := md5.New()

	h.Reset()
	bytes := []byte("8pgby")
	h.Write(bytes)
	saltArr := h.Sum(nil)

	h.Reset()
	//bytes := []byte("5069a865b421d1a0481bc6793f48c1a9")
	h.Write(saltArr)

	h.Write([]byte("111111"))
	arr := h.Sum(nil)
	//str := hex.EncodeToString(h.Sum(nil))

	for i := 0; i < 1023; i++ {
		h.Reset()
		h.Write(arr)
		arr = h.Sum(nil)
		//str = hex.EncodeToString(h.Sum(nil))
	}
	println(hex.EncodeToString(arr))
}

func TestOs(t *testing.T) {
	println(os.TempDir())
}

func TestUUID(t *testing.T) {
	uuids := uuid.Must(uuid.NewV4())
	println(uuids.String())
}
