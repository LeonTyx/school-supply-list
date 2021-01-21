import React, {useState} from "react";

function User(props) {
    const roles = props.roles;
    const [user, setUser] = useState(JSON.parse(JSON.stringify(props.user)))

    function deleteUser() {
        const Http = new XMLHttpRequest();
        Http.open("DELETE", "./api/v1/user/" + props.user.user_id, true);
        Http.send();
    }

    function updateRoles() {
        const Http = new XMLHttpRequest();
        Http.open("POST", "./api/v1/user/" + props.user.user_id, true);
        Http.send(JSON.stringify(user));
    }

    function updateCheckBox(id) {
        let userCopy = Object.assign({}, user);
        if (userCopy.roles[id] == null) {
            userCopy.roles[id] = roles[id].name
            setUser(userCopy)
        } else {
            delete userCopy.roles[id]
            setUser(userCopy)
        }
    }

    return (
        <div>
            {user.name}
            <div className="roles">
                {Object.keys(roles).map((roleID) =>
                    <label key={roleID}>
                        <input type="checkbox"
                               checked={user.roles[roleID] != null}
                               onChange={() => updateCheckBox(roleID)}/>
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
