import React, { FC, HTMLProps, JSX, useEffect, useMemo, useState } from "preact/compat";

import { Link, Outlet, useNavigate } from "react-router-dom";

import styled from "styled-components";

import { GiBookCover, GiChatBubble, GiServerRack } from "react-icons/gi";

import "../../style-web.scss";


export const NavButtonSize = 64;
export const NavButtonIconSize = 38;
export const NavBackgroundColor = "rgba(0, 0, 0, 0.5)";

interface NavCategoryProps {
    to: string;
    children: JSX.Element;
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
        setActive(category[0] === to);
    }, [category[0], to]);

    return <>
        <div onClick={() => {
            category[1](to);
            navigate(to);
        }} style={{
            width: NavButtonSize,
            height: NavButtonSize,
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

const NavCategoryLink: FC<HTMLProps<HTMLAnchorElement>> = ({
    children,
    ...props
}) => {
    return <Link {...props}>
        {children}
    </Link>
}



export const WebContentFrame: FC = ({children}) => {

    const category = useState("/chat");

    return (
        <div class="flex-container flex-min h-fill" style={{
            height: "100%"
        }}>
            <div class="flex-min" style={{
                display: "flex",
                flexDirection: "column",
                height: "100%",
                flex: `0 0 ${NavButtonSize}px`,
            }}>
                <NavCategory icon={<GiBookCover size={NavButtonIconSize}/>} to="/notes" category={category}>
                    <NavCategoryLink>
                        Test
                    </NavCategoryLink>
                </NavCategory>
                <NavCategory icon={<GiChatBubble size={NavButtonIconSize}/>} to="/chat" category={category}>
                    <h1>links to markdown documents in the pages section here</h1>
                </NavCategory>
                <NavCategory icon={<GiServerRack size={NavButtonIconSize}/>} to="/server" category={category}>                
                    <GiChatBubble/>
                </NavCategory>
            </div>
            <div class="flex-fill">
                <div style={{
                    height: NavButtonSize,
                    backgroundColor: NavBackgroundColor
                }}>
                    test
                </div>
                <Outlet/>
            </div>
        </div>
    );
};