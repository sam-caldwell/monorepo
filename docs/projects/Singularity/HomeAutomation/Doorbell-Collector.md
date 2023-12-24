## MVP Features
1. Listen on https://0.0.0.0:443
2. When HTTP/PUT is received...
   1. Verify the UUID is a legit doorbell device.
   2. Push the UUID doorbell notification message to Kafka doorbell-events topic
3. When HTTP/GET is received...
   1. Respond with HTTP/200 OK (healthcheck)
4. The microservice must communicate over TLS using PKI-issued certificates.

## Nice-to-Have (Future)
none

## Language
1. Prototype is written in C.
2. We need to move this off ESP-32 and implement as microservice on K8s cluster.

## ToDo List
1. We need to move the code base into the `monorepo`.
2. we need to document the solution in the `monorepo`.
