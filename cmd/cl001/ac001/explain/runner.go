package explain

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/pkg/action"
	"github.com/giantswarm/awscnfm/pkg/action/ac001"
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

	var e action.Explainer
	{
		c := ac001.ExplainerConfig{}

		e, err = ac001.NewExplainer(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	fmt.Println(strings.TrimSpace(e.Explain()))
	fmt.Println()

	return nil
}