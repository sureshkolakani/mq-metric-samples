package main

/*
  Copyright (c) IBM Corporation 2016

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific

   Contributors:
     Mark Taylor - Initial Contribution
*/

import (
	"bufio"
	"flag"
	"fmt"
	cf "github.com/ibm-messaging/mq-metric-samples/pkg/config"
	"os"
)

type mqTTYConfig struct {
	cf       cf.Config
	interval string
}

var config mqTTYConfig

/*
initConfig parses the command line parameters.
*/
func initConfig() error {
	var err error

	cf.InitConfig(&config.cf)

	flag.StringVar(&config.interval, "ibmmq.interval", "10s", "How long between each collection")

	flag.Parse()

	if len(flag.Args()) > 0 {
		err = fmt.Errorf("Extra command line parameters given")
		flag.PrintDefaults()
	}

	if err == nil {
		err = cf.VerifyConfig(&config.cf)
	}

	if err == nil {
		if config.cf.CC.UserId != "" && config.cf.CC.Password == "" {
			// TODO: If stdin is a tty, then disable echo. Done differently on Windows and Unix
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Printf("Enter password: \n")
			scanner.Scan()
			config.cf.CC.Password = scanner.Text()
		}
	}

	return err

}
