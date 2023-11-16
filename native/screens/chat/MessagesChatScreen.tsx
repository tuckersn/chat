import { Button, Pressable, ScrollView, Text, TextInput, View } from "react-native";
import { ScreenProps } from "../Screen";
import { GlobalState } from "../../Global";
import { ChatScreenState } from "../ChatScreen";

export interface MessageChatScreenState extends ChatScreenState{
}

export interface MessageChatScreenProps extends ScreenProps<MessageChatScreenState> {}

const PADDING = 12;
const MARGIN = 12;

function MessageComponent({
    content,
    authorId,
}: {
    content: string;
    authorId: string;
} & MessageChatScreenProps) {

    const selfAuthored = authorId === GlobalState.userId;

    return <View style={{
        maxWidth: '75%',
        alignSelf: selfAuthored ? 'flex-end' : 'flex-start',
    }}>
        <View style={{
            padding: PADDING,
            marginLeft: MARGIN,
            marginRight: MARGIN,
            backgroundColor: selfAuthored ? 'darkgrey' : '#444',
        }}>
            <Text style={{
                marginRight: MARGIN,
                color: selfAuthored ? 'black' : 'white',
            }}>
                {content}
            </Text>
        </View>
        <Text style={{
            marginRight: MARGIN,
            marginLeft: MARGIN,
            alignSelf: selfAuthored ? 'flex-end' : 'flex-start',
        }}>
            11/12/23 12:34
        </Text>
    </View>
}


export function MessageChatScreen(props : MessageChatScreenProps) {

    const {
        state,
        setState,
    } = props;

    const {
        lastChatId
    } = state;

    return <View style={{
        flex: 1,
    }}>
        <View style={{
            borderBottomColor: 'black',
            borderBottomWidth: 1,
            flexDirection: 'row',
        }}>
            <Button title="Back" onPress={() => {
                setState({
                    ...state,
                    subPage: 'overview',
                });
            }}/>
            <Text style={{
                fontSize: 28,
                fontWeight: 'bold',
                textAlign: 'center',
                flex: 1,
            }}>
                {lastChatId}
            </Text>
        </View>
        <ScrollView style={{
            flex: 1,
            paddingTop: PADDING,
        }}>
            <MessageComponent content="Example message content from a different user." authorId="NOT_YOU" {...props}/>
            <MessageComponent content="Example message content from a different user." authorId="NOT_YOU" {...props}/>
            <MessageComponent content="Example reply." authorId="" {...props}/>
        </ScrollView>
        <View style={{
            height: 64,
            borderTopColor: 'black',
            borderTopWidth: 1,
            flexDirection: 'row',
            padding: 0,
            margin: 0,
        }}>
            <TextInput style={{
                flex: 1,
                height: "100%",
                flexWrap: 'wrap',
                textAlign: 'justify',
                padding: 0,
            }} multiline={true}>
                <Text style={{
                    flexWrap: 'wrap',
                    textAlign: 'justify'
                }}>Chat</Text>
            </TextInput>
            <Pressable>
                <Text>Send</Text>
            </Pressable>
        </View>
    </View>
}