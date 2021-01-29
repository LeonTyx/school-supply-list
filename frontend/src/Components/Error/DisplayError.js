import React from 'react';
import './Error.scss'

function DisplayError(props) {
    return (
        <div className="error-message-prompt">
            {props.msg}
        </div>
    );

}

export default DisplayError;
