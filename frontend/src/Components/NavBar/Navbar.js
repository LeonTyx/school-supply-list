import React from 'react';
import './Navbar.scss'
import {Link} from "react-router-dom";

function Navbar() {
    return (
        <nav>
            <Link to="/users">Users</Link>
        </nav>
    );

}

export default Navbar;
