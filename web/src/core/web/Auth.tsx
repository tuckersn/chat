import { useEffect, useState } from "preact/hooks";
import { hashStr } from "../shared/utils";
import { redirectBrowser } from "./Redirects";



export interface ICookie {
    token: string;
    newAccount: boolean;
}

export function setCookie(cookie: Partial<ICookie>) {
    let originalCookie = JSON.parse
}


export interface TokenJSON {
    sub: number;
    exp: number;
    iat: number;
    iss: string;
    username: string;
    display_name: string;
    email: string;
    admin: boolean;
}

const STRING_FIELDS: (keyof TokenJSON)[] = ["iss", "username", "display_name", "email"];
const NUMERIC_FIELDS: (keyof TokenJSON)[] = ["sub", "exp", "iat"];


export function clearToken() {
    localStorage.removeItem('token');
}



export function parseToken(token: string) {
    token = localStorage.getItem('token')!;
    if (!token) {
        return null;
    }
    token = atob(token.split(".")[1]);
    
    const tokenJSON: Partial<TokenJSON> = JSON.parse(token);

    try {
        for (const field of STRING_FIELDS) {
            if (typeof tokenJSON[field] !== "string") {
                if (tokenJSON[field] === null) {
                    continue;
                }
                throw new Error(`Expected string for field ${field} in token`);
            }
        }
        for (const field of NUMERIC_FIELDS) {
            if (typeof tokenJSON[field] === "string") {
                // @ts-ignore
                tokenJSON[field] = parseInt(tokenJSON[field] as any);
            } else if (typeof tokenJSON[field] !== "number") {
                if (tokenJSON[field] === null) {
                    continue;
                }
                throw new Error(`Expected number for field ${field} in token`);
            }
        }
    } catch(e) {
        console.error("MALFORMED TOKEN!!!");
        clearToken();
        throw e;
    }

    if (tokenJSON.exp! < Date.now() / 1000) {
        clearToken();
        return null;
    }

    return tokenJSON as TokenJSON;
}

let lastToken: TokenJSON | null = null;
(() => {
    const token = localStorage.getItem('token');
    if (token) {
        lastToken = parseToken(token);
    }
})();

export function useToken() {
    const [tokenJSON, setTokenJSON] = useState<TokenJSON | null>(lastToken);
    
    let token = localStorage.getItem('token');
    if (token) {
        token = atob(token.split(".")[1]);
    }

    useEffect(() => {
        if (token) {
            const tokenJSON = parseToken(token);
            if (tokenJSON) {
                lastToken = tokenJSON;
                setTokenJSON(tokenJSON);
            }
        }
    }, [token]);

    return tokenJSON;
}


export function useTokenRequired() {
    const token = useToken();

    if(token === null || token === undefined) {
        console.error("Token required");
        redirectBrowser("/account/login");
    }

    return token!;
}