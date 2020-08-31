package main

import "fmt"

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776             (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424    (exceeds 1 << 64)
	YiB // 1208925819614629174706176
)

const (
	KB = 1000
	MB = KB * KB
	GB = MB * KB // 1073741824
	TB = GB * KB // 1099511627776             (exceeds 1 << 32)
	PB = TB * KB // 1125899906842624
	EB = PB * KB // 1152921504606846976
	ZB = EB * KB // 1180591620717411303424    (exceeds 1 << 64)
	YB = ZB * KB // 1208925819614629174706176
)

func main() {
	fmt.Println(YB)
}
