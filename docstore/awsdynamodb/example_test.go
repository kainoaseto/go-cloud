// Copyright 2019 The Go Cloud Development Kit Authors
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

package awsdynamodb_test

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/kainoaseto/go-cloud/docstore"
	"github.com/kainoaseto/go-cloud/docstore/awsdynamodb"
)

func ExampleOpenCollection() {
	// PRAGMA: This example is used on github.com/kainoaseto/go-cloud; PRAGMA comments adjust how it is shown and can be ignored.
	sess, err := session.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	coll, err := awsdynamodb.OpenCollection(
		dynamodb.New(sess), "docstore-test", "partitionKeyField", "", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer coll.Close()
}

func Example_openCollectionFromURL() {
	// PRAGMA: This example is used on github.com/kainoaseto/go-cloud; PRAGMA comments adjust how it is shown and can be ignored.
	// PRAGMA: On github.com/kainoaseto/go-cloud, add a blank import: _ "github.com/kainoaseto/go-cloud/docstore/awsdynamodb"
	// PRAGMA: On github.com/kainoaseto/go-cloud, hide lines until the next blank line.
	ctx := context.Background()

	// docstore.OpenCollection creates a *docstore.Collection from a URL.
	coll, err := docstore.OpenCollection(ctx, "dynamodb://my-table?partition_key=name")
	if err != nil {
		log.Fatal(err)
	}
	defer coll.Close()
}
