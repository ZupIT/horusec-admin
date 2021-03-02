package kubernetes

//import (
//	clientset "github.com/tiagoangelozup/horusec-admin/pkg/client/clientset/versioned"
//	"github.com/tiagoangelozup/horusec-admin/pkg/client/clientset/versioned/typed/install/v1alpha1"
//	"k8s.io/client-go/tools/clientcmd"
//	"log"
//	"os"
//)
//
//type Client struct {
//	v1alpha1.HorusecManagerInterface
//}
//
//func NewClient() v1alpha1.HorusecManagerInterface {
//	kubeconfig := os.Getenv("KUBECONFIG")
//	namespace := os.Getenv("NAMESPACE")
//
//	if len(kubeconfig) == 0 || len(namespace) == 0 {
//		log.Fatalf("Environment variables KUBECONFIG and NAMESPACE need to be set")
//	}
//
//	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
//	if err != nil {
//		log.Fatalf("Failed to create k8s rest client: %s", err)
//	}
//
//	c, err := clientset.NewForConfig(restConfig)
//	if err != nil {
//		log.Fatalf("Failed to create istio client: %s", err)
//	}
//
//	c2 := &Client{HorusecManagerInterface: c.InstallV1alpha1().HorusecManagers(namespace)}
//	c2.List()
//	return c2
//}
