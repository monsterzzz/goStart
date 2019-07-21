package mylib

type Huashidu float32
type Sheshidu float32

func S2h(s Sheshidu) Huashidu {
	var tmpS float32
	tmpS = float32(s)
	result := tmpS*1.8 + 32
	return Huashidu(result)
}
