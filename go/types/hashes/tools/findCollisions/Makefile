build:
	go build -o findCollisions findCollisions.go
	shasum -a 256 findCollisions

install:
	sudo mv findCollisions /usr/local/bin/

stop:
	( \
		for i in $$(ps aux | grep findCollisions | grep -v grep | awk '{print $$2}'); do echo "killing $$i"; kill -9 $$i; done;\
	)

run80:
	echo "starting all segments"
	nohup findCollisions -Segment 00 -SegmentCount 4 -Seed 10000000 > log.00.txt &
	nohup findCollisions -Segment 80 -SegmentCount 4 -Seed 10000000 > log.80.txt &
	nohup findCollisions -Segment 160 -SegmentCount 4 -Seed 10000000 > log.160.txt &
	nohup findCollisions -Segment 240 -SegmentCount 4 -Seed 10000000 > log.240.txt &
	nohup findCollisions -Segment 320 -SegmentCount 4 -Seed 10000000 > log.320.txt &
	echo "started..."
	for PID in $(pgrep findCollisions); do ls /proc/$PID/task | xargs renice -20; done
	echo "re-nice done..."

run16:
	echo "starting all segments"
	nohup findCollisions -Segment 00 -SegmentCount 2 -Seed 00 > log.00.txt &
	echo "started..."
	for PID in $(pgrep findCollisions); do ls /proc/$PID/task | xargs renice -20; done
	echo "re-nice done..."

version: build install
	findCollisions --version
