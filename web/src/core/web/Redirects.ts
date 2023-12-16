import { getAPIUrl } from "./API";



export function redirectBrowser(url: string) {
    console.log("Redirecting to: " + url);
    window.location.href = url;
}

export function redirectToAPI(apiPath: string) {
    console.log("Redirecting to: " + apiPath);
    window.location.href = getAPIUrl() + apiPath;
}