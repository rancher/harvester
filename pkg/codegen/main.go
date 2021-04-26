package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"

	cniv1 "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1"
	storagev1beta1 "github.com/kubernetes-csi/external-snapshotter/v2/pkg/apis/volumesnapshot/v1beta1"
	longhornv1 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta1"
	upgradev1 "github.com/rancher/system-upgrade-controller/pkg/apis/upgrade.cattle.io/v1"
	controllergen "github.com/rancher/wrangler/pkg/controller-gen"
	"github.com/rancher/wrangler/pkg/controller-gen/args"
	"github.com/sirupsen/logrus"
	kv1 "kubevirt.io/client-go/api/v1"
	cdiv1 "kubevirt.io/containerized-data-importer/pkg/apis/core/v1beta1"

	harvesterv1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
)

func main() {
	os.Unsetenv("GOPATH")
	controllergen.Run(args.Options{
		OutputPackage: "github.com/harvester/harvester/pkg/generated",
		Boilerplate:   "scripts/boilerplate.go.txt",
		Groups: map[string]args.Group{
			"harvesterhci.io": {
				Types: []interface{}{
					harvesterv1.KeyPair{},
					harvesterv1.Preference{},
					harvesterv1.Setting{},
					harvesterv1.Upgrade{},
					harvesterv1.User{},
					harvesterv1.VirtualMachineBackup{},
					harvesterv1.VirtualMachineBackupContent{},
					harvesterv1.VirtualMachineRestore{},
					harvesterv1.VirtualMachineImage{},
					harvesterv1.VirtualMachineTemplate{},
					harvesterv1.VirtualMachineTemplateVersion{},
					harvesterv1.SupportBundle{},
				},
				GenerateTypes:   true,
				GenerateClients: true,
			},
			kv1.SchemeGroupVersion.Group: {
				Types: []interface{}{
					kv1.VirtualMachine{},
					kv1.VirtualMachineInstance{},
					kv1.VirtualMachineInstanceMigration{},
				},
				GenerateTypes:   false,
				GenerateClients: true,
			},
			cdiv1.SchemeGroupVersion.Group: {
				Types: []interface{}{
					cdiv1.DataVolume{},
				},
				GenerateTypes:   false,
				GenerateClients: true,
			},
			cniv1.SchemeGroupVersion.Group: {
				Types: []interface{}{
					cniv1.NetworkAttachmentDefinition{},
				},
				GenerateTypes:   false,
				GenerateClients: true,
			},
			storagev1beta1.SchemeGroupVersion.Group: {
				Types: []interface{}{
					storagev1beta1.VolumeSnapshotClass{},
					storagev1beta1.VolumeSnapshot{},
					storagev1beta1.VolumeSnapshotContent{},
				},
				GenerateTypes:   false,
				GenerateClients: true,
			},
			longhornv1.SchemeGroupVersion.Group: {
				Types: []interface{}{
					longhornv1.Volume{},
					longhornv1.Setting{},
				},
			},
			upgradev1.SchemeGroupVersion.Group: {
				Types: []interface{}{
					upgradev1.Plan{},
				},
				GenerateTypes:   false,
				GenerateClients: true,
			},
		},
	})
	nadControllerInterfaceRefactor()
}

// NB(GC), nadControllerInterfaceRefactor modify the generated resource name of NetworkAttachmentDefinition controller using a dash-separator,
// the original code is generated by https://github.com/rancher/wrangler/blob/e86bc912dfacbc81dc2d70171e4d103248162da6/pkg/controller-gen/generators/group_version_interface_go.go#L82-L97
// since the NAD crd uses a varietal plurals name(i.e network-attachment-definitions), and the default resource name generated by wrangler is
// `networkattachementdefinitions` that will raises crd not found exception of the NAD controller.
func nadControllerInterfaceRefactor() {
	absPath, _ := filepath.Abs("pkg/generated/controllers/k8s.cni.cncf.io/v1/interface.go")
	input, err := ioutil.ReadFile(absPath)
	if err != nil {
		logrus.Fatalf("failed to read the network-attachment-definition file: %v", err)
	}

	output := bytes.Replace(input, []byte("networkattachmentdefinitions"), []byte("network-attachment-definitions"), -1)

	if err = ioutil.WriteFile(absPath, output, 0644); err != nil {
		logrus.Fatalf("failed to update the network-attachment-definition file: %v", err)
	}
}
