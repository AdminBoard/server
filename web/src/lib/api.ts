import { getOverride, HttpResponse, postOverride } from "./http"

export default {
    get,
    getData,
    post,
    postData,
}

function get(url: string): Promise<HttpResponse> {
    if (url.startsWith('/')) url = url.substring(1)
    if (url.startsWith('api/')) url = url.substring(4)
    return getOverride('/api/' + url)
}

function getData(url: string): Promise<any> {
    return new Promise<any | string>((resolve, reject) => {
        get(url).then(resp => {
            if (resp.status == 0) resolve(resp.data)
            else reject(resp.message)
        }).catch(e => reject(e))
    })
}

function post(url: string, body?: any): Promise<HttpResponse> {
    if (url.startsWith('/')) url = url.substring(1)
    if (url.startsWith('api/')) url = url.substring(4)
    return postOverride('/api/' + url, body)
}

function postData(url: string, body?: any): Promise<any> {
    return new Promise<any | string>((resolve, reject) => {
        post(url, body).then(resp => {
            if (resp.status == 0) resolve(resp.data)
            else reject(resp.message)
        }).catch(e => reject(e))
    })
}

