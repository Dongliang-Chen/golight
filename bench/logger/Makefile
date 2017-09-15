GOTEST_FLAGS=-cpu=4 -benchmem -benchtime=5s
#1,2,
TEXT_PKGS=Zerolog GolightLogging
JSON_PKGS=Zerolog GolightLogging
#TEXT_PKGS= GolightLogging
#JSON_PKGS= GolightLogging

TEXT_PKG_TARGETS=$(addprefix test-text-,$(TEXT_PKGS))
JSON_PKG_TARGETS=$(addprefix test-json-,$(JSON_PKGS))

.PHONY: all deps test test-text test-json $(TEXT_PKG_TARGETS) $(JSON_PKG_TARGETS)

all: deps test

deps:
	go get -u github.com/rs/zerolog
	go get -u github.com/dlmc/golight/logger

test: test-text test-json

test-text: $(TEXT_PKG_TARGETS)

$(TEXT_PKG_TARGETS): test-text-%:
	go test $(GOTEST_FLAGS) -bench "$*.*Text"

test-json: $(JSON_PKG_TARGETS)

$(JSON_PKG_TARGETS): test-json-%:
	go test $(GOTEST_FLAGS) -bench "$*.*JSON"