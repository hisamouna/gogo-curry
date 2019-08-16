#!/bin/bash

NOW_DATE=$(date "+%Y%m%d%H%M")

echo $NOW_DATE

git tag $NOW_DATE

git push origin $NOW_DATE
