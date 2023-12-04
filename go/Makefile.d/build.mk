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
		  build/go/tools/color \
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
	@color -green -lf "$@ OK"

build/go/ansi:
	@color -yellow -lf "##$@ not implemented"

build/go/communications:
	@color -yellow -lf "##$@ not implemented"

build/go/convert:
	@color -yellow -lf "##$@ not implemented"

build/go/counters:
	@color -yellow -lf "##$@ not implemented"

build/go/crsce:
	@color -yellow -lf "##$@ not implemented"

build/go/db/dbMigrations:
	@color -blue -lf -reset ">>start $@"
	@go build -o build/dbMigrations go/db/dbMigrations/main.go || {\
  		color -red -lf -reset ">>>failed $@";\
	};\
	color -green -lf "<<complete $@"

build/go/db/postgres:
	@color -yellow -lf "##$@ not implemented"

build/go/environment:
	@color -yellow -lf "##$@ not implemented"

build/go/exit:
	@color -yellow -lf "##$@ not implemented"

build/go/experiments:
	@color -yellow -lf "##$@ not implemented"

build/go/fs/file:
	@color -yellow -lf "##$@ not implemented"

build/go/fs/directory:
	@color -yellow -lf "##$@ not implemented"

build/go/keyvalue:
	@color -yellow -lf "##$@ not implemented"

build/go/lists:
	@color -yellow -lf "##$@ not implemented"

build/go/lock:
	@color -yellow -lf "##$@ not implemented"

build/go/logging:
	@color -yellow -lf "##$@ not implemented"

build/go/metrics:
	@color -yellow -lf "##$@ not implemented"

build/go/misc:
	@color -yellow -lf "##$@ not implemented"

build/go/moremath:
	@color -yellow -lf "##$@ not implemented"

build/go/net/calculateSubnets:
	@color -blue -lf ">>start $@"
	@go build -o build/calculateSubnets go/net/calculateSubnets/main.go
	@color -green -lf "<<complete $@"

build/go/runcommand:
	@color -yellow -lf "##$@ not implemented"

build/go/semver:
	@color -yellow -lf "##$@ not implemented"

build/go/sets:
	@color -yellow -lf "##$@ not implemented"

build/go/simpleArgs:
	@color -yellow -lf "##$@ not implemented"

build/go/simpleLogger:
	@color -yellow -lf "##$@ not implemented"

build/go/systemrecon:
	@color -yellow -lf "##$@ not implemented"

build/go/testing:
	@color -yellow -lf "##$@ not implemented"

build/go/tools/badgeMaker:
	@color -blue -lf ">>start $@"
	@go build -o build/badgeMaker go/tools/badgeMaker/main.go
	@color -green -lf "<<complete $@"

build/go/tools/color:
	@color -blue -lf ">>start $@"
	@go build -o build/color go/tools/color/main.go
	@color -green -lf "<<complete $@"

build/go/tools/numberCpuCores:
	@color -blue -lf ">>start $@"
	@go build -o build/numCpuCores go/tools/numberCpuCores/main.go
	@color -green -lf "<<complete $@"

build/go/tools/what:
	@color -blue -lf ">>start $@"
	@go build -o build/what go/tools/what/main.go
	@color -green -lf "<<complete $@"

build/go/trees:
	@color -yellow -lf "##$@ not implemented"

build/go/types/bitBlock:
	@color -yellow -lf "##$@ not implemented"

build/go/types/generic:
	@color -yellow -lf "##$@ not implemented"

build/go/types/hashes/tools/findCollisions:
	@color -blue -lf ">>start $@"
	@go build -o build/findCollisions go/types/hashes/tools/findCollisions/main.go
	@color -green -lf "<<complete $@"

build/go/types/hashes/tools/preCalculateHash:
	@color -blue -lf ">>start $@"
	@go build -o build/preCalculateHash go/types/hashes/tools/preCalculateHash/main.go
	@color -green -lf "<<complete $@"

build/go/types/hashes/tools/searchHashes:
	@color -blue -lf ">>start $@"
	@go build -o build/searchHashes go/types/hashes/tools/searchHashes/main.go
	@color -green -lf "<<complete $@"

build/go/version:
	@color -yellow -lf "##$@ not implemented"

build/go/wrappers:
	@color -yellow -lf "##$@ not implemented"
