package pipeline

import (
	"fmt"
	utils "json-pipeline/pkg"
)

var log = utils.GetLogger()

type Pipeline struct {
	steps []Step
}

func New(steps ...Step) *Pipeline {
	return &Pipeline{steps: steps}
}

func (p *Pipeline) Run() error {
	for i, step := range p.steps {
		fmt.Printf("Running step %d...\n", i+1)
		if err := step.Run(); err != nil {
			return fmt.Errorf("step %d failed: %w", i+1, err)
		}
	}
	return nil
}
