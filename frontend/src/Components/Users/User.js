import React, {useState, useEffect} from "react";

function User(props) {
    const roles = props.roles;
    const [userRoles, setUserRoles] = useState(JSON.parse(JSON.stringify(props.user.roles)))

    function deleteUser(){
        const Http = new XMLHttpRequest();
        Http.open("DELETE", "./api/v1/user/"+ props.user.user_id, true);
        Http.send();
    }

    function updateRoles(){
        const Http = new XMLHttpRequest();
        Http.open("POST", "./api/v1/user/"+ props.user.user_id, true);
        Http.send(JSON.stringify(userRoles));
    }

    function updateCheckBox(id){
        let userRolesCopy = Object.assign({}, userRoles);
        if(userRolesCopy[id] == null){
            userRolesCopy[id] = roles[id].name
            setUserRoles(userRolesCopy)
        }else{
            delete userRolesCopy[id]
            setUserRoles(userRolesCopy)
        }
    }

    return (
        <div>
            {props.user.name}
            <div className="roles">
                {Object.keys(roles).map((roleID) =>
                    <label key={roleID}>
                        <input type="checkbox"
                               checked={userRoles[roleID] != null}
                               onChange={()=>updateCheckBox(roleID)}/>
                        {roles[roleID].name}
                    </label>
                )}
            </div>
            <button onClick={updateRoles}>Save Changes</button>
            <button onClick={deleteUser}>Delete</button>
        </div>
    );

}

export default User;
