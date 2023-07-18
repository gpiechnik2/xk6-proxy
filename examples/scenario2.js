import proxy from 'k6/x/proxy';


export default function () {
    let currentEnvHTTP = proxy.getCurrentEnvHTTP() // 0.0.0.0:9000
    let currentEnvHTTPS = proxy.getCurrentEnvHTTPS() // 0.0.0.0:6000
    
    proxy.setEnvHTTP('0.0.0.0:14000') // set a new HTTP_PROXY environment variable
    proxy.setEnvHTTPS('0.0.0.0:10000') // set a new HTTPS_PROXY environment variable

    currentEnvHTTP = proxy.getCurrentEnvHTTP() // 0.0.0.0:14000
    currentEnvHTTPS = proxy.getCurrentEnvHTTPS() // 0.0.0.0:10000
}
