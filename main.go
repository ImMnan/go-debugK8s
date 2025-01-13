package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	log := logrus.New()
	//log.SetFormatter(&logrus.JSONFormatter{})
	// Configure log formatter
	log.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "time",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyLevel: "level",
		},
	})

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// this creates a clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
		//nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		//fmt.Printf("There are %d nodes in the cluster\n", len(nodes.Items))
		for i, nd := range nodes.Items {

			capacity := map[string]string{
				"cpu":     nd.Status.Capacity.Cpu().String(),
				"memory":  nd.Status.Capacity.Memory().String(),
				"storage": nd.Status.Capacity.StorageEphemeral().String(),
				"pods":    nd.Status.Capacity.Pods().String(),
			}

			allocatable := map[string]string{
				"cpu":     nd.Status.Allocatable.Cpu().String(),
				"memory":  nd.Status.Allocatable.Memory().String(),
				"storage": nd.Status.Allocatable.StorageEphemeral().String(),
				"pods":    nd.Status.Allocatable.Pods().String(),
			}

			log.WithFields(logrus.Fields{
				"time":        time.Now().Format(time.RFC3339),
				"node":        i + 1,
				"name":        nd.Name,
				"allocatable": allocatable,
				"capacity":    capacity,
			}).Info("details")
		}
		time.Sleep(30 * time.Second)
	}
}
