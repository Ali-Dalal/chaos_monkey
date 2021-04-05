
package main

import (
	"chaos_monkey/service"
	"chaos_monkey/utils"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"os"
)

func main() {
	utils.InitLogrusLog()
	err := godotenv.Load()
	if err!=nil {
		utils.LogError("Could not load env variables")
	}
	rate := os.Getenv("rate")
	namespace := os.Getenv("namespace")
	utils.LogInfo(fmt.Sprintf("Running chaos monkey for namespace %s and rate %s minute(s)",namespace, rate))
	k8sClient := service.GetClient()
	c := cron.New()
	_, _ = c.AddFunc(fmt.Sprintf("@every %sm", rate), func() {
		service.KillRandomPod(k8sClient, namespace)
	})
	c.Run()
}
