import React, {useContext} from 'react';
import './CreateSchool.scss'
import {userSession} from "../../UserSession";

function CreateSchool() {
    const [user] = useContext(userSession);

    return (
        <div>
            <h3></h3>
        </div>
    );

}

export default CreateSchool;
