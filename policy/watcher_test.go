package policy

import (
	"os"
	"testing"

	"github.com/leodotcloud/log"
	"github.com/rancher/go-rancher-metadata/metadata"
)

// Some of the tests can run only when in development,
// remember to disable this before commiting the code.
const inDevelopment = false

func setup1() *watcher {
	var err error
	mc, err := metadata.NewClientAndWait("http://192.168.236.1:20001/2016-07-29")
	//mc, err := metadata.NewClientAndWait("http://192.168.236.1:20002/2016-07-29")
	//mc, err := metadata.NewClientAndWait("http://192.168.236.1:20003/2016-07-29")
	if err != nil {
		log.Errorf("error creating metadata client: %v", err)
	}

	return &watcher{
		c:                  mc,
		shutdownInProgress: false,
		doCleanup:          false,
		exitCh:             make(chan int),
		signalCh:           make(chan os.Signal, 2),
	}
}

func setup2() *watcher {
	var err error
	//mc, err := metadata.NewClientAndWait("http://192.168.236.1:20001/2016-07-29")
	mc, err := metadata.NewClientAndWait("http://192.168.236.1:20002/2016-07-29")
	//mc, err := metadata.NewClientAndWait("http://192.168.236.1:20003/2016-07-29")
	if err != nil {
		log.Errorf("error creating metadata client: %v", err)
	}

	return &watcher{
		c:                  mc,
		shutdownInProgress: false,
		doCleanup:          false,
		exitCh:             make(chan int),
		signalCh:           make(chan os.Signal, 2),
	}
}

func setup3() *watcher {
	var err error
	//mc, err := metadata.NewClientAndWait("http://192.168.236.1:20001/2016-07-29")
	//mc, err := metadata.NewClientAndWait("http://192.168.236.1:20002/2016-07-29")
	mc, err := metadata.NewClientAndWait("http://192.168.236.1:20003/2016-07-29")
	if err != nil {
		log.Errorf("error creating metadata client: %v", err)
	}

	return &watcher{
		c:                  mc,
		shutdownInProgress: false,
		doCleanup:          false,
		exitCh:             make(chan int),
		signalCh:           make(chan os.Signal, 2),
	}
}

func TestBuildLinkedMappings(t *testing.T) {
	if !inDevelopment {
		t.Skip("not in development mode")
	}
	log.SetLevel(log.DebugLevel)
	mClient, err := metadata.NewClientAndWait("http://192.168.236.1:19999/2016-07-29")
	if err != nil {
		log.Errorf("error creating metadata client: %v", err)
	}

	//services, err := mClient.GetServices()
	//if err != nil {
	//	log.Errorf("Error getting services from metadata: %v", err)
	//	return
	//}

	//buildLinkedMappings(services)
	//buildServiceAliasesMap(services)

	containers, err := mClient.GetContainers()
	if err != nil {
		log.Errorf("Error getting containers from metadata: %v", err)
	}
	buildLinkedMappingsForContainers(containers)
}

func TestFetchInfoFromMD(t *testing.T) {
	if !inDevelopment {
		t.Skip("not in development mode")
	}
	log.SetLevel(log.DebugLevel)
	log.Debugf("TestFetchInfoFromMD")

	var err error
	w := setup1()

	err = w.fetchInfoFromMetadata()
	if err != nil {
		log.Errorf("error fetching information from metadata: %v", err)
	}
}

func TestBuildLinkedMappingsAA(t *testing.T) {
	if !inDevelopment {
		t.Skip("not in development mode")
	}
	log.SetLevel(log.DebugLevel)
	log.Debugf("TestBuildLinkedMappings")

	w1 := setup1()
	w1.fetchInfoFromMetadata()
	w1.buildLinkedMappings()

	log.Debugf("----------------")

	w2 := setup2()
	w2.fetchInfoFromMetadata()
	w2.buildLinkedMappings()

	log.Debugf("----------------")

	w3 := setup3()
	w3.fetchInfoFromMetadata()
	w3.buildLinkedMappings()
}
