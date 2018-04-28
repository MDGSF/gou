#!/bin/bash

echo "nanosecods   =" $(date +%s%N)
echo "microsecods  =" $[$(date +%s%N)/1000]
echo "milliseconds =" $[$(date +%s%N)/1000000]
echo "seconds      =" $[$(date +%s%N)/1000000000]

