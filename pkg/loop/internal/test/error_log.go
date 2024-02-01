package test

import (
	"context"
	"fmt"

	"github.com/O1MaGnUmO1/chainlink-common/pkg/types"
)

var _ types.ErrorLog = (*StaticErrorLog)(nil)

type StaticErrorLog struct{}

func (s *StaticErrorLog) SaveError(ctx context.Context, msg string) error {
	if msg != errMsg {
		return fmt.Errorf("expected %q but got %q", errMsg, msg)
	}
	return nil
}
