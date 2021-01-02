import React, {useContext} from 'react';
import './LoginButton.scss'
import {userSession} from "../UserSession";

export default function LoginButton() {
    const [user, setUser] = useContext(userSession)

    function logout(){
        setUser(null)
        localStorage.removeItem('user')
    }

    function login(){
        fetch("/oauth/v1/login")
            .then(res => res.json())
            .then(
                (result) => {
                    setUser(result);
                    localStorage.setItem("user", result)
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
