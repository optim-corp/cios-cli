# Data Store

## List

### List Data

> cios datastore list

```shell
$cios ds ls
*****************************************************************************


|Channel ID|  : xxxxxxxxxxxxxxxxxxxx
|Channel Name|: Sample

        |ID|            |Timestamp|       |Mime Type|
xxxxxxxxxxxxxxxxxxxx 1593330665040531144 application/json
xxxxxxxxxxxxxxxxxxxx 1593330665040511144 application/json
xxxxxxxxxxxxxxxxxxxx 1593330665040501144 application/json
xxxxxxxxxxxxxxxxxxxx 1593330665030531144 application/json

*****************************************************************************
```

### List Data on Channel

> cios datastore list --channel_id [Channel ID]

```shell
$cios ds ls -c xxxxxxxxxxxxxxxxxxxx
*****************************************************************************


|Channel ID|  : xxxxxxxxxxxxxxxxxxxx
|Channel Name|: Sample

        |ID|            |Timestamp|       |Mime Type|
xxxxxxxxxxxxxxxxxxxx 1593330665040531144 application/json
xxxxxxxxxxxxxxxxxxxx 1593330665040511144 application/json
xxxxxxxxxxxxxxxxxxxx 1593330665040501144 application/json
xxxxxxxxxxxxxxxxxxxx 1593330665030531144 application/json

*****************************************************************************
```

### List Data value

>cios datastore list --data

```shell
$cios ds ls -d -c xxxxxxxxxxxxxxxxxxxx
****************************************************************************************************************
{"aa1":2,"aa2":3,"aaa":1}
{"aa1":2,"aa2":3,"aaa":1}
{"aa1":2,"aa2":3,"aaa":1}


****************************************************************************************************************
```

### Save Datastore Value for local

> cios datastore ls --data --save --channel_id [Channel ID]

```shell
$cios ds ls -s -d
****************************************************************************************************************

 {"aa1":2,"aa2":3,"aaa":1}
Saved:  /home/hogehoge/.cios-cli/datastore/iris/name3___channelID/objectID.txt
 {"aa1":2,"aa2":3,"aaa":1}
Saved:  /home/hogehoge/.cios-cli/datastore/iris/name3___channelID/objectID.txt
 {"aa1":2,"aa2":3,"aaa":1}
Saved:  /home/hogehoge/.cios-cli/datastore/iris/name3___channelID/objectID.txt
****************************************************************************************************************
```

### Load local Datastore value paths

> cios datastore store

```shell
$cios ds store
|-name1___Channel_ID <Directory>
 |--ObjectID.txt <File>
 |--ObjectID.txt <File>
 |--ObjectID.txt <File>
|-name2___Channel_ID <Directory>
 |--ObjectID.txt <File>
 |--ObjectID.txt <File>
 |--ObjectID.txt <File>
|-name2___Channel_ID <Directory>
 |--ObjectID.txt <File>
 |--ObjectID.txt <File>
 |--ObjectID.txt <File>
|-name3___Channel_ID <Directory>
 |--ObjectID.txt <File>
 |--ObjectID.txt <File>
 |--ObjectID.txt <File>
```

## Delete

### Delete a Datastore value

> cios datastore delete --channel_id  [Channel ID] [Datastore ID...]

```shell
$cios ds del -c xxxxxxxxxxxxxxxxxxxx xxxxxxxxxxxxxxxxxxxx xxxxxxxxxxxxxxxxxxxx
Completed xxxxxxxxxxxxxxxxxxxx
Completed xxxxxxxxxxxxxxxxxxxx
```

### Delete Datastore values by Channel

> cios datastore delete --channel_id [Channel ID]

```shell
$cios ds del -c xxxxxxxxxxxxxxxxxxxx
```
