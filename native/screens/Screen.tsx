export interface ScreenProps<T = any> {
    state: T;
    setState: (state: T) => void;
    isKeyboardVisible: boolean;
}