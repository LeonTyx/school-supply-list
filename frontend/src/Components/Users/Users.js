import React, {useEffect, useState} from 'react';
import './Users.scss'
import Error from "../Error/Error";

function Users() {
    const [users, setUsers] = useState(null);
    const [roles, setRoles] = useState(null);
    const [error, setError] = useState(null);

    useEffect(() => {
        //Fetch user from api
        fetch("/api/v1/user")
            .then((res) => {
                if (res.ok) {
                    return res.json()
                } else {
                    setError("error");
                }
            })
            .then(
                (result) => {
                    setUsers(result);
                }, (error) => {
                    setUsers(null)
                    setError(error);
                }
            )

        fetch("/api/v1/role")
            .then((res) => {
                if (res.ok) {
                    return res.json()
                } else {
                    setError("error");
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
        error === null && users !== null && roles !== null ? (
            <div>
                {users.map((user) =>
                    <div>{user.name}</div>
                )}
                <div>
                    {roles.map((role) =>
                        <div>{role.name}</div>
                    )}
                </div>
            </div>
        ) : (
            <Error error_msg_str={error}/>
        )
    );
}

export default Users;
