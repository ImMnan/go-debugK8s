package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"golang.org/x/exp/slog"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// this creates a clientset which can be used to fetch details from the cluster.
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	// get node details
	listNodesDetails(clientset)
}

func listNodesDetails(clientset *kubernetes.Clientset) {
	// Configure slog logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Infinite loop to get node details every 30 seconds, until the pod/program is terminated.
	for {
		nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			fmt.Printf("Error getting nodes: %v", err.Error())
		}
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

			logger.Info("details",
				slog.String("time", time.Now().Format(time.RFC3339)),
				slog.Int("node", i+1),
				slog.String("name", nd.Name),
				slog.Any("allocatable", allocatable),
				slog.Any("capacity", capacity),
			)
		}
		time.Sleep(30 * time.Second)
	}
}
