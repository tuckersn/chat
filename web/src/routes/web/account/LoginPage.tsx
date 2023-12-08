import { FC } from "preact/compat";
import { redirectBrowser, redirectToAPI } from "../../../core/web/Redirects";





export const LoginPage: FC = () => {
    return <div>
        <h1>Login / Sign Up</h1>
        <div>
            <div style={{
                border: "1px solid black",
                padding: "4px",
            }} onClick={() => {
                redirectToAPI("/login/google")
            }}>
                Google
            </div>
            <div style={{
                border: "1px solid black",
                padding: "4px",
            }}>
                GitLab (WIP)
            </div>
            <div style={{
                border: "1px solid black",
                padding: "4px",
            }}>
                GitHub (TODO)
            </div>
        </div>
    </div>
}