## Objectives

1. To automate interactions with lights, locks and other systems.
2. To collect information about the home environment and detect failures/catastrophes quickly.
3. To improve home security by integrating security cameras, SIGINT data and other feeds into a common threat-intel
   view.

### Projects

| Project Name                                | Status  | In Repo | Description                                                        |
|---------------------------------------------|---------|---------|--------------------------------------------------------------------|
| [Doorbell](Doorbell.md)                     | p-type  | no      | ESP-32 device needed to operate doorbell.                          |
| [Doorbell-Collector](Doorbell-Collector.md) | p-type  | no      | Microservice which serves the doorbell devices                     |
| [Chime](Chime.md)                           | p-type  | no      | ESP-32 device & software to ring chime (e.g. for doorbell)         |
| Notify                                      | p-type  | no      | Microservice send message to subscribers                           |
| Flood detector                              | p-type  | no      | ESP-32: detect water in basement, etc.                       |
| Thermostat probe                            | p-type  | no      | ESP-32: measure and send temperature to Thermostat Collector. |
| Thermostat collector                        | p-type  | no      | Microservice to collect metrics from thermostat probe              |
| Thermostat actor                            | p-type  | no      | ESP-32 device to turn on/off HVAC unit                             |
| Joshua                                      | p-type  | no      | Voice Activated Human interface microservice                       |
| Joshua Ears                                 | p-type  | no      | ESP-32 device to listen for audio queues for Joshua                |
| Joshua Voice                                | p-type  | no      | Rock64 device to play sounds for Joshua voice/prompts              |
| Sigint-Wifi                                 | p-type  | no      | ESP-32 device to listen for WIFI-enabled devices                   |
| Sigint-Blue                                 | p-type  | no      | ESP-32 device to listen for Bluetooth-enable devices               |
| Sigint Collector                            | p-type  | no      | Rock64 device & software for collecting data from SIGINT devices   |
| Camera                                      | p-type  | no      | ESP-32-CAM devices for collecting motion-activated video captures  |
| Camera-Collector                            | p-type  | no      | Microservice for collecting camera feeds from Camera               |
| Plate-Reader                                | p-type  | no      | Microservice for finding & reading license plates from video streams |
| NVR-collector                               | p-type  | no      | Microservice for receiving video streams from ONVIF cameras        |
| NVR-logger                                  | p-type  | no      | Microservice for logging NVR video stream data with metadata tagging |
| NVR-viewer                                  | p-type  | no      | web UI for viewing video streams (NEEDS A LOT OF WORK!)            |
| Mosaic                                      | planned | no      | microservice for integrating sigint & camera intel                 |
| Weather                                     | p-type  | no      | ESP-32 device & software for collecting temperature, bar. pressure | 

