# telegrapher
Simple command line tool to send notifications via Telegram.

Sending message:
```bash
telegrapher -token <bot_token> -chat <chat_id> message_text
```
Example:
```bash
telegrapher -token 987654321:BBBBB-aaaaaaaabbbbbbbbccccccdddddee \
    -chat 123456789 \
    "Some process ended"
```


You can store token and chat in config file _~/.config/telegrapher.json_
```json
{
    "token": "987654321:BBBBB-aaaaaaaabbbbbbbbccccccdddddee",
    "chat": "123456789"
}
```
Example:
```bash
telegrapher "Message with chat and token in config"
```
