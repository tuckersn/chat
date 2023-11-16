
import { ScrollView, Text, View, ViewStyle } from "react-native";
import remarkGfm from "remark-gfm";
import { h1, h2, h3, h4, h5 } from "./markdown/MdHeadings";
import ReactMarkdown from "react-markdown";

export type MarkdownComponentProps<REPLACEE extends keyof JSX.IntrinsicElements> = JSX.IntrinsicElements[REPLACEE] & {
    children?: React.ReactNode;
    node: {
        children: [{
            "type": "text",
            "value": string,
        }]
    }
};

export const MARKDOWN_COMPONENTS: {
    [TagName in keyof Partial<JSX.IntrinsicElements>]: React.FC<MarkdownComponentProps<TagName>>
} = {
    h1,
    h2,
    h3,
    h4,
    h5
}


export interface MarkdownProps {
    markdown: string;
    // type of react native style
    style?: Partial<ViewStyle>;
}

export const MarkdownRenderer: React.FC<MarkdownProps> = ({
    markdown,
    style
}: MarkdownProps) => {
    return <View style={{
        flex: 1,
        ...style
    }}>
        <ReactMarkdown
            remarkPlugins={[remarkGfm]}
            components={MARKDOWN_COMPONENTS as any}>
            {markdown.replace("\\n", "\n")}
        </ReactMarkdown>
    </View>;
}