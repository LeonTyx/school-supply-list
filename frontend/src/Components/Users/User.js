import React, {useState} from "react";
import Error from "../Error/Error";

function User(props) {
    const roles = props.roles;
    const [updating, setUpdating] = useState(false)
    const [user, setUser] = useState(JSON.parse(JSON.stringify(props.user)))

    const [error, setError] = useState(null)
    function handleErrors(response) {
        if (!response.ok) {
            setError("Unable to retrieve request")
        }
        return response;
    }

    function deleteUser() {
        fetch("./api/v1/user/"+props.user.user_id, {method:"DELETE"})
            .then(handleErrors)
            .then(response => console.log("ok") )
            .catch(error => setError(error) );
    }

    function updateRoles() {
        setUpdating(true)
        fetch("./api/v1/user/"+ props.user.user_id, {
            method:"POST",
            body: JSON.stringify(user)
        })
            .then(handleErrors)
            .then(response => setUpdating(false) )
            .catch(error => setError(error) );
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
            {updating ? (
                <button disabled={true}>Saving</button>
            ) : (
                <button onClick={updateRoles} >Save Changes</button>
            )}
            <button onClick={deleteUser}>Delete</button>
            {error != null && <Error error_msg_str={error}/>}
        </div>
    );

}

export default User;
