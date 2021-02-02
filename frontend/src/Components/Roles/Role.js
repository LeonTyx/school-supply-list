import React, {useState} from 'react';
import DisplayError from "../Error/DisplayError";

function Role(props) {
    const [role, setRole] = useState(Object.assign({}, props.role))
    const [updating, setUpdating] = useState(false)
    const [isDeleted, setIsDeleted] = useState(false)

    const [error, setError] = useState(null)

    function handleErrors(response, errorMessage) {
        if (!response.ok) {
            setError(errorMessage)
        }
        return response;
    }

    function onTogglePolicy(resource, policyID) {
        let roleCopy = Object.assign({}, props.role)
        role.resources[resource].policy[policyID] = !role.resources[resource].policy[policyID]
        setRole(roleCopy)
    }

    function updateRole() {
        setUpdating(true)
        fetch("./api/v1/role/" + props.role.id, {method: "POST", body: JSON.stringify(role)})
            .then((resp) => handleErrors(resp, "Unable to update role"))
            .then(() => setUpdating(false))
            .catch(error => setError(error));
    }

    function deleteRole() {
        fetch("./api/v1/role/" + props.role.id, {method: "DELETE"})
            .then((resp) => handleErrors(resp, "Unable to delete role"))
            .then(() => setIsDeleted(true))
            .catch(error => setError(error));
    }

    return (
        !isDeleted &&
        <div className="role">
            <h3>{role.name}</h3>
            {role.resources != null &&
            Object.keys(role.resources).map((resourceKey) =>
                <div key={resourceKey}> Resource: {resourceKey}
                    <div>
                        <input type="checkbox"
                               checked={role.resources[resourceKey].policy.can_add}
                               onChange={() => onTogglePolicy(resourceKey, "can_add")}/>
                        <input type="checkbox"
                               checked={role.resources[resourceKey].policy.can_view}
                               onChange={() => onTogglePolicy(resourceKey, "can_view")}/>
                        <input type="checkbox"
                               checked={role.resources[resourceKey].policy.can_edit}
                               onChange={() => onTogglePolicy(resourceKey, "can_edit")}/>
                        <input type="checkbox"
                               checked={role.resources[resourceKey].policy.can_delete}
                               onChange={() => onTogglePolicy(resourceKey, "can_delete")}/>
                    </div>
                </div>
            )}
            {updating ? (
                <button disabled={true}>Saving</button>
            ) : (
                <button onClick={updateRole}>Save Changes</button>
            )}
            <button onClick={deleteRole}>Delete Role</button>
            {error !== null && <DisplayError error_msg_str={error}/>}
        </div>
    );

}

export default Role;
