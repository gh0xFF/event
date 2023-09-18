package utils

import (
	"io"
	"log"

	_ "embed"

	"github.com/klauspost/compress/zstd"
)

//go:embed dict.zdict
var dict []byte

func Compress(data []byte) ([]byte, error) {
	compressor, err := zstd.NewWriter(nil, zstd.WithEncoderDict(dict))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer compressor.Close()

	compressed := make([]byte, 0)
	compressor.Reset(io.Discard)

	return compressor.EncodeAll(data, compressed), nil
}

func Decompress(compressed []byte) ([]byte, error) {
	decompressor, err := zstd.NewReader(nil, zstd.WithDecoderDicts(dict))
	if err != nil {
		return nil, err
	}
	defer decompressor.Close()

	decompressed, err := decompressor.DecodeAll(compressed, nil)
	if err != nil {
		return nil, err
	}

	return decompressed, nil
}
