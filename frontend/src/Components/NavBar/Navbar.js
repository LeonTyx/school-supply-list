import React from 'react';
import './Navbar.scss'
import {NavLink} from "react-router-dom";

function Navbar() {
    return (
        <nav>
            <NavLink to="/users" activeClassName="active">Users</NavLink>
            <NavLink to="/roles" activeClassName="active">Roles</NavLink>
        </nav>
    );

}

export default Navbar;
