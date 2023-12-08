import { FC, useEffect } from "preact/compat";
import { useNavigate } from "react-router-dom";


export const GoogleReceiveTokenPage: FC = () => {

    const nav = useNavigate();

    useEffect(() => {
        console.log("GoogleReceiveTokenPage");
        const urlParams = new URLSearchParams(window.location.search);
        const token = urlParams.get("token");
        const newAccount = urlParams.get("newAccount");

        localStorage.setItem("token", token!);

        if(newAccount === "true") {
            nav("/account/new");
        } else {
            nav("/");
        }
    }, []);

    return <div>
        <h1>Google Receive Token Page</h1>
    </div>

}