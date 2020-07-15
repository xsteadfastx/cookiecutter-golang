#!/bin/sh

set -ex

git init
git config --add user.name "{{ cookiecutter.author_name }}"
git config --add user.email "{{ cookiecutter.author_email }}"
go mod init "{{ cookiecutter.repo }}"
go mod tidy
go mod vendor
go fmt .
git add -A
git commit -m "first commit"
