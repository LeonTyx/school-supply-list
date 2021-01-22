import React, {useState} from 'react';

function Role(props) {
    const [role, setRoles] = useState(Object.assign({}, props.role))

    return (
        <div>
            <h3>{role.name}</h3>
            {role.resources != null &&
                Object.keys(role.resources).map((resourceKey) =>
                <div>Resource: {resourceKey}
                    <div>
                        <input type="checkbox" checked={role.resources[resourceKey].policy.can_add}/>
                        <input type="checkbox" checked={role.resources[resourceKey].policy.can_view}/>
                        <input type="checkbox" checked={role.resources[resourceKey].policy.can_edit}/>
                        <input type="checkbox" checked={role.resources[resourceKey].policy.can_delete}/>
                    </div>
                </div>
            )}
        </div>
    );

}

export default Role;
