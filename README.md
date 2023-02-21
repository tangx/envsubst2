# envsubst2


## Usage

```
Usage:
  envsubst2 [flags]

Flags:
      --force-update    replace the placeholder, even if the environment value is empty (default true)
  -h, --help            help for envsubst2
      --input string    input file
      --output string   output file
```


## Why `envsubst2`

`envsubst2` 只支持 `${key1}` 这种变量格式的替换。 

1. 减少认识压力， 不用在担心 `$key1` 和 `${key1}` 的问题。
2. **默认** 替换所有 `${key1}` ， 不论其值是否为空。
  + 但是可以通过开关 `--force-update=false` 关闭。 当关闭状态时， `${key1}` 值为空时 **不替换** ， 即在文件中保留字面值。


## The `envsubst`

`envsubst` 很好用， 但是要么全部替换， 要么使用白名单替换。 当白名单数据较多的时候， 就非常麻烦了。

```bash
## 全部替换
envsubst < input-file.txt > output-file.txt

## 白名单模式
envsubst '${Key1} ${Key2}' < input-file.txt > output-file.txt
```
