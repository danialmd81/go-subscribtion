# `[Rule]`

Rule is a collection of rules which will apply on network traffic.

Traffic which match rules defined in [Rule] section will be redirected to specified proxy or proxy group.

Traffic will try to match rules in sequence according defined in the profile.

---

# Domain

## Sample

```ini
DOMAIN, www.apple.com, ProxyHTTP, force-remote-dns
DOMAIN-SUFFIX, apple.com, Proxy, enhanced-mode
DOMAIN-KEYWORD, google, Proxy
```

## Format

```ini
{type}, {domain rule}, {target proxy}[, force-remote-dns][, enhanced-mode]
```

## Param

| Name             | Value                                            | Mandatory | Note                                                                                                                  |
|------------------|--------------------------------------------------|-----------|-----------------------------------------------------------------------------------------------------------------------|
| type             | DOMAIN<br/>DOMAIN-SUFFIX<br/>DOMAIN-KEYWORD<br/> | true      | DOMAIN means exact matching<br/>DOMAIN-SUFFIX means suffix matching<br/>DOMAIN-KEYWORK means keyword matching         |
| domain rule      | -                                                | true      |                                                                                                                       |
| target proxy     | -                                                | true      | Specified proxy or proxy group must existed in profile                                                                |
| force-remote-dns | true<br/>false                                   | false     | Default value: false<br/>If set to true, DNS query will triggered in the remote proxy                                 |
| enhanced-mode    | true<br/>false                                   | false     | Default value: false<br/>If set to true, a fake ip will be returned in DNS query                                      |

---

# IP

## CIDR

Reference: <https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing>

### Sample

```ini
IP-CIDR,192.168.0.0/16,DIRECT
IP-CIDR6,2001:db8:abcd:8000::/50,DIRECT
```

### Format

```ini
{type}, {route}, {target proxy}
```

### Param

| Name         | Value                | Mandatory | Note                                                             |
|--------------|----------------------|-----------|------------------------------------------------------------------|
| type         | IP-CIDR<br/>IP-CIDR6 | true      | IP-CIDR works on IPv4 traffic<br/>IP-CIDR6 works on IPv6 traffic |
| route        | -                    | true      | Format: \{IP}/\{mask}, mask is in prefix format                  |
| target proxy | -                    | true      | Specified proxy or proxy group must existed in profile           |

> **caution:**
IPv6 is not supported by Surfboard currently, `IP-CIDR6` rules will be ignored.

## GEOIP

Reference:

- <https://en.wikipedia.org/wiki/Internet_geolocation>
- <https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2>

### Sample

```ini
GEOIP,US,REJECT
```

### Format

```ini
GEOIP,{country},REJECT
```

### Param

| Name         | Value | Mandatory | Note                                                             |
|--------------|-------|-----------|------------------------------------------------------------------|
| country      | -     | true      | Format: 2 Bit ISO country code                                   |
| target proxy | -     | true      | Specified proxy or proxy group must existed in profile           |

---

# Rule Set

Rule Set is a remote url configuration, whose content is a collection of Rule, but without target definition in it.

Use Rule Set can greatly simplify the content complexity of profile which has a lot of rules in it,
and also can reuse rules encapsulated by other contributors.

> **tip:**
Define a large number of rules in a rule set will significantly reduce the efficiency of rule matching.
In this scenario, we strongly recommend you switch to the [Domain Set](./domainset) standard.

## Sample

### Rule Set definition

```ini
RULE-SET,https://ruleset.com/cn,ProxyVMess
```

### Remote Rule Set content sample

```ini
DOMAIN,www.apple.com
DOMAIN,www.google.com
DOMAIN-SUFFIX,apple.com
DOMAIN-KEYWORD,google
IP-CIDR,192.168.0.0/16
GEOIP,US
```

You can see that there is no target definition in it, all matching traffic will use proxy 'ProxyVMess' as target

## Format

```ini
RULE-SET, {rule set url}, {target}
```

## Param

| Name         | Value | Mandatory | Note                                                   |
|--------------|-------|-----------|--------------------------------------------------------|
| rule set url | -     | true      |                                                        |
| target       | -     | true      | Specified proxy or proxy group must existed in profile |

---

# Final

Traffic doesn't match any other rules will match final rule if defined.

In general, a profile should only include one final rule, and place it as the last one in rule section.

## Sample

```ini
FINAL, DIRECT
```

## Format

```ini
FINAL, {target}
```

## Param

| Name         | Value | Mandatory | Note                                                   |
|--------------|-------|-----------|--------------------------------------------------------|
| target       | -     | true      | Specified proxy or proxy group must existed in profile |
