import React, {useContext} from 'react';
import './LoginButton.scss'
import {userSession} from "../UserSession";

export default function LoginButton() {
    const [user, setUser] = useContext(userSession)

    return (
        user == null ? (
            <button onClick={() => setUser({name: "Johnny Test"})}>Login</button>
        ) : (
            <button onClick={() => setUser({name: null})}>Log out</button>
        )
    );

}
