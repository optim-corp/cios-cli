# Channel

## List

### List Channels

> cios channel list

```shell
$cios ch ls
****************************************************************************************************************

        |ID|                    |Resource Owner Name|             |Name|                |Labels|
xxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxx     [{key value}]
xxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxx     [{key value}]
xxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxx     [{key value}]
xxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxx     [{key value}]
xxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxx     [{key value}]
xxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxx     [{key value}]
xxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxx     [{key value}]
xxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxx     [{key value}]
****************************************************************************************************************
```

### List Channels with Resource Owner Name (v2.0.0~)

> cios channel list --all

```shell
$cios ch ls -a
****************************************************************************************************************

        |ID|                    |Resource Owner Name|             |Name|                |Labels| : |Resource Owner Name|
xxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxx     [{key value}] : xxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxx     [{key value}] : xxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxx     [{key value}] : xxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxx     [{key value}] : xxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxx     [{key value}] : xxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxx     [{key value}] : xxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxx     [{key value}] : xxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxx     [{key value}] : xxxxxxxxxxxxxxxxxxxxxxxxx
****************************************************************************************************************
```

### List Channels detail

> cios channel list --detail

```shell
$cios ch ls -d
****************************************************************************************************************
{
    "id": "xxxxxxxxxxxxxxx",
    "created_at": "1584066093074121592",
    "updated_at": "1587027359586189127",
    "resource_owner_id": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    "name": "xxxxxxxxxxxxxxx",
    "description": "xxxxxxxxxxxxxxx",
    "labels": [
        {
        "key": "xxxxxxxxxxxxxxx",
        "value": "xxxxxxxxxxxxxxx"
        }
    ],
    "messaging_config": {
        "enabled": true,
        "persisted": true
    },
    "datastore_config": {
        "enabled": true,
        "max_size": "0",
        "max_count": "0"
    }
}
****************************************************************************************************************
```

### Filtered Resource Owner or Label

> cios channel list --resource_owner_id xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

```shell
$cios -r xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
****************************************************************************************************************

        |ID|                    |Resource Owner Name|             |Name|                |Labels|
xxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx    sample2 [{key value} {sample sample} {sample value}]

****************************************************************************************************************
```

> cios channel list --label key=value

```shell
$cios ch ls -l key=value
****************************************************************************************************************

        |ID|                    |Resource Owner Name|             |Name|                |Labels|
xxxxxxxxxxxxxxxxxxxx    xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx    sample2 [{key value} {sample sample} {sample value}]

****************************************************************************************************************
```

## Create

### Create a Channel on CLI questions

> cios channel create

```shell
$cios ch add
? name:  sample2
? description:  
? language:  ja
? is default Yes
? resource owner id:  xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
? label(key=value):
key=value
sample=value
sample=sample
? Messaging Enabled Yes
? Messaging Persisted Yes
? Data Store Enabled Yes
? Max Count:  0
? Max Size:  0  
Completed  sample2
```

### One Line

> cios channel create --name [Name] --label[Key=value,Key=value,...] --description [Description] --messaging_enabled[Bool] --messaging_persisted[Bool] --datastore_enabled [Bool] --datastore_max_count [Number Text] --datastore_max_size [Number Text]  --resource_owner_id [Resource Owner ID]

```shell
$cios ch add -n name2 -l key=value,keyy=valuee -d description
Completed
```

⚠nameが必ず要ります

## Delete

### Delete Channels

> cios channel delete [Channel ID...]

```shell
$cios ch del xxxxxxxxxxx xxxxxxxxxxxxxx xxxxxxxxxxxxx xxxxxxxxxxxx
Completed xxxxxxxxxxx
Completed xxxxxxxxxxx
Completed xxxxxxxxxxx
Completed xxxxxxxxxxx
```

## Patch

....
