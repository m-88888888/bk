# bk

bk bookmarks your directory.

## Install

```sh
$ brew tap m-88888888/bk
$ brew install bk
```

下記のコマンドを`.rc`ファイルにコピーしてください。

### bash or zsh
```
cd `bk show | peco`
```

### fish
```
cd (bk show | peco)
```

## Requirements
- peco


## Usage

- bookmark
```sh
$ bk save
```

- show your bookmarks
```sh
$ bk show
```

- delete your bookmark
```sh
$ bk delete
```

- jump your bookmarked directory
```sh
$ jp
```
