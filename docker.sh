#!/bin/bash

# You probably don't want to run this script lol (or do you?) ;) 

docker run --rm -it $(docker build -q .)
