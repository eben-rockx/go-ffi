DEPS:=rockxcrypto.h librockxcrypto.a

export CARGO_TARGET_DIR=target

all: $(DEPS)
.PHONY: all

clean:
	go clean -cache -testcache .
	rm -rf $(DEPS)
	cd rust && cargo clean && cd ..
.PHONY: clean

rockxcrypto:
	./install-rockxcrypto
.PHONY: rockxcrypto


shellcheck:
	shellcheck install-rockxcrypto


