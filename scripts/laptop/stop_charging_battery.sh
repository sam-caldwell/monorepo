#!/bin/bash -e

echo 30 > /sys/class/power_supply/BAT0/charge_start_threshold
echo 80 > /sys/class/power_supply/BAT0/charge_stop_threshold
