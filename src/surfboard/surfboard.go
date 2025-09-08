package surfboard

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// GenerateSSConfig creates a Surfboard config file from an ss.txt file.
func GenerateSSConfig(ssFile, outConf string) error {
	type Proxy struct {
		Name     string
		Host     string
		Port     string
		Method   string
		Password string
	}

	var proxies []Proxy
	file, err := os.Open(ssFile)
	if err != nil {
		return fmt.Errorf("failed to open ss.txt: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ssRE := regexp.MustCompile(`^ss://([A-Za-z0-9\-_]+={0,2})@([^\s:]+):(\d+)(?:#(.+))?$`)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || !strings.HasPrefix(line, "ss://") {
			continue
		}
		m := ssRE.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		userinfo, err := base64.RawURLEncoding.DecodeString(m[1])
		if err != nil {
			userinfo, err = base64.StdEncoding.DecodeString(m[1])
			if err != nil {
				continue
			}
		}
		userinfoParts := strings.SplitN(string(userinfo), ":", 2)
		if len(userinfoParts) != 2 {
			continue
		}
		method := userinfoParts[0]
		password := userinfoParts[1]
		name := m[4]
		if name == "" {
			name = fmt.Sprintf("%s-%s", m[2], m[3])
		}
		cleanName := regexp.MustCompile(`[^\w\-\.@]+`).ReplaceAllString(name, "")
		proxies = append(proxies, Proxy{
			Name:     cleanName,
			Host:     m[2],
			Port:     m[3],
			Method:   method,
			Password: password,
		})
	}

	out, err := os.Create(outConf)
	if err != nil {
		return fmt.Errorf("failed to create config: %w", err)
	}
	defer out.Close()

	out.WriteString(`[General]
dns-server = system, 8.8.8.8, 8.8.4.4, 1.1.1.1, 1.0.0.1
doh-server = https://doh.pub/dns-query
skip-proxy = 127.0.0.1, 192.168.0.0/16, 10.0.0.0/8, 172.16.0.0/12, 100.64.0.0/10, localhost, *.local
proxy-test-url = https://www.youtube.com/favicon.ico
internet-test-url = http://www.gstatic.cn/generate_204
always-real-ip = stun.l.google.com
udp-policy-not-supported-behaviour = DIRECT
test-timeout = 1

[Proxy]
`)
	var proxyNames []string
	for _, p := range proxies {
		out.WriteString(fmt.Sprintf("%s = ss, %s, %s, encrypt-method=%s, password=%s, udp-relay=false\n",
			p.Name, p.Host, p.Port, p.Method, p.Password))
		proxyNames = append(proxyNames, p.Name)
	}

	out.WriteString("\n[Proxy Group]\n")
	out.WriteString("myGroup = select")
	for _, name := range proxyNames {
		out.WriteString(", " + name)
	}
	out.WriteString("\n\n[Rule]\n")
	out.WriteString(`DOMAIN-SUFFIX,ir,DIRECT
GEOIP,IR,DIRECT
DOMAIN-SET,https://github.com/bootmortis/iran-hosted-domains/releases/latest/download/surge_domainset_ads.txt,REJECT,update-interval=432000
DOMAIN-SET,https://github.com/bootmortis/iran-hosted-domains/releases/latest/download/surge_domainset_other.txt,DIRECT,update-interval=432000
FINAL,myGroup
`)
	return nil
}

