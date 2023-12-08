import { FC } from "preact/compat"
import { useAccountRequired } from "../../../core/web/useAccount";


export const ViewProfilePage: FC = () => {

    const account = useAccountRequired();

    return <div>
        <h1>Hey {account.display_name}!</h1>
        <h4>Other info would be here.</h4>
    </div>
}