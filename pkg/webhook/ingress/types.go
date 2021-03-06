/*
Copyright 2020 Clastix Labs.

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

package ingress

import (
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	networkingv1 "k8s.io/api/networking/v1"
	networkingv1beta "k8s.io/api/networking/v1beta1"
)

const (
	annotationName = "kubernetes.io/ingress.class"
)

type Ingress interface {
	IngressClass() *string
	Namespace() string
	Hostnames() []string
}

type NetworkingV1 struct {
	*networkingv1.Ingress
}

func (n NetworkingV1) IngressClass() (res *string) {
	res = n.Spec.IngressClassName
	if res == nil {
		if a := n.GetAnnotations(); a != nil {
			if v, ok := a[annotationName]; ok {
				res = &v
			}
		}
	}
	return
}

func (n NetworkingV1) Namespace() string {
	return n.GetNamespace()
}

func (n NetworkingV1) Hostnames() []string {
	rules := n.Spec.Rules
	var hostnames []string
	for _, el := range rules {
		hostnames = append(hostnames, el.Host)
	}
	return hostnames
}

type NetworkingV1Beta1 struct {
	*networkingv1beta.Ingress
}

func (n NetworkingV1Beta1) IngressClass() (res *string) {
	res = n.Spec.IngressClassName
	if res == nil {
		if a := n.GetAnnotations(); a != nil {
			if v, ok := a[annotationName]; ok {
				res = &v
			}
		}
	}
	return
}

func (n NetworkingV1Beta1) Namespace() string {
	return n.GetNamespace()
}

func (n NetworkingV1Beta1) Hostnames() []string {
	rules := n.Spec.Rules
	var hostnames []string
	for _, rule := range rules {
		hostnames = append(hostnames, rule.Host)
	}
	return hostnames
}

type Extension struct {
	*extensionsv1beta1.Ingress
}

func (e Extension) IngressClass() (res *string) {
	res = e.Spec.IngressClassName
	if res == nil {
		if a := e.GetAnnotations(); a != nil {
			if v, ok := a[annotationName]; ok {
				res = &v
			}
		}
	}
	return
}

func (e Extension) Namespace() string {
	return e.GetNamespace()
}

func (e Extension) Hostnames() []string {
	rules := e.Spec.Rules
	var hostnames []string
	for _, el := range rules {
		hostnames = append(hostnames, el.Host)
	}
	return hostnames
}
