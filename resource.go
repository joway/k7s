package main

type Definition struct {
	Name string
	Spec string
}

type Resource struct {
	Name   string
	Extend Definition
	Spec   string
}

func NewDefinition(name string, spec string) *Definition {
	return &Definition{
		Name: name,
		Spec: spec,
	}
}

func NewResource(name string, spec string) *Resource {
	return &Resource{
		Name:   name,
		Spec:   spec,
		Extend: nil,
	}
}
