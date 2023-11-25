import React, { FC, HTMLProps, JSX, useEffect, useMemo, useState } from "preact/compat";

import { Link, Outlet, useNavigate } from "react-router-dom";

import styled from "styled-components";

import { GiChatBubble } from "react-icons/gi";

import "../../style-web.scss";



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
            backgroundColor: active ? "red" : "blue"
        }}>
            {icon}
        </div>
        <div class="flex-fill" style={{
            height: "100%",
            width: "100%",
            flex: 1,
            display: active ? "block" : "none",
            backgroundColor: "rgba(0, 0, 0, 0.5)"
        }}>
            {children}
        </div>
    </>;
}

export const WebContentFrame: FC = ({children}) => {

    const category = useState("/chat");

    return (
        <div class="flex-container">
            <div class="flex-min" style={{
                display: "flex",
                flexDirection: "column",
                height: "100%"
            }}>
                <NavCategory icon={<GiChatBubble/>} to="/chat" category={category}>
                    <h1>links to chat stuff here</h1>
                </NavCategory>
                <NavCategory icon={<GiChatBubble/>} to="/page" category={category}>
                    <h1>links to markdown documents in the pages section here</h1>
                </NavCategory>
                <NavCategory icon={<GiChatBubble/>} to="/server" category={category}>                
                    <GiChatBubble/>
                </NavCategory>
            </div>
            <div class="flex-fill">
                <Outlet/>
            </div>
        </div>
    );
};