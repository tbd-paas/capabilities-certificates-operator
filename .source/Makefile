VERSION ?= latest

#
# code generation
#
INIT_OPTS=init \
	--workload-config .source/workload.yaml \
   	--repo github.com/tbd-paas/capabilities-certificates-operator \
	--controller-image quay.io/tbd-paas/certificates-operator:$(VERSION) \
    --skip-go-version-check

CREATE_OPTS=create api \
	--workload-config .source/workload.yaml \
	--controller \
	--resource

# preserve manually mananged assets
preserve:
	cp -R ../config/install assets/
	cp ../README.md assets/
	cp ../Makefile assets/
	cp ../Dockerfile assets/
	cp ../go.mod assets/

# restore manually managed assets
restore:
	cp -R assets/install ../config
	cp assets/README.md ../
	cp assets/Makefile ../
	cp assets/Dockerfile ../
	cp assets/go.mod ../

operator-clean: preserve
	rm -rf ../*

operator-init:
	cd .. && operator-builder $(INIT_OPTS)

operator-create:
	cd .. && operator-builder $(CREATE_OPTS)

#
# project management
#
PROJECT ?=
project:
	@mkdir -p $(PROJECT)/config $(PROJECT)/vendor $(PROJECT)/manifests
	@touch $(PROJECT)/config/values.yaml $(PROJECT)/config/vendor.yaml $(PROJECT)/config/overlays.yaml

download:
	vendir sync --file $(PROJECT)/config/vendor.yaml --lock-file $(PROJECT)/config/vendor.yaml.lock

overlay:
	yot -I 2 -i $(PROJECT)/config/overlays.yaml -o $(PROJECT) -f $(PROJECT)/config/values.yaml
