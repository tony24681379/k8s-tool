package controller

import (
	"sync"

	"github.com/golang/glog"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
)

type DaemonSetController struct {
	sync.Mutex
	Pod       v1.PodInterface
	DaemonSet v1beta1.DaemonSetInterface
}

// RestartOnePod scale up the deployment and scale down the deployment
func (d *DaemonSetController) RestartOnePod(resourceName, podName string) error {
	if err := d.deletePod(podName); err != nil {
		glog.Error(err)
		return err
	}
	return nil
}

func (d *DaemonSetController) deletePod(podName string) error {
	if err := d.Pod.Delete(podName, &metav1.DeleteOptions{}); err != nil {
		glog.Error(err)
		return err
	}
	return nil
}
