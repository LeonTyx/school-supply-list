import React, {useState} from 'react';

function Role(props) {
    const [role, setRole] = useState(Object.assign({}, props.role))

    function updatePolicy(resource, policyID){
        let roleCopy = Object.assign({}, props.role)
        role.resources[resource].policy[policyID] = !role.resources[resource].policy[policyID]
        setRole(roleCopy)
    }

    return (
        <div>
            <h3>{role.name}</h3>
            {role.resources != null &&
                Object.keys(role.resources).map((resourceKey) =>
                <div key={resourceKey}> Resource: {resourceKey}
                    <div>
                        <input type="checkbox"
                               checked={role.resources[resourceKey].policy.can_add}
                               onChange={()=>updatePolicy(resourceKey, "can_add")}/>
                        <input type="checkbox"
                               checked={role.resources[resourceKey].policy.can_view}
                               onChange={()=>updatePolicy(resourceKey, "can_view")}/>
                        <input type="checkbox"
                               checked={role.resources[resourceKey].policy.can_edit}
                               onChange={()=>updatePolicy(resourceKey, "can_edit")}/>
                        <input type="checkbox"
                               checked={role.resources[resourceKey].policy.can_delete}
                               onChange={()=>updatePolicy(resourceKey, "can_delete")}/>
                    </div>
                </div>
            )}
        </div>
    );

}

export default Role;
