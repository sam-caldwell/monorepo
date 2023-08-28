Ansi Tester
===========

## Description
This project provides a simple wrapper for the `go test` `testing.T` facility to offer
stats tracking and color coded output.

## Usage
```go
//initialize the repotester...
test:=ansi.Test(t)
//defer stats printout when all else is done.
defer ansi.Test(t).Stats()
//report a passing test
ansi.Test(t).Pass("this works")
//Report a failed test
ansi.test(t).Fail("this was broken")
//Report a skip
ansi.test(t).Skip("this was skipped")
```

The tool will provide output similar to this proof of concept:
<div style="background-color:#000000;margin:4px 4px 4px 4px;padding:14px 14px 14px 14px">
--- PASS: TestCalcMaxColumnWidth (0.00s)<br/>
=== RUN   TestLeftPad<br/>
<span style="color:green;margin:0px 0px 0px 0px;padding:0px 0px 0px 0px;">
[ OK ] Incrementing padding-size test<br/>
[ OK ] same-length test<br/>
[ OK ] zero-length test<br/>
[ OK ] negative-length test<br/>
</span>
<span style="color:blue;margin:0px 0px 0px 0px;padding:0px 0px 0px 0px;">
---Test Statistics---<br/>
</span>
<span style="color:green;margin:0px 0px 0px 0px;padding:0px 0px 0px 0px;">
Pass: 4<br/>
</span>
<span style="color:red;margin:0px 0px 0px 0px;padding:0px 0px 0px 0px;">
Fail: 0<br/>
</span>
<span style="color:yellow;">
Skip: 0<br/>
</span>
<span style="color:blue;">
---------------------<br/>
</span>
--- PASS: TestLeftPad (0.00s)<br/>
=== RUN   TestNullObjectStructSize<br/>
--- PASS: TestNullObjectStructSize (0.00s)<br/>
</div>