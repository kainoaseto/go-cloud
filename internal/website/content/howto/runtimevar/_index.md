---
title: "Runtime Configuration"
date: 2019-07-11T12:00:00-07:00
lastmod: 2019-07-29T12:00:00-07:00
showInSidenav: true
toc: true
---

The [`runtimevar` package][] provides an easy and portable way to watch runtime
configuration variables. This guide shows how to work with runtime configuration
variables using the Go CDK.

<!--more-->

Subpackages contain driver implementations of runtimevar for various services,
including Cloud and on-prem solutions. You can develop your application locally
using [`filevar`][] or [`constantvar`][], then deploy it to multiple Cloud
providers with minimal initialization reconfiguration.

[`runtimevar` package]: https://godoc.org/github.com/kainoaseto/go-cloud/runtimevar
[`filevar`]: https://godoc.org/github.com/kainoaseto/go-cloud/runtimevar/filevar
[`constantvar`]: https://godoc.org/github.com/kainoaseto/go-cloud/runtimevar/constantvar

## Opening a Variable {#opening}

The first step in watching a variable is to instantiate a portable
[`*runtimevar.Variable`][] for your service.

The easiest way to do so is to use [`runtimevar.OpenVariable`][] and a service-specific URL pointing
to the variable, making sure you ["blank import"][] the driver package to link
it in.

```go
import (
	"github.com/kainoaseto/go-cloud/runtimevar"
	_ "github.com/kainoaseto/go-cloud/runtimevar/<driver>"
)
...
v, err := runtimevar.OpenVariable(context.Background(), "<driver-url>")
if err != nil {
    return fmt.Errorf("could not open variable: %v", err)
}
defer v.Close()
// v is a *runtimevar.Variable; see usage below
...
```

See [Concepts: URLs][] for general background and the [guide below][]
for URL usage for each supported service.

Alternatively, if you need fine-grained control
over the connection settings, you can call the constructor function in the
driver package directly (like `etcdvar.OpenVariable`).

```go
import "github.com/kainoaseto/go-cloud/runtimevar/<driver>"
...
v, err := <driver>.OpenVariable(...)
...
```

You may find the [`wire` package][] useful for managing your initialization code
when switching between different backing services.

See the [guide below][] for constructor usage for each supported service.

When opening the variable, you can provide a [decoder][] parameter (either as a
[query parameter][] for URLs, or explicitly to the constructor) to specify
whether the raw value stored in the variable is interpreted as a `string`, a
`[]byte`, or as JSON. Here's an example of using a JSON encoder:

{{< goexample src="github.com/kainoaseto/go-cloud/runtimevar.Example_jsonDecoder" imports="0" >}}

[`*runtimevar.Variable`]: https://godoc.org/github.com/kainoaseto/go-cloud/runtimevar#Variable
[`runtimevar.OpenVariable`]: https://godoc.org/github.com/kainoaseto/go-cloud/runtimevar#OpenVariable
["blank import"]: https://golang.org/doc/effective_go.html#blank_import
[Concepts: URLs]: {{< ref "/concepts/urls.md" >}}
[decoder]: https://godoc.org/github.com/kainoaseto/go-cloud/runtimevar#Decoder
[guide below]: {{< ref "#services" >}}
[query parameter]: https://godoc.org/github.com/kainoaseto/go-cloud/runtimevar#DecoderByName
[`wire` package]: http://github.com/google/wire

## Using a Variable {#using}

Once you have opened a `runtimevar.Variable` for the provider you want, you can
use it portably.

### Latest {#latest}

The easiest way to a `Variable` is to use the [`Variable.Latest`][] method. It
returns the latest good [`Snapshot`][] of the variable value, blocking if no
good value has *ever* been detected. The dynamic type of `Snapshot.Value`
depends on the decoder you provided when creating the `Variable`.

{{< goexample src="github.com/kainoaseto/go-cloud/runtimevar.ExampleVariable_Latest" imports="0" >}}

To avoid blocking, you can pass an already-`Done` context. You can also use
[`Variable.CheckHealth`][], which reports as healthy when `Latest` will
return a value without blocking.

[`Variable.Latest`]: https://godoc.org/github.com/kainoaseto/go-cloud/runtimevar#Variable.Latest
[`Variable.CheckHealth`]: https://godoc.org/github.com/kainoaseto/go-cloud/runtimevar#Variable.CheckHealth
[`Snapshot`]: https://godoc.org/github.com/kainoaseto/go-cloud/runtimevar#Snapshot

### Watch {#watch}

`Variable` also has a [`Watch`][] method for obtaining the value of a variable;
it has different semantics than `Latest` and may be useful in some scenarios. We
recommend starting with `Latest` as it's conceptually simpler to work with.

[`Watch`]: https://godoc.org/github.com/kainoaseto/go-cloud/runtimevar#Variable.Watch

## Other Usage Samples

