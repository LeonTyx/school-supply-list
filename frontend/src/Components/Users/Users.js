import React, {useEffect, useState} from 'react';
import './Users.scss'
import DisplayError from "../Error/DisplayError";
import User from "./User";

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
            <div className="users">
                {users.map((user) =>
                    <User user={user} roles={roles} key={user.user_id}/>
                )}
            </div>
        ) : (
            <DisplayError error_msg_str={error}/>
        )
    );
}

export default Users;
