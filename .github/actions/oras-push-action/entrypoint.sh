#!/bin/sh -l

gzip -k tektoncd/task.yaml
oras push --manifest-annotations tektoncd/manifest-annotations.json "$1" tektoncd/task.yaml.gz --username "$2" --password "$3"
echo "::set-output name=test::hello"
