#!/bin/bash -e

echo 1 > /sys/class/power_supply/BAT0/charge_start_threshold
echo 100 > /sys/class/power_supply/BAT0/charge_stop_threshold
