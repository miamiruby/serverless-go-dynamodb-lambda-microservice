#!/usr/bin/env bash

# Run the build process
./scripts/build.sh

# Deploy to Amazon Lambda
serverless deploy
