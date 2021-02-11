package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"

	cniv1 "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1"
	controllergen "github.com/rancher/wrangler/pkg/controller-gen"
	"github.com/rancher/wrangler/pkg/controller-gen/args"
	"github.com/sirupsen/logrus"
	kv1 "kubevirt.io/client-go/api/v1"
	cdiv1 "kubevirt.io/containerized-data-importer/pkg/apis/core/v1beta1"

	harv1 "github.com/rancher/harvester/pkg/apis/harvester.cattle.io/v1alpha1"
)

func main() {
	os.Unsetenv("GOPATH")
	controllergen.Run(args.Options{
		OutputPackage: "github.com/rancher/harvester/pkg/generated",
		Boilerplate:   "scripts/boilerplate.go.txt",
		Groups: map[string]args.Group{
			"harvester.cattle.io": {
				Types: []interface{}{
					harv1.VirtualMachineImage{},
					harv1.Setting{},
					harv1.KeyPair{},
					harv1.VirtualMachineTemplate{},
					harv1.VirtualMachineTemplateVersion{},
					harv1.User{},
					harv1.Preference{},
				},
				GenerateTypes:   true,
				GenerateClients: true,
			},
			kv1.GroupName: {
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
