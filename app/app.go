/*
 * Licensed to the AcmeStack under one or more contributor license
 * agreements. See the NOTICE file distributed with this work for
 * additional information regarding copyright ownership.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package app

import (
	"os"
	
	"github.com/acmestack/devstack/http"
	"github.com/acmestack/devstack/logging"
	"github.com/acmestack/devstack/settings"
)

type App struct {
	setting *settings.Setting
}

var app *App

func Run(enginePatch http.EnginePatchFunc, patch ...interface{}) {
	writer := logging.MultiWriter()
	cfg := settings.NewSetting(settings.AppConfigurations, patch...)
	
	logging.InitLogger(os.Getenv(settings.LogLevel))
	
	engine := http.InitEngine(cfg, writer)
	if enginePatch != nil {
		enginePatch(engine)
	}
	http.Run(engine)
	app = &App{setting: cfg}
}

func Settings() *settings.Setting {
	return app.setting
}
