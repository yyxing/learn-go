package main

import (
	"context"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func CreatePVC() {
	config, err := clientcmd.BuildConfigFromFlags("", "C:\\Users\\bodenai\\.kube\\config")
	if err != nil {
		panic(err)
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	storageClass := "csi-obs"
	pvc := &v1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				"everest.io/obs-volume-type": "STANDARD",
				"csi.storage.k8s.io/fstype":  "s3fs",
			},
			Name:      "test1-go-pvc",
			Namespace: "ouroboros",
		},
		Spec: v1.PersistentVolumeClaimSpec{
			AccessModes: []v1.PersistentVolumeAccessMode{"ReadWriteMany"},
			Resources: v1.ResourceRequirements{
				Requests: v1.ResourceList{v1.ResourceStorage: resource.MustParse("10Gi")},
			},
			StorageClassName: &storageClass,
		},
	}
	//job := v2alpha1.BatchV2alpha1Client{}
	_, err = client.CoreV1().PersistentVolumeClaims("ouroboros").
		Create(context.Background(), pvc, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
}
