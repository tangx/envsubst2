# envsubst2

## `envsubst2` 的使用方式

`envsubst2` 使用了 `${[a-zA-Z0-9_]+}` 作为正则匹配的 **模版**。 意味着， **只有类似 `${key1}` 这种带有花括号的变量才会被渲染**

这样做，

1. 兼容 `envsubst` 的使用使用， 迁移过来更方便。 
2. 减少认知压力， 不用在担心 `$key1` 和 `${key1}` 这样 **有没有花括号** 的问题了。

默认情况下， `envsubst2` 只会读取 **存在** 的 **环境变量** 进行替换。 如果 **变量存在， 且值为空** 也会替换。

**注意**： 变量 **不存在** 和 **值为空** 是两种不同的状态。

```
$ envsubst2 -h
version: v0.1.8+sha.5acb37b-devel

Usage:
  envsubst2 [flags]

Flags:
      --force-replace   replace all the placeholders, even if their value is empty (default true)
  -h, --help            help for envsubst2
      --input string    input file
      --output string   output file, os.Stdout if empty.
```

当前版本， 支持以上 3 个参数

1. `--input`: 输入文件（模版文件）。
2. `--outpu`: 输出文件， 内容保存的地方。 如果没有指定， 则为 **屏幕标准输出**。
3. `--force-replace`: 强制替换所有变量， **即使变量值为空** 或者 **没有被设置**。

2. **默认** 替换所有 `${key1}` ， 不论其值是否为空。
  + 但是可以通过开关 `--force-update=false` 关闭。 当关闭状态时， `${key1}` 值为空时 **不替换** ， 即在文件中保留字面值。


## Demo 案例

**模版**

```
My home path is ${HOME}.
My user is ${USER}.

This variables doesn't exist: ${DONT_EXIST}
This variables is empty: ${EMPTY}
```

**渲染**

```bash
$ EMPTY="" ./envsubst2 --input template.txt --force-replace=false
```

**结果**

```
My home path is /Users/devops.
My user is devops.

This variables doesn't exist: ${DONT_EXIST}
This variables is empty: 
```

## The `envsubst`

`envsubst` 很好用， 但是要么全部替换， 要么使用白名单替换。 当白名单数据较多的时候， 就非常麻烦了。

```bash
## 全部替换
envsubst < input-file.txt > output-file.txt

## 白名单模式
envsubst '${Key1} ${Key2}' < input-file.txt > output-file.txt
```

## 关于我


<a href="https://typonotes.com/"><img src="https://static.typonotes.com/mp/qrcode.png" alt="" width="500px"></a>

