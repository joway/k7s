package resource

type Definition struct {
	Name string
	Spec string
}

type Resource struct {
	Name string
	Extend Definition
	Spec string
}
