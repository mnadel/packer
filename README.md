Packer
======

A simple command-line utility to read/write [MessagePack](http://msgpack.org/index.html).

Underlying codec provided by [ugorji](https://github.com/ugorji/go).

Usage
=====

## Packing

    $ echo '{"a":1,"b":[2,3,4]}' | packer | tee packed
    ��a�?��b��@��@

## Unpacking

    $ cat packed | packer -u
    {"a":1,"b":[2,3,4]}

    $ packer -u -f packed
    {"a":1,"b":[2,3,4]}