package policy

import (
	"testing"

	"github.com/leodotcloud/log"
	//"github.com/rancher/go-rancher-metadata/metadata"
)

func init() {
	log.SetLevelString("debug")
}

func TestParseNetworkPolicyWithValidPolicies(t *testing.T) {
	var err error

	np1str := `
[
  {
    "within": "stack",
    "action": "allow"
  },
  {
    "within": "service",
    "action": "allow"
  },
  {
    "within": "linked",
    "action": "allow"
  },
  {
    "from": {
      "selector": "com.company.label1=value1"
    },
    "to": {
      "selector": "com.company.label2=value2"
    },
    "action": "deny"
  }
]
`
	log.Debugf("TestParseNetworkPolicy")

	_, err = ParseNetworkPolicyStr(np1str)
	if err != nil {
		log.Errorf("error parsing policy: %v", err)
	}

	noerrnp1str := `[]`
	_, err = ParseNetworkPolicyStr(noerrnp1str)
	if err != nil {
		t.Errorf("NOT expecting error but got: %v", err)
	}

	noerrnp2str := `
[
  {
    "from": {
      "selector": "com.alpha.label1=value1"
    },
    "to": {
      "selector": "com.bravo.label2=value2"
    },
    "ports": [
      "80/tcp",
      "8080"
    ],
    "action": "allow"
  },
  {
    "within": "stack",
    "action": "allow"
  }
]
`
	_, err = ParseNetworkPolicyStr(noerrnp2str)
	if err != nil {
		t.Errorf("NOT expecting error but got: %v", err)
	}

	noerrnp3str := `
[
  {
    "within": "stack",
    "action": "deny"
  }
]
`
	_, err = ParseNetworkPolicyStr(noerrnp3str)
	if err != nil {
		t.Errorf("NOT expecting error but got: %v", err)
	}

	noerrnp5str := `
[
  {
    "between": {
      "selector": "com.rancher.stack.location=sfo"
    },
    "action": "allow"
  }
]
`
	_, err = ParseNetworkPolicyStr(noerrnp5str)
	if err != nil {
		t.Errorf("not expecting got: %v", err)
	}

	noerrnp6str := `
[
  {
    "between": {
      "groupBy": "com.rancher.stack.location"
    },
    "action": "allow"
  }
]
`
	_, err = ParseNetworkPolicyStr(noerrnp6str)
	if err != nil {
		t.Errorf("not expecting got: %v", err)
	}

	noerrnp7str := `
[
  {
    "between": {
      "groupBy": "com.rancher.stack.location",
      "ports": [
        "80",
        "3000/tcp"
      ]
    },
    "action": "allow"
  }
]
`
	_, err = ParseNetworkPolicyStr(noerrnp7str)
	if err != nil {
		t.Errorf("not expecting got: %v", err)
	}

}

