package pl005

import (
	"time"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v12/pkg/table"
)

const (
	name  = "pl005"
	short = "Execute plan pl005 automatically."
)

type Config struct {
	Logger micrologger.Logger
}

func New(config Config) (*cobra.Command, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	f := &flag{}

	r := &runner{
		flag:   f,
		logger: config.Logger,
	}

	c := &cobra.Command{
		Use:   name,
		Short: short,
		Long:  mustLong(),
		RunE:  r.Run,
	}

	f.Init(c)

	return c, nil
}

func mustLong() string {
	var d time.Duration
	for _, s := range Plan {
		d += s.Backoff.Wait()
	}

	s := "Test plan pl005 launches a basic Tenant Cluster with the previous major version\n"
	s += "and upgrades to the lastest version. Plan execution might take up to " + d.String() + ".\n\n"

	t := [][]string{{"ACTION", "RETRY", "WAIT"}}

	for _, s := range Plan {
		t = append(t, []string{s.Action.String(), s.Backoff.Interval().String(), s.Backoff.Wait().String()})
	}

	colourized := table.Colourize(t)
	s += table.Format(colourized)

	return s
}
