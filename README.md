# shorten

Two ways short url algorithm, one is base62, another is MD5.

# Usage

```golang
var url = "http://www.xxx.com/xxxxxxxxxxxx" // long URL

m := shorten.MD5(url) // MD5() return four []byte, you can pick up any one as short url suffix.
m0 := shorten.MD50(url) // equals with m[0]
m1 := shorten.MD51(url) // equals with m[1]
m2 := shorten.MD52(url) // equals with m[2]
m3 := shorten.MD52(url) // equals with m[3]

var autoIncrementId uint64 = 123456789
shortURL := shorten.Base62Encode(autoIncrementId)
id, err := shorten.Base62Decode(shortURL)
```