package framework

import (
	"context"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (i *k8sInvocation) CreateSecret(secret *core.Secret) error {
	_, err := i.kubeClient.CoreV1().Secrets(i.Namespace()).Create(context.TODO(), secret, metav1.CreateOptions{})
	return err
}

func (i *k8sInvocation) DeleteSecret(name, ns string) error {
	return i.kubeClient.CoreV1().Secrets(ns).Delete(context.TODO(), name, *deleteInForeground())
}

func (i *k8sInvocation) GetSecret(name, ns string) (*core.Secret, error) {
	return i.kubeClient.CoreV1().Secrets(ns).Get(context.TODO(), name, metav1.GetOptions{})
}

func (i *k8sInvocation) GetTokenSecrets(ns string) (*core.SecretList, error) {
	return i.kubeClient.CoreV1().Secrets(ns).List(context.TODO(), metav1.ListOptions{FieldSelector: "type==kubernetes.io/service-account-token"})
}
