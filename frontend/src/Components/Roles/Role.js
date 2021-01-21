import React, {useState} from 'react';

function Role(props) {
    const [role, setRoles] = useState(Object.assign({}, props.role))

    return (
        <div>
            {role.name}
        </div>
    );

}

export default Role;
