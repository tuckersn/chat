

import { Button, Pressable, Text, TextInput, View } from "react-native";
import { ScreenProps } from "../Screen";
import { ChatScreenState } from "../ChatScreen";

export interface OverviewChatScreenState extends ChatScreenState{
}

export type OverviewChatScreenProps = ScreenProps<OverviewChatScreenState>;


function ChatListItem({
    chatId,
    content,
    timestamp,
    state,
    setState,
} : {
    chatId: string;
    content: string;
    timestamp: number;
} & OverviewChatScreenProps) {
    return <Pressable style={{
        backgroundColor: 'lightgrey',
        padding: 16,
        borderBottomColor: 'black',
        borderBottomWidth: 1,
    }} onPress={() => {
        setState({
            ...state,
            lastChatId: chatId,
            subPage: 'message',
        });
    }}>
        <View style={{
            flexDirection: 'row',
        }}>
            <Text style={{
                color: 'white',
                flex: 1,
            }}>
                {chatId}
            </Text>
            <Text style={{
                color: 'white',
            }}>
                {new Date(timestamp).toLocaleTimeString()}
            </Text>
        </View>
        <Text style={{
            color: 'white',
            fontStyle: 'italic',
        }}>
            {content.substring(0, content.length > 32 ? 32 : content.length) + (content.length >= 32 ? '...' : '')}
        </Text>
    </Pressable>;
}

export function OverviewChatScreen(props : OverviewChatScreenProps) {
    const {
        state,
        setState,
    } = props;

    return <View style={{
        flex: 1,
    }}>
        <View style={{
            flex: 1,
        }}>
            <ChatListItem
                chatId="U_UMEANS_USER"
                content="Hello world"
                timestamp={Date.now()}
                {...props}/>
            <ChatListItem
                chatId="B_MEANS_BOT"
                content="Hello world"
                timestamp={Date.now()}
                {...props}/>
            <ChatListItem
                chatId="G_MEANS_GROUP"
                content="this is a group chat example"
                timestamp={Date.now()}
                {...props}/>
            <ChatListItem
                chatId="B_MEANS_BOT_2"
                content="Hello world"
                timestamp={Date.now()}
                {...props}/>
            <ChatListItem
                chatId="G_MEANS_GROUP_2"
                content="this is a group chat example"
                timestamp={Date.now()}
                {...props}/>
            <ChatListItem
                chatId="B_MEANS_BOT_3"
                content="Hello world"
                timestamp={Date.now()}
                {...props}/>
            <ChatListItem
                chatId="G_MEANS_GROUP_3"
                content="this is a group chat example"
                timestamp={Date.now()}
                {...props}/>
        </View>
    </View>
}