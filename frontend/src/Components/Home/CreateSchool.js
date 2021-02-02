import React, {useContext, useState} from 'react';
import './CreateSchool.scss'
import {userSession} from "../../UserSession";
import {canCreate} from "../Permissions/Permissions";

function CreateSchool() {
    const [schoolName, setSchoolName] = useState("")
    const [updating, setUpdating] = useState(false)
    const [user] = useContext(userSession)

    const [error, setError] = useState(null)
    function handleErrors(response, errorMessage) {
        if (!response.ok) {
            setError(errorMessage)
        }
        return response.json();
    }

    function submitForm() {
        setUpdating(true)
        fetch("/api/v1/school", {
            method: "PUT",
            body: JSON.stringify({
                "school_name": schoolName,
            })
        })
            .then((resp) => handleErrors(resp, "Unable to create school"))
            .then(() => {
                setSchoolName("")
                setUpdating(false)
            })
            .catch(error => setError(error));
    }

    return (
        error === null &&
        canCreate("school", user) && (
            <form className="create-school"
                  onSubmit={(e) => e.preventDefault()}>
                <h2>Create School</h2>
                <label>
                    School Name
                    <input value={schoolName} onChange={(e)=>setSchoolName(e.target.value)}/>
                </label>
                {updating ? (
                    <button disabled={true}>Updating</button>
                ) : (
                    <button onClick={submitForm}>Submit</button>
                )}
            </form>
        )
    );

}

export default CreateSchool;
