import React, {useContext} from 'react';
import './LoginButton.scss'
import {userSession} from "../../UserSession";

export default function LoginButton() {
    const [user, setUser] = useContext(userSession)

    function logout(){
        fetch("/oauth/v1/logout")
            .then(res => res.json())
            .then(
                () => {
                    setUser(null);
                    localStorage.removeItem("user")
                }, (error) => {
                    setUser(null)
                    localStorage.removeItem("user")
                }
            )
    }

    return (
        user == null ? (
            <a href={"./oauth/v1/login"}>Login</a>
        ) : (
            <button onClick={() => logout()}>Log out</button>
        )
    );

}
