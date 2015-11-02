package walker

import "strings"

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
