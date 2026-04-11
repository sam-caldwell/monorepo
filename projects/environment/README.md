environment
===========

## Description

A reusable library of code for working with environment variables safely.

## Usage

> We create methods which will fetch a given environment variable
> and return the appropriate datatype.  Get<type> will get the type 
> or return the type's default value ("", 0, false, etc.)

* `environment.GetString(name string) (string, error)`
* `environment.GetStringp(name *string) (string, error)`
* `environment.RequireString(name string) (string, error)`
* `environment.RequireStringp(name *string) (string, error)`
* `environemnt.SanitizeString(name string, pattern string) (string, error)`
* `environemnt.SanitizeStringp(name *string, pattern *string) (string, error)`

* `environment.GetInt(name string) (int, error)`
* `environment.GetIntp(name *string) (int, error)`
* `environment.RequireInt(name string) (int, error)`
* `environment.RequireIntp(name *string) (int, error)`
* `environemnt.SanitizeInt(name string, min int, max int) (int, error)`
* `environemnt.SanitizeIntp(name *string, min int, max int) (int, error)`


* `environment.GetFloat(name string) (float64, error)`
* `environment.GetFloatp(name *string) (float64, error)`
* `environment.RequireFloat(name string) (float64, error)`
* `environment.RequireFloatp(name *string) (float64, error)`
* `environemnt.SanitizeFloat(name string, min float64, max float64) (float64, error)`
* `environemnt.SanitizeFloatp(name *string, min float64, max float64) (float64, error)`


* `environment.GetBool(name string) (bool, error)`
* `environment.GetBoolp(name *string) (bool, error)`
* `environment.RequireBool(name string) (bool, error)`
* `environment.RequireBoolp(name *string) (bool, error)`
* `environment.SanitizeBool(name string) (bool, error)`
* `environment.SanitizeBoolp(name *string) (bool, error)`
