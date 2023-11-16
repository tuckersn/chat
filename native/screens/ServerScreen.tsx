import { Button, Text, View } from "react-native";


export function ServerScreen() {
    return <View>
        <View style={{
            flexDirection: 'row',
        }}>
            <Button title="Ping" onPress={() => {
                //alert('Pong');
            }}/>
            <Button title="Reconnect to Live" onPress={() => {
                //alert('Back');
            }}/>
            <Button title="Clear Cache" onPress={() => {
                //alert('Back');
            }}/>
        </View>
        <Text style={{
            fontSize: 32,
            fontWeight: 'bold',
            textAlign: 'center',
        }}>
            Server Page
        </Text>
    </View>
}