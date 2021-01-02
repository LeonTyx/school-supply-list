import React, { useContext } from "react";
import {userContext} from './userContext';

function Main() {
    const user = useContext(userContext)
    return (
        <div>{user.user.name}</div>
    );

}

export default Main;
