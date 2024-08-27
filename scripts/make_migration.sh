#!/bin/bash
read -p "Migration name: " name
touch "./migrations/$(date +'%y_%m_%dT%H_%M_%S' -u) $name.sql"
