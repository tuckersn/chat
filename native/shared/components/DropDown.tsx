import { Pressable, ScrollView, Text, View } from "react-native";

export interface DropDownProps {
    children?: React.ReactNode;
}

export const DropDown: React.FC<DropDownProps> = ({
    children: chidren,
}: DropDownProps) => {
    return <View style={{
        flex: 1,
    }}>
        <Pressable>
            <Text>
                Dropdown here
            </Text>
        </Pressable>
        <View>
            <Text>
                {chidren}
            </Text>        
        </View>
    </View>;
}