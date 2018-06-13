space := $(empty) $(empty)
comma := ,
empty :=

import_path = k8s.io/helm/pkg/proto/hapi

dst 	= ../terminator/pkg
target  = python
plugins = grpc

chart_ias    = $(subst $(space),$(comma),$(addsuffix =$(import_path)/$(chart_pkg),$(addprefix M,$(chart_pbs))))
chart_pbs    = $(sort $(wildcard hapi/chart/*.proto))
chart_pkg    = chart

release_ias  = $(subst $(space),$(comma),$(addsuffix =$(import_path)/$(release_pkg),$(addprefix M,$(release_pbs))))
release_pbs  = $(sort $(wildcard hapi/release/*.proto))
release_pkg  = release

services_ias = $(subst $(space),$(comma),$(addsuffix =$(import_path)/$(services_pkg),$(addprefix M,$(services_pbs))))
services_pbs = $(sort $(wildcard hapi/services/*.proto))
services_pkg = services

version_ias    = $(subst $(space),$(comma),$(addsuffix =$(import_path)/$(version_pkg),$(addprefix M,$(version_pbs))))
version_pbs    = $(sort $(wildcard hapi/version/*.proto))
version_pkg    = version

google_deps	 = Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,Mgoogle/protobuf/any.proto=github.com/golang/protobuf/ptypes/any

.PHONY: all
all: chart release services version

chart:
	python -m grpc_tools.protoc -I. --$(target)_out=$(dst) --grpc_$(target)_out=$(dst) $(chart_pbs)

release:
	python -m grpc_tools.protoc -I. --$(target)_out=$(dst) --grpc_python_out=$(dst) $(release_pbs)

services:
	python -m grpc_tools.protoc -I. --$(target)_out=$(dst) --grpc_python_out=$(dst) $(services_pbs)

version:
	python -m grpc_tools.protoc -I. --$(target)_out=$(dst) --grpc_python_out=$(dst) $(version_pbs)

.PHONY: clean
clean:
	@rm -rf $(dst)/hapi 2>/dev/null
