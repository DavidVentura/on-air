The code at `client` will emit an event if you join a meeting, by detecting camera/microphone usage; works on Linux and Mac.

The code on `sign` is intended to be flashed on an ESP32, but is trivial to implement on any microcontroller.


```bash
cd sign
cmake . -DWIFI_SSID=<ssid> -DWIFI_PASS=<pass> -DMQTT_BROKER_URL=mqtt://your.broker
make flash && make monitor
```
