package ssh

import (
	pkg "golang.org/x/crypto/ssh"

	"github.com/mattn/anko/vm"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewPackage("ssh")
	m.Define("Dial", pkg.Dial)
	m.Define("ParsePrivateKey", pkg.ParsePrivateKey)
	m.DefineType("ClientConfig", make(pkg.ClientConfig))
	return m
}
