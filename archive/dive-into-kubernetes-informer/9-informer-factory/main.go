package main

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
)

func main() {
	informerFactory := informers.NewSharedInformerFactory(mustClientset(), 0)
	configMapsInformer := informerFactory.Core().V1().ConfigMaps().Informer()
	configMapsInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			configMap, ok := obj.(*corev1.ConfigMap)
			if !ok {
				return
			}
			fmt.Printf("created: %s\n", configMap.Name)
		},
	})

	stopper := make(chan struct{})
	defer close(stopper)

	fmt.Println("Start syncing....")

	go configMapsInformer.Run(stopper)
	runtime.HandleCrash()

	if !cache.WaitForCacheSync(stopper, configMapsInformer.HasSynced) {
		panic("timed out waiting for caches to sync")
	}

	<-stopper
}