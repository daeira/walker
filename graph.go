package walker

import (
	"errors"
	"fmt"
	"regexp"
)

type Dependency struct {
	DirRegexp *regexp.Regexp
	DepRegexp *regexp.Regexp
	Dirs      []string
	Deps      []string
}

func (d *Dependency) match(directory string) {
	if len(directory) == 0 {
		return
	}
	if d.DirRegexp.Match([]byte(directory)) {
		d.Dirs = append(d.Dirs, directory)
	}
	if d.DepRegexp.Match([]byte(directory)) {
		d.Deps = append(d.Deps, directory)
	}
}

func Create(directories []string, dependencies []Dependency) Graph {
	for _, dir := range directories {
		for i := range dependencies {
			dependencies[i].match(dir)
		}
	}

	graph := Graph{}
	for _, directory := range directories {
		graph = append(graph, createDeps(directory, dependencies))
	}
	return graph
}

func createDeps(directory string, dependencies []Dependency) Node {
	n := Node{
		Name: directory,
	}
	for _, d := range dependencies {
		for _, dir := range d.Dirs {
			if directory == dir {
				n.Edges = append(n.Edges, d.Deps...)
			}
		}
	}
	return n
}

type Node struct {
	Name  string
	Edges []string
}

type Graph []Node

func (g *Graph) AsMap() map[string][]string {
	m := map[string][]string{}
	for _, node := range *g {
		m[node.Name] = node.Edges
	}
	return m
}

func contains(items []string, m map[string][]string) bool {
	for _, item := range items {
		if _, ok := m[item]; ok {
			return true
		}
	}
	return false
}

func topoSort(g Graph) ([]string, error) {
	unsorted_map := g.AsMap()

	sorted := []string{}
	for {
		if len(unsorted_map) == 0 {
			break
		}
		acyclic := false
		for _, node := range g {
			if _, ok := unsorted_map[node.Name]; !ok {
				continue
			}
			if len(node.Edges) == 0 || !contains(node.Edges, unsorted_map) {
				acyclic = true
				fmt.Println("remove" + node.Name)
				delete(unsorted_map, node.Name)
				sorted = append(sorted, node.Name)
			}
		}
		if !acyclic {
			return nil, errors.New("acyclic")
		}
	}
	return sorted, nil
}
