import React, {useState, useEffect} from 'react';
import './School.scss'
import DisplayError from "../Error/DisplayError";

function School(props) {
    const [school, setSchool] = useState(null)

    const [error, setError] = useState(null)
    function handleErrors(response, errorMessage) {
        if (!response.ok) {
            setError(errorMessage)
        }
        return response.json();
    }

    useEffect(() => {
        //Fetch school from api
        fetch("/api/v1/school/" + props.match.params.id)
            .then((resp) => handleErrors(resp, "Unable to retrieve school information"))
            .then(
                (result) => {
                    setSchool(result);
                }, (error) => {
                    setSchool(null)
                    setError(error);
                }
            )

    }, [props.match.params.id])
    return (
        error == null ? (
            school != null && <div className="school">
                <div className="school-header">
                    <h2>{school.school_name}</h2>
                </div>
                <div className="supply-lists">
                    {school.supply_list != null ? (
                        school.supply_lists.map((list) =>
                                <div>List goes here</div>
                            )
                    ): (
                        <div>No lists yet!</div>
                    )}
                </div>
            </div>
        ):(
            <DisplayError msg={error}/>
        )
    );

}

export default School;
