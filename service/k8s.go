package service

import (
	"chaos_monkey/utils"
	"context"
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func GetClient() *kubernetes.Clientset {
	utils.LogInfo("Attempting to get a k8s client")
	// uncomment this code if you want to connect to k8s from outside a cluster
	//var kubeconfig *string
	//if home := homedir.HomeDir(); home != "" {
	//	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	//} else {
	//	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	//}
	//flag.Parse()
	//config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	//if err != nil {
	//	logrus.WithError(err).Fatal("could not get config")
	//	panic(err.Error())
	//}
	// use this code to connect to k8s from inside a cluster
	config, err := rest.InClusterConfig()
	if err != nil {
		logrus.WithError(err).Fatal("could not get k8s in-cluster config")
		panic(err.Error())
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		logrus.WithError(err).Fatal("could not get a new client")
		panic(err.Error())
	}
	utils.LogInfo("k8s client instantiated successfully")
	return client
}

func GetPods(client *kubernetes.Clientset, namespace string ) *v1.PodList  {
	pods, err := client.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logrus.WithError(err).Fatal("could not list pods")
	}
	return pods
}

func DeletePod(client *kubernetes.Clientset, name string, namespace string) error {
	err := client.CoreV1().Pods(namespace).Delete(context.TODO(),name,metav1.DeleteOptions{})
	return err
}