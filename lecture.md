---
marp: true
---

<!--
theme: gaia
template: default
page_number: true
paginate: true
style: |
    ul {
        font-size: 48px;
    }
    li li {
        font-size: 38px;
    }
    section.title *, h1 {
        text-align: center;
    }

-->

# Docker講習会

- 目的
  - Dockerすごいって思ってもらう
  - Dockerこわくないって思ってもらう
  - Docker使っても良いかもって思ってもらう

---

# Dockerとは？

- コンテナ型仮想環境の構築・管理ツール
- コンテナ型仮想化って？
  - 「コンテナ」という独立した区画で仮想環境を作る技術
  - VMほどしっかりと分離されていないよ

---

# 使えばわかる

```
(vm)$ sudo docker run -p 8080:80 nginx:1.17-alpine
```

ブラウザで，`localhost:8080`を開いてみよう

このnginxはどこで動いている？

---

# 仮想環境をなぜ使うのか

- ホスト環境との分離のため
  - 依存ソフトウェアのバージョン衝突防止
- 開発者間での環境共有のため
  - 仮想環境を渡すだけですぐ開発開始できる
- デプロイの容易さのため
  - 仮想環境ごとデプロイすれば楽

---
.
.
.
.
# 別にDockerじゃなくても普通のVMで良いのでは？

---

# Dockerをなぜ使うのか

- コンテナ型仮想化
  - 軽量（すぐ起動，すぐ終了，サイズ小）

- 強力なエコシステム
  - 多様な周辺ツール
  - クラウドによるマネージドサービス

---

# Dockerのキホン

- Dockerイメージ
  - コンテナのもと．
- Dockerコンテナ
  - 仮想環境を動かすもの．Dockerイメージから生成される．

この2つはオブジェクト指向でいうクラスとインスタンスの関係に似ている

---

# Dockerイメージを確認しよう

```
(vm)$ sudo docker images
```

```
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
nginx               1.17-alpine         48c8a7c47625        5 weeks ago         21.8MB
```

---


