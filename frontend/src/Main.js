import React, { useContext } from "react";
import {userSession} from './UserSession';

function Main() {
    const [user] = useContext(userSession)
    return (
        <div>{user.name}</div>
    );

}

export default Main;
