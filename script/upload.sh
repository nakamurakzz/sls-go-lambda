#!/bin/bash
BUCKET="dev-sls-s3-event-lambda"

# 対話によってAWS_PROFILEを変更する
echo "Input AWS_PROFILE"
read AWS_PROFILE

for i in {1..10}; do
  echo "Uploading file for loop $i"
  echo "Loop: $i" > test.txt
  aws s3 cp test.txt s3://${BUCKET}/hello/test_loop_${i}.txt --profile ${AWS_PROFILE}
  sleep 1
done