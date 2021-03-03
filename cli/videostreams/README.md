# Video Streaming

## List

### List Videos

> cios video list

```shell
$cios v ls
***********************************************************************
    |ID|                  |Resource Owner ID|              |device id|    |enabled|     |Name|
xxxxxxxxxxxxx   xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx     hogehoe1        true    hogehoe1
xxxxxxxxxxxxx   xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx     hogehoe2        true    hogehoe1
xxxxxxxxxxxxx   xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx     hogehoe3        true    hogehoe1

***********************************************************************
```

## List Videos by Resource Owner ID

> cios video list [Resource Owner ID ...]

```shell
$cios v ls xxxxxxxxxxxxxxxxxxxxxxxx xxxxxxxxxxxxxxxxxxxxxxxx xxxxxxxxxxxxxxxxxxxxxxxx
***********************************************************************
    |ID|           |Resource Owner ID|      |device id|    |enabled|     |Name|
xxxxxxxxxxxxx   xxxxxxxxxxxxxxxxxxxxxxxx     hogehoe1        true    hogehoe1
xxxxxxxxxxxxx   xxxxxxxxxxxxxxxxxxxxxxxx     hogehoe2        true    hogehoe1
xxxxxxxxxxxxx   xxxxxxxxxxxxxxxxxxxxxxxx     hogehoe3        true    hogehoe1

***********************************************************************
```
