import React, {useEffect, useState} from 'react';
import './Users.scss'

function Users() {
    const [users, setUsers] = useState(null);
    const [error, setError] = useState(null)

    useEffect(() => {
        //Fetch user from api
        fetch("/api/v1/user")
            .then((res) => {
                if (res.ok) {
                    return res.json()
                }
            })
            .then(
                (result) => {
                    setUsers(result);
                    console.log(result)
                }, (error) => {
                    setUsers(null)
                    setError(error);
                }
            )
    }, [])

    return (
        error != null &&
        <div>
            {users.map((user) =>
                <div>{user.name}</div>
            )}
        </div>
    );
}

export default Users;
