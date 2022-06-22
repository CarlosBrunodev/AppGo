package handler

import (
	"context"
	"fmt"
	"log"

	rbac "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func CreateRbacs(namespace string, rolename string, clientset *kubernetes.Clientset) {

	rolebinding := &rbac.RoleBinding{
		TypeMeta:   metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{Name: namespace, Namespace: namespace},
		Subjects:   []rbac.Subject{{APIGroup: "rbac.authorization.k8s.io", Kind: "Group", Name: fmt.Sprintf("runway:%s", rolename)}},
		RoleRef:    rbac.RoleRef{APIGroup: "rbac.authorization.k8s.io", Kind: "ClusterRole", Name: "cluster-admin"},
	}

	_, err := clientset.RbacV1().RoleBindings(namespace).Create(context.TODO(), rolebinding, metav1.CreateOptions{})
	if errors.IsAlreadyExists(err) {
		log.Print("Rbac already exists: ", err)
	} else if err != nil {
		log.Print("Failed to create rbac: ", err)
	} else {
		log.Printf("Rbac for %s successfully created using %s role", namespace, rolename)
	}

}
