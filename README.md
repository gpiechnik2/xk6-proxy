# xk6-proxy
[k6](https://github.com/grafana/k6) extension to define a separate (independent) proxy from the HTTP_PROXY variables for HTTP requests. Additionally, the extension allows to display and re-set a new proxy in tests. Implemented using the
[xk6](https://github.com/grafana/xk6) system.

## Build
```shell
xk6 build v0.36.0 --with github.com/gpiechnik2/xk6-proxy@latest
```

## Example
```javascript
import proxy from 'k6/x/proxy';

export default function () {
    const proxyRes = proxy.request('GET', 'https://k6.io', "http://0.0.0.0:8080", {
        headers: [
            {
                'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.9; rv:50.0) Gecko/20100101 Firefox/50.0',
                'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8'
            }
        ]
    });

    check(proxyRes, {
        'response code was 200': (proxyRes) => proxyRes.includes("200 OK") == true
    });
}
```

```javascript
import proxy from 'k6/x/proxy';

export default function () {
    let currentEnvHTTP = proxy.getCurrentEnvHTTP() // 0.0.0.0:9000
    let currentEnvHTTPS = proxy.getCurrentEnvHTTPS() // 0.0.0.0:6000
    
    proxy.setEnvHTTP('0.0.0.0:14000') // set a new HTTP_PROXY environment variable
    proxy.setEnvHTTPS('0.0.0.0:10000') // set a new HTTPS_PROXY environment variable

    currentEnvHTTP = proxy.getCurrentEnvHTTP() // 0.0.0.0:14000
    currentEnvHTTPS = proxy.getCurrentEnvHTTPS() // 0.0.0.0:10000
}
```

## Run sample script
```shell
./k6 run ./script.js
```
