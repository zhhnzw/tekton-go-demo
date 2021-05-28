package main

import (
	"tekton-go-demo/operator"
)

func main() {

	clients := operator.NewClients("Kubernetes", "default")
	// git-resource.yaml
	// 创建一个 PipelineResource CRD 对象来表示输入是从 git 仓库来获取代码
	operator.CreateGitResource(clients)

	// source-to-image.yaml
	// 创建一个 Task CRD 对象，功能是：构建新的镜像，并推送镜像到dockerhub
	// 创建的 Task 并不会立刻执行，在本例子中会通过 PipelineRun 来执行
	operator.CreateSource2Image(clients)

	// deploy-to-k8s.yaml
	// 创建一个 Task CRD 对象，功能是：部署到本地 k8s 集群
	operator.CreateDeploy2K8s(clients)

	// build-pipeline.yaml
	// 创建一个 Pipeline CRD 对象，功能是：组装流水线，把上面2个 Task 拼在一起
	operator.CreatePipeline(clients)

	// run.yaml
	// 创建一个 PipelineRun CRD 对象，功能是：执行上一步的流水线
	operator.Run(clients)
}
