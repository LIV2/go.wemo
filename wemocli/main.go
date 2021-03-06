//package main ...
/* Copyright 2014 Matt Ho
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
*/
package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "wemo"
	app.Usage = "command line interface wemo"
	app.Version = "0.1"
	app.Commands = []cli.Command{
		discoverCommand,
		onCommand,
		offCommand,
		toggleCommand,
		bulbCommand,
		bulbStatusCommand,
	}
	app.Run(os.Args)
}
