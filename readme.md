# Chaos Monkey k8s

Go application that targets a namespace and kill a pod randomly. Schedule rate and namespace are configurable values via  environment variables

# Requirements to run the application
- git cli
- kubectl cli
- helm chart
- minikube (for testing)

# Build the application
Docker image for this application is already available on docker registry
https://hub.docker.com/r/allloush92/chaos_monkey

You can build the image locally using the following command

```shell
git pull https://github.com/Ali-Dalal/chaos_monkey.git
docker build -t {TAG} . //replace tag with desired value
```

# Running the application
by default, helm chart will pull the image directly from docker registry, you can modify `chaos_monkey_helm_chart/templates/deployment.yml` to use a local image if built locally

to install the app, run the following command
```shell
helm install chaos-monkey chaos_monkey_helm_chart/
```

the previous command will install k8s templates that include:
- deployment
- role
- role binding

Role and role binding are important to allow chaos-monkey app to target and kill pods from a different namespace
without them, chaos-monkey will throw permission issue in console

Variables that control the deployment and env variables for chaos-monkey app can be found in `chaos_monkey_helm_chart/values.yml`. this includes env variables that set the namespace and cron time rate


```yaml
  env:
    - name: namespace
      value: {{ .Values.remoteNamespace }}
    - name: rate
      value: "{{ .Values.chaosMonkeyRate }}"
```

resources will be created in default namespace

to check if helm installs the app successfully, you can query the pods as follows:
```shell
kubectl get pods
```
the result will be like the following:
![get pods](./screenshots/chaos_monkey_pods.png?raw=true "get_pods")

### test the application

Chose-monkey runs a cron to target pods from a different namespace and kill a random pod. to check the result, you can query logs from chaos-monkey pod:
```shell
kubectl get pods
kubectl logs -f chaos-monkey-app-{POD_ID} //replcae {POD_ID} with the id from get pods command
```