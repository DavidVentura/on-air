```bash
cd sign
cmake . -DWIFI_SSID=<ssid> -DWIFI_PASS=<pass> -DMQTT_BROKER_URL=mqtt://your.broker
make flash && make monitor
```
