package clients

import (
	"context"

	"github.com/rancher/wrangler/pkg/clients"
	"github.com/rancher/wrangler/pkg/schemes"
	v1 "k8s.io/api/admissionregistration/v1"
	"k8s.io/client-go/rest"

	ctlcdiv1 "github.com/harvester/harvester/pkg/generated/controllers/cdi.kubevirt.io"
	ctlharvesterv1 "github.com/harvester/harvester/pkg/generated/controllers/harvesterhci.io"
	ctlcniv1 "github.com/harvester/harvester/pkg/generated/controllers/k8s.cni.cncf.io"
	ctlkubevirtv1 "github.com/harvester/harvester/pkg/generated/controllers/kubevirt.io"
)

type Clients struct {
	clients.Clients

	HarvesterFactory *ctlharvesterv1.Factory
	KubevirtFactory  *ctlkubevirtv1.Factory
	CNIFactory       *ctlcniv1.Factory

	CDIFactory *ctlcdiv1.Factory
}

func New(ctx context.Context, rest *rest.Config, threadiness int) (*Clients, error) {
	clients, err := clients.NewFromConfig(rest, nil)
	if err != nil {
		return nil, err
	}

	if err := schemes.Register(v1.AddToScheme); err != nil {
		return nil, err
	}

	harvesterFactory, err := ctlharvesterv1.NewFactoryFromConfigWithOptions(rest, clients.FactoryOptions)
	if err != nil {
		return nil, err
	}

	if err = harvesterFactory.Start(ctx, threadiness); err != nil {
		return nil, err
	}

	kubevirtFactory, err := ctlkubevirtv1.NewFactoryFromConfigWithOptions(rest, clients.FactoryOptions)
	if err != nil {
		return nil, err
	}

	if err = kubevirtFactory.Start(ctx, threadiness); err != nil {
		return nil, err
	}

	cdiFactory, err := ctlcdiv1.NewFactoryFromConfigWithOptions(rest, clients.FactoryOptions)
	if err != nil {
		return nil, err
	}

	if err = cdiFactory.Start(ctx, threadiness); err != nil {
		return nil, err
	}

	cniFactory, err := ctlcniv1.NewFactoryFromConfigWithOptions(rest, clients.FactoryOptions)
	if err != nil {
		return nil, err
	}

	if err = cniFactory.Start(ctx, threadiness); err != nil {
		return nil, err
	}

	return &Clients{
		Clients:          *clients,
		HarvesterFactory: harvesterFactory,
		KubevirtFactory:  kubevirtFactory,
		CNIFactory:       cniFactory,
		CDIFactory:       cdiFactory,
	}, nil
}
