package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type nodeAndPods struct {
	node string
	pods []string
}

type runningPod struct {
	pName string
	nName string
}

type response struct {
	Items []Item `json:"items"`
}

type Item struct {
	Metadata struct {
		Name string `json:"name"`
	} `json:"metadata"`
	Status struct {
		Phase string `json:"phase"`
	} `json:"status"`
	Spec struct {
		NodeName string `json:"nodeName"`
	} `json:"spec"`
}

func main() {

	res, err := exec.Command("kubectl", "get", "pods", "-o", "json").Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var resp response
	if err := json.Unmarshal(res, &resp); err != nil {
		log.Fatal(err)
	}

	uniqueNodes := map[string]int{}
	cnt_1 := 0
	for _, v := range resp.Items {
		if _, est := uniqueNodes[v.Spec.NodeName]; !est {
			uniqueNodes[v.Spec.NodeName] = cnt_1
			cnt_1++
		}
	}

	nodePods := make([]nodeAndPods, len(uniqueNodes))
	cnt_2 := 0
	for k, _ := range uniqueNodes {
		nodePods[cnt_2].node = k
		cnt_2++
	}

	runningPods := make([]runningPod, len(resp.Items))
	cnt_3 := 0
	for _, v := range resp.Items {
		if v.Status.Phase == "Running" {
			runningPods[cnt_3].nName = v.Spec.NodeName
			runningPods[cnt_3].pName = v.Metadata.Name
			cnt_3++
		}
	}

	for _, v := range runningPods {
		if v.nName != "" {
			nodePods[uniqueNodes[v.nName]].pods = append(nodePods[uniqueNodes[v.nName]].pods, strings.Split(v.pName, "-")[1])
		}
	}
	for k, _ := range nodePods {
		fmt.Println(strings.Split(nodePods[k].node, "-")[8], nodePods[k].pods)
	}
}
