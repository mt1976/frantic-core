# messageHelpers

`messageHelpers` defines shared message types for inter-service communication — user, session, behaviour, authority, and translation messages.

## Types

| Type | Description |
|---|---|
| `UserMessage` | User identity payload with Key, Code, Source, Locale, Theme, Timezone, Role, and Payload fields |
| `SessionMessage` | Session payload with SessionID, Expiry, User, Token, Locale, and Payload fields |
| `BehaviourMessage` | Behaviour/action payload with Key, Source, and Payload fields |
| `AuthorityMessage` | Authority envelope with Key, User, Behaviour, Payload, and Source fields |
| `GrantMessage` | Grant request/response containing User and Behaviour |
| `RevokeMessage` | Revoke request/response containing User and Behaviour |
| `DeclareMessage` | Domain behaviour declaration with Domain and Behaviour fields |
| `TranslationMessage` | Translation request/response with Text, Locale, Origin, Translation, and Payload fields |

## Methods

Each message type provides `Request(...)` and `Response(...)` builder methods.

- `UserMessage` and `BehaviourMessage` also have `Validate(*log.Logger) error`
- `TranslationMessage` has `ReponseWithPayload(...)` for payload-carrying responses

## Example

```go
import "github.com/mt1976/frantic-core/messageHelpers"

func main() {
    user := messageHelpers.UserMessage{}.Request("key", "code", "source")
    session := messageHelpers.SessionMessage{}.Request("sessionID", time.Now(), user)
    _ = session
}
```
