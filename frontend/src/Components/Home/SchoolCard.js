import React from 'react';
import './SchoolCard.scss'

function SchoolCard(props) {
    return (
        <div className="school-card">
            <h3>{props.school.name}</h3>
        </div>
    );

}

export default SchoolCard;
