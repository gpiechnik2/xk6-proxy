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
