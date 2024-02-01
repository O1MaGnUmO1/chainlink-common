package loop_test

import (
	"os/exec"
	"sync/atomic"
	"testing"
	"time"

	"github.com/O1MaGnUmO1/chainlink-common/pkg/logger"
	"github.com/O1MaGnUmO1/chainlink-common/pkg/loop"
	"github.com/O1MaGnUmO1/chainlink-common/pkg/loop/internal"
	"github.com/O1MaGnUmO1/chainlink-common/pkg/loop/internal/test"
	"github.com/O1MaGnUmO1/chainlink-common/pkg/services/servicetest"
)

func TestRelayerService(t *testing.T) {
	t.Parallel()
	relayer := loop.NewRelayerService(logger.Test(t), loop.GRPCOpts{}, func() *exec.Cmd {
		return NewHelperProcessCommand(loop.PluginRelayerName)
	}, test.ConfigTOML, test.StaticKeystore{})
	hook := relayer.XXXTestHook()
	servicetest.Run(t, relayer)

	t.Run("control", func(t *testing.T) {
		test.RunRelayer(t, relayer)
	})

	t.Run("Kill", func(t *testing.T) {
		hook.Kill()

		// wait for relaunch
		time.Sleep(2 * internal.KeepAliveTickDuration)

		test.RunRelayer(t, relayer)
	})

	t.Run("Reset", func(t *testing.T) {
		hook.Reset()

		// wait for relaunch
		time.Sleep(2 * internal.KeepAliveTickDuration)

		test.RunRelayer(t, relayer)
	})
}

func TestRelayerService_recovery(t *testing.T) {
	t.Parallel()
	var limit atomic.Int32
	relayer := loop.NewRelayerService(logger.Test(t), loop.GRPCOpts{}, func() *exec.Cmd {
		h := HelperProcessCommand{
			Command: loop.PluginRelayerName,
			Limit:   int(limit.Add(1)),
		}
		return h.New()
	}, test.ConfigTOML, test.StaticKeystore{})
	servicetest.Run(t, relayer)

	test.RunRelayer(t, relayer)
}
