package walker

import (
	"errors"
	"fmt"
	"strings"
)

type Dependency struct {
	Dir  string
	Dep  string
	Dirs []string
	Deps []string
}

func (d *Dependency) match(directory string) {
	if len(directory) == 0 {
		return
	}
	if strings.Contains(directory, d.Dep) {
		d.Deps = append(d.Deps, directory)
	}
	if strings.Contains(directory, d.Dir) {
		d.Dirs = append(d.Dirs, directory)
	}
}

func Create(directories []string, dependencies []Dependency) *map[string][]string {
	for _, dir := range directories {
		for i := range dependencies {
			dependencies[i].match(dir)
		}
	}

	graph := make(map[string][]string)
Loop:
	for _, directory := range directories {
		for _, d := range dependencies {
			for _, dir := range d.Dirs {
				if directory == dir {
					graph[directory] = append(graph[directory], d.Deps...)
					continue Loop
				}
			}
		}
		graph[directory] = []string{}
	}
	return &graph
}

type graph map[string][]string

func topSortDFS(g graph) (order, cyclic []string) {
	L := make([]string, len(g))
	i := len(L)
	temp := map[string]bool{}
	perm := map[string]bool{}
	var cycleFound bool
	var cycleStart string
	var visit func(string)
	visit = func(n string) {
		switch {
		case temp[n]:
			cycleFound = true
			cycleStart = n
			return
		case perm[n]:
			return
		}
		temp[n] = true
		for _, m := range g[n] {
			visit(m)
			if cycleFound {
				if cycleStart > "" {
					cyclic = append(cyclic, n)
					if n == cycleStart {
						cycleStart = ""
					}
				}
				return
			}
		}
		delete(temp, n)
		perm[n] = true
		i--
		L[i] = n
	}
	for n := range g {
		if perm[n] {
			continue
		}
		visit(n)
		if cycleFound {
			return nil, cyclic
		}
	}
	return L, nil
}

type Node struct {
	Name  string
	Edges []string
}

type Graph map[string][]string

func (g *Graph) Contains(edges []string) bool {
	m := *g
	for node, _ := range m {
		for _, edge := range edges {
			if edge == node {
				return true
			}
		}
	}
	return false
}

func bla(unsorted []Node) ([]string, error) {
	unsorted_map := make(map[string][]string)
	for _, n := range unsorted {
		unsorted_map[n.Name] = n.Edges
	}

	sorted := []string{}
	for {
		if len(unsorted_map) == 0 {
			break
		}
		acyclic := false
		fmt.Println(unsorted_map)
		for _, node := range unsorted {
			if _, ok := unsorted_map[node.Name]; !ok {
				continue
			}
			if contains(unsorted_map, node.Edges) {
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

func contains(list map[string][]string, items []string) bool {
	if len(items) == 0 {
		return true
	}
	for itm, _ := range list {
		for _, other := range items {
			if itm == other {
				return false
			}
		}
	}
	return true
}
