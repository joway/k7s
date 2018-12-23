package main

type DeployStep struct {
	Name string
	Spec string
}

const (
	OrderlyDeploy = "OrderlyDeploy"
	Deploy        = "Deploy"
)

type DeployPipelineSpec struct {
	action    string
	interrupt string
	steps     []DeployStep
}

type DeployPipeline struct {
	Name string
	Spec string
}
