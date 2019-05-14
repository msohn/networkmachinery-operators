package app

import (
	"github.com/networkmachinery/networkmachinery-operators/pkg/controllers"
	"github.com/networkmachinery/networkmachinery-operators/pkg/controllers/networkcontrol/controller"
	"github.com/networkmachinery/networkmachinery-operators/pkg/utils"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// NetworkMonitorCmdOptions necessary options to run the sFlowController
// the current context on a user's KUBECONFIG
type NetworkControlCmdOptions struct {
	*genericclioptions.ConfigFlags
	controllers.LeaderElectionOptions
	controllers.ControllerOptions
}

func (nm *NetworkControlCmdOptions) ApplyLeaderElection(mgr *manager.Options) *manager.Options {
	mgr.LeaderElectionID = nm.LeaderElectionID
	mgr.LeaderElectionNamespace = nm.LeaderElectionNamespace
	mgr.LeaderElection = nm.LeaderElection
	return mgr
}

func (nm *NetworkControlCmdOptions) InitConfig() *rest.Config {
	config, err := nm.ToRESTConfig()
	if err != nil {
		utils.LogErrAndExit(err, "Error getting config")
	}
	config.UserAgent = controller.Name
	return config
}
