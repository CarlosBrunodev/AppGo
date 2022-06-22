package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"main/handler"
	"os"
	"path/filepath"

	//"k8s.io/apimachinery/pkg/api/errors"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "eks-config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Printf("Could not get config: %v", err)
		config, err = rest.InClusterConfig() // returns a config object which uses the service account kubernetes gives to pods.
		if err != nil {
			log.Printf("Could not get inclusterconfig: %v", err)
		}
	}

	// Clientset will make the requests to cluster
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Printf("Could not create clientset: %v", err)
	}

	jsonFile, err := os.Open("../config/config.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	for key, value := range result {
		nsList := handler.NamespaceCheck(key, clientset)
		fmt.Println(nsList)
		for _, ns := range nsList {
			fmt.Println(ns)
			handler.CreateRbacs(ns, fmt.Sprintf("%v", value), clientset)
		}
	}

}
