export async function postOverride(url: string, body?: any, signal?: AbortSignal): Promise<HttpResponse> {
    const req: RequestInit = { method: 'POST' }

    if (body != null) req.body = JSON.stringify(body)
    if (signal != undefined) req.signal = signal
    return new Promise(resolve => {
        fetch(url, req)
            .then(resp => {
                resp.json()
                    .then(resp => resolve(resp))
                    .catch(_ => resolve({ status: resp.status, message: resp.statusText }))
            }).catch(e => resolve({ status: 99, message: e }))
    })
}

export async function getOverride(url: string, signal?: AbortSignal): Promise<HttpResponse> {
    const req: RequestInit = { method: 'GET' }

    if (signal != undefined) req.signal = signal
    return new Promise(resolve => {
        fetch(url, req)
            .then(resp => {
                if (resp.status == 200) {
                    resp.json()
                        .then(resp => resolve(resp))
                        .catch(e => resolve({ status: 99, message: 'Invalid server response' }))
                } else {
                    const status = resp.status
                    const msg = resp.statusText + ': '
                    resp.json().then(resp => resolve({ status: status, message: msg + resp.message }))
                }
            }).catch(e => resolve({ status: 99, message: e }))
    })
}

export async function post(url: string, body?: any): Promise<HttpResponse> {
    return postOverride(url, body)
}

export async function get(url: string): Promise<HttpResponse> {
    return getOverride(url)
}

export class HttpResponse {
    status: number = 0
    message: string = ''
    data?: any = null
}

export class Error {
    message: string | null = null;
}