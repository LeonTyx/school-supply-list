import React, {useEffect, useState} from 'react';
import './Role.scss'

function Role() {
    const [roles, setRoles] = useState(null);
    const [error, setError] = useState(null)

    useEffect(() => {
        //Fetch user from api
        fetch("/api/v1/roles")
            .then((res) => {
                if (res.ok) {
                    return res.json()
                }
            })
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
        error === null &&
        <div>
            {roles.map((role) =>
                <div>{role.role_name}</div>
            )}
        </div>
    );
}

export default Role;
