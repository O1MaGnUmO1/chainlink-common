package values

import (
	"github.com/O1MaGnUmO1/chainlink-common/pkg/values/pb"
)

type Int64 struct {
	Underlying int64
}

func NewInt64(i int64) (*Int64, error) {
	return &Int64{Underlying: i}, nil
}

func (i *Int64) Proto() (*pb.Value, error) {
	return pb.NewInt64Value(i.Underlying)
}

func (i *Int64) Unwrap() (any, error) {
	return i.Underlying, nil
}
