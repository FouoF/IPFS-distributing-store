package k8sclient

import (
	"context"
	"fmt"

	v1 "ipfs-store/internal/k8sclient/v1alpha1"

	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

var (
	// operatorScheme is the scheme used by the operator
	operatorScheme = runtime.NewScheme()
)

type K8sClient struct {
	client client.Client
}

func init() {
	// AddToScheme adds the types of this group-version to the given scheme.
	_ = v1.AddToScheme(operatorScheme)
}

func NewK8sClient() (*K8sClient, error) {
	// Create a client using the same scheme as the operator
	cfg, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	k8sClient, err := client.New(cfg, client.Options{Scheme: operatorScheme})
	if err != nil {
		return nil, err
	}

	return &K8sClient{client: k8sClient}, nil
}

func (c *K8sClient) ClusterScaleUP(name string, num int) error {
	// Fetch and update CRD
	instance := &v1.IpfsCluster{}
	err := c.client.Get(context.TODO(), client.ObjectKey{Namespace: "default", Name: name}, instance)
	if err != nil {
		return err
	}

	instance.Spec.Replicas = instance.Spec.Replicas + int32(num)

	// Update the CRD instance in the cluster
	err = c.client.Update(context.TODO(), instance)
	if err != nil {
		return err
	}

	fmt.Println("CRD instance updated successfully")
	return nil
}

func (c *K8sClient) ClusterScaleDown(name string, num int) error {
	// Fetch and update CRD
	instance := &v1.IpfsCluster{}
	err := c.client.Get(context.TODO(), client.ObjectKey{Namespace: "default", Name: name}, instance)
	if err != nil {
		return err
	}

	if num > int(instance.Spec.Replicas) {
		return fmt.Errorf("cannot scale down more than the current number of replicas")
	}
	instance.Spec.Replicas = instance.Spec.Replicas - int32(num)

	// Update the CRD instance in the cluster
	err = c.client.Update(context.TODO(), instance)
	if err != nil {
		return err
	}

	fmt.Println("CRD instance updated successfully")
	return nil
}
