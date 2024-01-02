import { FC, useEffect, useState } from "preact/compat"
import { apiRequest } from "../../../core/web/API";
import { useToken, useTokenRequired } from "../../../core/web/Auth";


interface Login {
    Ip: string;
    Origin: string;
    SessionStart: string;
}

export const ViewProfilePage: FC = () => {

    const account = useTokenRequired();
    console.log("Account", account);
    
    const [logins, setLogins] = useState<Login[]>([]);

    useEffect(() => {
        apiRequest("GET", "/login/recent").then(res => {
            let logins: Login[] = JSON.parse(res);
            for(let l of logins) {
                l.SessionStart = new Date(l.SessionStart).toLocaleString();
            }
            setLogins(logins);
            console.log("Logins", logins);
        })
    }, []);

    return <div>
        <h1>Hey {account.display_name}! ({account.username})</h1>
        <h4>Other info would be here.</h4>

        <div>
            <h1>
                Recent Logins
            </h1>
            <div>
                {
                    logins.map(login => {
                        return <div style={{
                            borderBottom: "1px solid white",
                        }}>
                            <div>
                                IP: {login.Ip}
                            </div>
                            <div>
                                Origin: {login.Origin}
                            </div>
                            <div>
                                When: {login.SessionStart}
                            </div>
                        </div>;
                    })
                }
            </div>
        </div>
    </div>
}