import { getAPIUrl } from "./API";



export function redirectBrowser(url: string) {
    window.location.href = url;
}

export function redirectToAPI(apiPath: string) {
    window.location.href = getAPIUrl() + apiPath;
}