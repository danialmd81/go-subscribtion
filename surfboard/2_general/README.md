# `[General]`

## `dns-server`

Specify DNS server used by the application.

### Sample

```ini
dns-server = system, 8.8.8.8, 8.8.4.4, 9.9.9.9:9953
```

### Format

```ini
dns-server = system, {dns server}[:port], ...
```

- Port definition is optional; default is `53`.
- `system` means use system DNS.

> **tip:**

- If `dns-server` is not assigned, `system` will be used.
- Currently [DoT](https://en.wikipedia.org/wiki/DNS_over_TLS) is not supported.

---

## `doh-server`

Specify [DoH (DNS over HTTPS)](https://en.wikipedia.org/wiki/DNS_over_HTTPS) server used by the application.

### Sample

```ini
doh-server = https://9.9.9.9/dns-query
```

### Format

```ini
doh-server = {doh_server}
```

> **tip:** Currently DoH query traffic will use [DIRECT](/docs/profile-format/proxy/built-in-proxy/direct) proxy by default. Please make sure you are using an available DoH server.

---

## `skip-proxy`

Specify route rule and domain rule. Matching traffic will not be redirected or rejected.

### Sample

```ini
skip-proxy = 127.0.0.1, 192.168.0.0/16, 10.0.0.0/8, 172.16.0.0/12, 100.64.0.0/10, localhost, *.local, www.baidu.com
```

### Format

```ini
skip-proxy = {ip}, {ip/mask}, {domain}, {wildcard domain}, ...
```

> **note:**
Due to system restriction, traffic matching `skip-proxy` will still be handled by VpnService, but treated like [DIRECT](/docs/profile-format/proxy/built-in-proxy/direct) rule.

---

## `proxy-test-url`

Test URL used by [url-test](/docs/profile-format/proxygroup/auto) and manual node speed test.

Non-direct proxy will use this URL. For [direct proxy](/docs/profile-format/proxy/built-in-proxy/direct), see [`internet-test-url`](./internet_test_url).

An `HTTP HEAD` request will be sent to this URL.

An `HTTP/1.1 204 No Content` response is expected.

### Sample

```ini
proxy-test-url = http://www.gstatic.com/generate_204
```

### Format

```ini
proxy-test-url = {http_url}
```

> **tip:** `url` should start with `http://`. `https://` or other scheme URIs are not supported.
> **tip:** You can emulate proxy test by using the command below:

```shell
curl -I http://www.gstatic.com/generate_204
```

:::

---

## `internet-test-url`

Test URL used by [DIRECT](/docs/profile-format/proxy/built-in-proxy/direct) proxy.

For non-direct proxy, see [`proxy-test-url`](./proxy_test_url).

An `HTTP HEAD` request will be sent to this URL.

An `HTTP/1.1 204 No Content` response is expected.

### Sample

```ini
internet-test-url = http://www.gstatic.cn/generate_204
```

### Format

```ini
internet-test-url = {http_url}
```

> **tip:** `url` should start with `http://`. `https://` or other scheme URIs are not supported.
> **tip:** You can emulate proxy test by using the command below:

```shell
curl -I http://www.gstatic.cn/generate_204
```

---

## `always-real-ip`

In some scenarios, domain DNS query will respond with a fake IP matching `198.18.0.0/16`. Generally, this will not cause issues.

If you encounter network issues due to this feature, you can specify `always-real-ip` to bypass this hack.

### Sample

```ini
always-real-ip = *.srv.nintendo.net, *.stun.playstation.net, xbox.*.microsoft.com, *.xboxlive.com
```

### Format

```ini
always-real-ip = {domain}, {wildcard domain}, ...
```

> **tip:** Currently, the Google Voice dialing problem can be resolved by using `always-real-ip`. Please try the sample below:

```ini
always-real-ip = stun.l.google.com
```

---

## `http-listen`

Establish an HTTP proxy server on your device and provide proxy service on the specified IP.

### References

- <https://en.wikipedia.org/wiki/HTTP_tunnel>
- <https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/CONNECT>
- <https://en.wikipedia.org/wiki/Proxy_server#Web_proxy_servers>

### Sample

```ini
http-listen = 0.0.0.0:1234
```

### Format

```ini
http-listen = {ip}:{port}
```

> **tip:** If you want to provide proxy service in your LAN, use `0.0.0.0` as the `ip` parameter. Use `127.0.0.1` to provide proxy service on your device only.

---

## `socks5-listen`

Establish a SOCKS5 proxy server on your device and provide proxy service on the specified IP.

### References

- <https://datatracker.ietf.org/doc/html/rfc1928>
- <https://datatracker.ietf.org/doc/html/rfc1929>

### Sample

```ini
socks5-listen = 127.0.0.1:1235
```

### Format

```ini
socks5-listen = {ip}:{port}
```

> **tip:** If you want to provide proxy service in your LAN, use `0.0.0.0` as the `ip` parameter. Use `127.0.0.1` to provide proxy service on your device only.

---

## `udp-policy-not-supported-behaviour`

If the proxy does not support UDP relay, use [DIRECT](/docs/profile-format/proxy/built-in-proxy/direct) or [REJECT](/docs/profile-format/proxy/built-in-proxy/reject) instead. Default value is `REJECT`.

### Sample

```ini
udp-policy-not-supported-behaviour = DIRECT
```

### Format

```ini
udp-policy-not-supported-behaviour = {DIRECT|REJECT}
```

> **note:**
Only DIRECT or REJECT is supported.

---

## `test-timeout`

Timeout used for all connectivity tests.

Unit: Seconds

Default value: 5

### Sample

```ini
test-timeout = 3
```

### Format

```ini
test-timeout = {timeout in seconds}
```
