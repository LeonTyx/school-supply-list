import React from 'react';
import './Error.scss'

function Error(props) {
    return (
        <div className="error-message-prompt">
            {props.error_msg_str}
        </div>
    );

}

export default Error;
