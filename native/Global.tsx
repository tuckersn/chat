
export interface IGlobalState {
    loggedIn: boolean;
    userId: string;
}

export const GlobalState: IGlobalState = {
    loggedIn: false,
    userId: '',
}