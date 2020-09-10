package vmtemplate

import (
	"github.com/rancher/apiserver/pkg/types"
	"github.com/rancher/harvester/pkg/api/store"
	"github.com/rancher/harvester/pkg/config"
	"github.com/rancher/steve/pkg/schema"
	"github.com/rancher/steve/pkg/server"
	"github.com/rancher/steve/pkg/stores/proxy"
)

const (
	templateSchemaID        = "harvester.cattle.io.virtualmachinetemplate"
	templateVersionSchemaID = "harvester.cattle.io.virtualmachinetemplateversion"
)

func RegisterSchema(scaled *config.Scaled, server *server.Server) {
	templates := scaled.HarvesterFactory.Harvester().V1alpha1().VirtualMachineTemplate()
	templateVersionCache := scaled.HarvesterFactory.Harvester().V1alpha1().VirtualMachineTemplateVersion().Cache()
	th := &templateLinkHandler{
		templateVersionCache: templateVersionCache,
	}

	templateVersionStore := &templateVersionStore{
		Store:                proxy.NewProxyStore(server.ClientFactory, server.AccessSetLookup),
		templateCache:        templates.Cache(),
		templateVersionCache: templateVersionCache,
		keyPairCache:         scaled.HarvesterFactory.Harvester().V1alpha1().KeyPair().Cache(),
	}

	t := []schema.Template{
		{
			ID:        templateSchemaID,
			Formatter: formatter,
			Customize: func(apiSchema *types.APISchema) {
				apiSchema.ByIDHandler = th.byIDHandler
			},
			Store: store.NamespaceStore{Store: proxy.NewProxyStore(server.ClientFactory, server.AccessSetLookup)},
		},
		{
			ID:        templateVersionSchemaID,
			Formatter: versionFormatter,
			Store:     store.NamespaceStore{Store: templateVersionStore},
		},
	}

	server.SchemaTemplates = append(server.SchemaTemplates, t...)
}
