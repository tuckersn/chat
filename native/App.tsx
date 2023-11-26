/**
 * Sample React Native App
 * https://github.com/facebook/react-native
 *
 * @format
 */

import React, { useEffect } from 'react';
import type {PropsWithChildren} from 'react';

import { NotesScreen } from './screens/NotesScreen';
import { ServerScreen } from './screens/ServerScreen';
import { Button, Keyboard, Pressable, Text, View } from 'react-native';
import { Styles } from './Style';
import { ChatScreen, ChatScreenState } from './screens/ChatScreen';


interface NavButtonProps {
	screenId: string;
	active: string;
	setActiveScreenId: (screenId: string) => void;
	children?: React.ReactNode;
}

const NavButton: React.FC<NavButtonProps> = ({
	screenId,
	active,
	setActiveScreenId,
	children
}: NavButtonProps) => {
	return <Pressable style={{
		...Styles.button,
		backgroundColor: screenId === active ? 'darkgray' : 'black',
	}} onTouchStart={() => {
		setActiveScreenId(screenId);
	}}>
		{children}
	</Pressable>;
}

export default function App() {

	const [activeScreen, setActiveScreen] = React.useState('chat');
	const [markdownState, setMarkdownState] = React.useState<any>({});
	const [chatState, setChatState] = React.useState<ChatScreenState>({
		lastChatId: '',
		subPage: 'overview',
	});
	const [settingsState, setSettingsState] = React.useState<any>({});

	const [isKeyboardVisible, setKeyboardVisible] = React.useState(false);

	useEffect(() => {
		const keyboardDidShowListener = Keyboard.addListener(
			'keyboardDidShow',
			() => {
				setKeyboardVisible(true); // or some other action
			}
		);
		const keyboardDidHideListener = Keyboard.addListener(
			'keyboardDidHide',
			() => {
				setKeyboardVisible(false); // or some other action
			}
		);

		return () => {
			keyboardDidHideListener.remove();
			keyboardDidShowListener.remove();
		};
	}, []);
	
	return <View style={{
		height: '100%',
	}}>
		<View style={{
			overflow: 'visible',
			flex: 1,
		}}>
			{
				activeScreen === 'pages' ? <NotesScreen/> : null
			}
			{
				activeScreen === 'server'	? <ServerScreen/> : null
			}
			{
				activeScreen === 'chat' ? <ChatScreen
					isKeyboardVisible={isKeyboardVisible}
					state={chatState}
					setState={setChatState}/> : null
			}
		</View>
		<View style={{
			flex: 0,
			flexDirection: 'row',
			display: isKeyboardVisible ? 'none' : 'flex',
		}}>
			<NavButton screenId="pages" active={activeScreen} setActiveScreenId={setActiveScreen}>
				<Text style={Styles.text}>Pages</Text>
			</NavButton>
			<NavButton screenId="chat" active={activeScreen} setActiveScreenId={setActiveScreen}>
				<Text style={Styles.text}>Chat</Text>
			</NavButton>
			<NavButton screenId="server" active={activeScreen} setActiveScreenId={setActiveScreen}>
				<Text style={Styles.text}>Server</Text>
			</NavButton>
		</View>
	</View>
}