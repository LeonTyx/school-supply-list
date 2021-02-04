import React from 'react';
import './SchoolCard.scss'
import {Link} from "react-router-dom";

function SchoolCard(props) {
    let school = props.school
    return (
        <div className="school-card">
            <Link to={"/school/" + school.school_id}>{school.school_name}</Link>
        </div>
    );

}

export default SchoolCard;
