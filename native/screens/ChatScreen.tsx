import { Text, TextInput, View } from "react-native";
import { ScreenProps } from "./Screen";
import { useEffect } from "react";
import { OverviewChatScreen } from "./chat/OverviewChatScreen";
import { MessageChatScreen } from "./chat/MessagesChatScreen";

export interface ChatScreenState {
    subPage: 'overview' | 'message' | 'search' | 'settings';
    lastChatId: string;
    messages?: {
        authorId: string;
        content: string;
        timestamp: number;
    }[]
}

export type ChatScreenProps = ScreenProps<ChatScreenState>;

export function ChatScreen(props : ChatScreenProps) {
    
    const {
        state,
        setState,
    } = props;
    const { subPage } = state;
    
    return <View style={{
        flex: 1,
    }}>
        {
            subPage === 'overview' ? <OverviewChatScreen {...props}/> : null
        }
        {
            subPage === 'message' ? <MessageChatScreen {...props}/> : null
        }
        {
            subPage === 'search' ? <Text>Search</Text> : null
        }
        {
            subPage === 'settings' ? <Text>Settings</Text> : null
        }
    </View>
}