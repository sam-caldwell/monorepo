---

excalidraw-plugin: parsed
tags: [excalidraw]

---

==⚠  Switch to EXCALIDRAW VIEW in the MORE OPTIONS menu of this document. ⚠==


# Text Elements
phase II starts ^xWyJtBFW

phase II ends ^fEKoGGOW

node requests
CA certificate chain
(root CA + ICA) ^MHVXcAit

PKI Server (Peer) receives request ^wy4ta4TO

[http/get] http://${pkiLocalLeader}/api/v1/ca ^NOrWANaU

is PKI Global Leader? ^kIgnjrMF

[http/get] http://${pkiGlobalVip}/api/v1/ca ^JgasGr1L

node requests
CA certificate chain
(root CA + ICA) ^t0MQADi3

PKI Server replies with root and intermediate CA certificate (ICA) ^0zOUJMqH

Certificate information is stored locally ^JTBt5ulv

Certificate information is stored locally ^hpYkmtna

PKI server signs the CSR to create the node ICA ^3wef01oc

HTTP/200 OK ^yRdpAvDC

[http/post] https://${pkiLocalLeader}/api/v1/ca ^ocJLmcs4

PKI server returns the node Intermediate Certificate Authority (ICA) ^wFuFRBTI

Generate Private Key and Certificate Signing Request (CSR) ^9jYaSHB8

New Node stores ICA Certificate in memory ^FPu6PrFd

HTTP/200 OK ^uWNLgcew

Request is authenticated with Pre-Shared Key (PSK) ^adNlWsId

Below this point we use HTTPS for all requests ^Rt6aJRrJ

Above this point we use HTTP for all requests ^iPtbjOoh

Yes ^MNH3eTFR

is PKI Local Leader? ^MaBv6rG2

Yes ^heL0DrP8

create self-signed root CA certificate ^PtyD9RWE

create ICA certificate 
from root CA ^DY1J9YvV

No ^KHsHIK4l

No ^zdQTbwZs

We need a new
root CA and
ICA for the new
cluster ^Nnqk8d0w

We only need to
initialize the node
in the trust chain ^5ZvpBRRi

We have a new AZ
and a new AZ PKI
and we just need
an initialization ^LWQLLahI

