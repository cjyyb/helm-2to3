/*
Copyright The Helm Authors.

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

package main

import (
	"log"
	"os"
	"fmt"

	"github.com/cjyyb/helm-2to3/cmd"
	"github.com/spf13/cobra"
	"helm.sh/helm/v3/pkg/action"

)

var settings = cmd.GetSettings()

func main() {
	actionConfig := new(action.Configuration)
	migrateCmd := cmd.NewRootCmd(actionConfig, os.Stdout, os.Args[1:])
	// run when each command's execute method is called
	cobra.OnInitialize(func() {
		if err := actionConfig.Init(settings.RESTClientGetter(), settings.Namespace(), "secrets", debug); err != nil {
			log.Fatal(err)
		}
	})
	if err := migrateCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func debug(format string, v ...interface{}) {
	if settings.Debug {
		format = fmt.Sprintf("[debug] %s\n", format)
		log.Output(2, fmt.Sprintf(format, v...))
	}
}
