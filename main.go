package main

import (
	"context"
	"gopro\pkg\kubernetes\client"
	"k8s.io/client-go/kubernetes"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

/*
 CoreV1() is the place where your basic resource lives like nodes, pods also
 Package v1 is the v1 version of the core API
 deployments lives under apps
 and all those

 How it works to find out which API
 1) Go to https://pkg.go.dev/k8s.io/client-go/kubernetes and Obeserv all the methods that
  available on " *Clienset "  Now say you want to get to the node then  Call "CoreV1()" on the
  *Clientset , it will return of type "corev1.CoreV1Interface"
  Now if you further go into that you will see that It has Many Interfaces .Lets Pick
  "NodesGetter" and if we further go we see that it is interface with method signature
  "Nodes() NodeInterface" --> NOw again NodeInterface is a Interface with method signature
  for "Create , update , delete, , list , Watch, Apply...etc". We can call any of these method
  as per our requirements

  context defines the Context type, which carries deadlines, cancellation signals,
  and other request-scoped values across API boundaries and between processes.

  ListOptions is the query options to a standard REST list call.

*/
func main() {
	clientset, err := client.NewK8sOutClusterClient()
	if err != nil {
		return nil, err
	}

	nodeList, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, node := nodeList.Items {
		fmt.Println(node.Name)
	}

}
