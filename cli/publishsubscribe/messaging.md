# Messaging

## Publish / Subscribe

### Publish and Subscribe

> cios ms [Channel ID....]

```shell
$cios ms xxxxxxxxxxxxxxxxxx xxxxxxxxxxxxxxxxxx xxxxxxxxxxxxxxxxxx
Connect: xxxxxxxxxxxxxxxxx
Connect: xxxxxxxxxxxxxxxxx
Connect: xxxxxxxxxxxxxxxxx
>>
```

## Subscribe

### Subscribe Messaging

> cios messaging --subscribe [Channel ID....]

```shell
$cios ms -s xxxxxxxxxxxxxxxxxx xxxxxxxxxxxxxxxxxx xxxxxxxxxxxxxxxxxx
Connect: xxxxxxxxxxxxxxxxx
Connect: xxxxxxxxxxxxxxxxx
Connect: xxxxxxxxxxxxxxxxx
```

## Publish

### Publish Messaging

> cios messaging --publish [Channel ID...]

```shell
$cios ms -p xxxxxxxxxxxxxxxxxx xxxxxxxxxxxxxxxxxx xxxxxxxxxxxxxxxxxx
>>
```

2回連続で改行を入力すればパブリッシュします。

## List

### List Channels

> cios messaging list

```shell
$cios ms ls
*****************************************************************************

        |ID|                    |Resource Owner Name|             |Name|                |Labels|
xxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx    name11   [{key value} {keyy value}]
xxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx    name2    [{key value} {keyy value}]
xxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx    name3    [{key value} {keyy value}]
xxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx    name4    [{key value} {keyy value}]
xxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx    name5    [{key value} {keyy value}]

*****************************************************************************
```
