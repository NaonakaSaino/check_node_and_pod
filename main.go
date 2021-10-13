package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type nodeAndPods struct {
	node string
	pods []string
}

func main() {
	podsStdOut, err := exec.Command("kubectl", "get", "pods", "-o=jsonpath='{.items[*].metadata.name}'").Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	podsStr := strings.Replace(string(podsStdOut), "'", "", -1)
	pods := strings.Split(podsStr, " ")
	fmt.Println("pods:", len(pods))

	nodesStdOut, err := exec.Command("kubectl", "get", "pods", "-o=jsonpath='{.items[*].spec.nodeName}'").Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	nodesStr := strings.Replace(string(nodesStdOut), "'", "", -1)
	nodes := strings.Split(nodesStr, " ")
	fmt.Println("nodes:", len(nodes))

	var uniqueNodes []string
	for _, v := range nodes {
		if !contains(uniqueNodes, v) {
			uniqueNodes = append(uniqueNodes, v)
		}
	}

	nodePods := make([]nodeAndPods, len(uniqueNodes))
	for i, v := range uniqueNodes {
		nodePods[i].node = v
		for j, n := range nodes {
			if n == v {
				nodePods[i].pods = append(nodePods[i].pods, strings.Split(pods[j], "-")[1])
			}
		}
	}
	for k, _ := range nodePods {
		fmt.Println(strings.Split(nodePods[k].node, "-")[8], nodePods[k].pods)
	}
}

func contains(slice []string, s string) bool {
	for _, v := range slice {
		if s == v {
			return true
		}
	}
	return false
}
