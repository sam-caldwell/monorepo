## MVP Features
1. Given a push button, when the circuit closes, a signal should be sent to the ESP-32 device.
2. The software running on the ESP-32 device should detect the signal and emit an `HTTP/PUT` to the `doorbell-collector`
   microservice.
3. Each doorbell will be identified by a UUID which will be sent with the `HTTP/PUT` request to identify which doorbell
   was pressed.
4. The `doorbell` software should rate limit the number of HTTP requests emitted to thwart assholes who (a) hold down
   the button or who repeatedly press the button.
5. Doorbell should use TLS through internal PKI.

## Nice-to-Have (Future)
1. ESP-32 should have connected speaker and sound file to provide feedback to a person pressing the button.
2. When a user presses the button, an audio file should play a polite feedback sound.
3. When user exceeds the rate-limit threshold, an audio file should tell them they are being an asshole.

## Language
1. Prototype is written in C.

## ToDo List
1. We need to implement a CI/CD deployment process.
2. We need to implement the test board for testing changes to Doorbell.
3. We need to move the code base into the `monorepo`.
4. we need to document the solution in the `monorepo`.
