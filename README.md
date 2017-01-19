# Description

Utlity which adds arbitrary distance to existing .tcx file. For example, you use stationary cycle trainer which has no wireless sensors, but you have separate cadense/heartrate sensors. You may record your training in strava/endomondo/wahoo/etc., and then use this utility to add distance ridden to recored track.

## Install

1. Get Go (`brew install go` on mac with HomeBrew)
2. `go get github.com/ivanzoid/fixIndoorTcx`
3. `go install github.com/ivanzoid/fixIndoorTcx`

## Usage

`fixIndoorTcx` `<tcxFile>` `<distanceInMeters>` `>` `<outputTcxFile>`

For example:

`fixIndoorTcx 20170115_144636.tcx 13320 > 20170115_144636-with-distance.tcx`

## Install
