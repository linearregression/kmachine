/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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
	"github.com/spf13/cobra"

	"k8s.io/helm/pkg/repo"
)

var serveDesc = `This command starts a local chart repository server that serves the charts saved in your $HELM_HOME/local/ directory.`

//TODO: add repoPath flag to be passed in in case you want
//  to serve charts from a different local dir

func init() {
	RootCommand.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start a local http web server",
	Long:  serveDesc,
	Run:   serve,
}

func serve(cmd *cobra.Command, args []string) {
	repo.StartLocalRepo(localRepoDirectory())
}
