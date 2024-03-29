package reportingplugins_test

import (
	"os/exec"
	"sync/atomic"
	"testing"
	"time"

	"github.com/O1MaGnUmO1/chainlink-common/pkg/logger"
	"github.com/O1MaGnUmO1/chainlink-common/pkg/loop"
	"github.com/O1MaGnUmO1/chainlink-common/pkg/loop/internal"
	"github.com/O1MaGnUmO1/chainlink-common/pkg/loop/internal/test"
	"github.com/O1MaGnUmO1/chainlink-common/pkg/loop/reportingplugins"
	"github.com/O1MaGnUmO1/chainlink-common/pkg/services/servicetest"
	"github.com/O1MaGnUmO1/chainlink-common/pkg/types"
)

type HelperProcessCommand test.HelperProcessCommand

func (h *HelperProcessCommand) New() *exec.Cmd {
	h.CommandLocation = "../internal/test/cmd/main.go"
	return (test.HelperProcessCommand)(*h).New()
}

func NewHelperProcessCommand(command string) *exec.Cmd {
	h := HelperProcessCommand{
		Command: command,
	}
	return h.New()
}

func TestLOOPPService(t *testing.T) {
	t.Parallel()

	tests := []struct {
		Plugin string
	}{
		// A generic plugin with a median provider
		{Plugin: test.ReportingPluginWithMedianProviderName},
		// A generic plugin with a plugin provider
		{Plugin: reportingplugins.PluginServiceName},
	}
	for _, ts := range tests {
		looppSvc := reportingplugins.NewLOOPPService(logger.Test(t), loop.GRPCOpts{}, func() *exec.Cmd {
			return NewHelperProcessCommand(ts.Plugin)
		}, types.ReportingPluginServiceConfig{}, test.MockConn{}, &test.StaticPipelineRunnerService{}, &test.StaticTelemetry{}, &test.StaticErrorLog{})
		hook := looppSvc.XXXTestHook()
		servicetest.Run(t, looppSvc)

		t.Run("control", func(t *testing.T) {
			test.ReportingPluginFactory(t, looppSvc)
		})

		t.Run("Kill", func(t *testing.T) {
			hook.Kill()

			// wait for relaunch
			time.Sleep(2 * internal.KeepAliveTickDuration)

			test.ReportingPluginFactory(t, looppSvc)
		})

		t.Run("Reset", func(t *testing.T) {
			hook.Reset()

			// wait for relaunch
			time.Sleep(2 * internal.KeepAliveTickDuration)

			test.ReportingPluginFactory(t, looppSvc)
		})
	}
}

func TestLOOPPService_recovery(t *testing.T) {
	t.Parallel()
	var limit atomic.Int32
	looppSvc := reportingplugins.NewLOOPPService(logger.Test(t), loop.GRPCOpts{}, func() *exec.Cmd {
		h := HelperProcessCommand{
			Command: test.ReportingPluginWithMedianProviderName,
			Limit:   int(limit.Add(1)),
		}
		return h.New()
	}, types.ReportingPluginServiceConfig{}, test.MockConn{}, &test.StaticPipelineRunnerService{}, &test.StaticTelemetry{}, &test.StaticErrorLog{})
	servicetest.Run(t, looppSvc)

	test.ReportingPluginFactory(t, looppSvc)
}
