#!/bin/bash
scriptdir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
basedir=$(echo "${scriptdir}" | grep -Po ".*(?=\/)")
tempdir="${basedir}/tmp"

for u in {00..03}; do
    for i in {00..03}; do
        snapfol="${tempdir}/testdata/user${u}/repo${i}/snapshots"
        for i in {00..05}; do
            tar="${snapfol}/${i}"
            mkdir -p "${snapfol}"
            echo tempfle >"${tar}"
            dat="$(date --date="-${i} day" "+%Y-%m-%d %H:%M:%S.%N")"
            touch -m -d "${dat}" "${tar}"
        done
    done
done
