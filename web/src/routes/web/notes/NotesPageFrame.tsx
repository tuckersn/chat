import { FC } from "preact/compat";
import { Outlet } from "react-router-dom";


export const NotesPageFrame: FC<{}> = () => {
    return <div>
        <Outlet/>
    </div>;
}