func TestParseNetworkPolicyWithInvalidPolicies(t *testing.T) {
	var err error
	errnp2str := `
[
  {
    "within": "stack",
    "action": "has_to_be_allow_or_deny"
  }
]
`
	_, err = ParseNetworkPolicyStr(errnp2str)
	if err == nil {
		t.Errorf("expecting error got nil")
	}

	errnp3str := `
[
  {
    "within": "has_to_be_stack_or_service_or_linked",
    "action": "deny"
  }
]
`
	_, err = ParseNetworkPolicyStr(errnp3str)
	if err == nil {
		t.Errorf("expecting error got nil")
	}

	errnp4str := `
[
  {
    "between": {
      "group_by": "com.rancher.stack.location"
    },
    "action": "allow"
  }
]
`
	_, err = ParseNetworkPolicyStr(errnp4str)
	if err == nil {
		t.Errorf("expecting error got nil")
	}

	errnp5str := `
[
  {
    "from": {
      "selector": "com.rancher.label=value"
    },
    "ports": [
      "80/tcp",
      "8080"
    ],
    "action": "allow"
  }
]
`
	_, err = ParseNetworkPolicyStr(errnp5str)
	if err == nil {
		t.Errorf("expecting error got nil")
	}

	errnp6str := `
[
  {
    "to": {
      "selector": "com.rancher.label=value"
    },
    "ports": [
      "80/tcp",
      "8080"
    ],
    "action": "allow"
  }
]
`
	_, err = ParseNetworkPolicyStr(errnp6str)
	if err == nil {
		t.Errorf("expecting error got nil")
	}

	errnp7str := `
[
  {
    "ports": [
      "80/tcp",
      "8080"
    ],
    "action": "allow"
  }
]
`
	_, err = ParseNetworkPolicyStr(errnp7str)
	if err == nil {
		t.Errorf("expecting error got nil")
	}

	errnp8str := `
[
  {
    "from": {
      "selector": "com.rancher.label1=value1"
    },
    "to": {
      "selector": "com.rancher.label2=value2"
    },
    "action": "has_to_be_allow_or_deny"
  }
]
`
	_, err = ParseNetworkPolicyStr(errnp8str)
	if err == nil {
		t.Errorf("expecting error got nil")
	}

	errnp9str := `
[
  {
    "between": {
      "selectr": "com.rancher.testlabel=testlabelvalue"
    },
    "action": "allow"
  }
]
`
	_, err = ParseNetworkPolicyStr(errnp9str)
	if err == nil {
		t.Errorf("expecting error got nil")
	}

	errnp10str := `
[
  {
    "from": {
      "selector": "com.rancher.label1=value1"
    },
    "to": {
      "selector": "com.rancher.label2=value2"
    },
    "ports": [
      "abcd"
    ],
    "action": "allow"
  }
]
`
	_, err = ParseNetworkPolicyStr(errnp10str)
	if err == nil {
		t.Errorf("expecting error got nil")
	}

	errnp11str := `
[
  {
    "from": {
      "selector": "com.rancher.label1=value1"
    },
    "to": {
      "selector": "com.rancher.label2=value2"
    },
    "ports": [
      "0"
    ],
    "action": "allow"
  }
]
`
	_, err = ParseNetworkPolicyStr(errnp11str)
	if err == nil {
		t.Errorf("expecting error got nil")
	}

	errnp12str := `
[
  {
    "from": {
      "selector": "com.rancher.label1=value1"
    },
    "to": {
      "selector": "com.rancher.label2=value2"
    },
    "ports": [
      "80000"
    ],
    "action": "allow"
  }
]
`
	_, err = ParseNetworkPolicyStr(errnp12str)
	if err == nil {
		t.Errorf("expecting error got nil")
	}

	errnp13str := `
[
  {
    "from": {
      "selector": "com.rancher.label1=value1"
    },
    "to": {
      "selector": "com.rancher.label2=value2"
    },
    "ports": [
      "1234/xyz"
    ],
    "action": "allow"
  }
]
`
	_, err = ParseNetworkPolicyStr(errnp13str)
	if err == nil {
		t.Errorf("expecting error got nil")
	}

	errnp14str := `
[
  {
    "from": {
      "selector": "com.rancher.label1=value1"
    },
    "to": {
      "selector": "com.rancher.label2=value2"
    },
    "ports": [
      "abcd/xyz"
    ],
    "action": "allow"
  }
]
`
	_, err = ParseNetworkPolicyStr(errnp14str)
	if err == nil {
		t.Errorf("expecting error got nil")
	}

}

// TODO: Comment out
// Works only inside a container with access to metadata
//func TestGetContainersGroupedBy(t *testing.T) {
//	mClient, err := metadata.NewClientAndWait("http://169.254.169.250/2016-07-29")
//	if err != nil {
//		log.Errorf("error creating metadata client: %v", err)
//	}
//
//	w := &watcher{
//		c: mClient,
//	}
//
//	err = w.fetchInfoFromMetadata()
//	if err != nil {
//		log.Errorf("error fetching information from metadata: %v", err)
//	}
//	w.getContainersGroupedBy("com.rancher.stack.location")
//}
//
//func TestMD(t *testing.T) {
//	mClient, err := metadata.NewClientAndWait("http://169.254.169.250/2016-07-29")
//	if err != nil {
//		log.Errorf("error creating metadata client: %v", err)
//	}
//
//	w := &watcher{
//		c: mClient,
//	}
//
//	defaultNetwork, err := w.getDefaultNetwork()
//	if err != nil {
//		log.Errorf("Error while finding default network: %v", err)
//	}
//	log.Debugf("defaultNetwork: %#v", defaultNetwork)
//}
