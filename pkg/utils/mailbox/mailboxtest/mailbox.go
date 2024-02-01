package mailboxtest

import (
	"testing"

	"github.com/O1MaGnUmO1/chainlink-common/pkg/logger"
	"github.com/O1MaGnUmO1/chainlink-common/pkg/utils/mailbox"
)

func NewMonitor(t testing.TB) *mailbox.Monitor {
	return mailbox.NewMonitor(t.Name(), logger.Named(logger.Test(t), "Mailbox"))
}
