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

package stub

import (
	ctx "context"
	"github.com/apache/camel-k/pkg/apis/camel/v1alpha1"
	"github.com/apache/camel-k/pkg/stub/action/platform"

	"github.com/apache/camel-k/pkg/stub/action/context"
	"github.com/apache/camel-k/pkg/stub/action/integration"
	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/sirupsen/logrus"
)

// NewHandler --
func NewHandler(ctx ctx.Context, namespace string) sdk.Handler {
	return &handler{
		integrationActionPool: []integration.Action{
			integration.NewInitializeAction(),
			integration.NewBuildAction(namespace),
			integration.NewDeployAction(),
			integration.NewMonitorAction(),
		},
		integrationContextActionPool: []context.Action{
			context.NewInitializeAction(),
			context.NewBuildAction(ctx),
			context.NewMonitorAction(),
		},
		integrationPlatformActionPool: []platform.Action{
			platform.NewInitializeAction(),
			platform.NewCreateAction(),
			platform.NewStartAction(),
		},
	}
}

type handler struct {
	integrationActionPool         []integration.Action
	integrationContextActionPool  []context.Action
	integrationPlatformActionPool []platform.Action
}

func (h *handler) Handle(ctx ctx.Context, event sdk.Event) error {
	switch o := event.Object.(type) {
	case *v1alpha1.Integration:
		for _, a := range h.integrationActionPool {
			if a.CanHandle(o) {
				logrus.Debug("Invoking action ", a.Name(), " on integration ", o.Name)
				if err := a.Handle(o); err != nil {
					return err
				}
			}
		}
	case *v1alpha1.IntegrationContext:
		for _, a := range h.integrationContextActionPool {
			if a.CanHandle(o) {
				logrus.Debug("Invoking action ", a.Name(), " on context ", o.Name)
				if err := a.Handle(o); err != nil {
					return err
				}
			}
		}
	case *v1alpha1.IntegrationPlatform:
		for _, a := range h.integrationPlatformActionPool {
			if a.CanHandle(o) {
				logrus.Debug("Invoking action ", a.Name(), " on platform ", o.Name)
				if err := a.Handle(o); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
