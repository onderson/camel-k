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

package kubernetes

import (
	routev1 "github.com/openshift/api/route/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// A Collection is a container of Kubernetes resources
type Collection struct {
	items []runtime.Object
}

// NewCollection creates a new empty collection
func NewCollection() *Collection {
	return &Collection{
		items: make([]runtime.Object, 0),
	}
}

// Items returns all resources belonging to the collection
func (c *Collection) Items() []runtime.Object {
	return c.items
}

// Add adds a resource to the collection
func (c *Collection) Add(resource runtime.Object) {
	c.items = append(c.items, resource)
}

// VisitDeployment executes the visitor function on all Deployment resources
func (c *Collection) VisitDeployment(visitor func(*appsv1.Deployment)) {
	c.Visit(func(res runtime.Object) {
		if conv, ok := res.(*appsv1.Deployment); ok {
			visitor(conv)
		}
	})
}

// GetDeployment returns a Deployment that matches the given function
func (c *Collection) GetDeployment(filter func(*appsv1.Deployment)bool) *appsv1.Deployment {
	var retValue *appsv1.Deployment
	c.VisitDeployment(func(re *appsv1.Deployment) {
		if filter(re) {
			retValue = re
		}
	})
	return retValue
}

// VisitConfigMap executes the visitor function on all ConfigMap resources
func (c *Collection) VisitConfigMap(visitor func(*corev1.ConfigMap)) {
	c.Visit(func(res runtime.Object) {
		if conv, ok := res.(*corev1.ConfigMap); ok {
			visitor(conv)
		}
	})
}

// GetConfigMap returns a ConfigMap that matches the given function
func (c *Collection) GetConfigMap(filter func(*corev1.ConfigMap)bool) *corev1.ConfigMap {
	var retValue *corev1.ConfigMap
	c.VisitConfigMap(func(re *corev1.ConfigMap) {
		if filter(re) {
			retValue = re
		}
	})
	return retValue
}

// VisitService executes the visitor function on all Service resources
func (c *Collection) VisitService(visitor func(*corev1.Service)) {
	c.Visit(func(res runtime.Object) {
		if conv, ok := res.(*corev1.Service); ok {
			visitor(conv)
		}
	})
}

// GetService returns a Service that matches the given function
func (c *Collection) GetService(filter func(*corev1.Service)bool) *corev1.Service {
	var retValue *corev1.Service
	c.VisitService(func(re *corev1.Service) {
		if filter(re) {
			retValue = re
		}
	})
	return retValue
}

// VisitRoute executes the visitor function on all Route resources
func (c *Collection) VisitRoute(visitor func(*routev1.Route)) {
	c.Visit(func(res runtime.Object) {
		if conv, ok := res.(*routev1.Route); ok {
			visitor(conv)
		}
	})
}

// GetRoute returns a Route that matches the given function
func (c *Collection) GetRoute(filter func(*routev1.Route)bool) *routev1.Route {
	var retValue *routev1.Route
	c.VisitRoute(func(re *routev1.Route) {
		if filter(re) {
			retValue = re
		}
	})
	return retValue
}

// VisitMetaObject executes the visitor function on all meta.Object resources
func (c *Collection) VisitMetaObject(visitor func(metav1.Object)) {
	c.Visit(func(res runtime.Object) {
		if conv, ok := res.(metav1.Object); ok {
			visitor(conv)
		}
	})
}

// Visit executes the visitor function on all resources
func (c *Collection) Visit(visitor func(runtime.Object)) {
	for _, res := range c.items {
		visitor(res)
	}
}
