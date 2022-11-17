package engine

/*
Provides arguments for input to the engine

*/

type OptFenFile string
type OptFenString string

type Options struct {
	FenFilePath string
	FenString   string
}

func ParseOptions(opts ...interface{}) (Options, error) {
	res := Options{}

	for _, opt := range opts {
		switch v := opt.(type) {
		case OptFenFile:
			res.FenFilePath = string(v)
		case OptFenString:
			res.FenString = string(v)
		}
	}

	return res, nil
}
