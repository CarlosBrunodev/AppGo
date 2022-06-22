package handler

import (
	"context"
	"log"
	"regexp"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func NamespaceCheck(namespaceRegex string, clientset *kubernetes.Clientset) []string {

	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if errors.IsNotFound(err) {
		log.Print("No namespace in the cluster: ", err)
	} else if err != nil {
		log.Print("Failed to fetch namespaces in the cluster: ", err)
	}

	var names []string

	nsregex := regexp.MustCompile(namespaceRegex)
	for _, ns := range namespaces.Items {
		if nsregex.MatchString(ns.Name) {
			names = append(names, ns.Name)
			log.Printf("Namespace %v matches the regex %v", ns, nsregex)
		}
	}
	return names
}
