# Resource Owner

## List

### List Resource Owners

> cios resourceowner list

```shell
$cios ro ls
****************************************************************************************************************
                |id|                            |group_id|                              |user_id|                        |author_id|                  |profile|
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx   HogeHoge1
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx   HogeHoge2
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx   HogeHoge3
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx   HogeHoge4
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx   HogeHoge5
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx   HogeHoge6
****************************************************************************************************************
```
## Delete

### Clean Geography, Channel, DataSore and Bucket

> cios resourceowner delete [Resource Owner ID...]

```shell
cios ro del xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx


Bucket Deleting...
Total: 1
Delete: xxxxxxxxxxxxxxxx
Total: 0


Channel Deleting...
Total: 1
Data Store Completed: xxxxxxxxxxxxxxxx
Deleting: xxxxxxxxxxxxxxxx
Total: 0


Geography Deleting...
Point
Total: 1
Delete: xxxxxxxxxxxxxxxx
Total: 0
Circles
Total: 1
Delete: xxxxxxxxxxxxxxxx
Total: 0
Route
Total: 1
Delete: xxxxxxxxxxxxxxxx
Total: 0
```


