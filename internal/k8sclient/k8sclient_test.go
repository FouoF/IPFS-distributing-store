package k8sclient

import (
	"testing"

	v1 "ipfs-store/internal/k8sclient/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestClusterScaleUP(t *testing.T) {
	// 设置
	scheme := runtime.NewScheme()
	_ = v1.AddToScheme(scheme)

	tests := []struct {
		name        string
		ipfsCluster *v1.IpfsCluster
		wantErr     bool
	}{
		{
			name: "ScaleUp_Success",
			ipfsCluster: &v1.IpfsCluster{
				Spec: v1.IpfsClusterSpec{
					Replicas: 1,
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-cluster",
					Namespace: "default",
				},
			},
			wantErr: false,
		},
		{
			name: "ScaleUp_GetError",
			ipfsCluster: &v1.IpfsCluster{
				Spec: v1.IpfsClusterSpec{
					Replicas: 1,
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-cluster",
					Namespace: "default",
				},
			},
			wantErr: true,
		},
		{
			name: "ScaleUp_UpdateError",
			ipfsCluster: &v1.IpfsCluster{
				Spec: v1.IpfsClusterSpec{
					Replicas: 1,
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-cluster",
					Namespace: "default",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 设置模拟客户端
			client := fake.NewClientBuilder().WithScheme(scheme).WithObjects(tt.ipfsCluster).Build()

			if tt.name == "ScaleUp_GetError" || tt.name == "ScaleUp_UpdateError" {
				client = fake.NewClientBuilder().WithScheme(scheme).Build() // 模拟获取错误
			}

			k8sClient := &K8sClient{client: client}

			err := k8sClient.ClusterScaleUP("test-cluster", 1)

			if tt.wantErr && err == nil {
				t.Errorf("ClusterScaleUP() expected error, got nil")
			} else if !tt.wantErr && err != nil {
				t.Errorf("ClusterScaleUP() unexpected error: %v", err)
			}
		})
	}
}

func TestClusterScaleDown(t *testing.T) {
	// 设置
	scheme := runtime.NewScheme()
	_ = v1.AddToScheme(scheme)

	tests := []struct {
		name        string
		ipfsCluster *v1.IpfsCluster
		wantErr     bool
	}{
		{
			name: "ScaleDown_Success",
			ipfsCluster: &v1.IpfsCluster{
				Spec: v1.IpfsClusterSpec{
					Replicas: 2,
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-cluster",
					Namespace: "default",
				},
			},
			wantErr: false,
		},
		{
			name: "ScaleDown_GetError",
			ipfsCluster: &v1.IpfsCluster{
				Spec: v1.IpfsClusterSpec{
					Replicas: 2,
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-cluster",
					Namespace: "default",
				},
			},
			wantErr: true,
		},
		{
			name: "ScaleDown_UpdateError",
			ipfsCluster: &v1.IpfsCluster{
				Spec: v1.IpfsClusterSpec{
					Replicas: 2,
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-cluster",
					Namespace: "default",
				},
			},
			wantErr: true,
		},
		{
			name: "ScaleDown_MoreThanCurrentReplicas",
			ipfsCluster: &v1.IpfsCluster{
				Spec: v1.IpfsClusterSpec{
					Replicas: 1,
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-cluster",
					Namespace: "default",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 设置模拟客户端
			client := fake.NewClientBuilder().WithScheme(scheme).WithObjects(tt.ipfsCluster).Build()

			if tt.name == "ScaleDown_GetError" || tt.name == "ScaleDown_UpdateError" {
				client = fake.NewClientBuilder().WithScheme(scheme).Build() // 模拟获取错误
			}

			k8sClient := &K8sClient{client: client}

			err := k8sClient.ClusterScaleDown("test-cluster", 1)

			if tt.wantErr && err == nil {
				t.Errorf("ClusterScaleDown() expected error, got nil")
			} else if !tt.wantErr && err != nil {
				t.Errorf("ClusterScaleDown() unexpected error: %v", err)
			}
		})
	}
}
