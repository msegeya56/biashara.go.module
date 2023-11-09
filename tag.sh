#!/bin/bash

# Get the latest tag
latest_tag=$(git describe --abbrev=0 --tags)

# Extract the tag number
tag_number=${latest_tag##*v}

# Increment the tag number
new_tag_number=$((tag_number + 1))

# Create the new tag with a 'v' prefix
new_tag="v$new_tag_number"

# Tag the current commit with the new tag
git tag "$new_tag"

# Push the new tag to the remote repository
git push origin "$new_tag"