%%
# Drawing
```json
{
	"type": "excalidraw",
	"version": 2,
	"source": "https://github.com/zsviczian/obsidian-excalidraw-plugin/releases/tag/2.0.13",
	"elements": [
		{
			"id": "Ugea7KMJV-9Jc45-E_z8T",
			"type": "rectangle",
			"x": -665.2550048828125,
			"y": -849.4909210205078,
			"width": 173,
			"height": 35,
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
			"seed": 340248924,
			"version": 381,
			"versionNonce": 1874978012,
			"isDeleted": false,
			"boundElements": [
				{
					"id": "xWyJtBFW",
					"type": "text"
				},
				{
					"id": "SQ1IjhLfrnQg4B9RZMYT7",
					"type": "arrow"
				}
			],
			"updated": 1704651740520,
			"link": null,
			"locked": false
		},
		{
			"id": "xWyJtBFW",
			"type": "text",
			"x": -660.2249374389648,
			"y": -844.4909210205078,
			"width": 162.9398651123047,
			"height": 25,
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
			"seed": 854947300,
			"version": 403,
			"versionNonce": 1365272292,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1704651740520,
			"link": null,
			"locked": false,
			"text": "phase II starts",
			"rawText": "phase II starts",
			"fontSize": 20,
			"fontFamily": 1,
			"textAlign": "center",
			"verticalAlign": "middle",
			"baseline": 18,
			"containerId": "Ugea7KMJV-9Jc45-E_z8T",
			"originalText": "phase II starts",
			"lineHeight": 1.25
		},
		{
			"type": "rectangle",
			"version": 421,
			"versionNonce": 1080617308,
			"isDeleted": false,
			"id": "uh1Sa6tVIjsygaX_OS9mS",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -962.6172485351562,
			"y": 661.8146362304688,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 442.6895751953125,
			"height": 35,
			"seed": 290460764,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 3
			},
			"boundElements": [
				{
					"id": "fEKoGGOW",
					"type": "text"
				},
				{
					"id": "j6DJunfGf0MYhN0rreJ3o",
					"type": "arrow"
				},
				{
					"id": "MRgUT4cK0RwuLeZLRKyS_",
					"type": "arrow"
				}
			],
			"updated": 1704651740520,
			"link": null,
			"locked": false
		},
		{
			"type": "text",
			"version": 448,
			"versionNonce": 100264548,
			"isDeleted": false,
			"id": "fEKoGGOW",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -810.8724060058594,
			"y": 666.8146362304688,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 139.19989013671875,
			"height": 25,
			"seed": 2139433180,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": [],
			"updated": 1704651740520,
			"link": null,
			"locked": false,
			"fontSize": 20,
			"fontFamily": 1,
			"text": "phase II ends",
			"rawText": "phase II ends",
			"textAlign": "center",
			"verticalAlign": "middle",
			"containerId": "uh1Sa6tVIjsygaX_OS9mS",
			"originalText": "phase II ends",
			"lineHeight": 1.25,
			"baseline": 18
		},
		{
			"id": "53fRAEIQ9J8LjlUDTAxZz",
			"type": "rectangle",
			"x": -701.18994140625,
			"y": -182.13897705078125,
			"width": 208,
			"height": 85,
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
			"seed": 440676836,
			"version": 433,
			"versionNonce": 1589337564,
			"isDeleted": false,
			"boundElements": [
				{
					"id": "MHVXcAit",
					"type": "text"
				},
				{
					"id": "QgqVN3S79Dq0Iyp5aPUO2",
					"type": "arrow"
				},
				{
					"id": "mS0TzpS5gPD9O2-u1l1uY",
					"type": "arrow"
				},
				{
					"id": "gz5fzxYhWwaMyvqdJ8DDF",
					"type": "arrow"
				},
				{
					"id": "Ever_ytXq0If902POMHSm",
					"type": "arrow"
				}
			],
			"updated": 1704651740520,
			"link": null,
			"locked": false
		},
		{
			"id": "MHVXcAit",
			"type": "text",
			"x": -696.1698608398438,
			"y": -177.13897705078125,
			"width": 197.9598388671875,
			"height": 75,
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
			"seed": 279133156,
			"version": 538,
			"versionNonce": 60227044,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1704651740520,
			"link": null,
			"locked": false,
			"text": "node requests\nCA certificate chain\n(root CA + ICA)",
			"rawText": "node requests\nCA certificate chain\n(root CA + ICA)",
			"fontSize": 20,
			"fontFamily": 1,
			"textAlign": "center",
			"verticalAlign": "middle",
			"baseline": 68,
			"containerId": "53fRAEIQ9J8LjlUDTAxZz",
			"originalText": "node requests\nCA certificate chain\n(root CA + ICA)",
			"lineHeight": 1.25
		},
		{
			"type": "rectangle",
			"version": 525,
			"versionNonce": 1674265180,
			"isDeleted": false,
			"id": "1OG3EZDz14xGaxnFGtQJ-",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": 243.861328125,
			"y": -173.58123779296875,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 208,
			"height": 85,
			"seed": 332104420,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 3
			},
			"boundElements": [
				{
					"id": "wy4ta4TO",
					"type": "text"
				},
				{
					"id": "YWmvlL8S9960xqc8SW0y8",
					"type": "arrow"
				},
				{
					"id": "Y805G7pdZRz6Jzh16uBQE",
					"type": "arrow"
				},
				{
					"id": "4ogDsWAX3gNfkICbNLCvD",
					"type": "arrow"
				}
			],
			"updated": 1704651740520,
			"link": null,
			"locked": false
		},
		{
			"type": "text",
			"version": 691,
			"versionNonce": 54894948,
			"isDeleted": false,
			"id": "wy4ta4TO",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": 259.44140625,
			"y": -156.08123779296875,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 176.83984375,
			"height": 50,
			"seed": 77694564,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": [],
			"updated": 1704651740520,
			"link": null,
			"locked": false,
			"fontSize": 20,
			"fontFamily": 1,
			"text": "PKI Server (Peer)\nreceives request",
			"rawText": "PKI Server (Peer) receives request",
			"textAlign": "center",
			"verticalAlign": "middle",
			"containerId": "1OG3EZDz14xGaxnFGtQJ-",
			"originalText": "PKI Server (Peer) receives request",
			"lineHeight": 1.25,
			"baseline": 43
		},
		{
			"id": "fxRUePEvKVlzkzBDGWPyc",
			"type": "rectangle",
			"x": -379.65472412109375,
			"y": -152.59326171875,
			"width": 469,
			"height": 35,
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
			"seed": 1831113188,
			"version": 138,
			"versionNonce": 306035940,
			"isDeleted": false,
			"boundElements": [
				{
					"id": "NOrWANaU",
					"type": "text"
				},
				{
					"id": "Y805G7pdZRz6Jzh16uBQE",
					"type": "arrow"
				},
				{
					"id": "gz5fzxYhWwaMyvqdJ8DDF",
					"type": "arrow"
				}
			],
			"updated": 1704652021294,
			"link": null,
			"locked": false
		},
		{
			"id": "NOrWANaU",
			"type": "text",
			"x": -374.5245361328125,
			"y": -147.59326171875,
			"width": 458.7396240234375,
			"height": 25,
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
			"seed": 1491921500,
			"version": 306,
			"versionNonce": 750551140,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1704652021294,
			"link": null,
			"locked": false,
			"text": "[http/get] http://${pkiLocalLeader}/api/v1/ca",
			"rawText": "[http/get] http://${pkiLocalLeader}/api/v1/ca",
			"fontSize": 20,
			"fontFamily": 1,
			"textAlign": "center",
			"verticalAlign": "middle",
			"baseline": 18,
			"containerId": "fxRUePEvKVlzkzBDGWPyc",
			"originalText": "[http/get] http://${pkiLocalLeader}/api/v1/ca",
			"lineHeight": 1.25
		},
		{
			"id": "YP4DDqKE7YZdgfiaEc7AB",
			"type": "diamond",
			"x": -674.8341064453125,
			"y": -758.8992004394531,
			"width": 178.92938232421875,
			"height": 170,
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
				"type": 2
			},
			"seed": 1972800356,
			"version": 383,
			"versionNonce": 16341468,
			"isDeleted": false,
			"boundElements": [
				{
					"type": "text",
					"id": "kIgnjrMF"
				},
				{
					"id": "SQ1IjhLfrnQg4B9RZMYT7",
					"type": "arrow"
				},
				{
					"id": "mS0TzpS5gPD9O2-u1l1uY",
					"type": "arrow"
				},
				{
					"id": "WPELjVUuqxUq9Hpjpag42",
					"type": "arrow"
				}
			],
			"updated": 1704651761229,
			"link": null,
			"locked": false
		},
		{
			"id": "kIgnjrMF",
			"type": "text",
			"x": -623.8517303466797,
			"y": -711.3992004394531,
			"width": 76.49993896484375,
			"height": 75,
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
			"seed": 742323684,
			"version": 388,
			"versionNonce": 206542948,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1704651740520,
			"link": null,
			"locked": false,
			"text": "is PKI\nGlobal\nLeader?",
			"rawText": "is PKI Global Leader?",
			"fontSize": 20,
			"fontFamily": 1,
			"textAlign": "center",
			"verticalAlign": "middle",
			"baseline": 68,
			"containerId": "YP4DDqKE7YZdgfiaEc7AB",
			"originalText": "is PKI Global Leader?",
			"lineHeight": 1.25
		},
		{
			"id": "Ze5chHn8eJ8GJf5RxYEYa",
			"type": "rectangle",
			"x": -179.88156127929688,
			"y": -409.2638854980469,
			"width": 478,
			"height": 35,
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
			"seed": 2070417628,
			"version": 360,
			"versionNonce": 1782846692,
			"isDeleted": false,
			"boundElements": [
				{
					"id": "JgasGr1L",
					"type": "text"
				},
				{
					"id": "YWmvlL8S9960xqc8SW0y8",
					"type": "arrow"
				},
				{
					"id": "nsi7bM43e5NmHShNN9ANt",
					"type": "arrow"
				}
			],
			"updated": 1704651826659,
			"link": null,
			"locked": false
		},
		{
			"type": "text",
			"version": 601,
			"versionNonce": 1647827044,
			"isDeleted": false,
			"id": "JgasGr1L",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -153.5413818359375,
			"y": -404.2638854980469,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 425.31964111328125,
			"height": 25,
			"seed": 2010624604,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": null,
			"updated": 1704651826659,
			"link": null,
			"locked": false,
			"fontSize": 20,
			"fontFamily": 1,
			"text": "[http/get] http://${pkiGlobalVip}/api/v1/ca",
			"rawText": "[http/get] http://${pkiGlobalVip}/api/v1/ca",
			"textAlign": "center",
			"verticalAlign": "middle",
			"containerId": "Ze5chHn8eJ8GJf5RxYEYa",
			"originalText": "[http/get] http://${pkiGlobalVip}/api/v1/ca",
			"lineHeight": 1.25,
			"baseline": 18
		},
		{
			"type": "rectangle",
			"version": 909,
			"versionNonce": 751371868,
			"isDeleted": false,
			"id": "BXxqV5UKC-OgIXRbHvJ7f",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -416.3863525390625,
			"y": -450.69390869140625,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 174.49291992187497,
			"height": 110,
			"seed": 2123124060,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 3
			},
			"boundElements": [
				{
					"id": "t0MQADi3",
					"type": "text"
				},
				{
					"id": "nsi7bM43e5NmHShNN9ANt",
					"type": "arrow"
				},
				{
					"id": "JyxolEYiJcEq1zeHSIDey",
					"type": "arrow"
				}
			],
			"updated": 1704651839537,
			"link": null,
			"locked": false
		},
		{
			"type": "text",
			"version": 1056,
			"versionNonce": 1624250852,
			"isDeleted": false,
			"id": "t0MQADi3",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -410.39984130859375,
			"y": -445.69390869140625,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 162.5198974609375,
			"height": 100,
			"seed": 1255293404,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": [],
			"updated": 1704651826673,
			"link": null,
			"locked": false,
			"fontSize": 20,
			"fontFamily": 1,
			"text": "node requests\nCA certificate\nchain\n(root CA + ICA)",
			"rawText": "node requests\nCA certificate chain\n(root CA + ICA)",
			"textAlign": "center",
			"verticalAlign": "middle",
			"containerId": "BXxqV5UKC-OgIXRbHvJ7f",
			"originalText": "node requests\nCA certificate chain\n(root CA + ICA)",
			"lineHeight": 1.25,
			"baseline": 93
		},
		{
			"id": "YWmvlL8S9960xqc8SW0y8",
			"type": "arrow",
			"x": 304.44243257997147,
			"y": -392.49884816931046,
			"width": 55.871046022569715,
			"height": 214.9515154056386,
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
			"seed": 608462692,
			"version": 1130,
			"versionNonce": 1263155036,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1704651939374,
			"link": null,
			"locked": false,
			"points": [
				[
					0,
					0
				],
				[
					49.59272367002842,
					25.636299341185463
				],
				[
					55.871046022569715,
					214.9515154056386
				]
			],
			"lastCommittedPoint": null,
			"startBinding": {
				"elementId": "Ze5chHn8eJ8GJf5RxYEYa",
				"gap": 6.323993859268342,
				"focus": -0.9043166519365995
			},
			"endBinding": {
				"elementId": "1OG3EZDz14xGaxnFGtQJ-",
				"gap": 3.966094970703125,
				"focus": 0.13275016526199337
			},
			"startArrowhead": null,
			"endArrowhead": "arrow"
		},
		{
			"id": "Y805G7pdZRz6Jzh16uBQE",
			"type": "arrow",
			"x": 94.41503906250003,
			"y": -134.02981105382625,
			"width": 139.65661621093747,
			"height": 0.931232641061257,
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
			"seed": 112341732,
			"version": 246,
			"versionNonce": 1054521188,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1704652021295,
			"link": null,
			"locked": false,
			"points": [
				[
					0,
					0
				],
				[
					139.65661621093747,
					-0.931232641061257
				]
			],
			"lastCommittedPoint": null,
			"startBinding": {
				"elementId": "fxRUePEvKVlzkzBDGWPyc",
				"gap": 5.06976318359375,
				"focus": 0.14448778054739056
			},
			"endBinding": {
				"elementId": "1OG3EZDz14xGaxnFGtQJ-",
				"gap": 9.7896728515625,
				"focus": 0.10739023592069895
			},
			"startArrowhead": null,
			"endArrowhead": "arrow"
		},
		{
			"type": "rectangle",
			"version": 584,
			"versionNonce": 626704860,
			"isDeleted": false,
			"id": "Ohk0ESIHy8GzL204S88Gc",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": 237.50732421875,
			"y": -12.49200439453125,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 208,
			"height": 110,
			"seed": 1081913828,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 3
			},
			"boundElements": [
				{
					"id": "0zOUJMqH",
					"type": "text"
				},
				{
					"id": "4ogDsWAX3gNfkICbNLCvD",
					"type": "arrow"
				},
				{
					"id": "ExGXKdvQaw0w6aZVuYqWQ",
					"type": "arrow"
				},
				{
					"id": "vtHkM6VN9MH4LnMAvWFWj",
					"type": "arrow"
				}
			],
			"updated": 1704651740520,
			"link": null,
			"locked": false
		},
		{
			"type": "text",
			"version": 867,
			"versionNonce": 68147684,
			"isDeleted": false,
			"id": "0zOUJMqH",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": 252.0474090576172,
			"y": -7.49200439453125,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 178.91983032226562,
			"height": 100,
			"seed": 22336868,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": [],
			"updated": 1704651740520,
			"link": null,
			"locked": false,
			"fontSize": 20,
			"fontFamily": 1,
			"text": "PKI Server replies\nwith root and\nintermediate CA\ncertificate (ICA)",
			"rawText": "PKI Server replies with root and intermediate CA certificate (ICA)",
			"textAlign": "center",
			"verticalAlign": "middle",
			"containerId": "Ohk0ESIHy8GzL204S88Gc",
			"originalText": "PKI Server replies with root and intermediate CA certificate (ICA)",
			"lineHeight": 1.25,
			"baseline": 93
		},
		{
			"id": "4ogDsWAX3gNfkICbNLCvD",
			"type": "arrow",
			"x": 354.2694659587129,
			"y": -84.56176757812499,
			"width": 1.542809322753783,
			"height": 57.24999999999996,
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
			"seed": 1048302180,
			"version": 331,
			"versionNonce": 1515723612,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1704651939380,
			"link": null,
			"locked": false,
			"points": [
				[
					0,
					0
				],
				[
					-1.542809322753783,
					57.24999999999996
				]
			],
			"lastCommittedPoint": null,
			"startBinding": {
				"elementId": "1OG3EZDz14xGaxnFGtQJ-",
				"gap": 4.01947021484375,
				"focus": -0.072868428694044
			},
			"endBinding": {
				"elementId": "Ohk0ESIHy8GzL204S88Gc",
				"gap": 14.819763183593764,
				"focus": 0.0885247687627333
			},
			"startArrowhead": null,
			"endArrowhead": "arrow"
		},
		{
			"type": "rectangle",
			"version": 525,
			"versionNonce": 943914340,
			"isDeleted": false,
			"id": "s2jNlBvEL2p8GCNL9mMwS",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -697.111083984375,
			"y": -6.84637451171875,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 208,
			"height": 85,
			"seed": 1365457764,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 3
			},
			"boundElements": [
				{
					"id": "JTBt5ulv",
					"type": "text"
				},
				{
					"id": "ExGXKdvQaw0w6aZVuYqWQ",
					"type": "arrow"
				},
				{
					"id": "QgqVN3S79Dq0Iyp5aPUO2",
					"type": "arrow"
				},
				{
					"id": "kXIfKR7O0jLhawHuXRd6Z",
					"type": "arrow"
				}
			],
			"updated": 1704651740521,
			"link": null,
			"locked": false
		},
		{
			"type": "text",
			"version": 675,
			"versionNonce": 528546524,
			"isDeleted": false,
			"id": "JTBt5ulv",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -660.0010223388672,
			"y": -1.84637451171875,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 133.77987670898438,
			"height": 75,
			"seed": 397291236,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": [],
			"updated": 1704651740521,
			"link": null,
			"locked": false,
			"fontSize": 20,
			"fontFamily": 1,
			"text": "Certificate\ninformation is\nstored locally",
			"rawText": "Certificate information is stored locally",
			"textAlign": "center",
			"verticalAlign": "middle",
			"containerId": "s2jNlBvEL2p8GCNL9mMwS",
			"originalText": "Certificate information is stored locally",
			"lineHeight": 1.25,
			"baseline": 68
		},
		{
			"id": "ExGXKdvQaw0w6aZVuYqWQ",
			"type": "arrow",
			"x": 227.6983642578125,
			"y": 39.017145762169605,
			"width": 707.3667602539062,
			"height": 0.4777438996402239,
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
			"seed": 1146769764,
			"version": 341,
			"versionNonce": 393930332,
			"isDeleted": false,
			"boundElements": [
				{
					"type": "text",
					"id": "yRdpAvDC"
				}
			],
			"updated": 1704651939397,
			"link": null,
			"locked": false,
			"points": [
				[
					0,
					0
				],
				[
					-707.3667602539062,
					-0.4777438996402239
				]
			],
			"lastCommittedPoint": null,
			"startBinding": {
				"elementId": "Ohk0ESIHy8GzL204S88Gc",
				"gap": 9.8089599609375,
				"focus": 0.06199328614432181
			},
			"endBinding": {
				"elementId": "s2jNlBvEL2p8GCNL9mMwS",
				"gap": 9.44268798828125,
				"focus": 0.06598879996475268
			},
			"startArrowhead": null,
			"endArrowhead": "arrow"
		},
		{
			"id": "yRdpAvDC",
			"type": "text",
			"x": -145.87098693847656,
			"y": 89.39835844155331,
			"width": 144.93991088867188,
			"height": 25,
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
			"seed": 2113779684,
			"version": 14,
			"versionNonce": 332742492,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1704651740521,
			"link": null,
			"locked": false,
			"text": "HTTP/200 OK",
			"rawText": "HTTP/200 OK",
			"fontSize": 20,
			"fontFamily": 1,
			"textAlign": "center",
			"verticalAlign": "middle",
			"baseline": 18,
			"containerId": "ExGXKdvQaw0w6aZVuYqWQ",
			"originalText": "HTTP/200 OK",
			"lineHeight": 1.25
		},
		{
			"id": "QgqVN3S79Dq0Iyp5aPUO2",
			"type": "arrow",
			"x": -596.0258652127359,
			"y": -90.910400390625,
			"width": 2.955372905642207,
			"height": 66.59674072265625,
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
			"seed": 2000005092,
			"version": 170,
			"versionNonce": 1380056924,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1704651939397,
			"link": null,
			"locked": false,
			"points": [
				[
					0,
					0
				],
				[
					-2.955372905642207,
					66.59674072265625
				]
			],
			"lastCommittedPoint": null,
			"startBinding": {
				"elementId": "53fRAEIQ9J8LjlUDTAxZz",
				"gap": 6.22857666015625,
				"focus": -0.03141595850002132
			},
			"endBinding": {
				"elementId": "s2jNlBvEL2p8GCNL9mMwS",
				"gap": 17.467285156250007,
				"focus": -0.0805708825220882
			},
			"startArrowhead": null,
			"endArrowhead": "arrow"
		},
		{
			"type": "rectangle",
			"version": 568,
			"versionNonce": 1848439772,
			"isDeleted": false,
			"id": "YigqX26mGlO6XIkqUv3cV",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -695.0313720703125,
			"y": 135.07635498046875,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 208,
			"height": 85,
			"seed": 814034268,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 3
			},
			"boundElements": [
				{
					"id": "hpYkmtna",
					"type": "text"
				},
				{
					"id": "MKY8XSVdgSux_7Vmvebap",
					"type": "arrow"
				},
				{
					"id": "kXIfKR7O0jLhawHuXRd6Z",
					"type": "arrow"
				},
				{
					"id": "SSvxqGz-xUvoVLePx2w7g",
					"type": "arrow"
				}
			],
			"updated": 1704651740521,
			"link": null,
			"locked": false
		},
		{
			"type": "text",
			"version": 716,
			"versionNonce": 238143460,
			"isDeleted": false,
			"id": "hpYkmtna",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -657.9213104248047,
			"y": 140.07635498046875,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 133.77987670898438,
			"height": 75,
			"seed": 179249628,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": [],
			"updated": 1704651740521,
			"link": null,
			"locked": false,
			"fontSize": 20,
			"fontFamily": 1,
			"text": "Certificate\ninformation is\nstored locally",
			"rawText": "Certificate information is stored locally",
			"textAlign": "center",
			"verticalAlign": "middle",
			"containerId": "YigqX26mGlO6XIkqUv3cV",
			"originalText": "Certificate information is stored locally",
			"lineHeight": 1.25,
			"baseline": 68
		},
		{
			"id": "MKY8XSVdgSux_7Vmvebap",
			"type": "arrow",
			"x": -479.28466796875,
			"y": 363.1704415579576,
			"width": 114.69305419921875,
			"height": 1.4123961040043014,
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
			"seed": 1951271004,
			"version": 788,
			"versionNonce": 60527196,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1704651939416,
			"link": null,
			"locked": false,
			"points": [
				[
					0,
					0
				],
				[
					114.69305419921875,
					1.4123961040043014
				]
			],
			"lastCommittedPoint": null,
			"startBinding": {
				"elementId": "1L3_pJftpKFGjv2yvbNIx",
				"gap": 7.630126953125,
				"focus": 0.06999736663622264
			},
			"endBinding": {
				"elementId": "hOcJIxxgIG0Q-0sXUIQeh",
				"gap": 8.20623779296875,
				"focus": -0.12886803577550585
			},
			"startArrowhead": null,
			"endArrowhead": "arrow"
		},
		{
			"type": "rectangle",
			"version": 752,
			"versionNonce": 1140271972,
			"isDeleted": false,
			"id": "g3djWeIsq17mnQSVWqPBX",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": 240.8623046875,
			"y": 314.60504150390625,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 208,
			"height": 110,
			"seed": 340379868,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 3
			},
			"boundElements": [
				{
					"id": "3wef01oc",
					"type": "text"
				},
				{
					"id": "MKY8XSVdgSux_7Vmvebap",
					"type": "arrow"
				},
				{
					"id": "vtHkM6VN9MH4LnMAvWFWj",
					"type": "arrow"
				},
				{
					"id": "mMavSjGTbfnus-luK8Br6",
					"type": "arrow"
				},
				{
					"id": "lWZOeont2ynVPno_TdvWO",
					"type": "arrow"
				}
			],
			"updated": 1704651740521,
			"link": null,
			"locked": false
		},
		{
			"type": "text",
			"version": 1119,
			"versionNonce": 2062462172,
			"isDeleted": false,
			"id": "3wef01oc",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": 250.7323760986328,
			"y": 332.10504150390625,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 188.25985717773438,
			"height": 75,
			"seed": 1381256540,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": [],
			"updated": 1704651740521,
			"link": null,
			"locked": false,
			"fontSize": 20,
			"fontFamily": 1,
			"text": "PKI server signs\nthe CSR to create\nthe node ICA",
			"rawText": "PKI server signs the CSR to create the node ICA",
			"textAlign": "center",
			"verticalAlign": "middle",
			"containerId": "g3djWeIsq17mnQSVWqPBX",
			"originalText": "PKI server signs the CSR to create the node ICA",
			"lineHeight": 1.25,
			"baseline": 68
		},
		{
			"id": "vtHkM6VN9MH4LnMAvWFWj",
			"type": "arrow",
			"x": 351.3845446246146,
			"y": 101.5313720703125,
			"width": 1.6709857299899227,
			"height": 200.88653564453136,
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
			"seed": 237547868,
			"version": 583,
			"versionNonce": 386392924,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1704651939404,
			"link": null,
			"locked": false,
			"points": [
				[
					0,
					0
				],
				[
					1.6709857299899227,
					200.88653564453136
				]
			],
			"lastCommittedPoint": null,
			"startBinding": {
				"elementId": "Ohk0ESIHy8GzL204S88Gc",
				"gap": 4.02337646484375,
				"focus": -0.08985722603717837
			},
			"endBinding": {
				"elementId": "g3djWeIsq17mnQSVWqPBX",
				"gap": 12.187133789062386,
				"focus": 0.08378615871991041
			},
			"startArrowhead": null,
			"endArrowhead": "arrow"
		},
		{
			"type": "rectangle",
			"version": 288,
			"versionNonce": 1522315612,
			"isDeleted": false,
			"id": "hOcJIxxgIG0Q-0sXUIQeh",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -356.3853759765625,
			"y": 333.59967041015625,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 518.627197265625,
			"height": 60,
			"seed": 1663575900,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 3
			},
			"boundElements": [
				{
					"id": "ocJLmcs4",
					"type": "text"
				},
				{
					"id": "MKY8XSVdgSux_7Vmvebap",
					"type": "arrow"
				},
				{
					"id": "mMavSjGTbfnus-luK8Br6",
					"type": "arrow"
				}
			],
			"updated": 1704651740521,
			"link": null,
			"locked": false
		},
		{
			"type": "text",
			"version": 480,
			"versionNonce": 1101439588,
			"isDeleted": false,
			"id": "ocJLmcs4",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -337.2915802001953,
			"y": 351.09967041015625,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 480.4396057128906,
			"height": 25,
			"seed": 585292764,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": [],
			"updated": 1704651740521,
			"link": null,
			"locked": false,
			"fontSize": 20,
			"fontFamily": 1,
			"text": "[http/post] https://${pkiLocalLeader}/api/v1/ca",
			"rawText": "[http/post] https://${pkiLocalLeader}/api/v1/ca",
			"textAlign": "center",
			"verticalAlign": "middle",
			"containerId": "hOcJIxxgIG0Q-0sXUIQeh",
			"originalText": "[http/post] https://${pkiLocalLeader}/api/v1/ca",
			"lineHeight": 1.25,
			"baseline": 18
		},
		{
			"id": "mMavSjGTbfnus-luK8Br6",
			"type": "arrow",
			"x": 166.466552734375,
			"y": 367.2789877666307,
			"width": 60.978515625,
			"height": 0.30268803733548566,
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
			"seed": 1509247068,
			"version": 662,
			"versionNonce": 212453212,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1704651939412,
			"link": null,
			"locked": false,
			"points": [
				[
					0,
					0
				],
				[
					60.978515625,
					0.30268803733548566
				]
			],
			"lastCommittedPoint": null,
			"startBinding": {
				"elementId": "hOcJIxxgIG0Q-0sXUIQeh",
				"gap": 4.2247314453125,
				"focus": 0.07578674606508833
			},
			"endBinding": {
				"elementId": "g3djWeIsq17mnQSVWqPBX",
				"gap": 13.417236328125,
				"focus": 0.02594780321680155
			},
			"startArrowhead": null,
			"endArrowhead": "arrow"
		},
		{
			"type": "rectangle",
			"version": 818,
			"versionNonce": 418166244,
			"isDeleted": false,
			"id": "qkkgx2XZaEl28jIQ-Ath2",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": 244.88592529296875,
			"y": 476.06170654296875,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 208,
			"height": 135,
			"seed": 1536423652,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 3
			},
			"boundElements": [
				{
					"id": "wFuFRBTI",
					"type": "text"
				},
				{
					"id": "lWZOeont2ynVPno_TdvWO",
					"type": "arrow"
				},
				{
					"id": "oVowkxlKyAJcIbTuMHMA_",
					"type": "arrow"
				}
			],
			"updated": 1704651740521,
			"link": null,
			"locked": false
		},
		{
			"type": "text",
			"version": 1054,
			"versionNonce": 1103956572,
			"isDeleted": false,
			"id": "wFuFRBTI",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": 254.99600982666016,
			"y": 481.06170654296875,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 187.7798309326172,
			"height": 125,
			"seed": 591100516,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": [],
			"updated": 1704651740521,
			"link": null,
			"locked": false,
			"fontSize": 20,
			"fontFamily": 1,
			"text": "PKI server returns\nthe node\nIntermediate\nCertificate\nAuthority (ICA)",
			"rawText": "PKI server returns the node Intermediate Certificate Authority (ICA)",
			"textAlign": "center",
			"verticalAlign": "middle",
			"containerId": "qkkgx2XZaEl28jIQ-Ath2",
			"originalText": "PKI server returns the node Intermediate Certificate Authority (ICA)",
			"lineHeight": 1.25,
			"baseline": 118
		},
		{
			"type": "rectangle",
			"version": 637,
			"versionNonce": 1428445540,
			"isDeleted": false,
			"id": "1L3_pJftpKFGjv2yvbNIx",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -694.914794921875,
			"y": 302.85626220703125,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 208,
			"height": 110,
			"seed": 467074012,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 3
			},
			"boundElements": [
				{
					"id": "9jYaSHB8",
					"type": "text"
				},
				{
					"id": "MKY8XSVdgSux_7Vmvebap",
					"type": "arrow"
				},
				{
					"id": "SSvxqGz-xUvoVLePx2w7g",
					"type": "arrow"
				},
				{
					"id": "f9iX31ZEthHqDvRUCMrIB",
					"type": "arrow"
				}
			],
			"updated": 1704651740521,
			"link": null,
			"locked": false
		},
		{
			"type": "text",
			"version": 855,
			"versionNonce": 1089691356,
			"isDeleted": false,
			"id": "9jYaSHB8",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -688.1447143554688,
			"y": 307.85626220703125,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 194.4598388671875,
			"height": 100,
			"seed": 2009276508,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": [],
			"updated": 1704651740521,
			"link": null,
			"locked": false,
			"fontSize": 20,
			"fontFamily": 1,
			"text": "Generate Private\nKey and Certificate\nSigning Request\n(CSR)",
			"rawText": "Generate Private Key and Certificate Signing Request (CSR)",
			"textAlign": "center",
			"verticalAlign": "middle",
			"containerId": "1L3_pJftpKFGjv2yvbNIx",
			"originalText": "Generate Private Key and Certificate Signing Request (CSR)",
			"lineHeight": 1.25,
			"baseline": 93
		},
		{
			"type": "arrow",
			"version": 1072,
			"versionNonce": 581478364,
			"isDeleted": false,
			"id": "mS0TzpS5gPD9O2-u1l1uY",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -684.118189929489,
			"y": -675.2281557170397,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 210.60833751707833,
			"height": 487.1768388180335,
			"seed": 1521729244,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 2
			},
			"boundElements": [
				{
					"type": "text",
					"id": "MNH3eTFR"
				}
			],
			"updated": 1704651939425,
			"link": null,
			"locked": false,
			"startBinding": {
				"elementId": "YP4DDqKE7YZdgfiaEc7AB",
				"gap": 9.378716771018446,
				"focus": 0.372014265329856
			},
			"endBinding": {
				"elementId": "lxpWAkkRsdBK97KEKwGyI",
				"gap": 8.599412602131224,
				"focus": -0.07454304223899487
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
					-190.43332374238605,
					58.417486771727226
				],
				[
					-210.60833751707833,
					487.1768388180335
				]
			]
		},
		{
			"id": "MNH3eTFR",
			"type": "text",
			"x": -830.0474193994514,
			"y": -560.6060469263765,
			"width": 32.119964599609375,
			"height": 25,
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
			"roundness": null,
			"seed": 1970520676,
			"version": 18,
			"versionNonce": 559620956,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1704651740521,
			"link": null,
			"locked": false,
			"text": "Yes",
			"rawText": "Yes",
			"fontSize": 20,
			"fontFamily": 1,
			"textAlign": "center",
			"verticalAlign": "middle",
			"baseline": 18,
			"containerId": "mS0TzpS5gPD9O2-u1l1uY",
			"originalText": "Yes",
			"lineHeight": 1.25
		},
		{
			"type": "arrow",
			"version": 344,
			"versionNonce": 482963804,
			"isDeleted": false,
			"id": "kXIfKR7O0jLhawHuXRd6Z",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -597.3839586213198,
			"y": 79.15362548828125,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 0.38780436316164923,
			"height": 52.80548077374698,
			"seed": 1846762588,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 2
			},
			"boundElements": [],
			"updated": 1704651939400,
			"link": null,
			"locked": false,
			"startBinding": {
				"elementId": "s2jNlBvEL2p8GCNL9mMwS",
				"gap": 1,
				"focus": 0.03789981315242619
			},
			"endBinding": {
				"elementId": "YigqX26mGlO6XIkqUv3cV",
				"gap": 3.1172487184405213,
				"focus": -0.06782917198033092
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
					-0.38780436316164923,
					52.80548077374698
				]
			]
		},
		{
			"type": "arrow",
			"version": 403,
			"versionNonce": 388475740,
			"isDeleted": false,
			"id": "SSvxqGz-xUvoVLePx2w7g",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -600.4705473006012,
			"y": 222.04205303937198,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 0.1766192560780837,
			"height": 75.45831298828125,
			"seed": 447080924,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 2
			},
			"boundElements": [],
			"updated": 1704651939416,
			"link": null,
			"locked": false,
			"startBinding": {
				"elementId": "YigqX26mGlO6XIkqUv3cV",
				"gap": 1.9656980589032287,
				"focus": 0.09167435685165433
			},
			"endBinding": {
				"elementId": "1L3_pJftpKFGjv2yvbNIx",
				"gap": 5.355896179378021,
				"focus": -0.08871578981432555
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
					0.1766192560780837,
					75.45831298828125
				]
			]
		},
		{
			"type": "arrow",
			"version": 243,
			"versionNonce": 706162276,
			"isDeleted": false,
			"id": "gz5fzxYhWwaMyvqdJ8DDF",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -488.6607328474762,
			"y": -137.37428076588424,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 100.99909850412496,
			"height": 0.5780657357307746,
			"seed": 1129680228,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 2
			},
			"boundElements": [],
			"updated": 1704652021295,
			"link": null,
			"locked": false,
			"startBinding": {
				"elementId": "53fRAEIQ9J8LjlUDTAxZz",
				"gap": 4.529208558773803,
				"focus": 0.07050425216689543
			},
			"endBinding": {
				"elementId": "fxRUePEvKVlzkzBDGWPyc",
				"gap": 8.006910222257488,
				"focus": 0.2254023967415792
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
					100.99909850412496,
					-0.5780657357307746
				]
			]
		},
		{
			"type": "arrow",
			"version": 1268,
			"versionNonce": 334866780,
			"isDeleted": false,
			"id": "nsi7bM43e5NmHShNN9ANt",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -240.79992431577986,
			"y": -391.72670327275387,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 54.14674254709371,
			"height": 0.7754069300964375,
			"seed": 263967068,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 2
			},
			"boundElements": [],
			"updated": 1704651939378,
			"link": null,
			"locked": false,
			"startBinding": {
				"elementId": "BXxqV5UKC-OgIXRbHvJ7f",
				"gap": 1.0935083014076667,
				"focus": 0.04804352855479294
			},
			"endBinding": {
				"elementId": "Ze5chHn8eJ8GJf5RxYEYa",
				"gap": 6.7716204893892495,
				"focus": -0.20705630873105463
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
					54.14674254709371,
					0.7754069300964375
				]
			]
		},
		{
			"type": "arrow",
			"version": 824,
			"versionNonce": 535601244,
			"isDeleted": false,
			"id": "lWZOeont2ynVPno_TdvWO",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": 350.35620143554183,
			"y": 429.8184403583408,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 0.38726807156842824,
			"height": 39.65817260742176,
			"seed": 1623776996,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 2
			},
			"boundElements": [],
			"updated": 1704651939414,
			"link": null,
			"locked": false,
			"startBinding": {
				"elementId": "g3djWeIsq17mnQSVWqPBX",
				"gap": 5.213398854434502,
				"focus": -0.058179255892544406
			},
			"endBinding": {
				"elementId": "qkkgx2XZaEl28jIQ-Ath2",
				"gap": 6.58509357720618,
				"focus": 0.0034354941870271835
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
					-0.38726807156842824,
					39.65817260742176
				]
			]
		},
		{
			"type": "rectangle",
			"version": 721,
			"versionNonce": 1846481116,
			"isDeleted": false,
			"id": "DokAJgGgtcjtmXQliR0SC",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -698.160400390625,
			"y": 487.36370849609375,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 208,
			"height": 110,
			"seed": 1385550940,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 3
			},
			"boundElements": [
				{
					"id": "FPu6PrFd",
					"type": "text"
				},
				{
					"id": "f9iX31ZEthHqDvRUCMrIB",
					"type": "arrow"
				},
				{
					"id": "oVowkxlKyAJcIbTuMHMA_",
					"type": "arrow"
				},
				{
					"id": "j6DJunfGf0MYhN0rreJ3o",
					"type": "arrow"
				}
			],
			"updated": 1704651740521,
			"link": null,
			"locked": false
		},
		{
			"type": "text",
			"version": 995,
			"versionNonce": 181140196,
			"isDeleted": false,
			"id": "FPu6PrFd",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -683.3803253173828,
			"y": 504.86370849609375,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 178.43984985351562,
			"height": 75,
			"seed": 303229148,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": [],
			"updated": 1704651740521,
			"link": null,
			"locked": false,
			"fontSize": 20,
			"fontFamily": 1,
			"text": "New Node stores\nICA Certificate in\nmemory",
			"rawText": "New Node stores ICA Certificate in memory",
			"textAlign": "center",
			"verticalAlign": "middle",
			"containerId": "DokAJgGgtcjtmXQliR0SC",
			"originalText": "New Node stores ICA Certificate in memory",
			"lineHeight": 1.25,
			"baseline": 68
		},
		{
			"type": "arrow",
			"version": 610,
			"versionNonce": 1534664028,
			"isDeleted": false,
			"id": "f9iX31ZEthHqDvRUCMrIB",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -601.8146208604874,
			"y": 416.35038916319604,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 2.261587418991553,
			"height": 60.04083251953125,
			"seed": 1649395684,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 2
			},
			"boundElements": [],
			"updated": 1704651939419,
			"link": null,
			"locked": false,
			"startBinding": {
				"elementId": "1L3_pJftpKFGjv2yvbNIx",
				"gap": 3.4941269561647914,
				"focus": 0.08198697728204514
			},
			"endBinding": {
				"elementId": "DokAJgGgtcjtmXQliR0SC",
				"gap": 10.972486813366459,
				"focus": -0.11690983023650633
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
					-2.261587418991553,
					60.04083251953125
				]
			]
		},
		{
			"type": "arrow",
			"version": 809,
			"versionNonce": 2002679388,
			"isDeleted": false,
			"id": "oVowkxlKyAJcIbTuMHMA_",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": 239.92313928726367,
			"y": 541.1885667899369,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 718.0457928532969,
			"height": 0.10876359076246445,
			"seed": 1652859492,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 2
			},
			"boundElements": [
				{
					"type": "text",
					"id": "uWNLgcew"
				}
			],
			"updated": 1704651939419,
			"link": null,
			"locked": false,
			"startBinding": {
				"elementId": "qkkgx2XZaEl28jIQ-Ath2",
				"gap": 4.962786005705084,
				"focus": 0.035393880974993165
			},
			"endBinding": {
				"elementId": "DokAJgGgtcjtmXQliR0SC",
				"gap": 12.03774682459175,
				"focus": -0.019063661345060493
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
					-718.0457928532969,
					0.10876359076246445
				]
			]
		},
		{
			"id": "uWNLgcew",
			"type": "text",
			"x": -137.97370428293948,
			"y": 649.2117916922793,
			"width": 144.93991088867188,
			"height": 25,
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
			"roundness": null,
			"seed": 1946791132,
			"version": 3,
			"versionNonce": 1957530076,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1704651740521,
			"link": null,
			"locked": false,
			"text": "HTTP/200 OK",
			"rawText": "HTTP/200 OK",
			"fontSize": 20,
			"fontFamily": 1,
			"textAlign": "center",
			"verticalAlign": "middle",
			"baseline": 18,
			"containerId": "oVowkxlKyAJcIbTuMHMA_",
			"originalText": "HTTP/200 OK",
			"lineHeight": 1.25
		},
		{
			"type": "arrow",
			"version": 1050,
			"versionNonce": 444698460,
			"isDeleted": false,
			"id": "j6DJunfGf0MYhN0rreJ3o",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -609.9195594004977,
			"y": 601.3295072630049,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 0.6079394792944868,
			"height": 55.68157958984375,
			"seed": 1189844316,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 2
			},
			"boundElements": [],
			"updated": 1704651939420,
			"link": null,
			"locked": false,
			"startBinding": {
				"elementId": "DokAJgGgtcjtmXQliR0SC",
				"gap": 3.9657987669111208,
				"focus": 0.15681527713181095
			},
			"endBinding": {
				"elementId": "uh1Sa6tVIjsygaX_OS9mS",
				"gap": 4.803549377620129,
				"focus": 0.5967627872454965
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
					0.6079394792944868,
					55.68157958984375
				]
			]
		},
		{
			"id": "adNlWsId",
			"type": "text",
			"x": -364.409423828125,
			"y": 408.80157470703125,
			"width": 524.259521484375,
			"height": 25,
			"angle": 0,
			"strokeColor": "#f08c00",
			"backgroundColor": "transparent",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"seed": 1824144100,
			"version": 149,
			"versionNonce": 1385688668,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1704651740521,
			"link": null,
			"locked": false,
			"text": "Request is authenticated with Pre-Shared Key (PSK)",
			"rawText": "Request is authenticated with Pre-Shared Key (PSK)",
			"fontSize": 20,
			"fontFamily": 1,
			"textAlign": "center",
			"verticalAlign": "top",
			"baseline": 18,
			"containerId": null,
			"originalText": "Request is authenticated with Pre-Shared Key (PSK)",
			"lineHeight": 1.25
		},
		{
			"type": "arrow",
			"version": 724,
			"versionNonce": 1406518620,
			"isDeleted": false,
			"id": "SQ1IjhLfrnQg4B9RZMYT7",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -581.8823407951248,
			"y": -804.8918788488044,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 0.5905367969329518,
			"height": 47.36528862366947,
			"seed": 581665508,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 2
			},
			"boundElements": [],
			"updated": 1704651939371,
			"link": null,
			"locked": false,
			"startBinding": {
				"elementId": "Ugea7KMJV-9Jc45-E_z8T",
				"gap": 9.599042171703559,
				"focus": 0.03216732404799747
			},
			"endBinding": {
				"elementId": "YP4DDqKE7YZdgfiaEc7AB",
				"gap": 1,
				"focus": 0.020722090707975514
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
					-0.5905367969329518,
					47.36528862366947
				]
			]
		},
		{
			"type": "text",
			"version": 223,
			"versionNonce": 1052885724,
			"isDeleted": false,
			"id": "Rt6aJRrJ",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -348.3146057128906,
			"y": 223.65631103515625,
			"strokeColor": "#f08c00",
			"backgroundColor": "transparent",
			"width": 480.25958251953125,
			"height": 25,
			"seed": 1658878556,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": [],
			"updated": 1704651740521,
			"link": null,
			"locked": false,
			"fontSize": 20,
			"fontFamily": 1,
			"text": "Below this point we use HTTPS for all requests",
			"rawText": "Below this point we use HTTPS for all requests",
			"textAlign": "center",
			"verticalAlign": "top",
			"containerId": null,
			"originalText": "Below this point we use HTTPS for all requests",
			"lineHeight": 1.25,
			"baseline": 18
		},
		{
			"type": "text",
			"version": 324,
			"versionNonce": 1498769636,
			"isDeleted": false,
			"id": "iPtbjOoh",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -344.6457977294922,
			"y": 167.8970947265625,
			"strokeColor": "#f08c00",
			"backgroundColor": "transparent",
			"width": 469.8595886230469,
			"height": 25,
			"seed": 1236795868,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": [],
			"updated": 1704651740521,
			"link": null,
			"locked": false,
			"fontSize": 20,
			"fontFamily": 1,
			"text": "Above this point we use HTTP for all requests",
			"rawText": "Above this point we use HTTP for all requests",
			"textAlign": "center",
			"verticalAlign": "top",
			"containerId": null,
			"originalText": "Above this point we use HTTP for all requests",
			"lineHeight": 1.25,
			"baseline": 18
		},
		{
			"id": "OZhwWBiQ6CKk_-2le34Xl",
			"type": "line",
			"x": -372.1258544921875,
			"y": 208.7493896484375,
			"width": 517.3455810546875,
			"height": 3.1221923828125,
			"angle": 0,
			"strokeColor": "#f08c00",
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
			"seed": 490448220,
			"version": 50,
			"versionNonce": 532481884,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1704651740521,
			"link": null,
			"locked": false,
			"points": [
				[
					0,
					0
				],
				[
					517.3455810546875,
					-3.1221923828125
				]
			],
			"lastCommittedPoint": null,
			"startBinding": null,
			"endBinding": null,
			"startArrowhead": null,
			"endArrowhead": null
		},
		{
			"type": "diamond",
			"version": 494,
			"versionNonce": 1147500516,
			"isDeleted": false,
			"id": "yHkkhbd38L1IRedreNoV7",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -671.1698303222656,
			"y": -479.85455322265625,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 178.92938232421875,
			"height": 170,
			"seed": 1845753308,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 2
			},
			"boundElements": [
				{
					"type": "text",
					"id": "MaBv6rG2"
				},
				{
					"id": "Ever_ytXq0If902POMHSm",
					"type": "arrow"
				},
				{
					"id": "WPELjVUuqxUq9Hpjpag42",
					"type": "arrow"
				},
				{
					"id": "JyxolEYiJcEq1zeHSIDey",
					"type": "arrow"
				}
			],
			"updated": 1704651836507,
			"link": null,
			"locked": false
		},
		{
			"type": "text",
			"version": 438,
			"versionNonce": 970701788,
			"isDeleted": false,
			"id": "MaBv6rG2",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -620.1874542236328,
			"y": -432.35455322265625,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 76.49993896484375,
			"height": 75,
			"seed": 776002140,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": [],
			"updated": 1704651740521,
			"link": null,
			"locked": false,
			"fontSize": 20,
			"fontFamily": 1,
			"text": "is PKI\nLocal\nLeader?",
			"rawText": "is PKI Local Leader?",
			"textAlign": "center",
			"verticalAlign": "middle",
			"containerId": "yHkkhbd38L1IRedreNoV7",
			"originalText": "is PKI Local Leader?",
			"lineHeight": 1.25,
			"baseline": 68
		},
		{
			"type": "arrow",
			"version": 316,
			"versionNonce": 1618110556,
			"isDeleted": false,
			"id": "Ever_ytXq0If902POMHSm",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -585.1437885009373,
			"y": -308.49908091260875,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 2.059416120934202,
			"height": 114.03655631526198,
			"seed": 1939993948,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 2
			},
			"boundElements": [
				{
					"type": "text",
					"id": "heL0DrP8"
				}
			],
			"updated": 1704651939422,
			"link": null,
			"locked": false,
			"startBinding": {
				"elementId": "yHkkhbd38L1IRedreNoV7",
				"gap": 3.3511628071232877,
				"focus": 0.021004187927299068
			},
			"endBinding": {
				"elementId": "53fRAEIQ9J8LjlUDTAxZz",
				"gap": 12.323547546565521,
				"focus": 0.08587264502703434
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
					-2.059416120934202,
					114.03655631526198
				]
			]
		},
		{
			"id": "heL0DrP8",
			"type": "text",
			"x": -601.9667920425534,
			"y": -239.8142396852374,
			"width": 32.119964599609375,
			"height": 25,
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
			"roundness": null,
			"seed": 1989957860,
			"version": 5,
			"versionNonce": 1999626332,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1704651740521,
			"link": null,
			"locked": false,
			"text": "Yes",
			"rawText": "Yes",
			"fontSize": 20,
			"fontFamily": 1,
			"textAlign": "center",
			"verticalAlign": "middle",
			"baseline": 18,
			"containerId": "Ever_ytXq0If902POMHSm",
			"originalText": "Yes",
			"lineHeight": 1.25
		},
		{
			"type": "rectangle",
			"version": 568,
			"versionNonce": 117192028,
			"isDeleted": false,
			"id": "lxpWAkkRsdBK97KEKwGyI",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -993.2294311523438,
			"y": -179.451904296875,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 208,
			"height": 85,
			"seed": 1304486236,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 3
			},
			"boundElements": [
				{
					"id": "Z3HEkdhLD-HDBRGNXQgPt",
					"type": "arrow"
				},
				{
					"type": "text",
					"id": "PtyD9RWE"
				},
				{
					"id": "mS0TzpS5gPD9O2-u1l1uY",
					"type": "arrow"
				}
			],
			"updated": 1704651793744,
			"link": null,
			"locked": false
		},
		{
			"id": "PtyD9RWE",
			"type": "text",
			"x": -985.7493438720703,
			"y": -161.951904296875,
			"width": 193.03982543945312,
			"height": 50,
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
			"roundness": null,
			"seed": 1945435228,
			"version": 4,
			"versionNonce": 820206180,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1704651740521,
			"link": null,
			"locked": false,
			"text": "create self-signed\nroot CA certificate",
			"rawText": "create self-signed root CA certificate",
			"fontSize": 20,
			"fontFamily": 1,
			"textAlign": "center",
			"verticalAlign": "middle",
			"baseline": 43,
			"containerId": "lxpWAkkRsdBK97KEKwGyI",
			"originalText": "create self-signed root CA certificate",
			"lineHeight": 1.25
		},
		{
			"type": "rectangle",
			"version": 583,
			"versionNonce": 2095851228,
			"isDeleted": false,
			"id": "hrosFFmtAWWNsKJFlqXCV",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -991.0551147460938,
			"y": -14.72674560546875,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 208,
			"height": 85,
			"seed": 185327324,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 3
			},
			"boundElements": [
				{
					"id": "Z3HEkdhLD-HDBRGNXQgPt",
					"type": "arrow"
				},
				{
					"id": "MRgUT4cK0RwuLeZLRKyS_",
					"type": "arrow"
				},
				{
					"type": "text",
					"id": "DY1J9YvV"
				}
			],
			"updated": 1704651740521,
			"link": null,
			"locked": false
		},
		{
			"id": "DY1J9YvV",
			"type": "text",
			"x": -952.1150588989258,
			"y": -9.72674560546875,
			"width": 130.11988830566406,
			"height": 75,
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
			"roundness": null,
			"seed": 254951772,
			"version": 4,
			"versionNonce": 879969508,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1704651740521,
			"link": null,
			"locked": false,
			"text": "create ICA\ncertificate\nfrom root CA",
			"rawText": "create ICA certificate \nfrom root CA",
			"fontSize": 20,
			"fontFamily": 1,
			"textAlign": "center",
			"verticalAlign": "middle",
			"baseline": 68,
			"containerId": "hrosFFmtAWWNsKJFlqXCV",
			"originalText": "create ICA certificate \nfrom root CA",
			"lineHeight": 1.25
		},
		{
			"type": "arrow",
			"version": 624,
			"versionNonce": 1661821148,
			"isDeleted": false,
			"id": "Z3HEkdhLD-HDBRGNXQgPt",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -896.3180087726371,
			"y": -91.56375170718597,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 0.20916292018750937,
			"height": 70.29687524280887,
			"seed": 1075133916,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 2
			},
			"boundElements": [],
			"updated": 1704651939428,
			"link": null,
			"locked": false,
			"startBinding": {
				"elementId": "lxpWAkkRsdBK97KEKwGyI",
				"gap": 2.888152589689014,
				"focus": 0.06677965312754476
			},
			"endBinding": {
				"elementId": "hrosFFmtAWWNsKJFlqXCV",
				"gap": 6.5401308589083555,
				"focus": -0.09236818850271065
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
					-0.20916292018750937,
					70.29687524280887
				]
			]
		},
		{
			"type": "arrow",
			"version": 740,
			"versionNonce": 1515706844,
			"isDeleted": false,
			"id": "MRgUT4cK0RwuLeZLRKyS_",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -899.8393808064043,
			"y": 77.56540407169888,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 2.797074925774723,
			"height": 568.6210329576527,
			"seed": 1468417500,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 2
			},
			"boundElements": [],
			"updated": 1704651939429,
			"link": null,
			"locked": false,
			"startBinding": {
				"elementId": "hrosFFmtAWWNsKJFlqXCV",
				"gap": 7.292149677167629,
				"focus": 0.12502939909837135
			},
			"endBinding": {
				"elementId": "uh1Sa6tVIjsygaX_OS9mS",
				"gap": 15.62819920111724,
				"focus": -0.7027334762339228
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
					2.797074925774723,
					568.6210329576527
				]
			]
		},
		{
			"type": "arrow",
			"version": 608,
			"versionNonce": 1086390620,
			"isDeleted": false,
			"id": "WPELjVUuqxUq9Hpjpag42",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -582.1267208659392,
			"y": -587.0932927228109,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 1.2053512983590053,
			"height": 103.15655951473565,
			"seed": 1103575772,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 2
			},
			"boundElements": [
				{
					"type": "text",
					"id": "KHsHIK4l"
				}
			],
			"updated": 1704651939423,
			"link": null,
			"locked": false,
			"startBinding": {
				"elementId": "YP4DDqKE7YZdgfiaEc7AB",
				"gap": 3.5427416713945803,
				"focus": -0.02490809737148328
			},
			"endBinding": {
				"elementId": "yHkkhbd38L1IRedreNoV7",
				"gap": 4.156740093823842,
				"focus": 0.0203953789863024
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
					1.2053512983590053,
					103.15655951473565
				]
			]
		},
		{
			"id": "KHsHIK4l",
			"type": "text",
			"x": -593.5093910751625,
			"y": -548.4735416248441,
			"width": 23.959976196289062,
			"height": 25,
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
			"roundness": null,
			"seed": 1202505188,
			"version": 3,
			"versionNonce": 1237665756,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1704651772811,
			"link": null,
			"locked": false,
			"text": "No",
			"rawText": "No",
			"fontSize": 20,
			"fontFamily": 1,
			"textAlign": "center",
			"verticalAlign": "middle",
			"baseline": 18,
			"containerId": "WPELjVUuqxUq9Hpjpag42",
			"originalText": "No",
			"lineHeight": 1.25
		},
		{
			"type": "arrow",
			"version": 289,
			"versionNonce": 251484764,
			"isDeleted": false,
			"id": "JyxolEYiJcEq1zeHSIDey",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -491.1322912591392,
			"y": -397.23422748828216,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "transparent",
			"width": 66.55232116037496,
			"height": 0.0336995707806409,
			"seed": 1160700516,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 2
			},
			"boundElements": [
				{
					"type": "text",
					"id": "zdQTbwZs"
				}
			],
			"updated": 1704651939423,
			"link": null,
			"locked": false,
			"startBinding": {
				"elementId": "yHkkhbd38L1IRedreNoV7",
				"gap": 2.4884629095697193,
				"focus": -0.028535728469787927
			},
			"endBinding": {
				"elementId": "BXxqV5UKC-OgIXRbHvJ7f",
				"gap": 8.193617559701764,
				"focus": 0.026493119462375276
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
					66.55232116037496,
					0.0336995707806409
				]
			]
		},
		{
			"id": "zdQTbwZs",
			"type": "text",
			"x": -469.83611877709626,
			"y": -409.71737770289184,
			"width": 23.959976196289062,
			"height": 25,
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
			"roundness": null,
			"seed": 1103968220,
			"version": 7,
			"versionNonce": 875154532,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1704651844947,
			"link": null,
			"locked": false,
			"text": "No",
			"rawText": "No",
			"fontSize": 20,
			"fontFamily": 1,
			"textAlign": "center",
			"verticalAlign": "middle",
			"baseline": 18,
			"containerId": "JyxolEYiJcEq1zeHSIDey",
			"originalText": "No",
			"lineHeight": 1.25
		},
		{
			"id": "Nnqk8d0w",
			"type": "text",
			"x": -825.587646484375,
			"y": -640.2745971679688,
			"width": 160.8798828125,
			"height": 100,
			"angle": 0,
			"strokeColor": "#f08c00",
			"backgroundColor": "transparent",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"seed": 2134308580,
			"version": 194,
			"versionNonce": 1920256100,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1704651908035,
			"link": null,
			"locked": false,
			"text": "We need a new\nroot CA and\nICA for the new\ncluster",
			"rawText": "We need a new\nroot CA and\nICA for the new\ncluster",
			"fontSize": 20,
			"fontFamily": 1,
			"textAlign": "center",
			"verticalAlign": "top",
			"baseline": 93,
			"containerId": null,
			"originalText": "We need a new\nroot CA and\nICA for the new\ncluster",
			"lineHeight": 1.25
		},
		{
			"type": "text",
			"version": 286,
			"versionNonce": 137723996,
			"isDeleted": false,
			"id": "5ZvpBRRi",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -528.7933883666992,
			"y": -547.8002319335938,
			"strokeColor": "#f08c00",
			"backgroundColor": "transparent",
			"width": 176.23985290527344,
			"height": 75,
			"seed": 1157372124,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": [],
			"updated": 1704651936073,
			"link": null,
			"locked": false,
			"fontSize": 20,
			"fontFamily": 1,
			"text": "We only need to\ninitialize the node\nin the trust chain",
			"rawText": "We only need to\ninitialize the node\nin the trust chain",
			"textAlign": "center",
			"verticalAlign": "top",
			"containerId": null,
			"originalText": "We only need to\ninitialize the node\nin the trust chain",
			"lineHeight": 1.25,
			"baseline": 68
		},
		{
			"type": "text",
			"version": 482,
			"versionNonce": 551329244,
			"isDeleted": false,
			"id": "LWQLLahI",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "dashed",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -549.7114562988281,
			"y": -312.65338134765625,
			"strokeColor": "#f08c00",
			"backgroundColor": "transparent",
			"width": 185.33990478515625,
			"height": 100,
			"seed": 1956518236,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": [],
			"updated": 1704652011092,
			"link": null,
			"locked": false,
			"fontSize": 20,
			"fontFamily": 1,
			"text": "We have a new AZ\nand a new AZ PKI\nand we just need\nan initialization",
			"rawText": "We have a new AZ\nand a new AZ PKI\nand we just need\nan initialization",
			"textAlign": "center",
			"verticalAlign": "top",
			"containerId": null,
			"originalText": "We have a new AZ\nand a new AZ PKI\nand we just need\nan initialization",
			"lineHeight": 1.25,
			"baseline": 93
		}
	],
	"appState": {
		"theme": "dark",
		"viewBackgroundColor": "#ffffff",
		"currentItemStrokeColor": "#f08c00",
		"currentItemBackgroundColor": "transparent",
		"currentItemFillStyle": "solid",
		"currentItemStrokeWidth": 2,
		"currentItemStrokeStyle": "dashed",
		"currentItemRoughness": 1,
		"currentItemOpacity": 100,
		"currentItemFontFamily": 1,
		"currentItemFontSize": 20,
		"currentItemTextAlign": "center",
		"currentItemStartArrowhead": null,
		"currentItemEndArrowhead": "arrow",
		"scrollX": 1274.5,
		"scrollY": 869.5390625,
		"zoom": {
			"value": 1
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