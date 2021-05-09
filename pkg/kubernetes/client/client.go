package client

import (
	"flag"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

/*
 Here we are creating a Kubenetes Clientset, to interact with our k8s cluster from Outside

1) First we use HomeDir() function from k8s.io/client-go/util/homedir" which returns the home
 directory for the current user.
	|-----------------------|
	|func HomeDir() string  |
    |-----------------------|

2) Second we use "Flag" as Package flag implements command-line flag parsing.
   We use a String() function from Flag as String() defines a string flag with specified name,
   default value, and usage string. The return value is the address of a string variable
   that stores the value of the flag.
   |--------------------------------------------------------------|
   |func String(name string, value string, usage string) *string  |
   |--------------------------------------------------------------|

   After all flags are defined, call
   |--------------|
   |flag.Parse()  |
   |--------------|
   to parse the command line into the defined flags.

3) Next we use BuildConfigFromFlags() function as BuildConfigFromFlags is a helper function
  that builds configs from a master url or a kubeconfig filepath. These are passed in as command
  line flags for cluster components. Warnings should reflect this usage. If neither masterUrl
  or kubeconfigPath are passed in we fallback to inClusterConfig. If inClusterConfig fails,
  we fallback to the default config.
   |----------------------------------------------------------------------------------------|
   |func BuildConfigFromFlags(masterUrl, kubeconfigPath string) (*restclient.Config, error) |
   |----------------------------------------------------------------------------------------|

4) Finally we use NewForConfig() function to create a Clientset(as it is group of clients/nodes in k8s)
   for the give config.
   NOTE: If config's RateLimiter is not set and QPS and Burst are acceptable, NewForConfig
   will generate a rate-limiter in configShallowCopy.
   |------------------------------------------------------|
   |func NewForConfig(c *rest.Config) (*Clientset, error) |
   |------------------------------------------------------|

*/

func NewK8sOutClusterClient() (*kubernetes.Clientset, error) {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig) // here we are defrencing the kubeconfig Pointer
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset

}
