package datavolume

import (
	"fmt"
	"strings"

	kv1alpha3 "kubevirt.io/client-go/api/v1alpha3"

	"github.com/rancher/apiserver/pkg/apierror"
	"github.com/rancher/apiserver/pkg/types"
	cdiv1beta1 "github.com/rancher/harvester/pkg/generated/controllers/cdi.kubevirt.io/v1beta1"
	"github.com/rancher/harvester/pkg/ref"
	"github.com/rancher/harvester/pkg/util"
	"github.com/rancher/wrangler/pkg/schemas/validation"
)

type dvStore struct {
	types.Store
	dvCache cdiv1beta1.DataVolumeCache
}

func (s *dvStore) Create(request *types.APIRequest, schema *types.APISchema, data types.APIObject) (types.APIObject, error) {
	util.SetHTTPSourceDataVolume(data.Data())
	return s.Store.Create(request, request.Schema, data)
}

func (s *dvStore) Update(request *types.APIRequest, schema *types.APISchema, data types.APIObject, id string) (types.APIObject, error) {
	util.SetHTTPSourceDataVolume(data.Data())
	return s.Store.Update(request, request.Schema, data, id)
}

func (s *dvStore) Delete(request *types.APIRequest, schema *types.APISchema, id string) (types.APIObject, error) {
	if err := s.canDelete(request.Namespace, request.Name); err != nil {
		return types.APIObject{}, apierror.NewAPIError(validation.ServerError, err.Error())
	}
	return s.Store.Delete(request, request.Schema, id)
}

func (s *dvStore) canDelete(namespace, name string) error {
	dv, err := s.dvCache.Get(namespace, name)
	if err != nil {
		return fmt.Errorf("failed to get dv %s, %v", name, err)
	}

	annotationSchemaOwners, err := ref.GetSchemaOwnersFromAnnotation(dv)
	if err != nil {
		return fmt.Errorf("failed to get schema owners from annotation: %w", err)
	}

	attachedList := annotationSchemaOwners.List(kv1alpha3.VirtualMachineGroupVersionKind.GroupKind())
	if len(attachedList) != 0 {
		return fmt.Errorf("can not delete the volume %s which is currently attached by these VMs: %s", name, strings.Join(attachedList, ","))
	}

	if len(dv.OwnerReferences) == 0 {
		return nil
	}

	var ownerList []string
	for _, owner := range dv.OwnerReferences {
		if owner.Kind == kv1alpha3.VirtualMachineGroupVersionKind.Kind {
			ownerList = append(ownerList, owner.Name)
		}
	}

	if len(ownerList) != 0 {
		return fmt.Errorf("can not delete the volume %s which is currently owned by these VMs: %s", name, strings.Join(ownerList, ","))
	}

	return nil
}
