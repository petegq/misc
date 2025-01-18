#!/bin/bash

source_dir="/home/username/Documents"
backup_dir="/media/username/Backup"
timestamp=$(date +%Y-%m-%d_%H-%M-%S)
backup_file="backup-$timestamp.tar.gz"
tar -czf $backup_dir/$backup_file $source_dir