build/go: build/go/ansi \
		  build/go/communications \
		  build/go/convert \
		  build/go/counters \
		  build/go/crsce \
		  build/go/db/dbMigrations \
		  build/go/db/postgres \
		  build/go/environment \
		  build/go/exit \
		  build/go/experiments \
		  build/go/fs/file \
		  build/go/fs/directory \
		  build/go/keyvalue \
		  build/go/lists \
		  build/go/lock \
		  build/go/logging \
		  build/go/metrics \
		  build/go/misc \
		  build/go/moremath \
		  build/go/net/calculateSubnets \
		  build/go/runcommand \
		  build/go/sets \
		  build/go/simpleArgs \
		  build/go/simpleLogger \
		  build/go/systemrecon \
		  build/go/testing \
		  build/go/tools/badgeMaker \
		  build/go/tools/numberCpuCores \
		  build/go/tools/what \
		  build/go/trees \
		  build/go/types/bitBlock \
		  build/go/types/generic \
		  build/go/types/hashes/tools/findCollisions \
		  build/go/types/hashes/tools/preCalculateHash \
		  build/go/types/hashes/tools/searchHashes \
		  build/go/version \
		  build/go/wrappers
	@echo ${GREEN}"$@ OK"${RESET}

build/go/ansi:
	@echo ${YELLOW}"--$@ not implemented"${RESET}

build/go/communications:
	@echo ${YELLOW}"--$@ not implemented"${RESET}

build/go/convert:
	@echo ${YELLOW}"--$@ not implemented"${RESET}

build/go/counters:
	@echo ${YELLOW}"--$@ not implemented"${RESET}

build/go/crsce:
	@echo ${YELLOW}"--$@ not implemented"${RESET}

build/go/db/dbMigrations:
	@echo ${BLUE}">>start $@"${RESET}
	go build -o build/dbMigrations go/db/dbMigrations/main.go
	@echo ${GREEN}"<<complete $@"${RESET}

build/go/db/postgres:
	@echo ${YELLOW}"--$@ not implemented"${RESET}

build/go/environment:
	@echo ${YELLOW}"--$@ not implemented"${RESET}

build/go/exit:
	@echo ${YELLOW}"--$@ not implemented"${RESET}

build/go/experiments:
	@echo ${YELLOW}"--$@ not implemented"${RESET}

build/go/fs/file:
	@echo ${YELLOW}"--$@ not implemented"${RESET}

build/go/fs/directory:
	@echo ${YELLOW}"--$@ not implemented"${RESET}

build/go/keyvalue:
	@echo ${YELLOW}"--$@ not implemented"${RESET}

build/go/lists:
	@echo ${YELLOW}"--$@ not implemented"${RESET}

build/go/lock:
	@echo ${YELLOW}"--$@ not implemented"${RESET}

build/go/logging:
	@echo ${YELLOW}"--$@ not implemented"${RESET}

build/go/metrics:
	@echo ${YELLOW}"--$@ not implemented"${RESET}

build/go/misc:
	@echo ${YELLOW}"--$@ not implemented"${RESET}

build/go/moremath:
	@echo ${YELLOW}"--$@ not implemented"${RESET}

build/go/net/calculateSubnets:
	@echo ${BLUE}">>start $@"${RESET}
	@go build -o build/calculateSubnets go/net/calculateSubnets/main.go
	@echo ${GREEN}"<<complete $@"${RESET}

build/go/runcommand:
	@echo ${YELLOW}"--$@ not implemented"${RESET}

build/go/semver:
	@echo ${YELLOW}"--$@ not implemented"${RESET}

build/go/sets:
	@echo ${YELLOW}"--$@ not implemented"${RESET}

build/go/simpleArgs:
	@echo ${YELLOW}"--$@ not implemented"${RESET}

build/go/simpleLogger:
	@echo ${YELLOW}"--$@ not implemented"${RESET}

build/go/systemrecon:
	@echo ${YELLOW}"--$@ not implemented"${RESET}

build/go/testing:
	@echo ${YELLOW}"--$@ not implemented"${RESET}

build/go/tools/badgeMaker:
	@echo ${BLUE}">>start $@"${RESET}
	@go build -o build/badgeMaker go/tools/badgeMaker/main.go
	@echo ${GREEN}"<<complete $@"${RESET}

build/go/tools/numberCpuCores:
	@echo ${BLUE}">>start $@"${RESET}
	@go build -o build/numCpuCores go/tools/numberCpuCores/main.go
	@echo ${GREEN}"<<complete $@"${RESET}

build/go/tools/what:
	@echo ${BLUE}">>start $@"${RESET}
	@go build -o build/what go/tools/what/main.go
	@echo ${GREEN}"<<complete $@"${RESET}

build/go/trees:
	@echo ${YELLOW}"--$@ not implemented"${RESET}

build/go/types/bitBlock:
	@echo ${YELLOW}"--$@ not implemented"${RESET}

build/go/types/generic:
	@echo ${YELLOW}"--$@ not implemented"${RESET}

build/go/types/hashes/tools/findCollisions:
	@echo ${BLUE}">>start $@"${RESET}
	@go build -o build/findCollisions go/types/hashes/tools/findCollisions/main.go
	@echo ${GREEN}"<<complete $@"${RESET}

build/go/types/hashes/tools/preCalculateHash:
	@echo ${BLUE}">>start $@"${RESET}
	@go build -o build/preCalculateHash go/types/hashes/tools/preCalculateHash/main.go
	@echo ${GREEN}"<<complete $@"${RESET}

build/go/types/hashes/tools/searchHashes:
	@echo ${BLUE}">>start $@"${RESET}
	@go build -o build/searchHashes go/types/hashes/tools/searchHashes/main.go
	@echo ${GREEN}"<<complete $@"${RESET}

build/go/version:
	@echo ${YELLOW}"--$@ not implemented"${RESET}

build/go/wrappers:
	@echo ${YELLOW}"--$@ not implemented"${RESET}
