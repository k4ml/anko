package agent

import (
	pkg "golang.org/x/crypto/ssh/agent"

	"github.com/mattn/anko/vm"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewPackage("agent")
	m.Define("NewClient", pkg.NewClient)
	return m
}
