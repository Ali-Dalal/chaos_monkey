package service

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	testclient "k8s.io/client-go/kubernetes/fake"
	"testing"
)

func TestGetPods(t *testing.T) {
	// use fake k8s client for testing
	client := testclient.NewSimpleClientset()
	_, err := GetPods(client, "default")
	if err != nil {
		t.Fatal("unable to get pods")
	}
}

func TestDeletePod(t *testing.T) {
	client := testclient.NewSimpleClientset()
	// create pod mocked object
	pod := v1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: "test-pod"},
	}
	// create mocked pod via fake client
	v1Pod,_ := client.CoreV1().Pods("default").Create(context.TODO(), &pod, metav1.CreateOptions{})
	err := DeletePod(client, v1Pod.Name, "default")
	if err != nil {
		t.Fatal("Unable to delete a pod")
	}
}

func TestKillRandomPod(t *testing.T) {
	// get fake client
	client := testclient.NewSimpleClientset()
	// create 2 fake pods
	pod1 := v1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: "test-pod"},
	}
	pod2 := v1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: "test-pod2"},
	}
	// create mocked pod via fake client
	client.CoreV1().Pods("default").Create(context.TODO(), &pod1, metav1.CreateOptions{})
	client.CoreV1().Pods("default").Create(context.TODO(), &pod2, metav1.CreateOptions{})
	// KillRandomPod will panic if anything breaks while executing thus the test will fail
	KillRandomPod(client,"default")
}
