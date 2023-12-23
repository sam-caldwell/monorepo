### Table: `PropertyKeys`
* Identify a set of property names, identified by an `entity`.`id` (UUID).
* The property values will be stored in type-specific tables (`stringProperties` and `numericProperties`).
### Table: `NumericProperties`
* Store a numeric value identified by a `name` from the `PropertyKeys` table.
* The numeric value is an `integer`.

### Table: `StringProperties`
* Store a string value identified by a `name` from the `PropertyKeys` table.
* The string value is a `text` type.

### Functions:
#### [`createPropertyKey()`](080-Properties.sql)
* Create a property key (`name`) only.  This should not really be used by the application tier.  Instead a property should be created using a type-specific function (below).
#### [`createIntegerProperty()`](./080-Properties.sql)
* Create a numeric property key and value.
#### [`createStringProperty()`](./080-properties.sql)
* Create a string property key and value.
#### [`deletePropertyByName()`](./080-Properties.sql)
* Given a property name, delete the value and key from all tables.
#### [`getIntegerProperty()`](./080-Properties.sql)
* Given a property name, return the matching integer value.
#### [`getStringProperty()`](./080-Properties.sql)
* Given a property name, return the matching string value.
