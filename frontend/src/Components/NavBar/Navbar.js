import React, {useContext} from 'react';
import './Navbar.scss'
import {NavLink} from "react-router-dom";
import {userSession} from "../../UserSession";

function Navbar() {
    const [user] = useContext(userSession)

    let canView = (key) => {
        if(user !== null && user !== undefined) {
            let resc = user.consolidated_resources
            if(resc[key] !== undefined && resc[key].policy.can_view){
                return true
            }
        }
        return false
    }

    return (
        <nav>
            <NavLink to="/" exact activeClassName="active">Home</NavLink>
            <React.Fragment>
                {canView("user") &&
                <NavLink to="/users" activeClassName="active">Users</NavLink>
                }
                {canView("role") !== undefined &&
                <NavLink to="/roles" activeClassName="active">Roles</NavLink>
                }
            </React.Fragment>
        </nav>
    );

}

export default Navbar;
