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

MANAGED_ASSETS=../config/install \
	../README.md \
	../Makefile \
	../Dockerfile \
	../go.mod

# preserve manually mananged assets
preserve:
	@mkdir -p assets; \
	for ASSET in $(MANAGED_ASSETS); do \
		if [ -f $$ASSET ]; then \
			echo "preserving file $$ASSET"; \
			cp -R $$ASSET asset; \
		elif [ -d $$ASSET ]; then \
			mkdir -p assets/$$ASSET; \
			echo "preserving directory $$ASSET"; \
			cp -R $$ASSET/* asset/$$ASSET; \
		else \
			echo "not preserving non-existing asset $$ASSET"; \
		fi \
	done

# restore manually managed assets
restore:
	@for ASSET in $(MANAGED_ASSETS); do \
		ASSET=`echo $$ASSET | sed 's/\.\.\///g'`; \
		if [ -e assets/$$ASSET ]; then \
			echo "restoring asset $$ASSET"; \
			cp -R assets/$$ASSET ..; \
		else \
			echo "not restoring non-existing asset $$ASSET"; \
		fi \
	done

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
