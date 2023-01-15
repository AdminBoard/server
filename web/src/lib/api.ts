import { getOverride, HttpResponse, postOverride } from "./http"

export default {
    get,
    post,
    postData
}

function postData(url: string, body?: any): Promise<any> {
    return new Promise<any>((resolve, reject) => {
        post(url, body).then(resp => {
            if (resp.status == 0) resolve(resp.data)
            else reject(resp.message)
        }).catch(e => reject(e))
    })
}

function get(url: string): Promise<HttpResponse> {
    if (url.startsWith('/')) url = url.substring(1)
    return getOverride('/api/' + url)
}

function post(url: string, body?: any): Promise<HttpResponse> {
    if (url.startsWith('/')) url = url.substring(1)
    return postOverride('/api/' + url, body)
}