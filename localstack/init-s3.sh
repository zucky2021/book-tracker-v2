#!/bin/bash

# This script is used to initialize the S3 bucket in LocalStack

echo "=== Initializing S3 bucket in LocalStack ==="

# Make bucket
awslocal s3 mb s3://local-book-tracker

# Upload files
awslocal s3 cp --recursive /init-data/ s3://local-book-tracker/

# Check
echo "Checking S3 bucket..."
awslocal s3 ls

echo "Checking S3 bucket contents..."
awslocal s3 ls --recursive s3://local-book-tracker/

echo "=== Completed initializing S3 bucket in LocalStack ==="