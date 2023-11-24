import { FC } from "preact/compat";
import { Outlet } from "react-router-dom";



export const ChatFrame: FC<{}> = () => {

    return <div>
        chat frame
        <Outlet/>
    </div>;

}