package guid

const (
	emptyGUID = "{00000000-0000-0000-0000-000000000000}"
	guidRegex = `^\{?([0-9A-Fa-f]{8})-?([0-9A-Fa-f]{4})-?([0-9A-Fa-f]{4})-?([0-9A-Fa-f]{4})-?([0-9A-Fa-f]{12})\}?$`
)

const (
	IidNull                   iIdName = "IID_NULL"                      // null Interface ID, used when no other Interface ID is known
	IUnknown                  iIdName = "IID_IUnknown"                  // for IUnknown interfaces
	IDispatch                 iIdName = "IID_IDispatch"                 // for IDispatch interfaces
	IEnumVariant              iIdName = "IID_IEnumVariant"              // for IEnumVariant interfaces
	IConnectionPointContainer iIdName = "IID_IConnectionPointContainer" // for IConnectionPointContainer interfaces
	IConnectionPoint          iIdName = "IID_IConnectionPoint"          // for IConnectionPoint interfaces
	IInspectable              iIdName = "IID_IInspectable"              // for IInspectable interfaces
	IProvideClassInfo         iIdName = "IID_IProvideClassInfo"         // for IProvideClassInfo interfaces
	IcomTestString            iIdName = "IID_ICOMTestString"            // for ICOMTestString interfaces
	IcomTestInt8              iIdName = "IID_ICOMTestInt8"              // for ICOMTestInt8 interfaces
	IcomTestInt16             iIdName = "IID_ICOMTestInt16"             // for ICOMTestInt16 interfaces
	IcomTestInt32             iIdName = "IID_ICOMTestInt32"             // for ICOMTestInt32 interfaces
	IcomTestInt64             iIdName = "IID_ICOMTestInt64"             // for ICOMTestInt64 interfaces
	IcomTestFloat             iIdName = "IID_ICOMTestFloat"             // for ICOMTestFloat interfaces
	IcomTestDouble            iIdName = "IID_ICOMTestDouble"            // for ICOMTestDouble interfaces
	IcomTestBoolean           iIdName = "IID_ICOMTestBoolean"           // for ICOMTestBoolean interfaces
	IcomEchoTestObject        iIdName = "IID_ICOMEchoTestObject"        // for ICOMEchoTestObject interfaces
	IcomTestTypes             iIdName = "IID_ICOMTestTypes"             // for ICOMTestTypes interfaces
	ClsIdComEchoTestObject    iIdName = "CLSID_COMEchoTestObject"       // for COMEchoTestObject class
	ClsIdComTestScalarClass   iIdName = "CLSID_COMTestScalarClass"      // for COMTestScalarClass class
)
