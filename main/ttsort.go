package main

import (
	"sort"
)

// sort interface for events
func (e Events) Len() int {
	return len(e.Events)
}
func (e Events) Less(i, j int) bool {
	return e.Events[i].Id > e.Events[j].Id
}
func (e Events) Swap(i, j int) {
	e.Events[i], e.Events[j] = e.Events[j], e.Events[i]
}

// sort interface for idmaps
func (e IdMaps) Len() int {
	return len(e.LidMap)
}
func (e IdMaps) Less(i, j int) bool {
	return e.LidMap[i].Lid > e.LidMap[j].Lid
}
func (e IdMaps) Swap(i, j int) {
	e.LidMap[i], e.LidMap[j] = e.LidMap[j], e.LidMap[i]
}

// sort interface for Workflows
func (e Workflows) Len() int {
	return len(e.Workflows)
}
func (e Workflows) Less(i, j int) bool {
	return e.Workflows[i].Pathway > e.Workflows[j].Pathway
}
func (e Workflows) Swap(i, j int) {
	e.Workflows[i], e.Workflows[j] = e.Workflows[j], e.Workflows[i]
}

// sort interface for Pathways
func (e Pwys) Len() int {
	return len(e.Pwy)
}
func (e Pwys) Less(i, j int) bool {
	return e.Pwy[i].Text < e.Pwy[j].Text
}
func (e Pwys) Swap(i, j int) {
	e.Pwy[i], e.Pwy[j] = e.Pwy[j], e.Pwy[i]
}

// sort interface for Subscriptions
func (e Subscriptions) Len() int {
	return len(e.Subscriptions)
}
func (e Subscriptions) Less(i, j int) bool {
	return e.Subscriptions[i].User < e.Subscriptions[j].User
}
func (e Subscriptions) Swap(i, j int) {
	e.Subscriptions[i], e.Subscriptions[j] = e.Subscriptions[j], e.Subscriptions[i]
}
func (i *Trans) sortSubscriptions() {
	data := i.Subscriptions.Subscriptions
	sort.SliceStable(data, func(i, j int) bool {
		return data[i].User < data[j].User
	})
	sort.SliceStable(data, func(i, j int) bool {
		if data[i].User == data[j].User {
			return data[i].Pathway < data[j].Pathway
		}
		return false
	})
	sort.SliceStable(data, func(i, j int) bool {
		if data[i].Pathway == data[j].Pathway {
			return data[i].Expression < data[j].Expression
		}
		return false
	})
}
