import './Account.scss'
import {useContext} from "react";
import {userSession} from "../../UserSession";

function Account() {
    const [user] = useContext(userSession)

    return (
        <div>
            <div>{user.name}</div>
            <div>{user.email}</div>
            <div>{user.picture}</div>
        </div>
    );

}

export default Account;
