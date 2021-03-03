# File

## File Upload

### Upload File

> cios file upload --bucket_id [Bucket ID] --node_id [Node ID]  [Local File Path...]

```shell
$cios file up -b bucket34dtu3er1er666 ./sample1.txt ./aaa/sample2.txt ./test/sample3.txt
Completed ./sample1.txt
Completed ./aaa/sample2.txt
Completed ./test/sample3.txt

$cios file up -b bucket34dtu3er1er666 -n nodeo34dtu6661ere666  ./sample/a.exe ./sample/b.exe
Completed  ./sample/a.exe
Completed  ./sample/b.exe
```

### Upload Directory

> cios file upload --bucket_id [Bucket ID] --node_id [Node ID]  --directory ./sample

```shell
$cios f up -b bucket34dtu3er1er666 -d ./sample/
Completed sample1.txt
Completed sample2.txt
Completed sample3.txt

$cios f up -b bucket34dtu3er1er666 -n nodeo34dtu6661ere666 -d ./sample/
Completed sample1.txt
Completed sample2.txt
Completed sample3.txt
```

## File Download

⚠ --ld を指定しないと 「./」になります。

### Download From Bucket
> cios file download -b [Bucket ID] [Node ID]

```shell
$cios f get -b bucket34dtu3er1er666
Completed  ./sample/t1.txt
Completed  ./sample/t2.txt
Completed  ./sample/t3.txt
$cios f get -b bucket34dtu3er1er666 nodeo34dtu6661ere666 nodeo34dtu6661ere667 nodeo34dtu6661ere668
```
