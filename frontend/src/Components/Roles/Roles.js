import React, {useEffect, useState} from 'react';
import Role from "./Role";
import './Roles.scss'
import CreateRole from "./CreateRole";

function Roles() {
    const [roles, setRoles] = useState(null)
    const [resources, setResources] = useState(null);

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
            .then((resp)=>handleErrors(resp, "Unable to retrieve roles"))
            .then(
                (result) => {
                    setRoles(result);
                }, (error) => {
                    setRoles(null)
                    setError(error);
                }
            )

        fetch("/api/v1/resource")
            .then((resp)=>handleErrors(resp, "Unable to retrieve resources"))
            .then(
                (result) => {
                    setResources(result);
                }, (error) => {
                    setResources(null)
                    setError(error);
                }
            )
    }, [])

    return (
        error === null && roles != null &&
        <div>
            <CreateRole resources={resources}/>
            {Object.keys(roles).map((roleID) =>
                <Role role={roles[roleID]} key={roleID}/>
            )}
        </div>
    );
}

export default Roles;
