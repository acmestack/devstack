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

package http

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	
	"github.com/gin-gonic/gin"
	
	"github.com/acmestack/devstack/logging"
	"github.com/acmestack/devstack/settings"
)

func InitEngine(cfg *settings.Setting, writer io.Writer) *gin.Engine {
	gin.SetMode(os.Getenv(gin.EnvGinMode))
	
	router := gin.Default()
	
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {
			// custom format
			return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
			)
		},
		Output:    writer,
		SkipPaths: nil,
	}))
	
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.CustomRecovery(func(context *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			logging.Logger.Errorf("error %v", err)
			context.JSON(http.StatusInternalServerError, http.StatusInternalServerError)
		}
		context.AbortWithStatus(http.StatusInternalServerError)
	}))
	
	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	
	return router
}
