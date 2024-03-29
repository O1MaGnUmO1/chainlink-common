package loop

import (
	"context"
	"fmt"

	"github.com/O1MaGnUmO1/chainlink-common/pkg/services"
	"github.com/O1MaGnUmO1/chainlink-common/pkg/types"
)

// RelayerExt is a subset of [loop.Relayer] for adapting [types.Relayer], typically with a Chain. See [RelayerAdapter].
type RelayerExt interface {
	types.ChainService
	ID() string
}

var _ Relayer = (*RelayerAdapter)(nil)

// RelayerAdapter adapts a [types.Relayer] and [RelayerExt] to implement [Relayer].
type RelayerAdapter struct {
	types.Relayer
	RelayerExt
}

func (r *RelayerAdapter) NewPluginProvider(ctx context.Context, rargs types.RelayArgs, pargs types.PluginArgs) (types.PluginProvider, error) {
	return nil, fmt.Errorf("unexpected call to NewPluginProvider: did you forget to wrap RelayerAdapter in a relayerServerAdapter?")
}

func (r *RelayerAdapter) Start(ctx context.Context) error {
	var ms services.MultiStart
	return ms.Start(ctx, r.RelayerExt, r.Relayer)
}

func (r *RelayerAdapter) Close() error {
	return services.CloseAll(r.Relayer, r.RelayerExt)
}

func (r *RelayerAdapter) Name() string {
	return r.Relayer.Name()
}

func (r *RelayerAdapter) Ready() (err error) {
	return r.Relayer.Ready()
}

func (r *RelayerAdapter) HealthReport() map[string]error {
	hr := make(map[string]error)
	services.CopyHealth(hr, r.Relayer.HealthReport())
	return hr
}
