package main

import (
	"github.com/gogf/gf/v2/crypto/gaes"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gres"
)

var (
	CryptoKey = []byte("x76cgqt36i9c863bzmotuf8626dxiwu0")
)

func main() {
	binContent := gfile.GetBytes("data.bin")
	binContent, err := gaes.Decrypt(binContent, CryptoKey)
	if err != nil {
		panic(err)
	}
	if err := gres.Add(binContent); err != nil {
		panic(err)
	}
	gres.Dump()
}
