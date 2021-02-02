import './Account.scss'
import {useContext} from "react";
import {userSession} from "../../UserSession";

function Account() {
    const [user] = useContext(userSession)

    return (
        <div className="account">
            <img src={user.picture} alt={"profile"}/>

            <div>{user.name}</div>
            <div>{user.email}</div>
        </div>
    );

}

export default Account;
