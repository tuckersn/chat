import { useEffect, useState } from "preact/hooks";
import { redirectBrowser } from "./Redirects";


let globalAccount: AccountJWT | null = getAccount();

export interface AccountJWT {
    user_id: string;
    email?: string;
    username: string;
    display_name: string;
    exp: number;
}

function getAccount(): AccountJWT | null {
    const token = localStorage.getItem('token');
    if (!token) {
        return null;
    }
    try {
        let jwt = localStorage.getItem('token');
        if (!jwt) {
            throw new Error("No token");
        }
        jwt = jwt.split('.')[1];
        jwt = atob(jwt);
        let token = JSON.parse(jwt);
        if (token) {
            if (token.exp < Date.now() / 1000) {
                throw new Error("Expired");
            }
        }
        return token;
    } catch (e) {
        localStorage.removeItem('token');
        return null;
    }
    return null;
}




export const useAccount = () => {
    const [account, setAccount] = useState<AccountJWT | null>(globalAccount);

    useEffect(() => {
        setAccount(globalAccount);
    }, [globalAccount]);

    console.log("Account", account);

    return account;
}

export const useAccountRequired = () => {
    const account = useAccount();

    if (account === null || account === undefined) {
        redirectBrowser("/account/login");
    }

    return account!;
}