package qrcode

type Version int8

func (v Version) PatternSize() int {
	return 4*int(v) + 17
}
