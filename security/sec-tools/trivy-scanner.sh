`#!/bin/bash`

# Presenting the options to the admin on what they would like to scan
echo "1. Image"
echo "2. Tar File"
echo "3. Repsoitory"
echo "4. File system"
read -p "Select one of the 4 options: " op
# Based on the option the user has chosen. We present then with the prompt to provide the parameter to what they would like to scan.i
if [ $op -eq 1 ]
then
read -p "Enter image name: " image
trivy image $image
elif [ $op -eq 2 ]
then
read -p "Enter path to tar file: " tar
trivy image -input $tar
elif [ $op -eq 3 ]
then
read -p "Enter repository URL: " repo
trivy repo $repo
elif [ $op -eq 4 ]
then
read -p "Enter path to unpacked container image filesystem: " fs
trivy fs $fs
fi`
