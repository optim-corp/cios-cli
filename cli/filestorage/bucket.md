# Bucket

## List

### List Buckets

> cios bucket list

```shell
$cios b ls
***********************************************************************
    |ID|                    |Resource Owner ID|              |Name|
xxxxxxxxxxxxx   xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx     hogehoe1
xxxxxxxxxxxxx   xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx     hogehoe2
xxxxxxxxxxxxx   xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx     hogehoe3

***********************************************************************
```

### List Buckets with Resorce Owner Name (v2.0.0~)

> cios bucket list --all

```shell
$cios b ls -a
*****************************************************************************************************
    |ID|                    |Resource Owner ID|              |Name| : |Resource Owner Name|
xxxxxxxxxxxxx   xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx     hogehoge1  :  resource_owner_name
xxxxxxxxxxxxx   xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx     hogehoge2  :  resource_owner_name
xxxxxxxxxxxxx   xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx     hogehoge3  :  resource_owner_name

*****************************************************************************************************
```

## Create

### Create Buckets on My Resource Owner

> cios bucket create [Name...]

```shell
$cios b add name1 name2
Completed name1
Completed name2
```

### Create Buckets on a Resource Owner

> cios bucket create --resource_owner_id [Resource Owner ID] [Name...]

```shell
$cios b add -r xxxxxxxx-xxxxxx-xxxxxx name1 name2
Completed name1
Completed name2
```

## Delete

### Delete Buckets

> cios bucket delete [Bucket ID....]

```shell
$cios b del xxxxxxx xxxxxxx xxxxxxx
Completed xxxxxxx
Completed xxxxxxx
Completed xxxxxxx
```

### Delete all Buckets

> cios bucket delete --all

☠ これを行うとすべてのBucketを消してしまいます。大変気を付けてお使いください。

```shell
$cios b del -a
Completed xxxxxxx
Completed xxxxxxx
Completed xxxxxxx
```

## Patch

### Update Bucket Name

> cios bucket update --bucket_id [Buket ID] --name [New Name]

```shell
$cios b update -b xxxxxxx -n "TEXT TEXT"
Completed xxxxxxx
```
