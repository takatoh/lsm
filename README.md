# Least Squares Method

ばらつきのあるデータから、モデル関数のパラメータを最小二乗法によって求めます。

## Install

`go.mod` ファイルに次の行を追記してください。

```
require github.com/takatoh/lsm v1.0.0
```

## Functions

### LSM1(xs, ys []float64) (float64, float64)

1次のモデル関数 `y = a * x + b` のパラメータ `a`、`b` を返します。

```go
a, b := lsm.LSM1(xs, ys)
```

`xs` と `ys` は一連のデータの `x` と `y` のスライスで、同じ長さでなければなりません。

## License

MIT License
