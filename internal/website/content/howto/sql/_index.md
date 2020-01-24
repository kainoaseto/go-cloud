---
title: "MySQL/PostgreSQL"
date: 2019-06-21T09:26:56-07:00
showInSidenav: true
toc: true
---

Connecting to Cloud providers' hosted database services requires additional
steps to ensure the security of the connection. For example, each of GCP,
AWS and Azure require the use of custom certificate authorities to be
configured in the client. GCP requires a custom proxy with authentication
credentials. The Go CDK makes opening these connections easier while still
using the standard [`*sql.DB`][] type.

[`*sql.DB`]: https://godoc.org/database/sql#DB

<!--more-->

## Local or On-Premise {#local}

The Go CDK uses the same [URL opener pattern][] as seen in other Go CDK APIs. It
differs from the standard library's `sql.Open` in that it automatically
instruments the connection with [OpenCensus metrics][].

The portable function for MySQL is [`mysql.Open`][]:

{{< goexample "github.com/kainoaseto/go-cloud/mysql.ExampleOpen" >}}

And the portable function for PostgreSQL is [`postgres.Open`][]:

{{< goexample "github.com/kainoaseto/go-cloud/postgres.ExampleOpen" >}}

[`mysql.Open`]: https://godoc.org/github.com/kainoaseto/go-cloud/mysql#Open
[OpenCensus metrics]: https://opencensus.io/integrations/sql/go_sql/
[`postgres.Open`]: https://godoc.org/github.com/kainoaseto/go-cloud/postgres#Open
[URL opener pattern]: {{< ref "/concepts/urls.md" >}}

## GCP {#gcp}

Users of [GCP Cloud SQL for MySQL][] should import the `github.com/kainoaseto/go-cloud/mysql/gcpmysql` package:

{{< goexample "github.com/kainoaseto/go-cloud/mysql/gcpmysql.Example" >}}

Users of [GCP Cloud SQL for PostgreSQL][] should import the `github.com/kainoaseto/go-cloud/mysql/gcppostgres` package:

{{< goexample "github.com/kainoaseto/go-cloud/postgres/gcppostgres.Example" >}}

[GCP Cloud SQL for MySQL]: https://cloud.google.com/sql/docs/mysql/
[GCP Cloud SQL for PostgreSQL]: https://cloud.google.com/sql/docs/postgres/

## AWS {#aws}

Users of [AWS RDS for MySQL][] should import the `github.com/kainoaseto/go-cloud/mysql/awsmysql` package:

{{< goexample "github.com/kainoaseto/go-cloud/mysql/awsmysql.Example" >}}

Users of [AWS RDS for PostgreSQL][] should import the `github.com/kainoaseto/go-cloud/postgres/awspostgres` package:

{{< goexample "github.com/kainoaseto/go-cloud/postgres/awspostgres.Example" >}}

[AWS RDS for MySQL]: https://aws.amazon.com/rds/mysql/
[AWS RDS for PostgreSQL]: https://aws.amazon.com/rds/postgresql/

## Azure {#azure}

Users of [Azure Database for MySQL][] should import the `github.com/kainoaseto/go-cloud/mysql/azuremysql` package:

{{< goexample "github.com/kainoaseto/go-cloud/mysql/azuremysql.Example" >}}

[Azure Database for MySQL]: https://azure.microsoft.com/en-us/services/mysql/

## Other Usage Samples

* [Guestbook sample](https://github.com/kainoaseto/go-cloud/tutorials/guestbook/)
