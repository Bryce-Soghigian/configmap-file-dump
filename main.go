package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func parseFile(configMapName, namespace, fileToDump string) {

	// We want to create a configmap from the file we have fileToDump, and output a ConfigMap Kubernetes version
	//
	fmt.Println("apiVersion: v1")
	fmt.Println("kind: ConfigMap")
	fmt.Println("metadata:")
	fmt.Println("  name:", configMapName)
	fmt.Println("  namespace:", namespace)
	fmt.Println("data:")
	fmt.Println("  ", fileToDump, ":", "|")

	file, err := os.Open(fileToDump)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf("    %s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func main() {
	configMapName := flag.String("configmap", "", "Name of the ConfigMap")
	namespace := flag.String("namespace", "", "Namespace for the ConfigMap")
	fileToDump := flag.String("file", "", "File to be dumped into the ConfigMap")

	flag.Parse()

	if *configMapName == "" || *namespace == "" || *fileToDump == "" {
		fmt.Println("Please provide all required flags: -configmap, -namespace, and -file")
		os.Exit(1)
	}

	parseFile(*configMapName, *namespace, *fileToDump)
}

