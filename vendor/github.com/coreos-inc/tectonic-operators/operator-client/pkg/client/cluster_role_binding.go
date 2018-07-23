package client

import (
	"fmt"

	"github.com/golang/glog"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// CreateClusterRoleBinding creates the clusterRoleBinding.
func (c *Client) CreateClusterRoleBinding(ig *rbacv1.ClusterRoleBinding) (*rbacv1.ClusterRoleBinding, error) {
	return c.RbacV1().ClusterRoleBindings().Create(ig)
}

// GetClusterRoleBinding returns the existing clusterRoleBinding.
func (c *Client) GetClusterRoleBinding(name string) (*rbacv1.ClusterRoleBinding, error) {
	return c.RbacV1().ClusterRoleBindings().Get(name, metav1.GetOptions{})
}

// DeleteClusterRoleBinding deletes the clusterRoleBinding.
func (c *Client) DeleteClusterRoleBinding(name string, options *metav1.DeleteOptions) error {
	return c.RbacV1().ClusterRoleBindings().Delete(name, options)
}

// UpdateClusterRoleBinding will update the given ClusterRoleBinding resource.
func (c *Client) UpdateClusterRoleBinding(crb *rbacv1.ClusterRoleBinding) (*rbacv1.ClusterRoleBinding, error) {
	glog.V(4).Infof("[UPDATE ClusterRoleBinding]: %s", crb.GetName())
	oldCrb, err := c.GetClusterRoleBinding(crb.GetName())
	if err != nil {
		return nil, err
	}
	patchBytes, err := createPatch(oldCrb, crb)
	if err != nil {
		return nil, fmt.Errorf("error creating patch for ClusterRoleBinding: %v", err)
	}
	return c.RbacV1().ClusterRoleBindings().Patch(crb.GetName(), types.StrategicMergePatchType, patchBytes)
}
