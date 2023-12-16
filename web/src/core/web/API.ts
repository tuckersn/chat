import config from "../../../config.json";


export function getAPIUrl() {
    return config.api_url;
}

export function apiRequest(method: "GET" | "POST", url: string, body?: any) {
    const req = new XMLHttpRequest();
    req.open(method, getAPIUrl() + url);
    req.setRequestHeader("Content-Type", "application/json");
    req.setRequestHeader("Accept", "application/json");
    req.setRequestHeader("Authorization", "Bearer " + localStorage.getItem("token"));
    req.send(JSON.stringify(body));
    return new Promise<XMLHttpRequest["response"]>((resolve, reject) => {
        req.onload = () => {
            if (req.status >= 200 && req.status < 300) {
                resolve(req.response);
            } else {
                reject(req.response);
            }
        }
    });
}