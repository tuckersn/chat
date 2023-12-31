import { render } from 'preact'
import './index.scss'
import { BrowserRouter, Outlet, Route, Routes } from 'react-router-dom';
import { WebContentFrame } from './routes/web/WebContentFrame.tsx';
import { ChatFrame } from './routes/web/chat/ChatFrame.tsx';
import { ChatMessagePage } from './routes/web/chat/ChatMessagePage.tsx';
import { ChatOverviewPage } from './routes/web/chat/ChatOverviewPage.tsx';
import { ServerFrame } from './routes/web/server/ServerFrame.tsx';
import { UserOverviewAdminPage } from './routes/web/server/UserOverviewAdminPage.tsx';
import { ServerOverviewAdminPage } from './routes/web/server/ServerOverviewAdminPage.tsx';
import { UserManagementPage } from './routes/web/server/UserManagementPage.tsx';
import { NotFoundPage } from './routes/web/NotFoundPage.tsx';
import { UserListAdminPage } from './routes/web/server/UserListAdminPage.tsx';
import { NotesPageFrame } from './routes/web/notes/NotesPageFrame.tsx';
import { NotesDocumentPage } from './routes/web/notes/NotesDocumentPage.tsx';
import { LoginPage } from './routes/web/account/LoginPage.tsx';
import { CreateAccountPage } from './routes/web/account/CreateAccountPage.tsx';
import { GoogleReceiveTokenPage } from './routes/web/account/GoogleReceiveTokenPage.tsx';
import { ViewProfilePage } from './routes/web/account/ViewProfilePage.tsx';

const root = document.getElementById('app')!;

render(<BrowserRouter>
    <Routes>
        <Route path="*" element={<WebContentFrame/>}>
            <Route path="app" element={<h1>Hello World</h1>}/>
            <Route path="account">
                <Route path="new" element={<CreateAccountPage/>}/>
                <Route path="login" element={<LoginPage/>}/>
                <Route path="oauth">
                    <Route path="google" element={<GoogleReceiveTokenPage/>}/>
                </Route>
                <Route path="profile" element={<ViewProfilePage/>}/>
                <Route path="*" element={<NotFoundPage/>}/>
            </Route>
            <Route path="chat" element={<ChatFrame/>}>
                <Route index element={<ChatOverviewPage/>}/>
                <Route path="id/:chatId" element={<ChatMessagePage/>}/>
                <Route path="*" element={<NotFoundPage/>}/>
            </Route>
            <Route path="notes" element={<NotesPageFrame/>}>
                <Route index element={<h1>Notes Overview</h1>}/>
                <Route path="*" element={<NotesDocumentPage/>}/>
            </Route>
            <Route path="server" element={<ServerFrame/>}>
                <Route index element={<UserOverviewAdminPage/>}/>
                <Route path="list" element={<h1>Server List</h1>}/>
                <Route path="user">
                    <Route path="overview" element={<ServerOverviewAdminPage/>}/>
                    <Route path="list" element={<UserListAdminPage/>}/>
                    <Route path="id/:userId" element={<UserManagementPage/>}/>
                    <Route path="*" element={<NotFoundPage/>}/>
                </Route>
                <Route path="*" element={<NotFoundPage/>}/>
            </Route>
            <Route path="*" element={<NotFoundPage/>}/>
        </Route>
        <Route path="*" element={<NotFoundPage/>}/>
    </Routes>
</BrowserRouter>, root);
