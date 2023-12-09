Oauth Authorization Flow
========================


```plantuml
@startuml
participant Browser as Browser
participant Server as Server
participant Identity Provider (IdP) as IdP

Browser->Server: GET:/my-data (response_type=code, client_id, redirect_uri, scope)
Server->Browser: Redirect to IdP (http/302 redirect_uri)

Browser->IdP: GET:/idp_uri (grant_type=authorization_code, code_from_prev_step, redirect_uri, client_id, client_secret)
IdP->Browser: Http Response code and authorization code once user is authenticated.

Browser->Server: Request for content (with authorization code from IdP)
Server->IdP: Http Request with authorization code (Is this valid?)
Idp->Server: Yes: Here's the access token

Server->Idp: Give me the user's profile (sent with access token)
Idp->Server: Here is the user profile

Server->Browser: Here is the data you requested.
@enduml
```
