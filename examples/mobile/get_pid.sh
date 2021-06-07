#!/bin/bash
adb shell pidof -s dev.fragoso.mustard
while [ $? -ne 0 ]; do
  adb shell pidof -s dev.fragoso.mustard
done
