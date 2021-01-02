import React, {useContext} from 'react';
import './LoginButton.scss'
import {userSession} from "../UserSession";

export default function LoginButton() {
    const user = useContext(userSession)

    return (
        user.user.name != null ? (
            <button>Login</button>
        ):(
            <button>Log out</button>
        )
    );

}
