import React, {useState} from 'react';
import './SchoolCard.scss'
import {Link} from "react-router-dom";

function SchoolCard(props) {
    const [updating, setUpdating] = useState(false)
    let school = props.school

    const [error, setError] = useState(null)

    function handleErrors(response, errorMessage) {
        if (!response.ok) {
            setError(errorMessage)
        }
        return response;
    }

    function deleteSchool(){
        setUpdating(true)
        fetch("/api/v1/school/"+school.school_id, {
            method: "DELETE"
        })
            .then((resp) => handleErrors(resp, "Unable to delete school"))
            .then(() => {
                setUpdating(false)
                props.removeFromList(school.school_id)
            })
            .catch(error => setError(error));
    }

    return (
        error === null &&
        <div className="school-card">
            <Link to={"/school/" + school.school_id}>{school.school_name}</Link>
            {updating ?
                <button disabled={true}>Removing...</button>
                :
                <button className={"delete"}
                        onClick={deleteSchool}>
                    Remove
                </button>
            }
        </div>
    );

}

export default SchoolCard;
