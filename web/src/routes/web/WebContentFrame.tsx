import React, { FC, HTMLProps, JSX, useEffect, useMemo, useState } from "preact/compat";

import { Link, LinkProps, Outlet, useNavigate } from "react-router-dom";

import styled from "styled-components";

import { GiAbstract008, GiAbstract076, GiAccordion, GiBookCover, GiBookshelf, GiChatBubble, GiOverdrive, GiPerson, GiServerRack, GiSettingsKnobs } from "react-icons/gi";
import { FaAccessibleIcon, FaBox, FaPersonBooth, FaUser } from "react-icons/fa"

import "../../style-web.scss";
import { SearchBar } from "../../components/SearchBar";
import { SiEngadget } from "react-icons/si";
import { useToken } from "../../core/web/Auth";

export const NavButtonHeight = 64;
export const NavButtonWidth = 128;
export const NavButtonIconSize = 38;
export const NavBackgroundColor = "rgba(0, 0, 0, 0.5)";

interface NavCategoryProps {
    to: string;
    children: React.ReactNode;
    category: [string, (to: string) => void];
    icon: JSX.Element;
}

const NavCategory: FC<NavCategoryProps> = ({
    to,
    children,
    category,
    icon
} : NavCategoryProps) => {

    const [active, setActive] = useState(false);
    const navigate = useNavigate();
    
    useEffect(() => {
        setActive(category[0].startsWith(to));
    }, [category[0], to]);

    return <>
        <div onClick={() => {
            category[1](to);
            navigate(to);
        }} style={{
            width: NavButtonWidth,
            height: NavButtonHeight,
            backgroundColor: active ? "red" : "blue",
            display: "flex",
            placeItems: "center",
            placeContent: "center",
            border: "3px solid black",
            boxSizing: "border-box"
        }}>
            {icon}
        </div>
        <div class="flex-fill" style={{
            height: "100%",
            width: "100%",
            flex: 1,
            display: active ? "block" : "none",
            backgroundColor: NavBackgroundColor
        }}>
            {children}
        </div>
    </>;
}

const NavCategoryLink: FC<LinkProps> = ({
    children,
    ...props
}) => {
    return <div>
        <Link {...props}>
            {children}
        </Link>
    </div>;
}



export const WebContentFrame: FC = ({children}) => {

    const category = useState("/chat");
    const account = useToken();


    useEffect(() => {
        category[1](window.location.pathname);
        if (window.location.pathname === "/") {
            window.location.pathname = "/account/profile";
        }
    }, [window.location.pathname]);

    return (
        <div class="flex-container flex-min h-fill" style={{
            height: "100%"
        }}>
            <div class="flex-min" style={{
                display: "flex",
                flexDirection: "column",
                height: "100%",
                flex: `0 0 ${NavButtonWidth}px`,
            }}>
                <NavCategory icon={<FaUser size={NavButtonIconSize}/>} to="/account" category={category}>
                {
                    account === null ? <>
                        <NavCategoryLink to="/account/profile">
                            Profile
                        </NavCategoryLink>
                        <button onClick={() => {
                            localStorage.removeItem("token");
                            window.location.reload();
                        }}>
                            Logout
                        </button>
                    </> : <>
                        <NavCategoryLink to="/account/login">
                            Login
                        </NavCategoryLink>
                    </>
                }
                </NavCategory>
                <NavCategory icon={<GiBookCover size={NavButtonIconSize}/>} to="/notes" category={category}>
                    <NavCategoryLink to="/notes/test">
                        Test
                    </NavCategoryLink>
                </NavCategory>
                <NavCategory icon={<GiChatBubble size={NavButtonIconSize}/>} to="/chat" category={category}>
                    <button>+</button>
                    <div>chat1</div>
                    <div>chat2</div>
                    <div>chat3</div>
                </NavCategory>
                <NavCategory icon={<GiServerRack size={NavButtonIconSize}/>} to="/server" category={category}>                                
                    <div>Overview</div>
                    <div>Users</div>
                    <div>Settings</div>
                    <div>Webhooks</div>
                    <div>Admin</div>
                </NavCategory>
            </div>
            <div class="flex-fill">
                <div style={{
                    height: NavButtonHeight,
                    width: "100%",
                    backgroundColor: NavBackgroundColor,
                    display: "flex",
                    flexDirection: "row",
                }}>
                    <div style={{
                        flex: 1,
                        height: "100%",
                        display: "flex",
                        placeItems: "center",
                        placeContent: "center"
                    }}>
                        toolbar section
                    </div>
                    <div style={{
                        flex: 1,
                        height: "100%",
                        display: "flex",
                        placeItems: "center",
                        placeContent: "center"
                    }}>
                        <SearchBar onSearch={(s) => {
                            console.log("search: ", s);
                        }}/>
                    </div>
                </div>
                <Outlet/>
            </div>
        </div>
    );
};