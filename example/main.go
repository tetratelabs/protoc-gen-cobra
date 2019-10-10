// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/tetratelabs/protoc-gen-cobra/example/cmd"
	"github.com/tetratelabs/protoc-gen-cobra/example/pb"
)

func init() {
	// Add client generated commands to cobra's root cmd.
	cmd.RootCmd.AddCommand(pb.BankClientCommand())
	cmd.RootCmd.AddCommand(pb.CacheClientCommand())
	cmd.RootCmd.AddCommand(pb.TimerClientCommand())
	cmd.RootCmd.AddCommand(pb.NestedMessagesClientCommand())
	cmd.RootCmd.AddCommand(pb.CRUDClientCommand())
}

func main() {
	cmd.Execute()
}
