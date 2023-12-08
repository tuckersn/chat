import { redirectBrowser } from "./Redirects";




export const useLoggedIn = () => {
    const token = localStorage.getItem('token');
    return !!token;
}

export function useMustBeLoggedIn() {
    const isLoggedIn = useLoggedIn();

    if (!isLoggedIn) {
        redirectBrowser("/account/login");
    }

    return isLoggedIn;
}