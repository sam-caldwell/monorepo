Breadcrumbs
===========

This is a class which can be used to build a delimited string
from a set of strings added/removed over time. For example,
the YAML Parser (`projects/Parser`) uses this to create flat
key names from a YAML hierarchy such as (object.child.item).

The `Breadcrumbs` class can be defined for use with any of the
following tested data types:

* `double`
* `float`
* `int`
* `string`
* `unsigned`

We specifically discourage the use of `bool` with this class.

## Usage

The following would define a `Breadcrumbs` object of type `string`
with a `-` (hyphen) delimiter:

```c++
Breadcrumbs<string> myData('-');
```

By default, the delimiter is a period (`.`), as follows:

```c++
Breadcrumbs<int> myData;
myData.push(192);
myData.push(168);
myData.push(1);
myData.push(1);
cout << myData.toString() << endl;
```

which would produce the following output:

```text
192.168.1.1
```