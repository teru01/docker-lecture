# Docker講習会

## 環境構築

環境構築について説明します.

1. Virtualbox, Vagrant, ハンズオン資料, 仮想マシンを下記リンクよりインストールします

        Virtualbox: https://www.virtualbox.org

        Vagrant: https://www.vagrantup.com

        資料，仮想マシン: https://drive.google.com/drive/folders/1qqkfNyrCHc_t8v8z7Em4LwA4aEiCg_G8?usp=sharing
        (4GBほどあるので時間かかります)

2. 本プロジェクトをcloneし，1．で取得した仮想マシンをそこに移動させます

    ```
    $ git clone https://github.com/teru01/docker-lecture.git
    $ cd docker-lecture
    $ mv path/to/lecture.box .
    ```

3. 仮想マシンをVagrantに追加します

    ```
    $ vagrant box add --name docker-lecture lecture.box
    ```

4. 仮想マシンを起動させます

    ```
    $ vagrant up
    ```

5. 仮想マシンにログインします．Dockerがインストールできていることを確認してください

    ```
    $ vagrant ssh
    $ sudo docker version

    Client: Docker Engine - Community
    Version:           19.03.6
    API version:       1.40
    Go version:        go1.12.16
    Git commit:        369ce74a3c
    Built:             Thu Feb 13 01:28:06 20
    ...
    ```