* [CLI Sample](https://github.com/google/go-cloud/tree/master/samples/gocdk-runtimevar)
* [Guestbook sample](https://github.com/kainoaseto/go-cloud/tutorials/guestbook/)
* [runtimevar package examples](https://godoc.org/github.com/kainoaseto/go-cloud/runtimevar#pkg-examples)

## Supported Services {#services}

### GCP Runtime Configurator {#gcprc}

To open a variable stored in [GCP Runtime Configurator][] via a URL, you can use
the `runtimevar.OpenVariable` function as shown in the example below.

[GCP Runtime Configurator]: https://cloud.google.com/deployment-manager/runtime-configurator/

`runtimevar.OpenVariable` will use Application Default Credentials; if you have
authenticated via [`gcloud auth login`][], it will use those credentials. See
[Application Default Credentials][GCP creds] to learn about authentication
alternatives, including using environment variables.

[GCP creds]: https://cloud.google.com/docs/authentication/production
[`gcloud auth login`]: https://cloud.google.com/sdk/gcloud/reference/auth/login

{{< goexample
"github.com/kainoaseto/go-cloud/runtimevar/gcpruntimeconfig.Example_openVariableFromURL" >}}

#### GCP Constructor {#gcprc-ctor}

The [`gcpruntimeconfig.OpenVariable`][] constructor opens a Runtime Configurator
variable.

{{< goexample "github.com/kainoaseto/go-cloud/runtimevar/gcpruntimeconfig.ExampleOpenVariable" >}}

[`gcpruntimeconfig.OpenVariable`]: https://godoc.org/github.com/kainoaseto/go-cloud/runtimevar/gcpruntimeconfig#OpenVariable

### AWS Parameter Store {#awsps}

To open a variable stored in [AWS Parameter Store][] via a URL, you can use the
`runtimevar.OpenVariable` function as shown in the example below.

[AWS Parameter Store]:
https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-parameter-store.html

`runtimevar.OpenVariable` will create a default AWS Session with the
`SharedConfigEnable` option enabled; if you have authenticated with the AWS CLI,
it will use those credentials. See [AWS Session][] to learn about authentication
alternatives, including using environment variables.

[AWS Session]: https://docs.aws.amazon.com/sdk-for-go/api/aws/session/

{{< goexample
"github.com/kainoaseto/go-cloud/runtimevar/awsparamstore.Example_openVariableFromURL" >}}

#### AWS Constructor {#awsps-ctor}

The [`awsparamstore.OpenVariable`][] constructor opens a Parameter Store
variable.

{{< goexample "github.com/kainoaseto/go-cloud/runtimevar/awsparamstore.ExampleOpenVariable" >}}

[`awsparamstore.OpenVariable`]:
https://godoc.org/github.com/kainoaseto/go-cloud/runtimevar/awsparamstore#OpenVariable

### etcd {#etcd}

To open a variable stored in [etcd][] via URL, you can use the
`runtimevar.OpenVariable` function as follows.

{{< goexample "github.com/kainoaseto/go-cloud/runtimevar/etcdvar.Example_openVariableFromURL" >}}

[etcd]: https://etcd.io/

#### etcd Constructor {#etcd-ctor}

The [`etcdvar.OpenVariable`][] constructor opens an `etcd` variable.

[`etcdvar.OpenVariable`]:
https://godoc.org/github.com/kainoaseto/go-cloud/runtimevar/etcdvar#OpenVariable

{{< goexample "github.com/kainoaseto/go-cloud/runtimevar/etcdvar.ExampleOpenVariable" >}}

### HTTP {#http}

`httpvar` supports watching a variable via an HTTP request. Use
`runtimevar.OpenVariable` with a regular URL starting with `http` or `https`.
`httpvar` will periodically make an HTTP `GET` request to that URL, with the
`decode` URL parameter removed (if present).

{{< goexample "github.com/kainoaseto/go-cloud/runtimevar/httpvar.Example_openVariableFromURL" >}}

#### HTTP Constructor {#http-ctor}

The [`httpvar.OpenVariable`][] constructor opens a variable with a `http.Client`
and a URL.

{{< goexample "github.com/kainoaseto/go-cloud/runtimevar/httpvar.ExampleOpenVariable" >}}

[`httpvar.OpenVariable`]: https://godoc.org/github.com/kainoaseto/go-cloud/runtimevar/httpvar#OpenVariable

### Blob {#blob}

`blobvar` supports watching a variable based on the contents of a
[Go CDK blob][]. Set the environment variable `BLOBVAR_BUCKET_URL` to the URL
of the bucket, and then use `runtimevar.OpenVariable` as shown below.
`blobvar` will periodically re-fetch the contents of the blob.

{{< goexample "github.com/kainoaseto/go-cloud/runtimevar/blobvar.Example_openVariableFromURL" >}}

[Go CDK blob]: https://github.com/kainoaseto/go-cloud/howto/blob/

You can also use [`blobvar.OpenVariable`][].

[`blobvar.OpenVariable`]: https://godoc.org/github.com/kainoaseto/go-cloud/runtimevar/blobvar#OpenVariable

### Local {#local}

You can create an in-memory variable (useful for testing) using `constantvar`:

{{< goexample "github.com/kainoaseto/go-cloud/runtimevar/constantvar.Example_openVariableFromURL" >}}

Alternatively, you can create a variable based on the contents of a file using
`filevar`:

{{< goexample "github.com/kainoaseto/go-cloud/runtimevar/filevar.Example_openVariableFromURL" >}}
