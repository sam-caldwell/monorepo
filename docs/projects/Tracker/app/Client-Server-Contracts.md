---

excalidraw-plugin: parsed
tags: [excalidraw]

---
==⚠  Switch to EXCALIDRAW VIEW in the MORE OPTIONS menu of this document. ⚠==


# Text Elements
API Request ^EMN7EErM

API Response ^yPn27QsQ

Tracker
API Server ^bholOcMk

Tracker
API Client
(WASM) ^WJDBYFUU

Common GoLang Code ^NjtkFjwO

The API client is a compiled WASM
binary which is built off of a common 
Golang code base.  This allows the
client-server contract to be maintained
with less effort. ^X0XZIgI8

The API Server is a more traditional
Golang binary which derives from the 
same code base as the WASM client
with less effort and time. ^E0rBnx6j

ReactJS User-Interface ^t17FvC07

The ReactJS UI can focus
on the graphic user experience
and call the WASM API client
to perform underlying business
logic. ^DJvzbV6m

cli tools 
(in Golang) ^k3vlRznz

Command-line tools to interact
natively with the Tracker system
can be built immediately off the
Tracker API client, where the
API client is compiled to native
binary formats (not WASM). ^8KsmXKP0

The Core of the Tracker business logic is 
expressed in a common code base of several 
Golang packages.  This allows the client and
server side of the contract one time.
 ^9ebx82FR

A single-language, single source for both the client
and server side of the contract eliminates the 
need for special tooling (e.g. OpenAPI/Swagger). ^cuaHkAyk

TrackerDB ^zLZgWMGS

The TrackerDB database is
implemented in PostgreSQL. API
Server calls to the DB invoke
functions or query views. Direct
access table access is not 
allowed. ^dA7mnlIo

