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

package mongodocstore_test

import (
	"context"
	"log"

	"github.com/kainoaseto/go-cloud/docstore"
	"github.com/kainoaseto/go-cloud/docstore/mongodocstore"
)

func ExampleOpenCollection() {
	// PRAGMA: This example is used on github.com/kainoaseto/go-cloud; PRAGMA comments adjust how it is shown and can be ignored.
	// PRAGMA: On github.com/kainoaseto/go-cloud, hide lines until the next blank line.
	ctx := context.Background()

	client, err := mongodocstore.Dial(ctx, "mongodb://my-host")
	if err != nil {
		log.Fatal(err)
	}
	mcoll := client.Database("my-db").Collection("my-coll")
	coll, err := mongodocstore.OpenCollection(mcoll, "userID", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer coll.Close()
}

func ExampleOpenCollectionWithIDFunc() {
	// PRAGMA: This example is used on github.com/kainoaseto/go-cloud; PRAGMA comments adjust how it is shown and can be ignored.
	// PRAGMA: On github.com/kainoaseto/go-cloud, hide lines until the next blank line.
	ctx := context.Background()
	type HighScore struct {
		Game   string
		Player string
	}

	client, err := mongodocstore.Dial(ctx, "mongodb://my-host")
	if err != nil {
		log.Fatal(err)
	}
	mcoll := client.Database("my-db").Collection("my-coll")

	// The name of a document is constructed from the Game and Player fields.
	nameFromDocument := func(doc docstore.Document) interface{} {
		hs := doc.(*HighScore)
		return hs.Game + "|" + hs.Player
	}

	coll, err := mongodocstore.OpenCollectionWithIDFunc(mcoll, nameFromDocument, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer coll.Close()
}

func Example_openCollectionFromURL() {
	// PRAGMA: This example is used on github.com/kainoaseto/go-cloud; PRAGMA comments adjust how it is shown and can be ignored.
	// PRAGMA: On github.com/kainoaseto/go-cloud, add a blank import: _ "github.com/kainoaseto/go-cloud/docstore/mongodocstore"
	// PRAGMA: On github.com/kainoaseto/go-cloud, hide lines until the next blank line.
	ctx := context.Background()

	// docstore.OpenCollection creates a *docstore.Collection from a URL.
	coll, err := docstore.OpenCollection(ctx, "mongo://my-db/my-collection?id_field=userID")
	if err != nil {
		log.Fatal(err)
	}
	defer coll.Close()
}
