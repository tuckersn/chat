import { FC } from "preact/compat";
import { Link, Outlet } from "react-router-dom";
import styled from "styled-components";


interface NavLinkProps {
    to: boolean;
}

const NavLink = styled.div<NavLinkProps>`
    padding: 0.5rem;
    margin: 0.5rem;
    border: 1px solid black;
    border-radius: 0.5rem;
`;

export const WebContentFrame: FC = ({children}) => {
    return (
        <div style={{
            display: "flex",
            flexDirection: "row",
            height: "100%",
            width: "100%" 
        }}>
            <div style={{
                flex: 0,
            }}>
                <h2>Nav</h2>
                <NavLink>
                    <Link to="/page">Home</Link>
                </NavLink>
                <NavLink>
                    <Link to="/server">Server</Link>
                </NavLink>
                <NavLink>
                    <Link to="/chat">Chat</Link>
                </NavLink>
            </div>
            <div style={{
                flex: 1
            }}>
                <Outlet/>
            </div>
        </div>
    );
};