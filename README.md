Packer
======

A simple command-line utility to read/write [MessagePack](http://msgpack.org/index.html) from/to STDIN.

Underlying codec provided by [ugorji](https://github.com/ugorji/go).

Usage
=====

## Packing

    $ echo '{"a":1}' | packer
    ��a�?�

## Unpacking

    $ echo '{"a":1}' | packer | packer -u
    {"a":1}