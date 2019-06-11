# Home API

- Heartbeat server
- Heartbeat client code

## Installation

- Create Heroku instance
- [Enable free tier LogDNA](https://elements.heroku.com/addons/logdna)
- [Create ifttt webhook](https://ifttt.com/maker_webhooks) then name it like `missing_heart_beat`.
- Create LogDNA Alert > Add Webhook > Absence > When less than `1` matches appear within `15 minutes`.
  Webhook url should be like `https://maker.ifttt.com/trigger/missing_heart_beat/with/key/************` Mind the `missing_heart_beat` part.

## Client Installation

- ssh to RPI
- `crontab -e` then add `* * * * * /path/to/hartbeat.sh`.
