import React, {useContext} from 'react';
import './LoginButton.scss'
import {Link} from "react-router-dom";
import {userSession} from "../../UserSession";
import useDropdownMenu from 'react-accessible-dropdown-menu-hook';

export default function LoginButton() {
    const [user, setUser] = useContext(userSession)
    const { buttonProps, itemProps, isOpen } = useDropdownMenu(3);

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
            <a href={"./oauth/v1/login"} className="login-button">Login</a>
        ) : (
            <div>
                <button {...buttonProps} type='button' id='menu-button'>
                    <img src={user.picture} alt="profile"/>
                    <span>Account</span>
                    <i className='fal fa-angle-down' />
                </button>

                <div className={isOpen ? 'visible' : ''} role='menu' id='menu'>
                    <Link {...itemProps[0]} to='./account' id='menu-item-1'>
                        View Account
                    </Link>

                    <a {...itemProps[1]} onClick={()=>logout()} id='menu-item-2'>
                        Log Out
                    </a>
                </div>
            </div>
        )
    );

}
