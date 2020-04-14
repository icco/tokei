#! /bin/zsh

cd ~/Projects/tokei
git pull origin master
go run ./cmd/ink