%%
# Drawing
```json
{
	"type": "excalidraw",
	"version": 2,
	"source": "https://github.com/zsviczian/obsidian-excalidraw-plugin/releases/tag/2.0.11",
	"elements": [
		{
			"id": "DFhbEQfUd_O2npsIhLqxn",
			"type": "rectangle",
			"x": -469.6971435546875,
			"y": -282.8040771484375,
			"width": 246.9844970703125,
			"height": 247.19171142578125,
			"angle": 0,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "#a5d8ff",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 3
			},
			"seed": 1013561967,
			"version": 155,
			"versionNonce": 555866543,
			"isDeleted": false,
			"boundElements": [
				{
					"id": "ealnSVPz8ygJtn0sM73gb",
					"type": "arrow"
				},
				{
					"id": "luJ0tR6evb5TXczE30Wg4",
					"type": "arrow"
				},
				{
					"id": "1ZZ8x-JHQ78EMJ0rHoDqu",
					"type": "arrow"
				},
				{
					"id": "SXuqwALW8DHOJYe8WuKgd",
					"type": "arrow"
				},
				{
					"id": "gZL1xK2KSVA2uzhIKE18W",
					"type": "arrow"
				}
			],
			"updated": 1703254785273,
			"link": null,
			"locked": false
		},
		{
			"id": "4ButMRsNXZHJs-n3fHUl1",
			"type": "rectangle",
			"x": 534.1141357421875,
			"y": -280.68272399902344,
			"width": 181.5476888020835,
			"height": 247.53236389160156,
			"angle": 0,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "#b2f2bb",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 3
			},
			"seed": 437059279,
			"version": 231,
			"versionNonce": 489475617,
			"isDeleted": false,
			"boundElements": [
				{
					"id": "ealnSVPz8ygJtn0sM73gb",
					"type": "arrow"
				},
				{
					"id": "luJ0tR6evb5TXczE30Wg4",
					"type": "arrow"
				},
				{
					"id": "5uvCaNjERtguYSTsrhfOp",
					"type": "arrow"
				},
				{
					"id": "bX9VqXXR4CPuscnobA8_H",
					"type": "arrow"
				},
				{
					"id": "Eoko89AegP3fkRkMBuZNv",
					"type": "arrow"
				}
			],
			"updated": 1703254593880,
			"link": null,
			"locked": false
		},
		{
			"id": "EfXWR_LtsoHtq5WeUVueO",
			"type": "rectangle",
			"x": -471.3558349609375,
			"y": 159.496337890625,
			"width": 1230.1334228515625,
			"height": 110.67279052734375,
			"angle": 0,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 3
			},
			"seed": 285302575,
			"version": 425,
			"versionNonce": 36789871,
			"isDeleted": false,
			"boundElements": [
				{
					"id": "5uvCaNjERtguYSTsrhfOp",
					"type": "arrow"
				},
				{
					"id": "1ZZ8x-JHQ78EMJ0rHoDqu",
					"type": "arrow"
				}
			],
			"updated": 1703254576882,
			"link": null,
			"locked": false
		},
		{
			"id": "ealnSVPz8ygJtn0sM73gb",
			"type": "arrow",
			"x": -233.44983673095703,
			"y": -223.69794654790022,
			"width": 758.861457824707,
			"height": 9.44468116759694,
			"angle": 0,
			"strokeColor": "#2f9e44",
			"backgroundColor": "transparent",
			"fillStyle": "solid",
			"strokeWidth": 4,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 2
			},
			"seed": 599344737,
			"version": 619,
			"versionNonce": 2062369793,
			"isDeleted": false,
			"boundElements": [
				{
					"type": "text",
					"id": "EMN7EErM"
				}
			],
			"updated": 1703254576882,
			"link": null,
			"locked": false,
			"points": [
				[
					0,
					0
				],
				[
					758.861457824707,
					-9.44468116759694
				]
			],
			"lastCommittedPoint": null,
			"startBinding": {
				"elementId": "WJDBYFUU",
				"gap": 28.191246032714844,
				"focus": -0.9181493631511545
			},
			"endBinding": {
				"elementId": "4ButMRsNXZHJs-n3fHUl1",
				"gap": 8.7025146484375,
				"focus": 0.6202295611574598
			},
			"startArrowhead": null,
			"endArrowhead": "arrow"
		},
		{
			"id": "EMN7EErM",
			"type": "text",
			"x": -167.18309020996094,
			"y": -470.63044520055826,
			"width": 127.11990356445312,
			"height": 25,
			"angle": 0,
			"strokeColor": "#2f9e44",
			"backgroundColor": "transparent",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"seed": 1872642479,
			"version": 22,
			"versionNonce": 1752807567,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1703254576882,
			"link": null,
			"locked": false,
			"text": "API Request",
			"rawText": "API Request",
			"fontSize": 20,
			"fontFamily": 1,
			"textAlign": "center",
			"verticalAlign": "middle",
			"baseline": 18,
			"containerId": "ealnSVPz8ygJtn0sM73gb",
			"originalText": "API Request",
			"lineHeight": 1.25
		},
		{
			"id": "luJ0tR6evb5TXczE30Wg4",
			"type": "arrow",
			"x": 529.6973876953125,
			"y": -91.75726227273506,
			"width": 760.5874710083008,
			"height": 7.989113348152074,
			"angle": 0,
			"strokeColor": "#2f9e44",
			"backgroundColor": "transparent",
			"fillStyle": "solid",
			"strokeWidth": 4,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 2
			},
			"seed": 129204495,
			"version": 537,
			"versionNonce": 1065819105,
			"isDeleted": false,
			"boundElements": [
				{
					"type": "text",
					"id": "yPn27QsQ"
				}
			],
			"updated": 1703254576882,
			"link": null,
			"locked": false,
			"points": [
				[
					0,
					0
				],
				[
					-760.5874710083008,
					-7.989113348152074
				]
			],
			"lastCommittedPoint": null,
			"startBinding": {
				"elementId": "4ButMRsNXZHJs-n3fHUl1",
				"gap": 4.416748046875,
				"focus": -0.5304629205113109
			},
			"endBinding": {
				"elementId": "WJDBYFUU",
				"gap": 30.750999450683594,
				"focus": 0.8496790340860947
			},
			"startArrowhead": null,
			"endArrowhead": "arrow"
		},
		{
			"id": "yPn27QsQ",
			"type": "text",
			"x": -167.54032135009766,
			"y": -329.0811767578125,
			"width": 134.6798858642578,
			"height": 25,
			"angle": 0,
			"strokeColor": "#2f9e44",
			"backgroundColor": "transparent",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"seed": 414076417,
			"version": 16,
			"versionNonce": 888262319,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1703254576882,
			"link": null,
			"locked": false,
			"text": "API Response",
			"rawText": "API Response",
			"fontSize": 20,
			"fontFamily": 1,
			"textAlign": "center",
			"verticalAlign": "middle",
			"baseline": 18,
			"containerId": "luJ0tR6evb5TXczE30Wg4",
			"originalText": "API Response",
			"lineHeight": 1.25
		},
		{
			"id": "bholOcMk",
			"type": "text",
			"x": 545.9243793549259,
			"y": -193.7180689252444,
			"width": 154.91258239746094,
			"height": 71.11804780505956,
			"angle": 0,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"seed": 936487937,
			"version": 405,
			"versionNonce": 1340862607,
			"isDeleted": false,
			"boundElements": [
				{
					"id": "bX9VqXXR4CPuscnobA8_H",
					"type": "arrow"
				},
				{
					"id": "Eoko89AegP3fkRkMBuZNv",
					"type": "arrow"
				}
			],
			"updated": 1703254584916,
			"link": null,
			"locked": false,
			"text": "Tracker\nAPI Server",
			"rawText": "Tracker\nAPI Server",
			"fontSize": 28.447219122023824,
			"fontFamily": 1,
			"textAlign": "center",
			"verticalAlign": "top",
			"baseline": 60,
			"containerId": null,
			"originalText": "Tracker\nAPI Server",
			"lineHeight": 1.25
		},
		{
			"id": "WJDBYFUU",
			"type": "text",
			"x": -444.0530090332031,
			"y": -226.69473662968596,
			"width": 182.41192626953125,
			"height": 135,
			"angle": 0,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"seed": 1035925569,
			"version": 310,
			"versionNonce": 37888207,
			"isDeleted": false,
			"boundElements": [
				{
					"id": "ealnSVPz8ygJtn0sM73gb",
					"type": "arrow"
				},
				{
					"id": "luJ0tR6evb5TXczE30Wg4",
					"type": "arrow"
				},
				{
					"id": "gZL1xK2KSVA2uzhIKE18W",
					"type": "arrow"
				},
				{
					"id": "SXuqwALW8DHOJYe8WuKgd",
					"type": "arrow"
				},
				{
					"id": "41oPMtF4HJFiXWnQyESPI",
					"type": "arrow"
				},
				{
					"id": "1vrW6FDHf6sfxI5AF49lw",
					"type": "arrow"
				}
			],
			"updated": 1703254576882,
			"link": null,
			"locked": false,
			"text": "Tracker\nAPI Client\n(WASM)",
			"rawText": "Tracker\nAPI Client\n(WASM)",
			"fontSize": 36,
			"fontFamily": 1,
			"textAlign": "center",
			"verticalAlign": "top",
			"baseline": 122,
			"containerId": null,
			"originalText": "Tracker\nAPI Client\n(WASM)",
			"lineHeight": 1.25
		},
		{
			"id": "NjtkFjwO",
			"type": "text",
			"x": -165.63040924072266,
			"y": 174.99169921875,
			"width": 656.281982421875,
			"height": 79.07115180189801,
			"angle": 0,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"seed": 283183009,
			"version": 348,
			"versionNonce": 1393055649,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1703254576882,
			"link": null,
			"locked": false,
			"text": "Common GoLang Code",
			"rawText": "Common GoLang Code",
			"fontSize": 63.25692144151841,
			"fontFamily": 1,
			"textAlign": "center",
			"verticalAlign": "top",
			"baseline": 55,
			"containerId": null,
			"originalText": "Common GoLang Code",
			"lineHeight": 1.25
		},
		{
			"id": "5uvCaNjERtguYSTsrhfOp",
			"type": "arrow",
			"x": 432.5153041879112,
			"y": 154.3935546875,
			"width": 100.05091692157151,
			"height": 170.70374242447275,
			"angle": 0,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 2
			},
			"seed": 688197185,
			"version": 489,
			"versionNonce": 33763055,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1703254576882,
			"link": null,
			"locked": false,
			"points": [
				[
					0,
					0
				],
				[
					100.05091692157151,
					-170.70374242447275
				]
			],
			"lastCommittedPoint": null,
			"startBinding": {
				"elementId": "EfXWR_LtsoHtq5WeUVueO",
				"gap": 5.102783203125,
				"focus": 0.3912643002188475
			},
			"endBinding": {
				"elementId": "4ButMRsNXZHJs-n3fHUl1",
				"gap": 16.911163330078125,
				"focus": 0.06068776667442528
			},
			"startArrowhead": null,
			"endArrowhead": "arrow"
		},
		{
			"id": "1ZZ8x-JHQ78EMJ0rHoDqu",
			"type": "arrow",
			"x": -109.482666015625,
			"y": 148.26507568359375,
			"width": 123.11279296875,
			"height": 178.03515625,
			"angle": 0,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 2
			},
			"seed": 1955667201,
			"version": 528,
			"versionNonce": 1983822721,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1703254576882,
			"link": null,
			"locked": false,
			"points": [
				[
					0,
					0
				],
				[
					-123.11279296875,
					-178.03515625
				]
			],
			"lastCommittedPoint": null,
			"startBinding": {
				"elementId": "EfXWR_LtsoHtq5WeUVueO",
				"focus": -0.31708436332420353,
				"gap": 11.23126220703125
			},
			"endBinding": {
				"elementId": "DFhbEQfUd_O2npsIhLqxn",
				"focus": -0.11534219728972858,
				"gap": 5.84228515625
			},
			"startArrowhead": null,
			"endArrowhead": "arrow"
		},
		{
			"id": "X0XZIgI8",
			"type": "text",
			"x": -531.3950999948381,
			"y": -432.05926513671875,
			"width": 411.7388916015625,
			"height": 132.3825051386165,
			"angle": 0,
			"strokeColor": "#f08c00",
			"backgroundColor": "#b2f2bb",
			"fillStyle": "solid",
			"strokeWidth": 4,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"seed": 1203811681,
			"version": 550,
			"versionNonce": 515733775,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1703254576882,
			"link": null,
			"locked": false,
			"text": "The API client is a compiled WASM\nbinary which is built off of a common \nGolang code base.  This allows the\nclient-server contract to be maintained\nwith less effort.",
			"rawText": "The API client is a compiled WASM\nbinary which is built off of a common \nGolang code base.  This allows the\nclient-server contract to be maintained\nwith less effort.",
			"fontSize": 21.18120082217864,
			"fontFamily": 1,
			"textAlign": "left",
			"verticalAlign": "top",
			"baseline": 124,
			"containerId": null,
			"originalText": "The API client is a compiled WASM\nbinary which is built off of a common \nGolang code base.  This allows the\nclient-server contract to be maintained\nwith less effort.",
			"lineHeight": 1.25
		},
		{
			"type": "text",
			"version": 872,
			"versionNonce": 813979361,
			"isDeleted": false,
			"id": "E0rBnx6j",
			"fillStyle": "solid",
			"strokeWidth": 4,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": 423.36262439546135,
			"y": -407.4229986979355,
			"strokeColor": "#f08c00",
			"backgroundColor": "#b2f2bb",
			"width": 393.31231689453125,
			"height": 105.90600411089319,
			"seed": 579351329,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": [],
			"updated": 1703254601234,
			"link": null,
			"locked": false,
			"fontSize": 21.18120082217864,
			"fontFamily": 1,
			"text": "The API Server is a more traditional\nGolang binary which derives from the \nsame code base as the WASM client\nwith less effort and time.",
			"rawText": "The API Server is a more traditional\nGolang binary which derives from the \nsame code base as the WASM client\nwith less effort and time.",
			"textAlign": "left",
			"verticalAlign": "top",
			"containerId": null,
			"originalText": "The API Server is a more traditional\nGolang binary which derives from the \nsame code base as the WASM client\nwith less effort and time.",
			"lineHeight": 1.25,
			"baseline": 97
		},
		{
			"type": "rectangle",
			"version": 320,
			"versionNonce": 1575583535,
			"isDeleted": false,
			"id": "J8zRWTOKp6v8SQuUjkCzy",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -910.21533203125,
			"y": -278.29590606689453,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "#ffec99",
			"width": 226.12939453125003,
			"height": 127.18629455566406,
			"seed": 246601249,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 3
			},
			"boundElements": [
				{
					"type": "text",
					"id": "t17FvC07"
				},
				{
					"id": "gZL1xK2KSVA2uzhIKE18W",
					"type": "arrow"
				},
				{
					"id": "SXuqwALW8DHOJYe8WuKgd",
					"type": "arrow"
				}
			],
			"updated": 1703254576882,
			"link": null,
			"locked": false
		},
		{
			"id": "t17FvC07",
			"type": "text",
			"x": -899.9945983886719,
			"y": -249.7027587890625,
			"width": 205.68792724609375,
			"height": 70,
			"angle": 0,
			"strokeColor": "#e03131",
			"backgroundColor": "#b2f2bb",
			"fillStyle": "solid",
			"strokeWidth": 4,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"seed": 2132401359,
			"version": 177,
			"versionNonce": 1859424065,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1703254576882,
			"link": null,
			"locked": false,
			"text": "ReactJS User-\nInterface",
			"rawText": "ReactJS User-Interface",
			"fontSize": 28,
			"fontFamily": 1,
			"textAlign": "center",
			"verticalAlign": "middle",
			"baseline": 60,
			"containerId": "J8zRWTOKp6v8SQuUjkCzy",
			"originalText": "ReactJS User-Interface",
			"lineHeight": 1.25
		},
		{
			"id": "gZL1xK2KSVA2uzhIKE18W",
			"type": "arrow",
			"x": -674.1082153320312,
			"y": -233.13582162301373,
			"width": 197.91333563750152,
			"height": 0.9645819764756425,
			"angle": 0,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "#ffec99",
			"fillStyle": "solid",
			"strokeWidth": 4,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 2
			},
			"seed": 514074305,
			"version": 338,
			"versionNonce": 725967887,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1703254786274,
			"link": null,
			"locked": false,
			"points": [
				[
					0,
					0
				],
				[
					197.91333563750152,
					-0.9645819764756425
				]
			],
			"lastCommittedPoint": null,
			"startBinding": {
				"elementId": "J8zRWTOKp6v8SQuUjkCzy",
				"focus": -0.2780202170845725,
				"gap": 9.97772216796875
			},
			"endBinding": {
				"elementId": "DFhbEQfUd_O2npsIhLqxn",
				"focus": 0.6081087292273974,
				"gap": 6.497736139842232
			},
			"startArrowhead": null,
			"endArrowhead": "arrow"
		},
		{
			"id": "SXuqwALW8DHOJYe8WuKgd",
			"type": "arrow",
			"x": -476.3070544288272,
			"y": -187.10071408566267,
			"width": 202.65280017041368,
			"height": 1.3496621095957266,
			"angle": 0,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "#ffec99",
			"fillStyle": "solid",
			"strokeWidth": 4,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 2
			},
			"seed": 1895171951,
			"version": 367,
			"versionNonce": 1476041505,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1703254780490,
			"link": null,
			"locked": false,
			"points": [
				[
					0,
					0
				],
				[
					-202.65280017041368,
					-1.3496621095957266
				]
			],
			"lastCommittedPoint": null,
			"startBinding": {
				"elementId": "DFhbEQfUd_O2npsIhLqxn",
				"focus": 0.21721895812641104,
				"gap": 6.6099108741397
			},
			"endBinding": {
				"elementId": "J8zRWTOKp6v8SQuUjkCzy",
				"focus": 0.39575379757443846,
				"gap": 5.1260829007591155
			},
			"startArrowhead": null,
			"endArrowhead": "arrow"
		},
		{
			"type": "text",
			"version": 721,
			"versionNonce": 1814609775,
			"isDeleted": false,
			"id": "DJvzbV6m",
			"fillStyle": "solid",
			"strokeWidth": 4,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -950.7839965820312,
			"y": -429.22421155368323,
			"strokeColor": "#f08c00",
			"backgroundColor": "#b2f2bb",
			"width": 315.20050048828125,
			"height": 132.3825051386165,
			"seed": 217537679,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": [],
			"updated": 1703254576882,
			"link": null,
			"locked": false,
			"fontSize": 21.18120082217864,
			"fontFamily": 1,
			"text": "The ReactJS UI can focus\non the graphic user experience\nand call the WASM API client\nto perform underlying business\nlogic.",
			"rawText": "The ReactJS UI can focus\non the graphic user experience\nand call the WASM API client\nto perform underlying business\nlogic.",
			"textAlign": "left",
			"verticalAlign": "top",
			"containerId": null,
			"originalText": "The ReactJS UI can focus\non the graphic user experience\nand call the WASM API client\nto perform underlying business\nlogic.",
			"lineHeight": 1.25,
			"baseline": 124
		},
		{
			"type": "rectangle",
			"version": 358,
			"versionNonce": 1465414401,
			"isDeleted": false,
			"id": "hOScbiSmlGW9SGp7Q3Fyh",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -913.1927490234375,
			"y": -132.06000518798828,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "#ffec99",
			"width": 226.12939453125003,
			"height": 127.18629455566406,
			"seed": 127350081,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 3
			},
			"boundElements": [
				{
					"type": "text",
					"id": "k3vlRznz"
				},
				{
					"id": "41oPMtF4HJFiXWnQyESPI",
					"type": "arrow"
				},
				{
					"id": "1vrW6FDHf6sfxI5AF49lw",
					"type": "arrow"
				}
			],
			"updated": 1703254576882,
			"link": null,
			"locked": false
		},
		{
			"type": "text",
			"version": 277,
			"versionNonce": 1054669199,
			"isDeleted": false,
			"id": "k3vlRznz",
			"fillStyle": "solid",
			"strokeWidth": 4,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -872.6480178833008,
			"y": -103.46685791015625,
			"strokeColor": "#e03131",
			"backgroundColor": "#b2f2bb",
			"width": 145.03993225097656,
			"height": 70,
			"seed": 476112161,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": [],
			"updated": 1703254576882,
			"link": null,
			"locked": false,
			"fontSize": 28,
			"fontFamily": 1,
			"text": "cli tools \n(in Golang)",
			"rawText": "cli tools \n(in Golang)",
			"textAlign": "center",
			"verticalAlign": "middle",
			"containerId": "hOScbiSmlGW9SGp7Q3Fyh",
			"originalText": "cli tools \n(in Golang)",
			"lineHeight": 1.25,
			"baseline": 60
		},
		{
			"type": "arrow",
			"version": 415,
			"versionNonce": 1858476303,
			"isDeleted": false,
			"id": "41oPMtF4HJFiXWnQyESPI",
			"fillStyle": "solid",
			"strokeWidth": 4,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -471.9071787727171,
			"y": -68.16008783235304,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "#ffec99",
			"width": 206.14252401989927,
			"height": 1.295172328837225,
			"seed": 425972161,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 2
			},
			"boundElements": [],
			"updated": 1703254774472,
			"link": null,
			"locked": false,
			"startBinding": {
				"elementId": "WJDBYFUU",
				"focus": -1.3263196301557143,
				"gap": 27.854169739513964
			},
			"endBinding": {
				"elementId": "hOScbiSmlGW9SGp7Q3Fyh",
				"focus": 0.036840099969699294,
				"gap": 9.013651699571142
			},
			"lastCommittedPoint": null,
			"startArrowhead": null,
			"endArrowhead": "arrow",
			"points": [
				[
					0,
					0
				],
				[
					-206.14252401989927,
					1.295172328837225
				]
			]
		},
		{
			"type": "arrow",
			"version": 361,
			"versionNonce": 559835151,
			"isDeleted": false,
			"id": "1vrW6FDHf6sfxI5AF49lw",
			"fillStyle": "solid",
			"strokeWidth": 4,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -678.4867826723773,
			"y": -100.57941375478214,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "#ffec99",
			"width": 203.5952108474005,
			"height": 2.072247368080781,
			"seed": 367888687,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 2
			},
			"boundElements": [],
			"updated": 1703254789890,
			"link": null,
			"locked": false,
			"startBinding": {
				"elementId": "hOScbiSmlGW9SGp7Q3Fyh",
				"focus": -0.47687015898997365,
				"gap": 8.57657181981017
			},
			"endBinding": {
				"elementId": "WJDBYFUU",
				"focus": -0.8081577030651503,
				"gap": 30.83856279177371
			},
			"lastCommittedPoint": null,
			"startArrowhead": null,
			"endArrowhead": "arrow",
			"points": [
				[
					0,
					0
				],
				[
					203.5952108474005,
					-2.072247368080781
				]
			]
		},
		{
			"type": "text",
			"version": 970,
			"versionNonce": 2005296833,
			"isDeleted": false,
			"id": "8KsmXKP0",
			"fillStyle": "solid",
			"strokeWidth": 4,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -958.7426452636719,
			"y": 9.847321649441767,
			"strokeColor": "#f08c00",
			"backgroundColor": "#b2f2bb",
			"width": 346.92816162109375,
			"height": 158.85900616633978,
			"seed": 1016079937,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": [],
			"updated": 1703254576882,
			"link": null,
			"locked": false,
			"fontSize": 21.18120082217864,
			"fontFamily": 1,
			"text": "Command-line tools to interact\nnatively with the Tracker system\ncan be built immediately off the\nTracker API client, where the\nAPI client is compiled to native\nbinary formats (not WASM).",
			"rawText": "Command-line tools to interact\nnatively with the Tracker system\ncan be built immediately off the\nTracker API client, where the\nAPI client is compiled to native\nbinary formats (not WASM).",
			"textAlign": "left",
			"verticalAlign": "top",
			"containerId": null,
			"originalText": "Command-line tools to interact\nnatively with the Tracker system\ncan be built immediately off the\nTracker API client, where the\nAPI client is compiled to native\nbinary formats (not WASM).",
			"lineHeight": 1.25,
			"baseline": 150
		},
		{
			"type": "text",
			"version": 1259,
			"versionNonce": 2005475791,
			"isDeleted": false,
			"id": "9ebx82FR",
			"fillStyle": "solid",
			"strokeWidth": 4,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -12.324676513671875,
			"y": 303.7035535574551,
			"strokeColor": "#f08c00",
			"backgroundColor": "#b2f2bb",
			"width": 475.9990539550781,
			"height": 132.3825051386165,
			"seed": 1682220321,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": [],
			"updated": 1703254576882,
			"link": null,
			"locked": false,
			"fontSize": 21.18120082217864,
			"fontFamily": 1,
			"text": "The Core of the Tracker business logic is \nexpressed in a common code base of several \nGolang packages.  This allows the client and\nserver side of the contract one time.\n",
			"rawText": "The Core of the Tracker business logic is \nexpressed in a common code base of several \nGolang packages.  This allows the client and\nserver side of the contract one time.\n",
			"textAlign": "left",
			"verticalAlign": "top",
			"containerId": null,
			"originalText": "The Core of the Tracker business logic is \nexpressed in a common code base of several \nGolang packages.  This allows the client and\nserver side of the contract one time.\n",
			"lineHeight": 1.25,
			"baseline": 124
		},
		{
			"type": "text",
			"version": 1608,
			"versionNonce": 2054022817,
			"isDeleted": false,
			"id": "cuaHkAyk",
			"fillStyle": "solid",
			"strokeWidth": 4,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -80.53303697563103,
			"y": -49.84330559884512,
			"strokeColor": "#f08c00",
			"backgroundColor": "#b2f2bb",
			"width": 510.52496337890625,
			"height": 76.63484765274751,
			"seed": 851795599,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": [],
			"updated": 1703254576882,
			"link": null,
			"locked": false,
			"fontSize": 20.435959374066,
			"fontFamily": 1,
			"text": "A single-language, single source for both the client\nand server side of the contract eliminates the \nneed for special tooling (e.g. OpenAPI/Swagger).",
			"rawText": "A single-language, single source for both the client\nand server side of the contract eliminates the \nneed for special tooling (e.g. OpenAPI/Swagger).",
			"textAlign": "left",
			"verticalAlign": "top",
			"containerId": null,
			"originalText": "A single-language, single source for both the client\nand server side of the contract eliminates the \nneed for special tooling (e.g. OpenAPI/Swagger).",
			"lineHeight": 1.25,
			"baseline": 68
		},
		{
			"id": "Dwb2tJ8hRmSnqIMPHrJpr",
			"type": "ellipse",
			"x": 795.9722162784133,
			"y": -261.9389001439689,
			"width": 221.0165337136362,
			"height": 205,
			"angle": 0,
			"strokeColor": "#2f9e44",
			"backgroundColor": "#ffec99",
			"fillStyle": "solid",
			"strokeWidth": 4,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 2
			},
			"seed": 574564975,
			"version": 300,
			"versionNonce": 481784527,
			"isDeleted": false,
			"boundElements": [
				{
					"type": "text",
					"id": "zLZgWMGS"
				},
				{
					"id": "bX9VqXXR4CPuscnobA8_H",
					"type": "arrow"
				},
				{
					"id": "Eoko89AegP3fkRkMBuZNv",
					"type": "arrow"
				}
			],
			"updated": 1703254596302,
			"link": null,
			"locked": false
		},
		{
			"id": "zLZgWMGS",
			"type": "text",
			"x": 836.7333583441691,
			"y": -204.41734521559005,
			"width": 139.2119598388672,
			"height": 90,
			"angle": 0,
			"strokeColor": "#2f9e44",
			"backgroundColor": "#ffec99",
			"fillStyle": "solid",
			"strokeWidth": 4,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"seed": 1778835201,
			"version": 361,
			"versionNonce": 100548847,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1703254596302,
			"link": null,
			"locked": false,
			"text": "Tracker\nDB",
			"rawText": "TrackerDB",
			"fontSize": 36,
			"fontFamily": 1,
			"textAlign": "center",
			"verticalAlign": "middle",
			"baseline": 77,
			"containerId": "Dwb2tJ8hRmSnqIMPHrJpr",
			"originalText": "TrackerDB",
			"lineHeight": 1.25
		},
		{
			"id": "bX9VqXXR4CPuscnobA8_H",
			"type": "arrow",
			"x": 719.5819285333953,
			"y": -196.36464973286712,
			"width": 73.45770818045867,
			"height": 0.3505141597699719,
			"angle": 0,
			"strokeColor": "#2f9e44",
			"backgroundColor": "#ffec99",
			"fillStyle": "solid",
			"strokeWidth": 4,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 2
			},
			"seed": 553371279,
			"version": 515,
			"versionNonce": 1326923009,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1703254605850,
			"link": null,
			"locked": false,
			"points": [
				[
					0,
					0
				],
				[
					73.45770818045867,
					-0.3505141597699719
				]
			],
			"lastCommittedPoint": null,
			"startBinding": {
				"elementId": "4ButMRsNXZHJs-n3fHUl1",
				"gap": 3.9201039891243,
				"focus": -0.3144957257601399
			},
			"endBinding": {
				"elementId": "Dwb2tJ8hRmSnqIMPHrJpr",
				"gap": 9.756734586658581,
				"focus": 0.36894695474323264
			},
			"startArrowhead": null,
			"endArrowhead": "arrow"
		},
		{
			"id": "Eoko89AegP3fkRkMBuZNv",
			"type": "arrow",
			"x": 800.0392066332741,
			"y": -115.29085627345327,
			"width": 79.41204237053194,
			"height": 0.898839092975237,
			"angle": 0,
			"strokeColor": "#2f9e44",
			"backgroundColor": "#ffec99",
			"fillStyle": "solid",
			"strokeWidth": 4,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 2
			},
			"seed": 118277825,
			"version": 460,
			"versionNonce": 1905963201,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1703254605850,
			"link": null,
			"locked": false,
			"points": [
				[
					0,
					0
				],
				[
					-79.41204237053194,
					0.898839092975237
				]
			],
			"lastCommittedPoint": null,
			"startBinding": {
				"elementId": "Dwb2tJ8hRmSnqIMPHrJpr",
				"gap": 6.012512094405338,
				"focus": -0.4187787242248105
			},
			"endBinding": {
				"elementId": "4ButMRsNXZHJs-n3fHUl1",
				"gap": 4.965339718471228,
				"focus": 0.34944226333582146
			},
			"startArrowhead": null,
			"endArrowhead": "arrow"
		},
		{
			"type": "text",
			"version": 1336,
			"versionNonce": 241273071,
			"isDeleted": false,
			"id": "dA7mnlIo",
			"fillStyle": "solid",
			"strokeWidth": 4,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": 760.404254407995,
			"y": -40.853210792371726,
			"strokeColor": "#f08c00",
			"backgroundColor": "#b2f2bb",
			"width": 298.5704650878906,
			"height": 142.942184785621,
			"seed": 372084239,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": [],
			"updated": 1703254725887,
			"link": null,
			"locked": false,
			"fontSize": 19.058957971416135,
			"fontFamily": 1,
			"text": "The TrackerDB database is\nimplemented in PostgreSQL. API\nServer calls to the DB invoke\nfunctions or query views. Direct\naccess table access is not \nallowed.",
			"rawText": "The TrackerDB database is\nimplemented in PostgreSQL. API\nServer calls to the DB invoke\nfunctions or query views. Direct\naccess table access is not \nallowed.",
			"textAlign": "left",
			"verticalAlign": "top",
			"containerId": null,
			"originalText": "The TrackerDB database is\nimplemented in PostgreSQL. API\nServer calls to the DB invoke\nfunctions or query views. Direct\naccess table access is not \nallowed.",
			"lineHeight": 1.25,
			"baseline": 135
		}
	],
	"appState": {
		"theme": "dark",
		"viewBackgroundColor": "#ffffff",
		"currentItemStrokeColor": "#2f9e44",
		"currentItemBackgroundColor": "#ffec99",
		"currentItemFillStyle": "solid",
		"currentItemStrokeWidth": 4,
		"currentItemStrokeStyle": "solid",
		"currentItemRoughness": 1,
		"currentItemOpacity": 100,
		"currentItemFontFamily": 1,
		"currentItemFontSize": 36,
		"currentItemTextAlign": "left",
		"currentItemStartArrowhead": null,
		"currentItemEndArrowhead": "arrow",
		"scrollX": 1088.250692055454,
		"scrollY": 602.8005913155613,
		"zoom": {
			"value": 1.05
		},
		"currentItemRoundness": "round",
		"gridSize": null,
		"gridColor": {
			"Bold": "#C9C9C9FF",
			"Regular": "#EDEDEDFF"
		},
		"currentStrokeOptions": null,
		"previousGridSize": null,
		"frameRendering": {
			"enabled": true,
			"clip": true,
			"name": true,
			"outline": true
		}
	},
	"files": {}
}
```
%%