build: build/findCollisions \
	   build/preCalculateHash
	   @echo "ok"


build/findCollisions:
	go build -o findCollisions.bin findCollisions/findCollisions.go
	shasum -a 256 findCollisions.bin

build/preCalculateHash:
	go build -o preCalculateHash.bin preCalculateHash/preCalculateHash.go
	shasum -a 256 preCalculateHash.bin

install:
	sudo mv findCollisions.bin /usr/local/bin/
	sudo mv generatePrecalculated.bin /usr/local/bin/
	sudo mv preCalculateHash.bin /usr/local/bin/

stop:
	( \
		for i in $$(ps aux | grep findCollisions.bin | grep -v grep | awk '{print $$2}'); do echo "killing $$i"; kill -9 $$i; done;\
	)

run80:
	echo "starting all segments"
	nohup findCollisions.bin -Segment 00 -SegmentCount 4 -Seed 10000000 > log.00.txt &
	nohup findCollisions.bin -Segment 80 -SegmentCount 4 -Seed 10000000 > log.80.txt &
	nohup findCollisions.bin -Segment 160 -SegmentCount 4 -Seed 10000000 > log.160.txt &
	nohup findCollisions.bin -Segment 240 -SegmentCount 4 -Seed 10000000 > log.240.txt &
	nohup findCollisions.bin -Segment 320 -SegmentCount 4 -Seed 10000000 > log.320.txt &
	echo "started..."
	(\
		for PID in $(pgrep findCollisions.bin); do ls /proc/$PID/task | xargs sudo renice -20; done; \
	)
	echo "re-nice done..."

run16:
	echo "starting all segments"
	nohup findCollisions -Segment 00 -SegmentCount 2 -Seed 00 > log.00.txt &
	echo "started..."
	for PID in $(pgrep findCollisions); do ls /proc/$PID/task | xargs renice -20; done
	echo "re-nice done..."

version: build install
	findCollisions.bin --version
