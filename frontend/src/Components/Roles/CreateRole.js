import React, {useState} from 'react';
import './CreateRole.scss'

function CreateRole(props) {
    const [roleName, setRoleName] = useState("")
    const [roleDesc, setRoleDesc] = useState("")
    const [updating, setUpdating] = useState(false)
    const [roleResources, setRoleResources] = useState(Object.assign({}, props.resources))

    const [error, setError] = useState(null)

    function handleErrors(response, errorMessage) {
        if (!response.ok) {
            setError(errorMessage)
        }
        return response;
    }


    function updateResources(resourceKey, policyKey) {
        let resourceCopy = Object.assign({}, props.resources)
        resourceCopy[resourceKey].policy[policyKey] = !resourceCopy[resourceKey].policy[policyKey]

        setRoleResources(resourceCopy)
    }

    function submitForm() {
        setUpdating(true)
        fetch("/api/v1/role/", {
            method: "PUT",
            body: JSON.stringify({
                "name": roleName,
                "desc": roleDesc,
                "resources": roleResources
            })
        })
            .then((resp) => handleErrors(resp, "Unable to create role"))
            .then(response => {
                setRoleName("")
                setRoleDesc("")
                setRoleResources(Object.assign({}, props.resources))
                setUpdating(false)
            })
            .catch(error => setError(error));
    }

    return (
        error === null &&
        <form className="create-role"
              onSubmit={(e) => e.preventDefault()}>
            <h2>Create Role</h2>
            <label>
                Role name
                <input type="text"
                       value={roleName}
                       onChange={(e) => {
                           setRoleName(e.target.value)
                       }}/>
            </label>

            <label>
                Role description
                <textarea value={roleDesc}
                          onChange={(e) => {
                              setRoleDesc(e.target.value)
                          }}/>
            </label>

            <div className="resources">
                Resources

                {Object.keys(roleResources).map((key) =>
                    <div className="policy" key={key}>
                        <div className="resource">{key}</div>
                        <label>
                            <input type="checkbox" checked={roleResources[key].policy.can_add}
                                   onChange={() => {
                                       updateResources(key, "can_add")
                                   }}/>
                            Create
                        </label>

                        <label>
                            <input type="checkbox" checked={roleResources[key].policy.can_view}
                                   onChange={() => {
                                       updateResources(key, "can_view")
                                   }}/>

                            Read
                        </label>

                        <label>
                            <input type="checkbox" checked={roleResources[key].policy.can_edit}
                                   onChange={() => {
                                       updateResources(key, "can_edit")
                                   }}/>

                            Update
                        </label>

                        <label>
                            <input type="checkbox" checked={roleResources[key].policy.can_delete}
                                   onChange={() => {
                                       updateResources(key, "can_delete")
                                   }}/>

                            Delete
                        </label>
                    </div>
                )}
            </div>
            {updating ? (
                <button disabled={true}>Saving</button>
            ) : (
                <button onClick={submitForm}>Create Role</button>
            )}
        </form>
    );

}

export default CreateRole;
