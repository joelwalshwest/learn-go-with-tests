#!/bin/bash

# Build the image with secrets mounted
docker build -t learn-go-with-tests-image .

# Run the image
# - Run in an integrated terminal
# - Remove image after running
# - Name of the new container
# - Mount all local files to a shared volume in \code
# - Expose ports for running and debugging
# - Pass local secrets
# - Image to run
docker run -it \
  --rm \
  --name learn-go-with-tests-image-container \
  -v $(pwd):/code \
  learn-go-with-tests-image
