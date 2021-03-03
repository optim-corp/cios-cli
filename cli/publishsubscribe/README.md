# PubSub API

* [Channel](./channel.md)

    ```shell
    NAME:
        cios channel - cios channel | ch

    USAGE:
        cios channel command [command options] [arguments...]

    COMMANDS:
        create, cre, add, sub
        delete, del, remove, rm
        list, ls, show
        update, patch, up, p, edit, ed, et
        help, h                             Shows a list of commands or help for one command

    OPTIONS:
        --help, -h     show help (default: false)
        --version, -v  print the version (default: false)
    ```

* [Data Store](./datastore.md)

    ```shell
    NAME:
        cios datastore - cios datastore | ds

    USAGE:
        cios datastore command [command options] [arguments...]

    COMMANDS:
        create, cre, add, sub
        delete, del, remove, rm
        list, ls, show
        sore, backup, s
        help, h                  Shows a list of commands or help for one command

    OPTIONS:
        --help, -h     show help (default: false)
        --version, -v  print the version (default: false)
    ```

* [Messaging](./messaging.md)

    ```shell
    NAME:
        cios messaging - cios messaging | ms

    USAGE:
        cios messaging command [command options] [arguments...]

    COMMANDS:
        ls, list, show  cios messaging  ls | ms ls
        bot, b          cios messaging  bot | ms b
        help, h         Shows a list of commands or help for one command

    OPTIONS:
        --subscribe, -s                                  (default: false)
        --publish, -p                                    (default: false)
        --resource_owner_id value, -r value, --ro value
        --help, -h                                       show help (default: false)
        --version, -v                                    print the version (default: false)
    ```