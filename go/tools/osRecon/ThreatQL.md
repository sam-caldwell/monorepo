ThreatQL
========

## Objectives:
* Define a Threat Intelligence Query Language (ThreatQL).

## Assumptions
1. The query language exists as a set of pre-defined binary query patterns.
2. The query language returns a set of pre-defined binary result-set patterns.
3. The query language is statically defined in the client binary, and its integrity
   can be verified by executable hash.

## Query Pattern (Theory)
The theoretical pattern for a ThreatQL query is:

   ```text
    <table> where <operator> <arguments>
   ```

where `<arguments>` may be--

   ```
      <table>.<field> <operator> <value>
   ```

or--

   ```
      <table>.<field> regex <regular expression>
   ```

and--

1. A `<table>` is a "virtual table" that maps to some real resource on the system (e.g. 
   process table) and a `<field>` is a "virtual field" in that table mapping to a 
   specific value in that resource (e.g. process id -- PID).
2. An operator is `==`, `!=`, `<`, `>`, `<=`, `>=`, `contains`, `not contains` or
   `regex`

## Result Sets (Theory)
The result set for a query from ThreatQL is the entire record for a resource.
Thus, if we query for a process with a given process id, the query in theory would be--

```text
process where process.id == 1234
```

but the result set would not be able to specify anything less than the full process
record.  This is a bit more expensive, but it would require a threat actor to exert
more effort to deliver misleading data.

## Query Pattern (Actual)
Let us assume that--
1. Because all queries are a subset of data from a "virtual table" we know that the actual
   query structure can be stated as a numeric *resource identifier* and some conditional 
   structure.
2. Because each field identifier of the resource is fairly static (differing only across
   operating systems), the field identifier can also be expressed as a numeric value.
3. Likewise, the operator can be expressed as a numeric value.
4. This leaves only the comparison value which may be expressed as an arbitrary value.

This means, the actual query is a struct:

```text
type ResourceId uint32

type FieldId uint32

type OperatorId uint8  //e.g. 0=equals, 1=notEquals, ... n=regex

type ValueType uint8 //e.g. 0=string, 1=int, 2=bool

type OperandType struct {
    Type ValueType
    Value any
}

type Query struct {
    Resource ResourceId
    Field FieldId
    Operator OperatorId
    Operand OperandType
}
```