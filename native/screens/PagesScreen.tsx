import React from 'react';
import type {PropsWithChildren} from 'react';
import {
	ScrollView,
	StatusBar,
	StyleSheet,
	Text,
	useColorScheme,
	View,
} from 'react-native';
import { ScreenProps } from './Screen';
import { MarkdownRenderer } from '../shared/components/MarkdownRenderer';

export interface MarkdownScreenState {
	path: string;
	scroll: number;
	history: string[];
}

export type MarkdownScreenProps = ScreenProps<MarkdownScreenState>;

export function PagesScreen(): JSX.Element {

	return (
		<View style={{
			flex: 1
		}}>
			<View style={{
				borderBottomColor: 'black',
				borderBottomWidth: 1,
			}}>
				<Text style={{
					fontSize: 32,
					fontWeight: 'bold',
					textAlign: 'center',
					marginTop: 16,
				}}>Search</Text>
			</View>
			<MarkdownRenderer style={{flex: 1}} markdown="# Hello World\n- test 1\n-test 2"/>
		</View>
	);
}
