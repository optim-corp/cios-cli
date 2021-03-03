# Node

## List

### List Nodes

> cios node list

```shell
$cios n ls
****************************************************************************************************************

--------------------------------------------------------------------------------------------------------------

|Bucket ID        |: xxxxxxxxxxxxxxxxxx
|Bucket Name      |: device-model
|Resource Owner ID|: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

            |ID|              |Parent Node ID|  |Is directory|      |Name : Key|
    xxxxxxxxxxxxxxxxxxxx    null                    false           sample.txt : /sample.txt
    xxxxxxxxxxxxxxxxxxxx    null                    false           sample.txt : /sample.txt
    xxxxxxxxxxxxxxxxxxxx    null                    false           sample.txt : /sample.txt
    xxxxxxxxxxxxxxxxxxxx    null                    false           sample.txt : /sample.txt
    xxxxxxxxxxxxxxxxxxxx    null                    false           sample.txt : /sample.txt
    xxxxxxxxxxxxxxxxxxxx    null                    false           sample.txt : /sample.txt
    xxxxxxxxxxxxxxxxxxxx    null                    false           sample.txt : /sample.txt
    xxxxxxxxxxxxxxxxxxxx    null                    false           sample.txt : /sample.txt

--------------------------------------------------------------------------------------------------------------

****************************************************************************************************************
```

> cios node ls --all

☠ 下の階層もリクエストするため、大量の内容になる可能性があります。

```shell
$cios node ls -a
****************************************************************************************************************

--------------------------------------------------------------------------------------------------------------

|Bucket ID        |: xxxxxxxxxxxxxxxxxx
|Bucket Name      |: device-model
|Resource Owner ID|: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

            |ID|              |Parent Node ID|  |Is directory|      |Name : Key|
    xxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxxxx    false           sample.txt : /sample.txt
    xxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxxxx    false           sample.txt : /sample.txt
    xxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxxxx    false           sample.txt : /sample.txt
    xxxxxxxxxxxxxxxxxxxx    null                    false           sample.txt : /sample.txt
    xxxxxxxxxxxxxxxxxxxx    null                    false           sample.txt : /sample.txt
    xxxxxxxxxxxxxxxxxxxx    null                    false           sample.txt : /sample.txt
    xxxxxxxxxxxxxxxxxxxx    null                    false           sample.txt : /sample.txt
    xxxxxxxxxxxxxxxxxxxx    null                    false           sample.txt : /sample.txt

--------------------------------------------------------------------------------------------------------------

****************************************************************************************************************
```

### List Nodes on a Bucket

> cios node list [Bucket ID....]

```shell
$cios node ls xxxxxxxxxxxxxxxxxx xxxxxxxxxxxxxxxxxx xxxxxxxxxxxxxxxxxx
****************************************************************************************************************

--------------------------------------------------------------------------------------------------------------

|Bucket ID        |: xxxxxxxxxxxxxxxxxx
|Bucket Name      |: device-model
|Resource Owner ID|: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

            |ID|              |Parent Node ID|  |Is directory|      |Name : Key|
    xxxxxxxxxxxxxxxxxxxx    null                    false           sample.txt : /sample.txt
    xxxxxxxxxxxxxxxxxxxx    null                    false           sample.txt : /sample.txt
    xxxxxxxxxxxxxxxxxxxx    null                    false           sample.txt : /sample.txt
    xxxxxxxxxxxxxxxxxxxx    null                    false           sample.txt : /sample.txt
    xxxxxxxxxxxxxxxxxxxx    null                    false           sample.txt : /sample.txt
    xxxxxxxxxxxxxxxxxxxx    null                    false           sample.txt : /sample.txt
    xxxxxxxxxxxxxxxxxxxx    null                    false           sample.txt : /sample.txt
    xxxxxxxxxxxxxxxxxxxx    null                    false           sample.txt : /sample.txt

--------------------------------------------------------------------------------------------------------------
...

****************************************************************************************************************
```

## Create

### Create Node Directories on a Bucket

> cios node create --bucket_id [Bucket ID] [Name...]

```shell
$cios n add -b xxxxxxxxx sample1 sample2
Completed sample1
Completed sample2
```

## Delete

### Delete Nodes on a Bucket

> cios node delete --bucket_id [Bucket ID] [Node ID...]

```shell
$cios n del -b xxxxxxxxxxxxxxxxx xxxxxxxxxxxx xxxxxxxxxxxxxxxx
Completed xxxxxxxxxxxxxxxxx
Completed xxxxxxxxxxxxxxxxx
```

## Rename

### Rename a Node

> cios node rename --bucket_id [Bucket ID] --node_id[Node ID] --name [New Name]

```shell
$cios n rename -b xxxxxxxxxxxxx -n xxxxxxxxxxxxxxxxxxx --nm "New Name"
Completed  New Name
```

## Move

### Move Nodes

> cios node move --bucket_id [Bucket ID] --dest_bucket_id [Bucket ID] --parent_node_id [Node ID] [Node ID...]

```shell
$cios node mv -b xxxxxxxxxxx -d xxxxxxxxxxx -p xxxxxxxxxxxxx xxxxxxxxxxxxxxxxxx xxxxxxxxxxxxxxx xxxxxxxxxxxxx xxxxxxxxxxxx
Completed xxxxxxxxxxxxx
Completed xxxxxxxxxxxxx
Completed xxxxxxxxxxxxx
Completed xxxxxxxxxxxxx
```

## Copy

### Copy Nodes

> cios node copy --bucket_id [Bucket ID] --dest_bucket_id [Bucket ID] --parent_node_id [Node ID] [Node ID...]

```shell
$cios copy -b xxxxxxxxxxx -d xxxxxxxxxxxxx -p xxxxxxxxxxxxx xxxxxxxxxxxxx xxxxxxxxxxxxxx xxxxxxxxxxx xxxxxxxxxxx 
Completed xxxxxxxxxxxxx
Completed xxxxxxxxxxxxx
Completed xxxxxxxxxxxxx
Completed xxxxxxxxxxxxx
```
