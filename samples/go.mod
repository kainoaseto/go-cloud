// Copyright 2018-2019 The Go Cloud Development Kit Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

module github.com/kainoaseto/go-cloud/samples

go 1.12

require (
	contrib.go.opencensus.io/exporter/stackdriver v0.12.1
	github.com/Azure/azure-pipeline-go v0.2.1
	github.com/Azure/azure-storage-blob-go v0.8.0
	github.com/aws/aws-sdk-go v1.19.45
	github.com/go-sql-driver/mysql v1.4.1
	github.com/google/go-cmdtest v0.1.0
	github.com/google/go-cmp v0.3.1
	github.com/google/subcommands v1.0.1
	github.com/google/uuid v1.1.1
	github.com/google/wire v0.3.0
	github.com/gorilla/mux v1.7.2
	github.com/kainoaseto/go-cloud v0.18.0
	github.com/kainoaseto/go-cloud/docstore/mongodocstore v0.18.0
	github.com/kainoaseto/go-cloud/pubsub/kafkapubsub v0.18.0
	github.com/kainoaseto/go-cloud/pubsub/natspubsub v0.18.0
	github.com/kainoaseto/go-cloud/pubsub/rabbitpubsub v0.18.0
	github.com/kainoaseto/go-cloud/runtimevar/etcdvar v0.18.0
	github.com/kainoaseto/go-cloud/secrets/hashivault v0.18.0
	github.com/streadway/amqp v0.0.0-20190827072141-edfb9018d271
	go.opencensus.io v0.22.2
	google.golang.org/genproto v0.0.0-20191108220845-16a3f7862a1a
	gopkg.in/pipe.v2 v2.0.0-20140414041502-3c2ca4d52544
)

replace github.com/kainoaseto/go-cloud => ../

replace github.com/kainoaseto/go-cloud/docstore/mongodocstore => ../docstore/mongodocstore

replace github.com/kainoaseto/go-cloud/pubsub/kafkapubsub => ../pubsub/kafkapubsub

replace github.com/kainoaseto/go-cloud/pubsub/natspubsub => ../pubsub/natspubsub

replace github.com/kainoaseto/go-cloud/pubsub/rabbitpubsub => ../pubsub/rabbitpubsub

replace github.com/kainoaseto/go-cloud/runtimevar/etcdvar => ../runtimevar/etcdvar

replace github.com/kainoaseto/go-cloud/secrets/hashivault => ../secrets/hashivault
