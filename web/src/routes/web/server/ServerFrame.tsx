import { FC } from "preact/compat";
import { Outlet } from "react-router-dom";



export const ServerFrame: FC<{}> = () => {

    return <div style={{
        display: "flex",
        flexDirection: "row",
        height: "100%",
        width: "100%" 
    }}>
        <div style={{
            flex: 1
        }}>
            <h2>Server</h2>
        </div>
        <Outlet/>
    </div>;

}