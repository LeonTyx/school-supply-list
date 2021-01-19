import React from 'react';

function User(props) {
    const user = props.user;
    function deleteUser(){
        const Http = new XMLHttpRequest();
        Http.open("DELETE", "./api/v1/user/"+user.user_id, false);
        Http.send();
    }

    return (
        <div>
            {user.name}
            <button onClick={deleteUser}>Delete</button>
        </div>
    );

}

export default User;
