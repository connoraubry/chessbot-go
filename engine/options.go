package engine

import "fmt"

/*
Provides arguments for input to the engine

*/

type OptBitboardDirectory string
type OptFenFile string

type Options struct {
	BitboardDirectory string
	FenFilePath       string
}

func ParseOptions(opts ...interface{}) (Options, error) {
	res := Options{}

	for _, opt := range opts {
		fmt.Println(opt)
		switch v := opt.(type) {
		case OptBitboardDirectory:
			res.BitboardDirectory = string(v)
		case OptFenFile:
			res.FenFilePath = string(v)
		}
	}

	return res, nil
}
