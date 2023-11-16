import { Text, View } from "react-native"
import { MarkdownComponentProps } from "../MarkdownRenderer";

export const h1 = (props: MarkdownComponentProps<'h1'>) => {
    console.log(props.node);
    return <View>
        {
            props.node.children.map((child, index) => {
                return <Text key={index}>
                    {child.value}
                </Text>
            })
        }
    </View>
};

export const h2 = (props: MarkdownComponentProps<'h2'>) => {
    return <View>
        <Text>
            H2: {props.children}
        </Text>
    </View>;
}

export const h3 = (props: MarkdownComponentProps<'h3'>) => {
    return <Text>
        {props.children}
    </Text>
}

export const h4 = (props: MarkdownComponentProps<'h4'>) => {
    return <Text>
        {props.children}
    </Text>
}

export const h5 = (props: MarkdownComponentProps<'h5'>) => {
    return <Text>
        {props.children}
    </Text>
}