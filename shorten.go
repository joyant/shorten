package shorten

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strconv"
)

var base62 = [62]byte{
	'0', '1', '2', '3', '4', '5', '6', '7',
	'8', '9', 'A', 'B', 'C', 'D', 'E', 'F',
	'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N',
	'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V',
	'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd',
	'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l',
	'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
	'u', 'v', 'w', 'x', 'y', 'z',
}

var base32 = [32]byte {
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h',
	'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p',
	'q', 'r', 's', 't', 'u', 'v', 'w', 'x',
	'y', 'z', '0', '1', '2', '3', '4', '5',
}

func Base62Encode(n uint64) []byte {
	var b []byte
	for {
		i := n / 62
		j := n % 62
		b = append(b, base62[j])
		if i == 0 {
			break
		}
		n = i
	}
	for i := 0; len(b) > 1 && i < len(b)/2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}
	return b
}

func Base62Decode(encoded []byte) (uint64, error) {
	var sum uint64
	for _, e := range encoded {
		var num byte
		switch {
		case e >= '0' && e <= '9':
			num = e - '0'
		case e >= 'A' && e <= 'Z':
			num = e - 'A' + 10
		case e >= 'a' && e <= 'z':
			num = e - 'a' + 26 + 10
		default:
			return 0, errors.New("unknown encoded")
		}
		sum = sum * 62 + uint64(num)
	}
	return sum, nil
}

func MD5(rawURL string) [4][]byte {
	h := md5.New()
	h.Write([]byte(rawURL))
	m := hex.EncodeToString(h.Sum(nil))
	var bs [4][]byte
	for i := 0; i < 4; i++ {
		u, err := strconv.ParseUint(m[i*8:i*8+8], 16, 32)
		if err != nil {
			panic(err)
		}
		num30 := u & uint64(0x3FFFFFFF)
		var b [6]byte
		for j := 0; j < 6; j++ {
			b[j] = base32[num30 & 31]
			num30 >>= 5
		}
		bs[i] = b[:]
	}
	return bs
}

func _MD5ByIndex(rawURL string, index int) []byte {
	if index > 3 {
		return nil
	}
	h := md5.New()
	h.Write([]byte(rawURL))
	m := hex.EncodeToString(h.Sum(nil))
	u, err := strconv.ParseUint(m[index*8:index*8+8], 16, 32)
	if err != nil {
		panic(err)
	}
	num30 := u & uint64(0x3FFFFFFF)
	var b [6]byte
	for j := 0; j < 6; j++ {
		b[j] = base32[num30 & 31]
		num30 >>= 5
	}
	return b[:]
}

func MD50(rawURL string) []byte {
	return _MD5ByIndex(rawURL, 0)
}

func MD51(rawURL string) []byte {
	return _MD5ByIndex(rawURL, 1)
}

func MD52(rawURL string) []byte {
	return _MD5ByIndex(rawURL, 2)
}

func MD53(rawURL string) []byte {
	return _MD5ByIndex(rawURL, 3)
}