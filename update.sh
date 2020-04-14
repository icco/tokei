#! /bin/zsh

cd ~/Projects/tokei
curl -s -O https://rsms.me/inter/font-files/Inter-Regular.otf
git pull origin master
go run ./cmd/ink
