package util

import (
  _ "embed"
  "encoding/base64"
  "github.com/develar/errors"
  "github.com/klauspost/compress/zstd"
)

//go:embed zstd.dictionary
var ZstdDictionary []byte

func DecodeQuery(encoded string) ([]byte, error) {
  compressed, err := base64.RawURLEncoding.DecodeString(encoded)
  if err != nil {
    return nil, errors.WithStack(err)
  }

  reader, err := zstd.NewReader(nil, zstd.WithDecoderConcurrency(0), zstd.WithDecoderDicts(ZstdDictionary))
  if err != nil {
    return nil, errors.WithStack(err)
  }
  defer reader.Close()

  decompressed, err := reader.DecodeAll(compressed, nil)
  if err != nil {
    return nil, errors.WithStack(err)
  }
  return decompressed, nil
}