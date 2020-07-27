package qrcode

import (
	"fmt"
	"testing"
	"github.com/juliankoehn/mchurl/qrcode/bitset"
)

func TestBuildRegularSymbol(t *testing.T) {
	for k := 0; k <= 7; k++ {
		v := getQRCodeVersion(Low, 1)

		data := bitset.New()
		for i := 0; i < 26; i++ {
			data.AppendNumBools(8, false)
		}

		s, err := buildRegularSymbol(*v, k, data, false)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			_ = s
			//fmt.Print(m.string())
		}
	}
}