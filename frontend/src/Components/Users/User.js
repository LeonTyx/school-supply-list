import React from 'react';

function User(props) {
    const user = props.user;
    const roles = props.roles;

    function deleteUser(){
        const Http = new XMLHttpRequest();
        Http.open("DELETE", "./api/v1/user/"+user.user_id, false);
        Http.send();
    }

    function updateRoles(){
        const Http = new XMLHttpRequest();
        Http.open("POST", "./api/v1/user/"+ user.user_id, false);
        Http.send();
    }

    return (
        <div>
            {user.name}
            <div className="roles">
                {roles.map((role) =>
                    <label key={role.id}>
                        <input type="checkbox" value={role.name}/>
                        {role.name}
                    </label>
                )}
            </div>
            <button onClick={updateRoles}>Save Changes</button>
            <button onClick={deleteUser}>Delete</button>
        </div>
    );

}

export default User;
