package scope

import (
	"context"
	"github.com/kyma-project/cloud-resources/components/kcp/pkg/common/actions/focal"
	"github.com/kyma-project/cloud-resources/components/lib/composed"
)

func New(stateFactory StateFactory) composed.Action {
	return func(ctx context.Context, st composed.State) (error, context.Context) {
		state := stateFactory.CreateState(st.(focal.State))

		return composed.ComposeActions(
			"scopeMain",
			whenNoScopeRef,
			whenNoScopeCreated,
		)(ctx, state)
	}
}
