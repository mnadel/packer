Packer
======

A simple command-line utility to read/write [MessagePack](http://msgpack.org/index.html).

Underlying codec provided by [ugorji](https://github.com/ugorji/go).

Usage
=====

## Packing

    $ echo '{"a":1}' | packer
    ��a�?�

    $ echo '{"a":1}' > json
    $ packer -f json > packed
    $ cat packed
    ��a�?�


## Unpacking

    $ cat packed | packer -u
    {"a":1}

    $ packer -u -f packed
    {"a":1}