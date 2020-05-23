package controller

import (
	"github.com/kruczjak/pods-operator/pkg/controller/deploymentmanager"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, deploymentmanager.Add)
}
