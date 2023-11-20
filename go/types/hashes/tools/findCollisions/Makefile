build:
	go build -o find_collisions find_collisions.go

install:
	sudo mv find_collisions /usr/local/bin/

run80:
	echo "starting all segments"
	nohup find_collisions -Segment 00 -SegmentCount 8 -Seed 10000000 > log.00.txt &
	nohup find_collisions -Segment 80 -SegmentCount 8 -Seed 10000000 > log.80.txt &
	nohup find_collisions -Segment 160 -SegmentCount 8 -Seed 10000000 > log.160.txt &
	nohup find_collisions -Segment 240 -SegmentCount 8 -Seed 10000000 > log.240.txt &
	nohup find_collisions -Segment 320 -SegmentCount 8 -Seed 10000000 > log.320.txt &
	nohup find_collisions -Segment 400 -SegmentCount 8 -Seed 10000000 > log.400.txt &
	nohup find_collisions -Segment 480 -SegmentCount 8 -Seed 10000000 > log.480.txt &
	nohup find_collisions -Segment 560 -SegmentCount 8 -Seed 10000000 > log.560.txt &
	echo "started..."
