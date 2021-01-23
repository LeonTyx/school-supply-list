import React, {useContext} from 'react';
import './Navbar.scss'
import {NavLink} from "react-router-dom";
import {userSession} from "../../UserSession";

function Navbar() {
    const [user] = useContext(userSession)
    let userResources = () => {
        if(user !== null && user !== undefined) {
            return user.consolidated_resources
        }
        return undefined
    }

    return (
        <nav>
            <NavLink to="/" exact activeClassName="active">Home</NavLink>
            {userResources() !== undefined &&
            <React.Fragment>
                {userResources.user.policy.can_view !== undefined &&
                <NavLink to="/users" activeClassName="active">Users</NavLink>
                }
                {userResources.user.policy.can_view !== undefined &&
                <NavLink to="/roles" activeClassName="active">Roles</NavLink>
                }
            </React.Fragment>
            }
        </nav>
    );

}

export default Navbar;
