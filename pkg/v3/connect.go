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

package v3

import (
	"fmt"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"log"
)

// GetActionConfig returns action configuration based on Helm env
func GetActionConfig(namespace string, kubeConfig *cli.EnvSettings) (*action.Configuration, error) {
	actionConfig := new(action.Configuration)

	// Add kube config settings passed by user
	err := actionConfig.Init(kubeConfig.RESTClientGetter(), namespace, "secrets", debug(kubeConfig))
	if err != nil {
		return nil, err
	}

	return actionConfig, err
}

func debug(setting *cli.EnvSettings) func(format string, v ...interface{}) {
	return func(format string, v ...interface{}) {
		if setting.Debug {
			format = fmt.Sprintf("[debug] %s\n", format)
			log.Output(2, fmt.Sprintf(format, v...))
		}
	}
}
