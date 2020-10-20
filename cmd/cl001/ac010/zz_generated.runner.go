package ac010

import (
	"context"
	"io"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v12/pkg/action"
	"github.com/giantswarm/awscnfm/v12/pkg/action/cl001/ac010"
	"github.com/giantswarm/awscnfm/v12/pkg/config"
	"github.com/giantswarm/awscnfm/v12/pkg/env"
)

type runner struct {
	flag   *flag
	logger micrologger.Logger
	stdout io.Writer
	stderr io.Writer
}

func (r *runner) Run(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	err := r.flag.Validate()
	if err != nil {
		return microerror.Mask(err)
	}

	err = r.run(ctx, cmd, args)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

func (r *runner) run(ctx context.Context, cmd *cobra.Command, args []string) error {
	var err error

	var e action.Executor
	{
		c := ac010.ExecutorConfig{
			Command: cmd,
			Logger:  r.logger,

			Scope:         "cl001",
			TenantCluster: config.Cluster("cl001", env.TenantCluster()),
		}

		e, err = ac010.NewExecutor(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	err = e.Execute(ctx)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
