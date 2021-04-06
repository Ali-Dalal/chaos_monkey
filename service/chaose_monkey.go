package service

import (
	"chaos_monkey/utils"
	"fmt"
	"k8s.io/apimachinery/pkg/util/rand"
	"k8s.io/client-go/kubernetes"
)

func KillRandomPod(client kubernetes.Interface, namespace string){
	pods, _ := GetPods(client, namespace)
	randomPodIndex := rand.Int63nRange(0, int64(len(pods.Items)-1))
	selectedPod := pods.Items[randomPodIndex]
	utils.LogInfo(fmt.Sprintf("attempting to kill pod number %d, name: %s", randomPodIndex +1, selectedPod.Name))
	err :=DeletePod(client,selectedPod.Name, namespace)
	if err != nil {
		utils.LogError(fmt.Sprintf("Unable to delete pod %s ", selectedPod.Name))
		utils.LogError(err.Error())
		return
	}
	utils.LogInfo(fmt.Sprintf("Pod %s Deleted successfully",selectedPod.Name))
}
