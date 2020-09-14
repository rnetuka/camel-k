// Copyright 2019 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package crane

import (
	"net/http"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

type options struct {
	name   []name.Option
	remote []remote.Option
}

func makeOptions(opts ...Option) options {
	opt := options{
		remote: []remote.Option{
			remote.WithAuthFromKeychain(authn.DefaultKeychain),
		},
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

// Option is a functional option for crane.
type Option func(*options)

// WithTransport is a functional option for overriding the default transport
// for remote operations.
func WithTransport(t http.RoundTripper) Option {
	return func(o *options) {
		o.remote = append(o.remote, remote.WithTransport(t))
	}
}

// Insecure is an Option that allows image references to be fetched without TLS.
func Insecure(o *options) {
	o.name = append(o.name, name.Insecure)
}