// GenerateVMessConfig creates a Surfboard config file from a vmess.txt file.
func GenerateVMessConfig(vmessFile, outConf string) error {
	type Proxy struct {
		Name   string
		Host   string
		Port   string
		UUID   string
		WSPath string
		TLS    string
	}

	var proxies []Proxy
	file, err := os.Open(vmessFile)
	if err != nil {
		return fmt.Errorf("failed to open vmess.txt: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	vmessRE := regexp.MustCompile(`^vmess://([A-Za-z0-9\-_]+={0,2})(?:#(.+))?$`)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || !strings.HasPrefix(line, "vmess://") {
			continue
		}
		m := vmessRE.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		raw, err := base64.RawStdEncoding.DecodeString(m[1])
		if err != nil {
			raw, err = base64.StdEncoding.DecodeString(m[1])
			if err != nil {
				continue
			}
		}
		vm := make(map[string]interface{})
		_ = jsonUnmarshal(raw, &vm)
		host := fmt.Sprintf("%v", vm["add"])
		port := fmt.Sprintf("%v", vm["port"])
		uuid := fmt.Sprintf("%v", vm["id"])
		wsPath := fmt.Sprintf("%v", vm["path"])
		tls := fmt.Sprintf("%v", vm["tls"])
		name := m[2]
		if name == "" {
			name = fmt.Sprintf("%s-%s", host, port)
		}
		cleanName := regexp.MustCompile(`[^\w\-\.@]+`).ReplaceAllString(name, "")
		proxies = append(proxies, Proxy{
			Name:   cleanName,
			Host:   host,
			Port:   port,
			UUID:   uuid,
			WSPath: wsPath,
			TLS:    tls,
		})
	}

	out, err := os.Create(outConf)
	if err != nil {
		return fmt.Errorf("failed to create config: %w", err)
	}
	defer out.Close()

	out.WriteString(`[General]
dns-server = system, 8.8.8.8, 8.8.4.4, 1.1.1.1, 1.0.0.1
doh-server = https://doh.pub/dns-query
skip-proxy = 127.0.0.1, 192.168.0.0/16, 10.0.0.0/8, 172.16.0.0/12, 100.64.0.0/10, localhost, *.local
proxy-test-url = https://www.youtube.com/favicon.ico
internet-test-url = http://www.gstatic.cn/generate_204
always-real-ip = stun.l.google.com
udp-policy-not-supported-behaviour = DIRECT
test-timeout = 1

[Proxy]
`)
	var proxyNames []string
	for _, p := range proxies {
		out.WriteString(fmt.Sprintf("%s = vmess, %s, %s, username=%s, ws-path=%s, tls=%s, udp-relay=false\n",
			p.Name, p.Host, p.Port, p.UUID, p.WSPath, p.TLS))
		proxyNames = append(proxyNames, p.Name)
	}

	out.WriteString("\n[Proxy Group]\n")
	out.WriteString("myGroup = select")
	for _, name := range proxyNames {
		out.WriteString(", " + name)
	}
	out.WriteString("\n\n[Rule]\n")
	out.WriteString(`DOMAIN-SUFFIX,ir,DIRECT
GEOIP,IR,DIRECT
DOMAIN-SET,https://github.com/bootmortis/iran-hosted-domains/releases/latest/download/surge_domainset_ads.txt,REJECT,update-interval=432000
DOMAIN-SET,https://github.com/bootmortis/iran-hosted-domains/releases/latest/download/surge_domainset_other.txt,DIRECT,update-interval=432000
FINAL,myGroup
`)
	return nil
}

// GenerateHysteriaConfig creates a Surfboard config file from a hysteria.txt file.
func GenerateHysteriaConfig(hysteriaFile, outConf string) error {
	type Proxy struct {
		Name string
		Host string
	}

	var proxies []Proxy
	file, err := os.Open(hysteriaFile)
	if err != nil {
		return fmt.Errorf("failed to open hysteria.txt: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	hysteriaRE := regexp.MustCompile(`^hysteria://([^\s#]+)(?:#(.+))?$`)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || !strings.HasPrefix(line, "hysteria://") {
			continue
		}
		m := hysteriaRE.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		name := m[2]
		if name == "" {
			name = m[1]
		}
		cleanName := regexp.MustCompile(`[^\w\-\.@]+`).ReplaceAllString(name, "")
		proxies = append(proxies, Proxy{
			Name: cleanName,
			Host: m[1],
		})
	}

	out, err := os.Create(outConf)
	if err != nil {
		return fmt.Errorf("failed to create config: %w", err)
	}
	defer out.Close()

	out.WriteString(`[General]
dns-server = system, 8.8.8.8, 8.8.4.4, 1.1.1.1, 1.0.0.1
doh-server = https://doh.pub/dns-query
skip-proxy = 127.0.0.1, 192.168.0.0/16, 10.0.0.0/8, 172.16.0.0/12, 100.64.0.0/10, localhost, *.local
proxy-test-url = https://www.youtube.com/favicon.ico
internet-test-url = http://www.gstatic.cn/generate_204
always-real-ip = stun.l.google.com
udp-policy-not-supported-behaviour = DIRECT
test-timeout = 1

[Proxy]
`)
	var proxyNames []string
	for _, p := range proxies {
		out.WriteString(fmt.Sprintf("%s = hysteria, %s, udp-relay=false\n",
			p.Name, p.Host))
		proxyNames = append(proxyNames, p.Name)
	}

	out.WriteString("\n[Proxy Group]\n")
	out.WriteString("myGroup = select")
	for _, name := range proxyNames {
		out.WriteString(", " + name)
	}
	out.WriteString("\n\n[Rule]\n")
	out.WriteString(`DOMAIN-SUFFIX,ir,DIRECT
GEOIP,IR,DIRECT
DOMAIN-SET,https://github.com/bootmortis/iran-hosted-domains/releases/latest/download/surge_domainset_ads.txt,REJECT,update-interval=432000
DOMAIN-SET,https://github.com/bootmortis/iran-hosted-domains/releases/latest/download/surge_domainset_other.txt,DIRECT,update-interval=432000
FINAL,myGroup
`)
	return nil
}

// Helper for VMess JSON parsing (no external deps)
func jsonUnmarshal(data []byte, v *map[string]interface{}) error {
	dec := json.NewDecoder(strings.NewReader(string(data)))
	dec.UseNumber()
	return dec.Decode(v)
}
