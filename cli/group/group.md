# Group

## List

* `$cios gp ls`

    ```shell
    ##################################################################################################################################
                |id|                            |parent_group_id|                       |resource_owner_id|            |type|         |name / tags|
    xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx                                            xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx      Group        sample / [a2 a1]
    ##################################################################################################################################
    ```

## Create

* `$cios gp add`

    ```shell
    ? name:
    ? parent group id:
    ? tags(tag1,tag2,tag3):
    ```

* `$cios gp add -n sample -tag a1,a2`

    ```shell
    Completed sample
    ```

## Delete

* `$cios gp del group_id1 group_id2`

    ```shell
    Completed group_id1
    Completed group_id2
    ```
