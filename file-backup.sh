#!/bin/bash

read -p "Source directory:" source_dir
read -p "Backup directory:" backup_dir
timestamp=$(date +%Y-%m-%d_%H-%M-%S)
backup_file="backup-$timestamp.tar.gz"
tar -czf $backup_file $source_dir
mv $source_dir/$backup_file $backup_dir