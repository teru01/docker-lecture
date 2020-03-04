#!/bin/bash

cd work
mdview-init.sh document

# 必要ファイルをdocumentへコピー
cp jsbook.cls reviewmacro.sty review-custom.sty document/sty/
cp first.re second.re config.yml catalog.yml review-ext.rb document/

# bookの作成
cd document && review-pdfmaker config.yml
cd ..
mv document/book.pdf .
rm -rf document
