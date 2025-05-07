#!/bin/bash

# This script is used to initialize the S3 bucket in LocalStack

# Exit if an error occurs
set -e

echo "=== Initializing S3 bucket in LocalStack ==="

# Make bucket
awslocal s3 mb s3://local-book-tracker

# Upload files
awslocal s3 cp --recursive /init-data/ s3://local-book-tracker/

# Check
echo "Checking S3 bucket..."
awslocal s3 ls || {
  echo "Failed to list buckets"
  exit 1
}

echo "Checking S3 bucket contents..."
awslocal s3 ls --recursive s3://local-book-tracker/ || {
  echo "Failed to list bucket contents"
  exit 1
}

echo "=== Completed initializing S3 bucket in LocalStack ==="