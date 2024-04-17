# Support Services Operator

A Kubernetes operator to manage platform support services.

A Kubernetes cluster needs support services for tenant workloads.  The support
services operator exists to facilitate installation of these components.

## Quickstart

### Install

Install the support services operator:

```bash
kubectl apply -f config/samples/install
```

### Create Support Services

Set up the support services collection.  This resource logically associates the
support services components:

```bash
kubectl apply -f config/samples/setup_v1alpha1_supportservices.yaml
```

Install the certificates component which consists of cert-manager:

```bash
kubectl apply -f config/samples/platform_v1alpha1_certificatescomponent.yaml
```

Install the ingress component which, by default, consists of the kong ingress
controller and external DNS:

```bash
kubectl apply -f config/samples/platform_v1alpha1_ingresscomponent.yaml
```

## Local Development & Testing

To install the custom resource/s for this operator, make sure you have a
kubeconfig set up for a test cluster, then run:

    make install

To run the controller locally against a test cluster:

    make run

You can then test the operator by creating the sample manifest/s:

    kubectl apply -f config/samples

To clean up:

    make uninstall

### Common Changes

Following is a cheat sheet of operations for common changes to the operator:

In `.operator-builder` directory:

1. Make updates to yot overlays for project being changed.
2. Run `make overlays` and check the output manifests match your expectations.
3. Run `make operator-clean && make operator-init && make operator-create` to
   update the operator source code.
4. Run `make restore` to restore the static assets that were preserved, e.g.
   README.

In root directory:

1. Run `make install` or `make manifests` to generate CRD manifests.  (Check
   go.sum for changes if missing go.sum entry errors occur.)
2. Run `make deploy` to update deployment kustomize overlays.
3. If releasing a new version of support-services-operator, update the image
   version in `config/install/support-services-operator.yaml`.
4. If CRD updates were made, run `make crd-static-manifest` to update the static
   install CRD manifest.

In `.operator-builder` directory:

1. Run `make preserve` to preserve static assets.

In root directory:

1. Add and commit changes
2. If releasing a new version of support-services-operator, tag the commit.
3. Push changes.

## Deploy the Controller Manager

First, set the image:

    export IMG=myrepo/myproject:v0.1.0

Now you can build and push the image:

    make docker-build
    make docker-push

Then deploy:

    make deploy

To clean up:

    make undeploy

## Companion CLI

To build the companion CLI:

    make build-cli

The CLI binary will get saved to the bin directory.  You can see the help
message with:

    ./bin/ssctl help
