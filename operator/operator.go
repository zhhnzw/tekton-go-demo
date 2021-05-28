package operator

import (
	"context"
	"fmt"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	typedv1alpha1 "github.com/tektoncd/pipeline/pkg/client/clientset/versioned/typed/pipeline/v1alpha1"
	resourceversioned "github.com/tektoncd/pipeline/pkg/client/resource/clientset/versioned"
	resourcev1alpha1 "github.com/tektoncd/pipeline/pkg/client/resource/clientset/versioned/typed/resource/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	knativetest "knative.dev/pkg/test"
	"log"
	"os/user"
	"time"
)

type Clients struct {
	KubeClient             *kubernetes.Clientset
	PipelineClient         typedv1alpha1.PipelineInterface
	ClusterTaskClient      typedv1alpha1.ClusterTaskInterface
	TaskClient             typedv1alpha1.TaskInterface
	TaskRunClient          typedv1alpha1.TaskRunInterface
	PipelineRunClient      typedv1alpha1.PipelineRunInterface
	PipelineResourceClient resourcev1alpha1.PipelineResourceInterface
	ConditionClient        typedv1alpha1.ConditionInterface
	RunClient              typedv1alpha1.RunInterface
}

func NewClients(clusterName, namespace string) *Clients {
	var err error
	c := &Clients{}
	configPath := fmt.Sprintf("%s/.kube/config", GetHomePath())
	// 使用 kubectl 默认配置 ~/.kube/config
	k8sConfig, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		log.Fatalf("%v", err)
	}
	// 创建 k8s 客户端
	c.KubeClient, err = kubernetes.NewForConfig(k8sConfig)
	if err != nil {
		log.Fatalf("failed to create kubeclient from config file at %s: %s", configPath, err)
	}

	cfg, err := knativetest.BuildClientConfig(configPath, clusterName)
	if err != nil {
		log.Fatalf("failed to create configuration obj from %s for cluster %s: %s", configPath, clusterName, err)
	}

	cs, err := versioned.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("failed to create pipeline clientset from config file at %s: %s", configPath, err)
	}
	rcs, err := resourceversioned.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("failed to create pipeline clientset from config file at %s: %s", configPath, err)
	}
	c.PipelineClient = cs.TektonV1alpha1().Pipelines(namespace)
	c.ClusterTaskClient = cs.TektonV1alpha1().ClusterTasks()
	c.TaskClient = cs.TektonV1alpha1().Tasks(namespace)
	c.TaskRunClient = cs.TektonV1alpha1().TaskRuns(namespace)
	c.PipelineRunClient = cs.TektonV1alpha1().PipelineRuns(namespace)
	c.PipelineResourceClient = rcs.TektonV1alpha1().PipelineResources(namespace)
	c.ConditionClient = cs.TektonV1alpha1().Conditions(namespace)
	c.RunClient = cs.TektonV1alpha1().Runs(namespace)
	return c
}

func GetHomePath() string {
	u, err := user.Current()
	if err == nil {
		return u.HomeDir
	}
	return ""
}

const timeout = time.Second * 3

func CreateGitResource(clients *Clients) {

}

func CreateSource2Image(clients *Clients) {

}

func CreateDeploy2K8s(clients *Clients) {

}

func CreatePipeline(clients *Clients) {

}

func Run(clients *Clients) {
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	run := &v1alpha1.PipelineRun{
		ObjectMeta: metav1.ObjectMeta{
			Name: "generic-pipeline-run",
		},
		Spec: v1alpha1.PipelineRunSpec{
			PipelineRef: &v1alpha1.PipelineRef{Name: "build-pipeline"},
			Resources: []v1alpha1.PipelineResourceBinding{
				{Name: "git-source", ResourceRef: &v1alpha1.PipelineResourceRef{Name: "git-source"}},
			},
			Params: []v1alpha1.Param{
				{Name: "imageUrl", Value: v1alpha1.ArrayOrString{StringVal: "2804696160/tekton-demo", Type: v1alpha1.ParamTypeString}},
				{Name: "imageTag", Value: v1alpha1.ArrayOrString{StringVal: "v0.1", Type: v1alpha1.ParamTypeString}},
				{Name: "pathToDockerFile", Value: v1alpha1.ArrayOrString{StringVal: "Dockerfile", Type: v1alpha1.ParamTypeString}},
				{Name: "pathToYamlFile", Value: v1alpha1.ArrayOrString{StringVal: "deployment.yaml", Type: v1alpha1.ParamTypeString}},
			},
			ServiceAccountName: "tekton-test",
		},
	}
	run, err := clients.PipelineRunClient.Create(ctx, run, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("run err: %v", err)
	}
}
