# Kubernetes Generic Service Broker

![](https://github.com/spjmurray/service-broker/workflows/Build%20and%20Test/badge.svg)

Open Service Broker API driven templating engine for Kubernetes.
The Kubernetes Service Broker conforms to the [Open Service Broker Specification](https://github.com/openservicebrokerapi/servicebroker/blob/v2.13/spec.md) version 2.13.

## What are Service Brokers?

Open Service Brokers are an abstraction layer around a service that allow users to create service instances and bind applications to them.
Service instances are abstractions of a service that are controlled by a centralized authority.
This leads to a number of benefits for application developers:

* Service instances are under centralized administrative control
  * Fine graned control over what service instances can be provisioned
    * Security compliance
    * Resource constraining
    * Cost constraining
  * Single pane of glass monitoring and reporting
* Service instances are abstractions
  * No domain specific knowledge necessary to provision and manage services for application developers
  * Consume URIs and credentials
* Enhanced security
  * Elevated privileges to provision service instances are delegated to the broker

## What does the Service Broker do?

The Kubernetes Service Broker was originally designed to deploy databases.
Through evolution, it was possible to abstract away all the domain specific knowledge and provide a generic service broker implmenetation, that still supported our original goals.

### Security Model

The Service Broker is designed to be used with the [Kubernetes Service Catalog](https://kubernetes.io/docs/concepts/extend-kubernetes/service-catalog/) which provides Kubernetes native bindings in the form of `ServiceInstance` and `ServiceBinding` resources.
Using Kubernetes RBAC controls, platorm administrators can control precicely what users can provision and where.
Due to how the service catalog works, a binding must reside in the same namespace as the service instance.
This supports:

* Self-service
  * Users can provision both service instances and bind their applications to them
* Shared services
  * Administrators can provision service instances, and users can bind to and consume them 

The Service Broker is flexible enough so that resources created to realize a service instance can be located in the same namespace as the service instance resource, or in a hard coded namespace:

* Namespaced service instances
  * Users may be able to see, and modify, underlying resources, depending on RBAC rules
* Hard-coded namespaced service instances
  * Underlying resources are hidden from users, thus protecting sensitive configuration

The Service Broker may be deployed in its own namespace to keep its own configuration and runtime data separate and secured from other users.
Depending on how you wish to configure the Service Broker, it may only require permissions to create resources in its own namespace, or if provisioning resources in other namespaces, cluster wide permissions.

## Building

### Building an Official Container from Release Archives

Official releases are avaliable to download from [GitHub](https://github.com/spjmurray/service-broker/releases).
They contain the service broker binary, an example docker file and example YAML.

Download the package, decompress it, then from the root director of the archive you can build the container image:

```bash
$ sudo tar xf -C / couchbase-service-broker-0.0.0-99999.tar.gz
$ cd /usr/local/share/couchbase-service-broker
$ docker build . -t couchbase/service-broker:0.0.0
```

This will create the container image `couchbase/service-broker:0.0.0`.

### Building A Container Image from Source

To build a container from source you can use the following command:

```bash
$ make container -e APPLICATION=acme-service-broker IMAGE=acme/service-broker VERSION=1.0.0 REVISION=beta1
```

This allows you to change the application and image's name and the version.
This will require modification to the example files.

### Building a Release from Source

To build a release from source:

```bash
$ make archive -e APPLICATION=acme-service-broker VERSION=1.0.0 REVISION=beta1 PREFIX=/usr DESTDIR=/tmp/archive
```

Or for Red Hat RPMs:

```bash
$ make rpm -e APPLICATION=acme-service-broker VERSION=1.0.0 REVISION=beta1 PREFIX=/usr
```

Or for debian DEBs:

```bash
$ make deb -e APPLICATION=acme-service-broker VERSION=1.0.0 REVISION=beta1 PREFIX=/usr
```

## Installation

Ensure the [Kubernetes Service Catalog is installed](https://svc-cat.io/docs/install/).

Change to the install directory:

```bash
$ cd /usr/local/share/couchbase-service-broker
```

Install the custom resource definition:

```bash
$ kubectl create -f crds
```

Select a configuration template to use.
These define the permissions that are required by the service broker to deploy the service instances as defined in the configuration:

```bash
$ kubectl create -f example/configurations/couchbase-server/broker.yaml
```

Install the service broker, ensuring the service broker deployment is running:

```bash
$ kubectl create -f example/broker.yaml
$ kubectl wait --for=condition=Available deployment/couchbase-service-broker
```

Register the service broker with the service catalog, ensuring it is ready:

```bash
$ kubectl create -f example/clusterservicebroker.yaml
$ svcat get brokers
```

Finally you can test the broker configuration by creating a service instance:

```bash
$ kubectl create -f example/configurations/couchbase-server/serviceinstance.yaml
```

And get access to a secret containing credentials:

```bash
$ kubectl create -f example/configurations/couchbase-server/servicebinding.yaml
```

## Architecture

### Templating Engine

The core of the Service Broker is a flexible and generic templating engine.
A service instance or binding is conceptually an ordered list of templates of Kubernetes resources.

Upon creation of an instance the templates are first rendered to apply dynamic configuration from both the environment and the request.
All template rendering operations are carried out in JSON, using [JSON Pointer](https://tools.ietf.org/html/rfc6902) and [JSON Patch](https://tools.ietf.org/html/rfc6902) operations.
Once rendered the resources are then committed to the Kubernetes API.

## Contributing

### Testing

All code submissions must include sufficient tests to check correctness.
All tests must pass, and do so consistently.
Tests can be run with the following command:

```bash
$ make test
```

You can run individual tests or groups of tests while debugging with the following command:

```bash
$ go test -v -race ./test -run TestConnect -args -logtostderr -v 1
```

Code coverage is run as part of the test command and -- although not enforced, it is watched -- should be checked:

```bask
$ make cover
```

Any code that is added (and not auto-generated) must be covered by testing.
