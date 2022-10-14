/*
Copyright 2021 The Pixiu Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package exec

import (
	osexec "os/exec"
)

// ErrExecutableNotFound is returned if the executable is not found.
var ErrExecutableNotFound = osexec.ErrNotFound

type Cmd interface {
	Run() error
	CombinedOutput() ([]byte, error)
}

type Interface interface {
	// LookPath wraps os/exec.LookPath
	LookPath(cmd string) (string, error)
	// Command is part of the Interface to Run.
	Command(cmd string, args ...string) Cmd
}

// Implements Interface in terms of really exec()ing.
type executor struct{}

// New returns a new Interface which will os/exec to run commands.
func New() Interface {
	return &executor{}
}

// Command is part of the Interface.
func (executor *executor) Command(cmd string, args ...string) Cmd {
	return osexec.Command(cmd, args...)
}

// LookPath is part of the Interface
func (executor *executor) LookPath(cmd string) (string, error) {
	return osexec.LookPath(cmd)
}
