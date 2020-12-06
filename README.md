# tdd-sample

テスト駆動開発 (Kent Beck 著, 和田卓人 訳, 2017)のサンプルコードです．

## money(第1部)
golangを使っています．

```sh
/money$ docker-compose run --rm money bash
```

コンテナ内で`go test`などを実行できます．

## xunit(第2部)
書籍通りpythonを使っています・

```sh
/xunit$ docker-compose run --rm xunit sh
```

コンテナ内のシェルで`python xunit.py`を実行します．
