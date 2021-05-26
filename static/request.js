//运维端的请求
function request(params, header) {
    const service = axios.create({
        timeout: 500000
    })

    service.interceptors.request.use(
        config => {
            if (header && header.headers) {
                var xmcode = header.headers.xmcode;
                var secretkey = header.headers.secretkey;
                var token = localStorage.getItem('maque-token');
                config.headers['xmcode'] = xmcode
                config.headers['secretkey'] = secretkey
                config.headers['mq-token'] = token
            }
            return config
        },
        error => {
            console.log(error) // for debug
            return Promise.reject(error)
        }
    )
    service.interceptors.response.use(
        response => {
            const res = response.data
            return res
        },
        error => {

            layer.closeAll()
            console.log('err' + error)
            var result = $.ShowResponseErr(error)
            if (result != "") {
                layer.msg(result, {icon: 2});
            }
            return Promise.reject(error)
        }
    )
    return service(params);
}

//项目端的请求
function requestXm(params, header) {
    const service = axios.create({
        timeout: 500000
    })

    service.interceptors.request.use(
        config => {
            if (header && header.headers) {
                var xmcode = header.headers.xmcode;
                var secretkey = header.headers.secretkey;
                var token = localStorage.getItem('mq-token');
                config.headers['xmcode'] = xmcode
                config.headers['secretkey'] = secretkey
                config.headers['mq-token'] = token
            }
            return config
        },
        error => {
            console.log(error) // for debug
            return Promise.reject(error)
        }
    )
    service.interceptors.response.use(
        response => {
            const res = response.data
            return res
        },
        error => {

            layer.closeAll()
            console.log('err' + error)
            var result = $.ShowResponseErr(error)
            if (result != "") {
                layer.msg(result, {icon: 2});
            }
            return Promise.reject(error)
        }
    )
    return service(params);
}