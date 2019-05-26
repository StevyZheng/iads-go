package api_v1_0

import (
	"flag"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"path/filepath"
)

type K8sOp struct {
	kubeconfig *string
	clientset  *kubernetes.Clientset
}

func (e *K8sOp) Init() {
	if home := homedir.HomeDir(); home != "" {
		e.kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		e.kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *e.kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	e.clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
}

func (e *K8sOp) GetPodInfo() *v1.PodList {
	pods, err := e.clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		log.Println(err.Error())
	}
	return pods
}

func (e *K8sOp) GetNodeInfo() *v1.NodeList {
	nodes, err := e.clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		log.Println(err.Error())
	}
	return nodes
}
