import React, {useEffect, useState} from 'react';
import Role from "./Role";
import './Roles.scss'

function Roles() {
    const [roles, setRoles] = useState(null);
    const [error, setError] = useState(null)
    function handleErrors(response, errorMessage) {
        if (!response.ok) {
            setError(errorMessage)
        }
        return response.json();
    }
    useEffect(() => {
        //Fetch roles from api
        fetch("/api/v1/role")
            .then((resp)=>handleErrors(resp, "Unable to delete user"))
            .then(
                (result) => {
                    setRoles(result);
                }, (error) => {
                    setRoles(null)
                    setError(error);
                }
            )
    }, [])

    return (
        error === null && roles != null &&
        <div>
            {Object.keys(roles).map((roleID) =>
                <Role role={roles[roleID]} key={roleID}/>
            )}
        </div>
    );
}

export default Roles;
