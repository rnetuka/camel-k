/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package log

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/apache/camel-k/pkg/apis/camel/v1alpha1"
	"github.com/arsham/blush/blush"
)

// Print prints integrations logs to the stdout
func Print(ctx context.Context, integration *v1alpha1.Integration) error {
	scraper := NewSelectorScraper(integration.Namespace, "camel.apache.org/integration="+integration.Name)
	reader := scraper.Start(ctx)

	b := &blush.Blush{
		Finders: []blush.Finder{
			blush.NewExact("FATAL", blush.Red),
			blush.NewExact("ERROR", blush.Red),
			blush.NewExact("WARN", blush.Yellow),
			blush.NewExact("INFO", blush.Green),
			blush.NewExact("DEBUG", blush.Colour{
				Foreground: blush.RGB{R: 170, G: 170, B: 170},
				Background: blush.NoRGB,
			}),
			blush.NewExact("TRACE", blush.Colour{
				Foreground: blush.RGB{R: 170, G: 170, B: 170},
				Background: blush.NoRGB,
			}),
		},
		Reader: ioutil.NopCloser(reader),
	}

	if _, err := io.Copy(os.Stdout, b); err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
