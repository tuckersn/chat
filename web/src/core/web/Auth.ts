export interface ICookie {
    token: string;
    newAccount: boolean;
}

export function setCookie(cookie: Partial<ICookie>) {
    let originalCookie = JSON.parse
}


export function clearToken() {
    localStorage.removeItem('token');
}